package shell

import (
	"fmt"
	"math"
	"os"

	"gonum.org/v1/gonum/mat"
)

// Model is structural calculation model
type Model struct {
	// Points is slice of point coordinate
	//
	//	[0] - X coordinate
	//	[1] - Y coordinate
	//
	Points [][2]float64

	// Beams is slice of point index and beam property
	Beams []BeamProp

	// Pins is slice of pins for beams in local system coordinate.
	// Len of support must be same amount of beam.
	// Or if len is zero, then all DoF(degree of freedom) is rigid.
	//
	// first index is point index
	//
	//	[0] - X on start point
	//	[1] - Y on start point
	//	[2] - M on start point
	//	[3] - X on end point
	//	[4] - Y on end point
	//	[5] - M on end point
	//
	// if `true` then free degree of freedom
	// Pins [][6]bool

	// Supports is slice of fixed supports.
	// Len of support must be same amount of Points
	//
	// first index is point index
	//
	//	[0] - X
	//	[1] - Y
	//	[2] - M
	//
	Supports [][3]bool

	Ln []LoadNode
}

// BeamProp is beam property
type BeamProp struct {
	// Start and end point index
	//
	//	[0] - start of beam
	//	[1] - end of beam
	//
	N [2]int

	Mat int
	T   float64

	E1, E2, G, nu1, nu2 float64

	// A cross-section area
	// Unit : sq. meter.
	// A float64

	// J is moment inertia
	// Unit : meter^4
	// J float64

	// E is modulus of elasticity
	// Unit : Pa
	// E float64
}

// LoadNode is node load on specific point in global system coordinate
type LoadNode struct {
	// N is point index
	N int

	// Forces is node loads on each direction
	//
	//	[0] - X , Unit: N. Positive direction from left to right.
	//	[1] - Y , Unit: N. Positive direction from down to top.
	//	[2] - M , Unit: N*m. Positive direction is counter-clockwise direction.
	//
	Forces [3]float64 // orientation Fw=0, Fu=1, Mpho=2
}

func (m Model) check_elem_data() {
	for i := range m.Beams {
		if m.Points[m.Beams[i].N[0]][1] > m.Points[m.Beams[i].N[1]][1] {
			m.Beams[i].N[0], m.Beams[i].N[1] = m.Beams[i].N[1], m.Beams[i].N[0]
		}
	}
}

// computes material stiffness matrix of elemen
func (m Model) get_D_matrix(i int) (D *mat.Dense) {
	D = mat.NewDense(5, 5, nil)
	var (
		E1   = m.Beams[i].E1
		E2   = m.Beams[i].E2
		G    = m.Beams[i].G
		nu1  = m.Beams[i].nu1
		nu2  = m.Beams[i].nu2
		t    = m.Beams[i].T
		mult = t / (1.0 - nu1*nu2)
	)
	D.Set(0, 0, E1*mult)
	D.Set(0, 1, nu2*mult)
	D.Set(1, 0, nu2*mult)
	D.Set(1, 1, E2*mult)
	D.Set(2, 2, E1*t*t/12.*mult)
	D.Set(3, 3, E2*t*t/12.*mult)
	D.Set(2, 3, nu2*(E1*t*t)/12.*mult)
	D.Set(3, 2, nu2*(E1*t*t)/12.*mult)
	D.Set(4, 4, 5.0/6.0*G/t)
	return D
}

// get_B_matrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:704
func (m Model) get_B_matrix(i int) (B *mat.Dense, L float64, R float64) {
	B = mat.NewDense(5, 6, nil)
	var (
		dx = m.Points[m.Beams[i].N[1]][0] - m.Points[m.Beams[i].N[0]][0]
		dy = m.Points[m.Beams[i].N[1]][1] - m.Points[m.Beams[i].N[0]][1]
	)
	L = math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	R = 0.5 * (m.Points[m.Beams[i].N[1]][0] + m.Points[m.Beams[i].N[0]][0])

	var (
		S = -1 * dx / L
		C = -1 * dy / L
	)
	// B matrix:
	B.Set(0, 0, -1.*C/L)
	B.Set(0, 1, -1.*S/L)
	B.Set(0, 3, 1.*C/L)
	B.Set(0, 4, 1.*S/L)
	B.Set(1, 1, 1./(2*R))
	B.Set(1, 4, 1./(2.*R))
	B.Set(2, 2, -1./L)
	B.Set(2, 5, 1./L)
	B.Set(3, 2, S/(2.*R))
	B.Set(3, 5, S/(2.*R))
	B.Set(4, 0, -1.*S/L)
	B.Set(4, 1, 1.*C/L)
	B.Set(4, 2, 1./2.)
	B.Set(4, 3, 1.*S/L)
	B.Set(4, 4, -1.*C/L)
	B.Set(4, 5, 1./2.)
	return B, L, R
}

// creates stiffness matrix
func (m Model) get_matrix() (K, F, u *mat.Dense) {
	n_n := len(m.Points)
	K = mat.NewDense(n_n*3, n_n*3, nil)
	u = mat.NewDense(n_n*3, 1, nil)
	F = mat.NewDense(n_n*3, 1, nil)
	for i := 0; i < len(m.Beams); i++ {
		// material stiffness matrix D:
		D := m.get_D_matrix(i)
		// B matrix
		B, L, R := m.get_B_matrix(i)
		// transpose of B
		Bt := B.T()

		// matrix multiplications (Bt*D*B):
		BtD := mat.NewDense(6, 5, nil)
		BtD.Mul(Bt, D)

		// => Ke  without L*R
		Ke := mat.NewDense(6, 6, nil)
		Ke.Mul(BtD, B)

		// element stifness matrix Ke:
		Ke.Scale(R*L, Ke)

		// localisation to "K":
		for j := 1; j <= 6; j++ {
			var posj int
			var posk int
			if j < 4 {
				posj = m.Beams[i].N[0]*3 + j
			} else {
				posj = m.Beams[i].N[1]*3 + j - 3
			}
			for k := 1; k <= 6; k++ {
				if k < 4 {
					posk = m.Beams[i].N[0]*3 + k
				} else {
					posk = m.Beams[i].N[1]*3 + k - 3
				}
				// femMatPutAdd((&K), posj, posk, femMatGet((&Ke), j, k), 1)
				K.Set(posj-1, posk-1, K.At(posj-1, posk-1)+Ke.At(j-1, k-1))
			}
		}
	}
	return
}

func ZeroCol(m *mat.Dense, pos int) {
	v, _ := m.Caps()
	for i := 0; i < v; i++ {
		m.Set(i, pos, 0)
	}
}

func ZeroRow(m *mat.Dense, pos int) {
	_, v := m.Caps()
	for i := 0; i < v; i++ {
		m.Set(pos, i, 0)
	}
}

// applies supports in nodes
func (m Model) get_loads_and_supports(K, F, u *mat.Dense) {

	for i := range m.Ln {
		for g := range m.Ln[i].Forces {
			w := m.Ln[i].N*3 + g
			F.Set(w, 0, F.At(w, 0)+m.Ln[i].Forces[g])
		}
	}

	for n := range m.Supports {
		for d := range m.Supports[n] {
			if !m.Supports[n][d] {
				continue
			}

			d_n := n
			d_dir := d

			pos := d_n*3 + d_dir
			ZeroCol(K, pos)
			ZeroRow(K, pos)
			u.Set(pos, 0, 0)
			F.Set(pos, 0, 0)
			K.Set(pos, pos, 1)
		}
	}

	// 	for i = 0; i < n_d; i++ {
	// 		// 		if d_dir[i] > 2 {
	// 		// 			// stifnesses
	// 		// 			pos = d_n[i]*3 + d_dir[i] - 2
	// 		// 			femMatPutAdd((&K), pos, pos, d_val[i], 1)
	// 		// 		} else {
	// 		// displacements
	// 		pos = d_n[i]*3 + d_dir[i] + 1
	// 		if math.Abs(d_val[i]) <= 1e-07 {
	// 			femMatSetZeroCol((&K), pos)
	// 			femMatSetZeroRow((&K), pos)
	// 			femVecPutAdd((&u), pos, 0, 0)
	// 			// yes, it deletes force in support
	// 			femVecPutAdd((&F), pos, 0, 0)
	// 			femMatPutAdd((&K), pos, pos, 1, 0)
	// 		} else {
	// 			for j = 1; j <= n_n*3; j++ {
	// 				femVecPutAdd((&F), j, -1*femMatGet((&K), j, pos)*d_val[i], 1)
	// 			}
	// 			femMatSetZeroCol((&K), pos)
	// 			femMatSetZeroRow((&K), pos)
	// 			femVecPutAdd((&u), pos, d_val[i], 0)
	// 			femMatPutAdd((&K), pos, pos, femVecGet((&F), pos)/d_val[i], 0)
	// 		}
	// 		// 	}
	// 	}
}

// computes internal force is nodes
//	el element number <0..n_e-1>
//	N1 meridian force
//	N2 perpendicular force
//	M1 meridian moment
//	M2 perpendicular force
//	Q tangent force
func (m Model) get_int_forces(el int, u *mat.Dense) (N1, N2, M1, M2, Q float64) {

	ue := mat.NewDense(6, 1, nil)
	Fe := mat.NewDense(5, 1, nil)

	// get local stiffness vector
	for j := 1; j <= 6; j++ {
		var posj int
		if j < 4 {
			posj = m.Beams[el].N[0]*3 + j
		} else {
			posj = m.Beams[el].N[1]*3 + j - 3
		}
		ue.Set(j-1, 0, ue.At(j-1, 0)+u.At(posj-1, 0))
		// femVecPutAdd((&ue), j, femVecGet((&u), posj), 0)
	}

	// get B and D
	D := m.get_D_matrix(el)
	B, _, _ := m.get_B_matrix(el)

	DB := mat.NewDense(5, 6, nil)
	DB.Mul(D, B)
	Fe.Mul(DB, ue)

	N1 = Fe.At(0, 0)
	N2 = Fe.At(1, 0)
	M1 = Fe.At(2, 0)
	M2 = Fe.At(3, 0)
	Q = Fe.At(4, 0)
	return
}

func (m Model) print_result(u *mat.Dense) {
	fw := os.Stdout

	n_n := len(m.Points)
	n_e := len(m.Beams)

	fmt.Fprintf(fw, "#  X     Y        w            u           angle            N1          N2           M1          M2          Q\n")
	for i := 0; i < n_n; i++ {
		var (
			sN1   = 0.0
			sN2   = 0.0
			sM1   = 0.0
			sM2   = 0.0
			sQ    = 0.0
			count = 0
		)
		for j := 0; j < n_e; j++ {
			if m.Beams[j].N[0] == i || m.Beams[j].N[1] == i {
				// internal forces in centroid
				N1, N2, M1, M2, Q := m.get_int_forces(j, u)
				sN1 += N1
				sN2 += N2
				sM1 += M1
				sM2 += M2
				sQ += Q
				count++
			}
		}
		if count > 0 {
			sN1 /= float64(count)
			sN2 /= float64(count)
			sM1 /= float64(count)
			sM2 /= float64(count)
			sQ /= float64(count)
		}
		fmt.Fprintf(fw, string("%2.3f %2.3f %e %e %e %e %e %e %e %e\n"), m.Points[i][0], m.Points[i][1],
			u.At(3*i+1-1, 0),
			u.At(3*i+2-1, 0),
			u.At(3*i+3-1, 0),
			sN1, sN2, sM1, sM2, sQ)
	}
}
