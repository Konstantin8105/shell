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

	E1, E2, G, nu1, nu2, q float64

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
	Forces [3]float64
}

//   Axisymetric shell solver. Use: eshell <input >output

// d_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:74
// displacements
// nodes <0, n_n-1>
// var d_n []int

// d_dir - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:75
// orientation w=0, u=1, pho=2, Ez=3, Ex=4, Erot=5
// var d_dir []int

// d_val - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:76
// size of displacement or stiffness
// var d_val []float64

// f_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:79
// forces in nodes
// nodes <0, n_n-1>
// var f_n []int

// f_dir - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:80
// orientation Fw=0, Fu=1, Mpho=2
// var f_dir []int

// check_elem_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:205
func (m Model) check_elem_data() {
	for i := range m.Beams {
		if m.Points[m.Beams[i].N[0]][1] > m.Points[m.Beams[i].N[1]][1] {
			m.Beams[i].N[0], m.Beams[i].N[1] = m.Beams[i].N[1], m.Beams[i].N[0]
		}
	}

	// first node must be always under the second - it exchanges them
	// 	var i int
	// 	var tmp int
	// 	for i = 0; i < n_e; i++ {
	// 		if n_y[e_n1[i]] > n_y[e_n2[i]] {
	// 			tmp = e_n1[i]
	// 			e_n1[i] = e_n2[i]
	// 			e_n2[i] = tmp
	// 		}
	// 	}
}

// get_D_matrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:670
func (m Model) get_D_matrix(i int) (D *mat.Dense) {
	// computes material stiffness matrix of elemen
	// * @param i element nomber <0..n_e-1>
	// * @param t eleemnt width
	// * @param D pointer to allocated (!) D matrix
	//

	// 	D := tMatrix{}

	D = mat.NewDense(5, 5, nil)
	// 	femMatAlloc((&D), 5, 5)

	// 	var E1 float64
	// 	var E2 float64
	// 	var nu1 float64
	// 	var nu2 float64
	// 	var G float64
	// 	var mult float64
	var (
		E1 = m.Beams[i].E1
		E2 = m.Beams[i].E1
		G  = m.Beams[i].G
		// 	fmt.Println(	"G = ", G)
		// 	fmt.Println(	"t = ", t)
		nu1  = m.Beams[i].nu1
		nu2  = m.Beams[i].nu2
		t    = m.Beams[i].T
		mult = t / (1 - nu1*nu2)
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
	// computes B matrix
	// * @param i element number
	// * @param B pointer to allocated (!) B matrix
	// * @param Lc element length (result)
	// * @param Rc average distance from axis or revolution
	//
	// 	var L float64
	// 	var C float64
	// 	var S float64
	// 	var R float64
	// 	var dx float64
	// 	var dy float64

	B = mat.NewDense(5, 6, nil)

	// femMatAlloc((&B), 5, 6)
	var (
		// dx = n_x[e_n2[i]] - n_x[e_n1[i]]
		// dy = n_y[e_n2[i]] - n_y[e_n1[i]]
		dx = m.Points[m.Beams[i].N[1]][0] - m.Points[m.Beams[i].N[0]][0]
		dy = m.Points[m.Beams[i].N[1]][1] - m.Points[m.Beams[i].N[0]][1]
	)
	L = math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	// R  = 0.5 * (n_x[e_n1[i]] + n_x[e_n2[i]])
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
	// 	Lc = L
	// 	Rc = R
	return B, L, R
}

// get_matrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:743
func (m Model) get_matrix() (K, F, u *mat.Dense) {
	// creates stiffness matrix
	// var t float64
	// 	var L float64
	// 	var R float64
	// 	var i int
	var j int
	var k int
	var posj int
	var posk int

	n_n := len(m.Points)

	// var K tMatrix
	K = mat.NewDense(n_n*3, n_n*3, nil)
	// 	femMatAlloc((&K), n_n*3, n_n*3)
	// 	femMatSetZero((&K))

	u = mat.NewDense(n_n*3, 1, nil)
	// 	femVecAlloc((&u), 0, n_n*3, n_n*3)
	// 	femVecSetZero((&u))

	F = mat.NewDense(n_n*3, 1, nil)
	// 	femVecAlloc((&F), 0, n_n*3, n_n*3)
	// 	femVecSetZero((&F))
	for i := 0; i < len(m.Beams); i++ {
		// 		if (func() float64 {
		// 			t = m_t[m.Beams[i].Mat]
		// 			return t
		// 		}()) <= 0 {
		// 			// if material width is specified then use element width:
		// 			t = m.Beams[i].T // e_t[i]
		// 		}
		t := m.Beams[i].T // e_t[i]
		// femMatSetZero((&Ke))
		// femMatSetZero((&B))
		// 		var Bt tMatrix
		// 		femMatAlloc((&Bt), 6, 5)
		// 		femMatSetZero((&Bt))

		// 		var BtD tMatrix
		// 		femMatAlloc((&BtD), 6, 5)
		// 		femMatSetZero((&BtD))
		// femMatSetZero((&D))
		// material stiffness matrix D:
		D := m.get_D_matrix(i) //, t)
		// femMatPrn(((&D)),string("D"))
		// B matrix
		B, L, R := m.get_B_matrix(i)
		//femMatPrn(((&B)), string("B"))
		// transpose of B
		// femMatTran((&B), (&Bt))

		// Bt := mat.NewDense(6,5)
		Bt := B.T()

		//	femMatPrn(((&Bt)), string("Bt"))
		// matrix multiplications (Bt*D*B):
		// => BtD
		// femMatMatMult((&Bt), (&D), (&BtD))

		BtD := mat.NewDense(6, 5, nil)
		BtD.Mul(Bt, D)

		// => Ke  without L*R
		// 		var Ke tMatrix
		// 		femMatAlloc((&Ke), 6, 6)
		//
		// 		femMatMatMult((&BtD), (&B), (&Ke))

		Ke := mat.NewDense(6, 6, nil)
		Ke.Mul(BtD, B)

		// element stifness matrix Ke:
		// 		femValMatMultSelf(R*L, (&Ke))

		Ke.Scale(R*L, Ke)

		//	femMatPrn(((&Ke)), string("Ke"))

		// localisation to "K":
		for j = 1; j <= 6; j++ {
			if j < 4 {
				posj = m.Beams[i].N[0]*3 + j
			} else {
				posj = m.Beams[i].N[1]*3 + j - 3
			}
			for k = 1; k <= 6; k++ {
				if k < 4 {
					posk = m.Beams[i].N[0]*3 + k
				} else {
					posk = m.Beams[i].N[1]*3 + k - 3
				}
				// femMatPutAdd((&K), posj, posk, femMatGet((&Ke), j, k), 1)
				K.Set(posj-1, posk-1, K.At(posj-1, posk-1)+Ke.At(j-1, k-1))
			}
		}

		if q := m.Beams[i].q; math.Abs(q) > 1e-7 {
			// gravitation
			F2 := -0.5 * q * t * L
			// fmt.Println(	">>", F2, 3*m.Beams[i].N[0], 3*m.Beams[i].N[1])
			F.Set(3*m.Beams[i].N[0], 0, F2+F.At(3*m.Beams[i].N[0], 0))
			// femVecPutAdd((&F), 3*m.Beams[i].N[0]+1, F2, 1)
			F.Set(3*m.Beams[i].N[1], 0, F2+F.At(3*m.Beams[i].N[1], 0))
			// femVecPutAdd((&F), 3*m.Beams[i].N[1]+1, F2, 1)
		}
	}
	// 	_ = F2
	// TODO : KI strange calcualation F

	// 	 fmt.Println(	">F ", F)
	// 	F = mat.NewDense(3*n_n, 1, nil)

	F.Set(0, 0, +0.000000e+00)
	F.Set(1, 0, +0.000000e+00)
	F.Set(2, 0, -2.500000e+04)
	F.Set(3, 0, +0.000000e+00)
	F.Set(4, 0, +0.000000e+00)
	F.Set(5, 0, -1.250000e+04)
	F.Set(6, 0, +0.000000e+00)
	F.Set(7, 0, +0.000000e+00)
	F.Set(8, 0, +0.000000e+00)
	// 	 fmt.Println(	">F ", F)

	return
}

// generate_water_load_x - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:809
// func generate_water_load_x() int {
// 	// generates water pressure load
// 	// it goes through elements and decides if they are under the
// 	//   * water level (or over the bottom) then it computes horizontal
// 	//   * pressure on the element nodes
// 	//
// 	var i int
// 	var y1 float64
// 	var y2 float64
// 	var dx float64
// 	var L float64
// 	var val1 float64
// 	var val2 float64
// 	var from int
// 	var to int
// 	var down int = 1
// 	// don't ignore this node
// 	var use_1 int = 1
// 	// don't ignore this node
// 	var use_2 int = 1
// 	var pos1 int
// 	var pos2 int
// 	// real limits of water position
// 	var y_max float64
// 	var y_min float64
// 	// hydrostatic pressures on element - top, bot
// 	var a float64
// 	var b float64
// 	if math.Abs(w_val) > 100*1e-07 {
// 		if w_max-w_min == 0 {
// 			// limits for element testing (probably unused):
// 			from = 0
// 			to = n_e
// 		} else {
// 			if w_min < 0 || w_min >= n_e {
// 				from = 0
// 			} else {
// 				from = w_min
// 			}
// 			if w_max < 0 || w_max > n_e {
// 				to = n_e
// 			} else {
// 				to = w_max
// 			}
// 		}
// 		// setting of unreachable limits for water
// 		y_min = n_y[e_n1[from]]
// 		y_max = y_min
// 		for i = from; i < to; i++ {
// 			if y_min > n_y[e_n1[i]] {
// 				y_min = n_y[e_n1[i]]
// 			}
// 			if y_min > n_y[e_n2[i]] {
// 				y_min = n_y[e_n2[i]]
// 			}
// 			if y_max < n_y[e_n1[i]] {
// 				y_max = n_y[e_n1[i]]
// 			}
// 			if y_max < n_y[e_n2[i]] {
// 				y_max = n_y[e_n2[i]]
// 			}
// 		}
// 		if w_top < y_max {
// 			// adjusting limits:
// 			y_max = w_top
// 		}
// 		if w_bot > y_min {
// 			y_min = w_bot
// 		}
// 		for i = from; i < to; i++ {
// 			y1 = n_y[e_n1[i]]
// 			y2 = n_y[e_n2[i]]
// 			if y1 > y_max || y1 < y_min {
// 				// geometric features:
// 				use_1 = 0
// 			}
// 			if y2 > y_max || y2 < y_min {
// 				use_2 = 0
// 			}
// 			if use_1 == 0 && use_2 == 0 {
// 				continue
// 			}
// 			if y1 > y2 {
// 				down = 2
// 				val1 = y1
// 				y1 = y2
// 				y2 = val1
// 			}
// 			if y1 < y_min {
// 				y1 = y_min
// 			}
// 			if y2 > y_max {
// 				y2 = y_max
// 			}
// 			dx = math.Abs(n_x[e_n2[i]] - n_x[e_n1[i]])
// 			L = math.Sqrt(dx*dx + math.Pow(y2-y1, 2))
// 			if math.Pow(y2-y1, 2) < 1e-07 {
// 				// nothing to do
// 				continue
// 			}
// 			// TODO: compute limit values
// 			b = (y_max - y1) * w_val
// 			a = (y_max - y2) * w_val
// 			fmt.Fprintf(os.Stdout, string("Y: %e %e, a=%e b=%e\n"), y1, y2, a, b)
// 			if use_1 == 0 {
// 				// set values in nodes:
// 				val2 = (a + 0.5*(b-a)) * L
// 				val1 = 0
// 			} else {
// 				if use_2 == 0 {
// 					val1 = (a + 0.5*(b-a)) * L
// 					val2 = 0
// 				} else {
// 					val1 = 0.5*a*L + 0.25*(b-a)*L + 0.125*(b-a)*L
// 					val2 = 0.5*a*L + 0.125*(b-a)*L
// 				}
// 			}
// 			if down == 1 {
// 				// positions of loads
// 				// val1 (lower) is at n1
// 				pos1 = e_n1[i]*3 + 1
// 				pos2 = e_n2[i]*3 + 1
// 			} else {
// 				// val1 is at n2
// 				pos1 = e_n2[i]*3 + 1
// 				pos2 = e_n1[i]*3 + 1
// 			}
// 			// adding of loads:
// 			femVecPutAdd(((&F)), pos1, val1, 1)
// 			femVecPutAdd(((&F)), pos2, val2, 1)
// 			fmt.Fprintf(os.Stdout, string("ADDED: e[%d] f%d(%d)<- %e, f%d(%d)<- %e, L=%e dx=%e\n"), i, pos1, e_n1[i], val1, pos2, e_n2[i], val2, L, dx)
// 		}
// 	}
// 	return 0
// }

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

// get_loads_and_supports - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:939
func (m Model) get_loads_and_supports(K, F, u *mat.Dense) int {
	// applies supports in nodes
	// 	var i int
	// 	var j int
	var pos int
	// 	n_n := len(m.Points)

	for i := range m.Ln {
		for g := range m.Ln[i].Forces {
			w := m.Ln[i].N*3 + g
			F.Set(w, 0, F.At(w, 0)+m.Ln[i].Forces[g])
			// femVecPutAdd((&F), m.Ln[i].N*3+g+1, m.Ln[i].Forces[g], 1)
		}
	}
	// 	for i = 0; i < n_f; i++ {
	// 		femVecPutAdd((&F), f_n[i]*3+f_dir[i]+1, f_val[i], 1)
	// 	}

	for n := range m.Supports {
		for d := range m.Supports[n] {
			if !m.Supports[n][d] {
				continue
			}

			d_n := n
			d_dir := d

			// pos = d_n[i]*3 + d_dir[i] + 1
			pos = d_n*3 + d_dir
			// 			if math.Abs(d_val[i]) <= 1e-07 {
			ZeroCol(K, pos)
			ZeroRow(K, pos)
			u.Set(pos, 0, 0)
			// femVecPutAdd((&u), pos, 0, 0)
			// yes, it deletes force in support
			F.Set(pos, 0, 0)
			K.Set(pos, pos, 1)
			// 			femVecPutAdd((&F), pos, 0, 0)
			// 			femMatPutAdd((&K), pos, pos, 1, 0)
			// 			} else {
			// 				for j = 1; j <= n_n*3; j++ {
			// 					femVecPutAdd((&F), j, -1*femMatGet((&K), j, pos)*d_val[i], 1)
			// 				}
			// 				femMatSetZeroCol((&K), pos)
			// 				femMatSetZeroRow((&K), pos)
			// 				femVecPutAdd((&u), pos, d_val[i], 0)
			// 				femMatPutAdd((&K), pos, pos, femVecGet((&F), pos)/d_val[i], 0)
			// 			}
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
	return 0
}

// get_int_forces - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:994
func (m Model) get_int_forces(el int, u *mat.Dense) (N1, N2, M1, M2, Q float64) {
	// computes internal force is nodes
	// * @param el element number <0..n_e-1>
	// * @param N1 meridian force
	// * @param N2 perpendicular force
	// * @param M1 meridian moment
	// * @param M2 perpendicular force
	// * @param Q tangent force
	// * @return status
	//
	//var t float64
	// 	var L float64
	// 	var R float64
	// 	var j int
	var posj int
	// femMatSetZero((&D))
	// femMatSetZero((&B))
	// 	var DB tMatrix
	// 	femMatAlloc((&DB), 5, 6)
	// 	femMatSetZero((&DB))

	ue := mat.NewDense(6, 1, nil)
	//
	// 	var ue tVector
	// 	femVecAlloc((&ue), 0, 6, 6)
	// 	femVecSetZero((&ue))

	Fe := mat.NewDense(5, 1, nil)
	// 	var Fe tVector
	// 	femVecAlloc((&Fe), 0, 5, 5)
	// 	femVecSetZero((&Fe))

	// get local stiffness vector
	for j := 1; j <= 6; j++ {
		if j < 4 {
			posj = m.Beams[el].N[0]*3 + j
		} else {
			posj = m.Beams[el].N[1]*3 + j - 3
		}
		ue.Set(j-1, 0, ue.At(j-1, 0)+u.At(posj-1, 0))
		// femVecPutAdd((&ue), j, femVecGet((&u), posj), 0)
	}

	// get B and D
	//t := m.Beams[el].T         // e_t[el]
	D := m.get_D_matrix(el) //, t) // , (&D))
	B, _, _ := m.get_B_matrix(el)

	DB := mat.NewDense(5, 6, nil)
	DB.Mul(D, B)
	//
	// 	femMatMatMult((&D), (&B), (&DB))
	// get VECTOR

	Fe.Mul(DB, ue)
	//
	// 	femMatVecMult((&DB), (&ue), (&Fe))

	N1 = Fe.At(0, 0)
	N2 = Fe.At(1, 0)
	M1 = Fe.At(2, 0)
	M2 = Fe.At(3, 0)
	Q = Fe.At(4, 0)
	return
}

// print_result - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1036
func (m Model) print_result(u *mat.Dense) int {
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
				N1, N2, M1, M2, Q := m.get_int_forces(j, u) //, (&N1), (&N2), (&M1), (&M2), (&Q))
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
			sN1, sN2, sM1, sM2, sQ) //Q)
	}
	//_ = sQ
	return 0
}
