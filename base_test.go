package shell

import (
	"fmt"
	"os"
)

func Example() {

	model := Model{
		Points: [][2]float64{
			{10, 0},
			{10, 5},
			{10, 10},
		},
		Beams: []BeamProp{
			{N: [2]int{0, 1}, T: 0.2, E: 20e9, nu: 0.2},
			{N: [2]int{1, 2}, T: 0.2, E: 20e9, nu: 0.2},
		},
		Supports: [][3]bool{
			{true, true, false},
			{false, false, false},
			{false, true, false},
		},
		Ln: []LoadNode{
			{N: 1, Forces: [3]float64{0, 11.899e6, 0}},
		},
	}
	model.check_elem_data()
	K, F, u := model.get_matrix()
	model.get_loads_and_supports(K, F, u)
	if err := u.Solve(K, F); err != nil {
		fmt.Fprintf(os.Stdout, "%v", err)
		return
	}
	model.print_result(u)

	// Output:
	// #  X     Y        w            u           angle            N1          N2           M1          M2          Q
	// 10.000 0.000 0.000000e+00 0.000000e+00 -4.710195e-03 -1.746230e-10 2.358866e+06 1.308388e+04 1.308388e+04 5.233550e+03
	// 10.000 5.000 5.897164e-04 1.179433e-02 -1.420700e-18 -8.731149e-11 2.358866e+06 1.308388e+04 1.308388e+04 -2.773959e-11
	// 10.000 10.000 1.179433e-03 0.000000e+00 4.710195e-03 0.000000e+00 2.358866e+06 1.308388e+04 1.308388e+04 -5.233550e+03
}

func ExamplePlate() {
	// O.C. ZIENKIEWICZ, J. BAUER, K. MORGAN AND E. ONATE
	// A SIMPLE AND EFFICIENT ELEMENT FOR AXISYMMETRIC SHELLS

	for _, nodes := range []int{2 + 1, 4 + 1, 6 + 1} {
		points := make([][2]float64, nodes)
		R := 10.0
		for i := range points {
			points[i][0] = R * float64(i) / float64(nodes-1)
		}
		//fmt.Println("p ", points)

		beams := make([]BeamProp, nodes-1)
		for i := range beams {
			beams[i] = BeamProp{
				N:  [2]int{i, i + 1},
				T:  0.1,
				E:  1e7,
				nu: 0.3,
			}
		}
		//fmt.Println("b ", beams)

		sups := make([][3]bool, nodes)
		sups[0][0], sups[0][2] = true, true
		sups[nodes-1][0], sups[nodes-1][1], sups[nodes-1][2] = true, true, true
		//fmt.Println("s ", sups)

		loads := make([]LoadNode, nodes)
		for i := range loads {
			loads[i].N = i
			loads[i].Forces = [3]float64{0, -1.0 * R / float64(nodes-1), 0}
			if i == 0 {
				loads[i].Forces[1] /= 2.0
			}
		}
		//fmt.Println("l ", loads)

		model := Model{Points: points, Beams: beams, Supports: sups, Ln: loads}

		model.check_elem_data()
		K, F, u := model.get_matrix()
		model.get_loads_and_supports(K, F, u)
		if err := u.Solve(K, F); err != nil {
			fmt.Fprintf(os.Stdout, "%v", err)
			return
		}
		model.print_result(u)
	}
	// Output:
// #  X     Y        w            u           angle            N1          N2           M1          M2          Q
// 0.000 0.000 0.000000e+00 -1.750000e-06 0.000000e+00 -2.375000e-01 -7.625000e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 5.000 0.000 0.000000e+00 -1.706250e-06 -0.000000e+00 -3.250000e-01 -5.000000e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 10.000 0.000 0.000000e+00 0.000000e+00 0.000000e+00 -4.125000e-01 -2.375000e-01 0.000000e+00 0.000000e+00 0.000000e+00
// #  X     Y        w            u           angle            N1          N2           M1          M2          Q
// 0.000 0.000 0.000000e+00 -8.750000e-07 0.000000e+00 1.041667e-02 -1.010417e+00 0.000000e+00 0.000000e+00 0.000000e+00
// 2.500 0.000 0.000000e+00 -1.658854e-06 0.000000e+00 -7.708333e-02 -7.479167e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 5.000 0.000 0.000000e+00 -1.611458e-06 0.000000e+00 -2.520833e-01 -3.979167e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 7.500 0.000 0.000000e+00 -9.953125e-07 -0.000000e+00 -3.979167e-01 -2.520833e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 10.000 0.000 0.000000e+00 0.000000e+00 0.000000e+00 -4.562500e-01 -1.937500e-01 0.000000e+00 0.000000e+00 0.000000e+00
// #  X     Y        w            u           angle            N1          N2           M1          M2          Q
// 0.000 0.000 0.000000e+00 -5.833333e-07 0.000000e+00 1.533333e-01 -1.153333e+00 0.000000e+00 0.000000e+00 0.000000e+00
// 1.667 0.000 0.000000e+00 -1.415556e-06 0.000000e+00 6.583333e-02 -8.908333e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 3.333 0.000 0.000000e+00 -1.693611e-06 0.000000e+00 -1.091667e-01 -5.408333e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 5.000 0.000 0.000000e+00 -1.592500e-06 0.000000e+00 -2.550000e-01 -3.950000e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 6.667 0.000 0.000000e+00 -1.238611e-06 0.000000e+00 -3.570833e-01 -2.929167e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 8.333 0.000 0.000000e+00 -6.951389e-07 0.000000e+00 -4.358333e-01 -2.141667e-01 0.000000e+00 0.000000e+00 0.000000e+00
// 10.000 0.000 0.000000e+00 0.000000e+00 0.000000e+00 -4.708333e-01 -1.791667e-01 0.000000e+00 0.000000e+00 0.000000e+00
}
