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
		// 	e_n1[0], e_n2[0], e_mat[0], e_t[0] = 0, 1, 0, 0.2
		// 	e_n1[1], e_n2[1], e_mat[1], e_t[1] = 1, 2, 0, 0.2
		Beams: []BeamProp{
			{N: [2]int{0, 1}, Mat: 0, T: 0.2, E1: 20e9, E2: 20e9, G: 20e9 / (2 * (1 + 0.2)), nu1: 0.2, nu2: 0.2, q: 25000},
			{N: [2]int{1, 2}, Mat: 0, T: 0.2, E1: 20e9, E2: 20e9, G: 20e9 / (2 * (1 + 0.2)), nu1: 0.2, nu2: 0.2, q: 25000},
		},

		// 	d_n[0], d_dir[0], d_val[0] = 0, 0, 0
		// 	d_n[1], d_dir[1], d_val[1] = 0, 1, 0
		// 	d_n[2], d_dir[2], d_val[2] = 2, 1, 0
		Supports: [][3]bool{
			{true, true, false},
			{false, false, false},
			{false, true, false},
		},

		// f_n[0], f_dir[0], f_val[0] = 1, 1, 11.899e6
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
	// 10.000 0.000 0.000000e+00 0.000000e+00 -4.590364e-03 6.776264e-21 2.361648e+06 1.290723e+04 2.581446e+03 4.162893e+03
	// 10.000 5.000 2.833978e-14 1.133591e-02 5.623920e-05 6.776264e-21 2.361648e+06 1.259473e+04 2.518946e+03 -3.750000e+02
	// 10.000 10.000 5.667956e-14 0.000000e+00 4.477843e-03 6.776264e-21 2.361648e+06 1.228223e+04 2.456446e+03 -4.912893e+03
}
