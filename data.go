package main

func read_input_data() {

	var i int
	n_m = 1

	n_n = 3

	n_e = 2

	n_d = 3

	n_f = 1

	m_E1 = femDblAlloc(n_m)

	m_E2 = femDblAlloc(n_m)

	m_G = femDblAlloc(n_m)

	m_nu1 = femDblAlloc(n_m)

	m_nu2 = femDblAlloc(n_m)

	m_q = femDblAlloc(n_m)

	m_vp = femDblAlloc(n_m)

	m_t = femDblAlloc(n_m)

	n_x = femDblAlloc(n_n)

	n_y = femDblAlloc(n_n)

	e_n1 = femIntAlloc(n_e)

	e_n2 = femIntAlloc(n_e)

	e_mat = femIntAlloc(n_e)

	e_t = femDblAlloc(n_e)

	d_n = femIntAlloc(n_d)

	d_dir = femIntAlloc(n_d)

	d_val = femDblAlloc(n_d)

	f_n = femIntAlloc(n_f)

	f_dir = femIntAlloc(n_f)

	f_val = femDblAlloc(n_f)

	// en_num = femIntAlloc(n_n)

	//en_frm = femIntAlloc(n_n)

	m_E1[0], m_E2[0], m_G[0], m_nu1[0], m_nu2[0], m_q[0], m_vp[0], m_t[0] = 20e9, 0, 0, 0.2, 0, 25000, 1000, 0

	for i = 0; i < n_m; i++ {

		if m_E1[i] == m_E2[i] || m_E2[i] <= 0 {

			m_E2[i] = m_E1[i]
			m_nu2[i] = m_nu1[i]
			if m_G[i] <= 0 {
				m_G[i] = m_E1[i] / (2 * (1 + m_nu1[i]))

			}
		}
	}

	n_x[0], n_y[0] = 10, 0
	n_x[1], n_y[1] = 10, 5
	n_x[2], n_y[2] = 10, 10

	e_n1[0], e_n2[0], e_mat[0], e_t[0] = 0, 1, 0, 0.2
	e_n1[1], e_n2[1], e_mat[1], e_t[1] = 1, 2, 0, 0.2

	d_n[0], d_dir[0], d_val[0] = 0, 0, 0
	d_n[1], d_dir[1], d_val[1] = 0, 1, 0
	d_n[2], d_dir[2], d_val[2] = 2, 1, 0

	f_n[0], f_dir[0], f_val[0] = 1, 1, 11.899e6

	// 	w_top, w_bot, w_val, w_min, w_max = 0, 0, 0, 0, 0

	check_elem_data()

	fail_type = 1

	n_fail = 2

	fail_data = femDblAlloc(n_fail)

	fail_data[0] = 20e6
	fail_data[1] = 1e6

	n_r_inp = 1

	rand_type = femIntAlloc(n_r_inp)

	rand_pos = femIntAlloc(n_r_inp)

	rand_indx = femIntAlloc(n_r_inp)

	rand_type[0], rand_pos[0], rand_indx[0] = 4, 0, 0

	// n_r_opt = 0

}
