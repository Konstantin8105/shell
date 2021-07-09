package shell

// main - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1744
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

	_ = model

	// 	var i int
	// 	n_m = 1

	// 	n_n = 3

	// 	n_e = 2

	// 	n_d = 3

	// n_f = 1

	// 	m_E1 = make([]float64, n_m)
	//
	// 	m_E2 = make([]float64, n_m)
	//
	// 	m_G = make([]float64, n_m)
	//
	// 	m_nu1 = make([]float64, n_m)
	//
	// 	m_nu2 = make([]float64, n_m)
	//
	// 	m_q = make([]float64, n_m)
	//
	// 	m_vp = make([]float64, n_m)
	//
	// 	m_t = make([]float64, n_m)

	// 	n_x = make([]float64, n_n)
	//
	// 	n_y = make([]float64, n_n)

	// 	e_n1 = make([]int, n_e)
	//
	// 	e_n2 = make([]int, n_e)
	//
	// 	e_mat = make([]int, n_e)
	//
	// 	e_t = make([]float64, n_e)

	// 	d_n = make([]int, n_d)
	//
	// 	d_dir = make([]int, n_d)
	//
	// 	d_val = make([]float64, n_d)

	// 	f_n = make([]int, n_f)
	//
	// 	f_dir = make([]int, n_f)
	//
	// 	f_val = make([]float64, n_f)

	// en_num = make([]int,n_n)

	//en_frm = make([]int,n_n)

	// 	m_E1[0], m_E2[0], m_G[0], m_nu1[0], m_nu2[0], m_q[0], m_vp[0], m_t[0] = 20e9, 0, 0, 0.2, 0, 25000, 1000, 0
	//
	// 	for i = 0; i < n_m; i++ {
	//
	// 		if m_E1[i] == m_E2[i] || m_E2[i] <= 0 {
	//
	// 			m_E2[i] = m_E1[i]
	// 			m_nu2[i] = m_nu1[i]
	// 			if m_G[i] <= 0 {
	// 				m_G[i] = m_E1[i] / (2 * (1 + m_nu1[i]))
	//
	// 			}
	// 		}
	// 	}

	// 	w_top, w_bot, w_val, w_min, w_max = 0, 0, 0, 0, 0

	model.check_elem_data()

	// 	fail_type = 1
	//
	// 	n_fail = 2
	//
	// 	fail_data = make([]float64, n_fail)
	//
	// 	fail_data[0] = 20e6
	// 	fail_data[1] = 1e6
	//
	// 	n_r_inp = 1
	//
	// 	rand_type = make([]int,n_r_inp)
	//
	// 	rand_pos = make([]int,n_r_inp)
	//
	// 	rand_indx = make([]int,n_r_inp)
	//
	// 	rand_type[0], rand_pos[0], rand_indx[0] = 4, 0, 0

	// n_r_opt = 0

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
	// read_input_data() //noarch.Stdin)
	//stat +=
	// model.alloc_solver_data()
	// stat += optim_replace_data(opt_data)
	// 	if write_only == 1 {
	// write_input_data() //noarch.Stdout)
	// 		return
	// 	}
	// 	if solution_only == 1 {
	//stat +=
	K, F, u := model.get_matrix()
	// stat += generate_water_load_x()
	//stat +=
	model.get_loads_and_supports(K, F, u)

	// 	fmt.Printf("K = %.2e\n", mat.Formatted(K, mat.Prefix("    "), mat.Squeeze()))
	// 	fmt.Println("FF", F)
	//stat =
	// n_n := len(model.Points)

	// err :
	_ = u.Solve(K, F)

	// 	fmt.Println(">>>>>>>>>>>>>>> ", err)
	//
	// 	fmt.Println("uu ", u)

	// 	femEqsCGwJ((&K), (&F), (&u), 1e-09, 6*3*n_n)
	// 	}
	// 	if n_r_inp > 0 && random_only == 1 {
	// 		if solution_only != 0 {
	// 			print_result(noarch.Stdout)
	// 		}
	// 		generate_rand_input_file(noarch.Stdout)
	// 		generate_rand_out_file(noarch.Stdout)
	// 	} else {
	// 	if solution_only == 1 {
	model.print_result(u) //noarch.Stdout)
	// 	}
	// 	}
	// 	if solution_only == 1 {
	// 	if fail_test() != 0 {
	// 		fmt.Fprintf(os.Stdout, string("# Structure FAILED\n"))
	// 	}
	// 	}
	// 	if price_only == 1 {
	// 		if solution_only == 1 {
	// 	fmt.Fprintf(os.Stdout, string("# Price is %f\n"), compute_price())
	// 		} else {
	// 			fmt.Fprintf(noarch.Stdout, string("%e\n"), compute_price())
	// 		}
	// 	}
	// return

	// Output:
	// #  X     Y        w            u           angle            N1          N2           M1          M2          Q
	// 10.000 0.000 0.000000e+00 0.000000e+00 -4.590364e-03 6.776264e-21 2.361648e+06 1.290723e+04 2.581446e+03 4.162893e+03
	// 10.000 5.000 2.833978e-14 1.133591e-02 5.623920e-05 6.776264e-21 2.361648e+06 1.259473e+04 2.518946e+03 -3.750000e+02
	// 10.000 10.000 5.667956e-14 0.000000e+00 4.477843e-03 6.776264e-21 2.361648e+06 1.228223e+04 2.456446e+03 -4.912893e+03
}
