package shell

import (
	"fmt"
	"os"
)

// main - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1744
func Example() {
	// 	argc := int(len(os.Args))
	// 	argv := []string{}
	// 	for _, argvSingle := range os.Args {
	// 		argv = append(argv, string(argvSingle))
	// 	}
	// 	defer noarch.AtexitRun()
	// main() routine for standalone program only.
	// var stat int
	// os.Stdout = noarch.Stdout
	// 	cmd_param(argc, argv)
	read_input_data() //noarch.Stdin)
	//stat +=
	alloc_solver_data()
	// stat += optim_replace_data(opt_data)
	// 	if write_only == 1 {
	write_input_data() //noarch.Stdout)
	// 		return
	// 	}
	// 	if solution_only == 1 {
	//stat +=
	get_matrix()
	// stat += generate_water_load_x()
	//stat +=
	get_loads_and_supports()
	//stat =
	femEqsCGwJ((&K), (&F), (&u), 1e-09, 6*3*n_n)
	// 	}
	// 	if n_r_inp > 0 && random_only == 1 {
	// 		if solution_only != 0 {
	// 			print_result(noarch.Stdout)
	// 		}
	// 		generate_rand_input_file(noarch.Stdout)
	// 		generate_rand_out_file(noarch.Stdout)
	// 	} else {
	// 	if solution_only == 1 {
	print_result() //noarch.Stdout)
	// 	}
	// 	}
	// 	if solution_only == 1 {
	if fail_test() != 0 {
		fmt.Fprintf(os.Stdout, string("# Structure FAILED\n"))
	}
	// 	}
	// 	if price_only == 1 {
	// 		if solution_only == 1 {
	fmt.Fprintf(os.Stdout, string("# Price is %f\n"), compute_price())
	// 		} else {
	// 			fmt.Fprintf(noarch.Stdout, string("%e\n"), compute_price())
	// 		}
	// 	}
	// return

	// Output:
	// 1 3 2 3 1
	//  2.000000e+10 2.000000e+10 8.333333e+09 2.000000e-01 2.000000e-01 2.500000e+04 1.000000e+03 0.000000e+00
	// 1.000000e+01 0.000000e+00
	// 1.000000e+01 5.000000e+00
	// 1.000000e+01 1.000000e+01
	// 0 1 0 2.000000e-01
	// 1 2 0 2.000000e-01
	// 0 0 0.000000e+00
	// 0 1 0.000000e+00
	// 2 1 0.000000e+00
	// 1 1 1.189900e+07
	// 1
	// 2
	//  2.000000e+07 1.000000e+06
	// [ ]   linear step 1
	// ro = 1011.833528
	// alpha = 1.001683
	// [i] Convergence test 20913196.986873 < 0.102458 (step 1 from 54)
	// [ ]   linear step 2
	// ro = 1007.616880
	// beta = 0.995833
	// alpha = 132.973046
	// [i] Convergence test 20912.455361 < 13.858395 (step 2 from 54)
	// [ ]   linear step 3
	// ro = 0.000840
	// beta = 0.000001
	// alpha = 0.555587
	// [i] Convergence test 8536.877007 < 13.858395 (step 3 from 54)
	// [ ]   linear step 4
	// ro = 0.000093
	// beta = 0.111095
	// alpha = 7031.747778
	// [ ]  linear solution done in 4 iterations!
	// #  X     Y        w            u           angle            N1          N2           M1          M2          Q
	// 10.000 0.000 0.000000e+00 0.000000e+00 -4.590364e-03 -4.610823e-02 2.361648e+06 1.290723e+04 2.581446e+03 4.162893e+03
	// 10.000 5.000 5.535822e-11 1.133591e-02 5.623920e-05 -2.307347e-02 2.361648e+06 1.259473e+04 2.518946e+03 -4.912893e+03
	// 10.000 10.000 5.543302e-11 0.000000e+00 4.477843e-03 -3.871466e-05 2.361648e+06 1.228223e+04 2.456446e+03 -4.912893e+03
	// [0] fc = [2.000000e+07 1.000000e+06], ft = 1.000000e+06
	// [0] s1: 1.188569e+07, s2: 3.872167e+05
	// sm: 4.090967e+06 I1: 1.227290e+07, J2: 4.560571e+13, J3: -2.886969e+13
	// [0] alpha: 7.365453e+00, beta: 7.293421e+00, lambda: 3.338609e+01 cos3f: -2.435372e-07
	// c1: 3.855094e+01, c2: 9.972800e-01 => fc: 1.658849e+01
	// Element 0 FAILED : 1658.85 procent
	// # Structure FAILED
	// # Price is 125663.706144

}
