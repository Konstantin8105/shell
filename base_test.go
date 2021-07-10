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
	// 10.000 0.000 0.000000e+00 0.000000e+00 -4.530268e-03 6.776264e-21 2.359665e+06 1.258408e+04 2.516815e+03 5.033631e+03
	// 10.000 5.000 2.831599e-14 1.132639e-02 1.244024e-15 6.776264e-21 2.359665e+06 1.258408e+04 2.516815e+03 -4.523372e-09
	// 10.000 10.000 5.663197e-14 0.000000e+00 4.530268e-03 6.776264e-21 2.359665e+06 1.258408e+04 2.516815e+03 -5.033631e+03
}
