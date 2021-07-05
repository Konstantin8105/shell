package main

import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"github.com/Konstantin8105/c4go/noarch"
)

var msgout *noarch.File

func femIntAlloc(length int32) (c4goDefaultReturn []int32) {

	var field []int32
	var i int32
	if length < 1 {
		return nil
	}
	if len((func() []int32 {
		field = (*[1000000]int32)(unsafe.Pointer(uintptr(func() int64 {
			c4go_temp_name := make([]byte, uint32(length)*uint32(1))
			return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
		}())))[:]
		return field
	}())) == 0 {
		return nil
	} else {
		for i = 0; i < length; i++ {
			field[i] = 0
		}
		return field
	}
	return
}

func femIntFree(field []int32) int32 {
	_ = field

	field = nil
	return 0
}

func femDblAlloc(length int32) (c4goDefaultReturn []float64) {

	var field []float64
	var i int32
	if length < 1 {
		return nil
	}
	if len((func() []float64 {
		field = (*[1000000]float64)(unsafe.Pointer(uintptr(func() int64 {
			c4go_temp_name := make([]byte, uint32(length)*uint32(1))
			return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
		}())))[:]
		return field
	}())) == 0 {
		return nil
	} else {
		for i = 0; i < length; i++ {
			field[i] = 0
		}
		return field
	}
	return
}

func femDblFree(field []float64) int32 {
	_ = field

	field = nil
	return 0
}

type _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_47 struct {
	type_   int32
	rows    int32
	cols    int32
	len_    int32
	pos     []int32
	data    []float64
	frompos []int32
	defpos  []int32
}

type tMatrix = _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_47

type _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_60 struct {
	type_ int32
	rows  int32
	len_  int32
	pos   []int32
	data  []float64
}

type tVector = _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_60

var n_m int32

var n_n int32

var n_e int32

var n_d int32

var n_f int32

var n_r_inp int32

var n_r_opt int32

var m_E1 []float64

var m_E2 []float64

var m_G []float64

var m_nu1 []float64

var m_nu2 []float64

var m_q []float64

var m_vp []float64

var m_t []float64

var n_x []float64

var n_y []float64

var e_n1 []int32

var e_n2 []int32

var e_mat []int32

var e_t []float64

var d_n []int32

var d_dir []int32

var d_val []float64

var f_n []int32

var f_dir []int32

var f_val []float64

var w_top float64

var w_bot float64

var w_val float64

var w_min int32 = -1

var w_max int32 = -1

var rand_type []int32

var rand_pos []int32

var rand_indx []int32

var opt_type []int32

var opt_pos []int32

var opt_indx []int32

var opt_data []float64

var fail_type int32

var n_fail int32

var fail_data []float64

var K tMatrix

var u tVector

var F tVector

var Ke tMatrix

var D tMatrix

var B tMatrix

var Bt tMatrix

var BtD tMatrix

var DB tMatrix

var Fe tVector

var ue tVector

var n_en int32

var en_num []int32

var en_frm []int32

var en_pos []int32

var solution_only int32 = 1

var random_only int32 = 1

var price_only int32 = 1

var write_only int32

func free_input_data() {
	if len(m_E1) != 0 {

		femDblFree(m_E1)
	}
	if len(m_E2) != 0 {
		femDblFree(m_E2)
	}
	if len(m_G) != 0 {
		femDblFree(m_G)
	}
	if len(m_nu1) != 0 {
		femDblFree(m_nu1)
	}
	if len(m_nu2) != 0 {
		femDblFree(m_nu2)
	}
	if len(m_q) != 0 {
		femDblFree(m_q)
	}
	if len(m_vp) != 0 {
		femDblFree(m_vp)
	}
	if len(m_t) != 0 {
		femDblFree(m_t)
	}
	if len(n_x) != 0 {
		femDblFree(n_x)
	}
	if len(n_y) != 0 {
		femDblFree(n_y)
	}
	if len(e_n1) != 0 {
		femIntFree(e_n1)
	}
	if len(e_n2) != 0 {
		femIntFree(e_n2)
	}
	if len(e_mat) != 0 {
		femIntFree(e_mat)
	}
	if len(e_t) != 0 {
		femDblFree(e_t)
	}
	if len(d_n) != 0 {
		femIntFree(d_n)
	}
	if len(d_dir) != 0 {
		femIntFree(d_dir)
	}
	if len(d_val) != 0 {
		femDblFree(d_val)
	}
	if n_f > 0 {
		if len(f_n) != 0 {
			femIntFree(f_n)
		}
		if len(f_dir) != 0 {
			femIntFree(f_dir)
		}
		if len(f_val) != 0 {
			femDblFree(f_val)
		}
	}
	if n_r_inp > 0 {
		if len(rand_type) != 0 {
			femIntFree(rand_type)
		}
		if len(rand_pos) != 0 {
			femIntFree(rand_pos)
		}
		if len(rand_indx) != 0 {
			femIntFree(rand_indx)
		}
	}
	if n_r_opt > 0 {
		if len(opt_type) != 0 {
			femIntFree(opt_type)
		}
		if len(opt_pos) != 0 {
			femIntFree(opt_pos)
		}
		if len(opt_indx) != 0 {
			femIntFree(opt_indx)
		}
		if len(opt_data) != 0 {
			femDblFree(opt_data)
		}
	}
	if n_en > 0 {
		if len(en_num) != 0 {
			femIntFree(en_num)
		}
		if len(en_frm) != 0 {
			femIntFree(en_frm)
		}
		if len(en_pos) != 0 {
			femIntFree(en_pos)
		}
	}
	if n_fail > 0 {
		if len(fail_data) != 0 {
			femDblFree(fail_data)
		}
	}
	n_m = 0
	n_n = 0
	n_e = 0
	n_d = 0
	n_f = 0
	n_r_inp = 0
	n_r_opt = 0
	n_en = 0
	fail_type = 0
	n_fail = 0
}

func check_elem_data() {

	var i int32
	var tmp int32
	for i = 0; i < n_e; i++ {
		if n_y[e_n1[i]] > n_y[e_n2[i]] {
			tmp = e_n1[i]
			e_n1[i] = e_n2[i]
			e_n2[i] = tmp
		}
	}
}

func get_enode_fields() int32 {

	var i int32
	var j int32
	if len(en_num) == 0 {
		return -3
	}
	if len(en_frm) == 0 {
		return -3
	}
	for i = 0; i < n_e; i++ {
		en_num[e_n1[i]]++
		en_num[e_n2[i]]++
	}
	n_en = 0
	for i = 0; i < n_n; i++ {
		en_frm[i] = n_en
		n_en += en_num[i]
	}
	if len((func() []int32 {
		en_pos = femIntAlloc(n_en)
		return en_pos
	}())) == 0 {
		goto memFree
	}
	for i = 0; i < n_en; i++ {
		en_pos[i] = -1
	}
	for i = 0; i < n_e; i++ {
		for j = 0; j < en_num[e_n1[i]]; j++ {
			if en_pos[en_frm[e_n1[i]]+j] == -1 {
				en_pos[en_frm[e_n1[i]]+j] = i
				break
			}
		}
		for j = 0; j < en_num[e_n2[i]]; j++ {
			if en_pos[en_frm[e_n2[i]]+j] == -1 {
				en_pos[en_frm[e_n2[i]]+j] = i
				break
			}
		}
	}
	return 0
memFree:
	;
	return -4
}

func read_input_data(fw *noarch.File) int32 {

	var i int32
	if noarch.Fscanf(fw, []byte("%li\x00"), c4goUnsafeConvert_int32(&n_m)) <= 0 {
		goto memFree
	}
	if n_m < 1 {
		noarch.Fprintf(msgout, []byte("Invalid number of materials!\n\x00"))
		goto memFree
	}
	if noarch.Fscanf(fw, []byte("%li\x00"), c4goUnsafeConvert_int32(&n_n)) <= 0 {
		goto memFree
	}
	if n_n < 2 {
		noarch.Fprintf(msgout, []byte("Invalid number of nodes!\n\x00"))
		goto memFree
	}
	if noarch.Fscanf(fw, []byte("%li\x00"), c4goUnsafeConvert_int32(&n_e)) <= 0 {
		goto memFree
	}
	if n_e < 2 {
		noarch.Fprintf(msgout, []byte("Invalid number of elements!\n\x00"))
		goto memFree
	}
	if noarch.Fscanf(fw, []byte("%li\x00"), c4goUnsafeConvert_int32(&n_d)) <= 0 {
		goto memFree
	}
	if n_d < 3 {
		noarch.Fprintf(msgout, []byte("Invalid number of supports!\n\x00"))
		goto memFree
	}
	if noarch.Fscanf(fw, []byte("%li\x00"), c4goUnsafeConvert_int32(&n_f)) < 0 {
		goto memFree
	}
	if n_f < 0 {
		noarch.Fprintf(msgout, []byte("Invalid number of forces!\n\x00"))
		goto memFree
	}
	if len((func() []float64 {
		m_E1 = femDblAlloc(n_m)
		return m_E1
	}())) == 0 {

		goto memFree
	}
	if len((func() []float64 {
		m_E2 = femDblAlloc(n_m)
		return m_E2
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		m_G = femDblAlloc(n_m)
		return m_G
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		m_nu1 = femDblAlloc(n_m)
		return m_nu1
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		m_nu2 = femDblAlloc(n_m)
		return m_nu2
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		m_q = femDblAlloc(n_m)
		return m_q
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		m_vp = femDblAlloc(n_m)
		return m_vp
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		m_t = femDblAlloc(n_m)
		return m_t
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		n_x = femDblAlloc(n_n)
		return n_x
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		n_y = femDblAlloc(n_n)
		return n_y
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		e_n1 = femIntAlloc(n_e)
		return e_n1
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		e_n2 = femIntAlloc(n_e)
		return e_n2
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		e_mat = femIntAlloc(n_e)
		return e_mat
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		e_t = femDblAlloc(n_e)
		return e_t
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		d_n = femIntAlloc(n_d)
		return d_n
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		d_dir = femIntAlloc(n_d)
		return d_dir
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		d_val = femDblAlloc(n_d)
		return d_val
	}())) == 0 {
		goto memFree
	}
	if n_f > 0 {
		if len((func() []int32 {
			f_n = femIntAlloc(n_f)
			return f_n
		}())) == 0 {
			goto memFree
		}
		if len((func() []int32 {
			f_dir = femIntAlloc(n_f)
			return f_dir
		}())) == 0 {
			goto memFree
		}
		if len((func() []float64 {
			f_val = femDblAlloc(n_f)
			return f_val
		}())) == 0 {
			goto memFree
		}
	}
	if len((func() []int32 {
		en_num = femIntAlloc(n_n)
		return en_num
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		en_frm = femIntAlloc(n_n)
		return en_frm
	}())) == 0 {
		goto memFree
	}
	{

		for i = 0; i < n_m; i++ {
			if noarch.Fscanf(fw, []byte(" %f %f %f %f %f %f %f %f\x00\x00\x00\x00\x00\x00\x00\x00\x00"), m_E1[i:], m_E2[i:], m_G[i:], m_nu1[i:], m_nu2[i:], m_q[i:], m_vp[i:], m_t[i:]) <= 0 {
				goto memFree
			}
			if m_E1[i] == m_E2[i] || m_E2[i] <= 0 {

				m_E2[i] = m_E1[i]
				m_nu2[i] = m_nu1[i]
				if m_G[i] <= 0 {
					m_G[i] = m_E1[i] / (2 * (1 + m_nu1[i]))
				}
			} else {
				if m_E1[i] <= 0 || m_E2[i] <= 0 || m_G[i] <= 0 || m_nu1[i] <= 0 || m_nu2[i] <= 0 {
					noarch.Fprintf(msgout, []byte("Invalid or incomplete data for material %li\n\x00"), i)
					goto memFree
				}
			}
		}
	}
	{

		for i = 0; i < n_n; i++ {
			if noarch.Fscanf(fw, []byte("%f %f\x00\x00\x00"), n_x[i:], n_y[i:]) <= 0 {
				goto memFree
			}
		}
	}
	{

		for i = 0; i < n_e; i++ {
			if noarch.Fscanf(fw, []byte("%li %li %li %f\x00\x00"), e_n1[i:], e_n2[i:], e_mat[i:], e_t[i:]) <= 0 {
				goto memFree
			}
			if e_n1[i] < 0 || e_n1[i] >= n_n {
				noarch.Fprintf(msgout, []byte("Invalid n1 in element %li\n\x00"), i)
				goto memFree
			}
			if e_n2[i] < 0 || e_n2[i] >= n_n {
				noarch.Fprintf(msgout, []byte("Invalid n2 in element %li\n\x00"), i)
				goto memFree
			}
			if e_n1[i] == e_n2[i] {
				noarch.Fprintf(msgout, []byte("Invalid nodes in element %li\n\x00"), i)
				goto memFree
			}
			if e_mat[i] < 0 || e_mat[i] >= n_m {
				noarch.Fprintf(msgout, []byte("Invalid material in element %li\n\x00"), i)
				goto memFree
			}
			if e_t[i] <= 0 {
				noarch.Fprintf(msgout, []byte("Invalid width in element %li\n\x00"), i)
				goto memFree
			}
		}
	}
	{

		for i = 0; i < n_d; i++ {
			if noarch.Fscanf(fw, []byte("%li %li %f\x00\x00"), d_n[i:], d_dir[i:], d_val[i:]) <= 0 {
				goto memFree
			}
			if d_n[i] < 0 || d_n[i] >= n_n {
				noarch.Fprintf(msgout, []byte("Invalid node in support %li\n\x00"), i)
				goto memFree
			}
			if d_dir[i] < 0 || d_dir[i] >= 6 {
				noarch.Fprintf(msgout, []byte("Invalid direction in support %li\n\x00"), i)
				goto memFree
			}
			if d_dir[i] > 2 && d_val[i] < 0 {
				noarch.Fprintf(msgout, []byte("Invalid stiffness in support %li\n\x00"), i)
				goto memFree
			}
		}
	}
	{

		for i = 0; i < n_f; i++ {
			if noarch.Fscanf(fw, []byte("%li %li %f\x00\x00"), f_n[i:], f_dir[i:], f_val[i:]) <= 0 {
				goto memFree
			}
			if f_n[i] < 0 || f_n[i] >= n_n {
				noarch.Fprintf(msgout, []byte("Invalid node for force %li\n\x00"), i)
				goto memFree
			}
			if f_dir[i] < 0 || f_dir[i] >= 3 {
				noarch.Fprintf(msgout, []byte("Invalid direction for force %li\n\x00"), i)
				goto memFree
			}
		}
	}
	if noarch.Fscanf(fw, []byte("%f %f %f %li %li\x00\x00\x00\x00"), c4goUnsafeConvert_float64(&w_top), c4goUnsafeConvert_float64(&w_bot), c4goUnsafeConvert_float64(&w_val), c4goUnsafeConvert_int32(&w_min), c4goUnsafeConvert_int32(&w_max)) <= 0 {

		goto memFree
	}

	check_elem_data()
	if get_enode_fields() != 0 {
		goto memFree
	}
	if noarch.Fscanf(fw, []byte("%li\x00"), c4goUnsafeConvert_int32(&fail_type)) <= 0 {

		fail_type = 0
		n_fail = 0
	} else {
		if fail_type > 0 {
			if noarch.Fscanf(fw, []byte("%li\x00"), c4goUnsafeConvert_int32(&n_fail)) <= 0 {
				fail_type = 0
			} else {
				if len((func() []float64 {
					fail_data = femDblAlloc(n_fail)
					return fail_data
				}())) == 0 {
					noarch.Fprintf(msgout, []byte("Cannot allocate memory for failure data!\n\x00"))
					goto memFree
				}
				for i = 0; i < n_fail; i++ {
					if noarch.Fscanf(fw, []byte("%f\x00\x00"), fail_data[i:]) <= 0 {
						noarch.Fprintf(msgout, []byte("Invalid failure data!\n\x00"))
						goto memFree
					}
				}
			}
		}
	}
	if noarch.Fscanf(fw, []byte("%li\x00"), c4goUnsafeConvert_int32(&n_r_inp)) <= 0 {

		n_r_inp = 0

		return 0
	}
	if n_r_inp < 1 {
		return 0
	}
	if len((func() []int32 {
		rand_type = femIntAlloc(n_r_inp)
		return rand_type
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		rand_pos = femIntAlloc(n_r_inp)
		return rand_pos
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		rand_indx = femIntAlloc(n_r_inp)
		return rand_indx
	}())) == 0 {
		goto memFree
	}
	for i = 0; i < n_r_inp; i++ {
		if noarch.Fscanf(fw, []byte("%li %li %li\x00"), rand_type[i:], rand_pos[i:], rand_indx[i:]) <= 0 {
			goto memFree
		}
	}
	if noarch.Fscanf(fw, []byte("%li\x00"), c4goUnsafeConvert_int32(&n_r_opt)) <= 0 {

		n_r_opt = 0

		return 0
	}
	if n_r_opt < 1 {
		noarch.Fprintf(msgout, []byte("Invalid number of optim. inputs!\n\x00"))
		return 0
	}
	if len((func() []int32 {
		opt_type = femIntAlloc(n_r_opt)
		return opt_type
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		opt_pos = femIntAlloc(n_r_opt)
		return opt_pos
	}())) == 0 {
		goto memFree
	}
	if len((func() []int32 {
		opt_indx = femIntAlloc(n_r_opt)
		return opt_indx
	}())) == 0 {
		goto memFree
	}
	if len((func() []float64 {
		opt_data = femDblAlloc(n_r_opt)
		return opt_data
	}())) == 0 {
		goto memFree
	}
	for i = 0; i < n_r_opt; i++ {
		if noarch.Fscanf(fw, []byte("%li %li %li\x00"), opt_type[i:], opt_pos[i:], opt_indx[i:]) <= 0 {
			goto memFree
		}
	}
	for i = 0; i < n_r_opt; i++ {
		if noarch.Fscanf(fw, []byte("%f\x00\x00"), opt_data[i:]) <= 0 {
			femDblFree(opt_data)
			femIntFree(opt_indx)
			femIntFree(opt_pos)
			femIntFree(opt_type)
			n_r_opt = 0
		}
	}
	return 0
memFree:
	;
	free_input_data()
	noarch.Fprintf(msgout, []byte("Error when reading input!\n\x00"))
	return -2
}

func write_input_data(fw *noarch.File) int32 {

	var i int32

	noarch.Fprintf(fw, []byte("%li %li %li %li %li\n\x00"), n_m, n_n, n_e, n_d, n_f)
	{

		for i = 0; i < n_m; i++ {
			noarch.Fprintf(fw, []byte(" %e %e %e %e %e %e %e %e\n\x00"), m_E1[i], m_E2[i], m_G[i], m_nu1[i], m_nu2[i], m_q[i], m_vp[i], m_t[i])
		}
	}
	{

		for i = 0; i < n_n; i++ {
			noarch.Fprintf(fw, []byte("%e %e\n\x00"), n_x[i], n_y[i])
		}
	}
	{

		for i = 0; i < n_e; i++ {
			noarch.Fprintf(fw, []byte("%li %li %li %e\n\x00"), e_n1[i], e_n2[i], e_mat[i], e_t[i])
		}
	}
	{

		for i = 0; i < n_d; i++ {
			noarch.Fprintf(fw, []byte("%li %li %e\n\x00"), d_n[i], d_dir[i], d_val[i])
		}
	}
	{

		for i = 0; i < n_f; i++ {
			noarch.Fprintf(fw, []byte("%li %li %e\n\x00"), f_n[i], f_dir[i], f_val[i])
		}
	}

	noarch.Fprintf(fw, []byte("%e %e %e %li %li\n\x00"), w_top, w_bot, w_val, w_min, w_max)

	noarch.Fprintf(fw, []byte("%li\n\x00"), fail_type)
	if fail_type > 0 {
		noarch.Fprintf(fw, []byte("%li\n\x00"), n_fail)
		for i = 0; i < n_fail; i++ {
			noarch.Fprintf(fw, []byte("%e \x00"), fail_data[i])
		}
		noarch.Fprintf(fw, []byte("\n\x00"))
	}
	return 0
}

func free_solver_data() {

	femMatFree((*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:])
	femMatFree((*[1000000]tMatrix)(unsafe.Pointer(&D))[:])
	femMatFree((*[1000000]tMatrix)(unsafe.Pointer(&B))[:])
	femMatFree((*[1000000]tMatrix)(unsafe.Pointer(&Bt))[:])
	femMatFree((*[1000000]tMatrix)(unsafe.Pointer(&BtD))[:])
	femMatFree((*[1000000]tMatrix)(unsafe.Pointer(&DB))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&ue))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&Fe))[:])
	femMatFree((*[1000000]tMatrix)(unsafe.Pointer(&K))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&u))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&F))[:])
}

func alloc_solver_data() int32 {

	var i int32
	var j int32
	var n_field []int32
	var alloc_field []int32
	femMatNull((*[1000000]tMatrix)(unsafe.Pointer(&K))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&u))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&F))[:])
	femMatNull((*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:])
	femMatNull((*[1000000]tMatrix)(unsafe.Pointer(&D))[:])
	femMatNull((*[1000000]tMatrix)(unsafe.Pointer(&B))[:])
	femMatNull((*[1000000]tMatrix)(unsafe.Pointer(&Bt))[:])
	femMatNull((*[1000000]tMatrix)(unsafe.Pointer(&BtD))[:])
	femMatNull((*[1000000]tMatrix)(unsafe.Pointer(&DB))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&Fe))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&ue))[:])
	if femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:], 0, 6, 6, 0, nil) != 0 {
		goto memFree
	}
	if femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&D))[:], 0, 5, 5, 0, nil) != 0 {
		goto memFree
	}
	if femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&B))[:], 0, 5, 6, 0, nil) != 0 {
		goto memFree
	}
	if femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&Bt))[:], 0, 6, 5, 0, nil) != 0 {
		goto memFree
	}
	if femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&BtD))[:], 0, 6, 5, 0, nil) != 0 {
		goto memFree
	}
	if femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&DB))[:], 0, 5, 6, 0, nil) != 0 {
		goto memFree
	}
	if femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 0, 5, 5) != 0 {
		goto memFree
	}
	if femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&ue))[:], 0, 6, 6) != 0 {
		goto memFree
	}
	if len((func() []int32 {
		n_field = femIntAlloc(n_n)
		return n_field
	}())) == 0 {

		goto memFree
	}
	if len((func() []int32 {
		alloc_field = femIntAlloc(n_n * 3)
		return alloc_field
	}())) == 0 {
		goto memFree
	}
	for i = 0; i < n_n; i++ {
		for j = 0; j < n_e; j++ {
			if e_n1[j] == i {
				n_field[i]++
			}
			if e_n2[j] == i {
				n_field[i]++
			}
		}
	}
	for i = 0; i < n_n; i++ {
		for j = 0; j < 3; j++ {

			alloc_field[3*i+j] = 3 * 6 * n_field[i]
		}
	}
	if femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], 1, n_n*3, n_n*3, 0, alloc_field) != 0 {

		goto memFree
	}
	if femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&F))[:], 0, n_n*3, n_n*3) != 0 {
		goto memFree
	}
	if femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&u))[:], 0, n_n*3, n_n*3) != 0 {
		goto memFree
	}
	femIntFree(alloc_field)
	femIntFree(n_field)
	return 0
memFree:
	;
	if len(alloc_field) != 0 {
		femIntFree(alloc_field)
	}
	if len(n_field) != 0 {
		femIntFree(n_field)
	}
	free_solver_data()
	noarch.Fprintf(msgout, []byte("Out of memory!\x00"))
	return -4
}

func get_D_matrix(i int32, t float64, D []tMatrix) {

	var E1 float64
	var E2 float64
	var nu1 float64
	var nu2 float64
	var G float64
	var mult float64
	E1 = m_E1[e_mat[i]]
	E2 = m_E2[e_mat[i]]
	G = m_G[e_mat[i]]
	nu1 = m_nu1[e_mat[i]]
	nu2 = m_nu2[e_mat[i]]
	mult = t / (1 - nu1*nu2)
	femMatPutAdd(D, 1, 1, E1*mult, 0)
	femMatPutAdd(D, 1, 2, nu2*mult, 0)
	femMatPutAdd(D, 2, 1, nu2*mult, 0)
	femMatPutAdd(D, 2, 2, E2*mult, 0)
	femMatPutAdd(D, 3, 3, E1*t*t/12*mult, 0)
	femMatPutAdd(D, 4, 4, E2*t*t/12*mult, 0)
	femMatPutAdd(D, 3, 4, nu2*(E1*t*t)/12*mult, 0)
	femMatPutAdd(D, 4, 3, nu2*(E1*t*t)/12*mult, 0)
	femMatPutAdd(D, 5, 5, 5/6*G/t, 0)
}

func get_B_matrix(i int32, B []tMatrix, Lc []float64, Rc []float64) {

	var L float64
	var C float64
	var S float64
	var R float64
	var dx float64
	var dy float64
	dx = n_x[e_n2[i]] - n_x[e_n1[i]]
	dy = n_y[e_n2[i]] - n_y[e_n1[i]]
	L = math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	R = 0.5 * (n_x[e_n1[i]] + n_x[e_n2[i]])
	S = -1 * dx / L
	C = -1 * dy / L

	femMatPutAdd(B, 1, 1, -1*C/L, 0)
	femMatPutAdd(B, 1, 2, -1*S/L, 0)
	femMatPutAdd(B, 1, 4, 1*C/L, 0)
	femMatPutAdd(B, 1, 5, 1*S/L, 0)
	femMatPutAdd(B, 2, 2, 1/(2*R), 0)
	femMatPutAdd(B, 2, 5, 1/(2*R), 0)
	femMatPutAdd(B, 3, 3, -1/L, 0)
	femMatPutAdd(B, 3, 6, 1/L, 0)
	femMatPutAdd(B, 4, 3, S/(2*R), 0)
	femMatPutAdd(B, 4, 6, S/(2*R), 0)
	femMatPutAdd(B, 5, 1, -1*S/L, 0)
	femMatPutAdd(B, 5, 2, 1*C/L, 0)
	femMatPutAdd(B, 5, 3, 1/2, 0)
	femMatPutAdd(B, 5, 4, 1*S/L, 0)
	femMatPutAdd(B, 5, 5, -1*C/L, 0)
	femMatPutAdd(B, 5, 6, 1/2, 0)
	Lc[0] = L
	Rc[0] = R
}

func get_matrix() int32 {

	var t float64
	var L float64
	var R float64
	var F2 float64
	var q float64
	var i int32
	var j int32
	var k int32
	var posj int32
	var posk int32
	femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&K))[:])
	femVecSetZero((*[1000000]tVector)(unsafe.Pointer(&u))[:])
	femVecSetZero((*[1000000]tVector)(unsafe.Pointer(&F))[:])
	for i = 0; i < n_e; i++ {
		if (func() float64 {
			t = m_t[e_mat[i]]
			return t
		}()) <= 0 {

			t = e_t[i]
		}
		t = e_t[i]
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:])
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&B))[:])
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&Bt))[:])
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&BtD))[:])
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&D))[:])

		get_D_matrix(i, t, (*[1000000]tMatrix)(unsafe.Pointer(&D))[:])

		get_B_matrix(i, (*[1000000]tMatrix)(unsafe.Pointer(&B))[:], c4goUnsafeConvert_float64(&L), c4goUnsafeConvert_float64(&R))

		femMatTran((*[1000000]tMatrix)(unsafe.Pointer(&B))[:], (*[1000000]tMatrix)(unsafe.Pointer(&Bt))[:])

		femMatMatMult((*[1000000]tMatrix)(unsafe.Pointer(&Bt))[:], (*[1000000]tMatrix)(unsafe.Pointer(&D))[:], (*[1000000]tMatrix)(unsafe.Pointer(&BtD))[:])

		femMatMatMult((*[1000000]tMatrix)(unsafe.Pointer(&BtD))[:], (*[1000000]tMatrix)(unsafe.Pointer(&B))[:], (*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:])

		femValMatMultSelf(R*L, (*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:])
		{

			for j = 1; j <= 6; j++ {
				if j < 4 {
					posj = e_n1[i]*3 + j
				} else {
					posj = e_n2[i]*3 + j - 3
				}
				for k = 1; k <= 6; k++ {
					if k < 4 {
						posk = e_n1[i]*3 + k
					} else {
						posk = e_n2[i]*3 + k - 3
					}
					femMatPutAdd((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], posj, posk, femMatGet((*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:], j, k), 1)
				}
			}
		}
		if math.Abs((func() float64 {
			q = m_q[e_mat[i]]
			return q
		}())) > 1e-07 {

			F2 = -0.5 * q * t * L
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], 3*e_n1[i], F2, 1)
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], 3*e_n2[i], F2, 1)
		}
	}
	return 0
}

func generate_water_load_x() int32 {

	var i int32
	var y1 float64
	var y2 float64
	var dx float64
	var L float64
	var val1 float64
	var val2 float64
	var from int32
	var to int32
	var down int32 = 1

	var use_1 int32 = 1

	var use_2 int32 = 1
	var pos1 int32
	var pos2 int32

	var y_max float64
	var y_min float64

	var a float64
	var b float64
	if math.Abs(w_val) > 100*1e-07 {
		if w_max-w_min == 0 {

			from = 0
			to = n_e
		} else {
			if w_min < 0 || w_min >= n_e {
				from = 0
			} else {
				from = w_min
			}
			if w_max < 0 || w_max > n_e {
				to = n_e
			} else {
				to = w_max
			}
		}

		y_min = n_y[e_n1[from]]
		y_max = y_min
		for i = from; i < to; i++ {
			if y_min > n_y[e_n1[i]] {
				y_min = n_y[e_n1[i]]
			}
			if y_min > n_y[e_n2[i]] {
				y_min = n_y[e_n2[i]]
			}
			if y_max < n_y[e_n1[i]] {
				y_max = n_y[e_n1[i]]
			}
			if y_max < n_y[e_n2[i]] {
				y_max = n_y[e_n2[i]]
			}
		}
		if w_top < y_max {

			y_max = w_top
		}
		if w_bot > y_min {
			y_min = w_bot
		}
		for i = from; i < to; i++ {
			y1 = n_y[e_n1[i]]
			y2 = n_y[e_n2[i]]
			if y1 > y_max || y1 < y_min {

				use_1 = 0
			}
			if y2 > y_max || y2 < y_min {
				use_2 = 0
			}
			if use_1 == 0 && use_2 == 0 {
				continue
			}
			if y1 > y2 {
				down = 2
				val1 = y1
				y1 = y2
				y2 = val1
			}
			if y1 < y_min {
				y1 = y_min
			}
			if y2 > y_max {
				y2 = y_max
			}
			dx = math.Abs(n_x[e_n2[i]] - n_x[e_n1[i]])
			L = math.Sqrt(dx*dx + math.Pow(y2-y1, 2))
			if math.Pow(y2-y1, 2) < 1e-07 {

				continue
			}

			b = (y_max - y1) * w_val
			a = (y_max - y2) * w_val
			if use_1 == 0 {

				val2 = (a + 0.5*(b-a)) * L
				val1 = 0
			} else {
				if use_2 == 0 {
					val1 = (a + 0.5*(b-a)) * L
					val2 = 0
				} else {
					val1 = 0.5*a*L + 0.25*(b-a)*L + 0.125*(b-a)*L
					val2 = 0.5*a*L + 0.125*(b-a)*L
				}
			}
			if down == 1 {

				pos1 = e_n1[i]*3 + 1
				pos2 = e_n2[i]*3 + 1
			} else {

				pos1 = e_n2[i]*3 + 1
				pos2 = e_n1[i]*3 + 1
			}

			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], pos1, val1, 1)
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], pos2, val2, 1)
		}
	}
	return 0
}

func get_loads_and_supports() int32 {

	var i int32
	var j int32
	var pos int32
	for i = 0; i < n_f; i++ {
		femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], f_n[i]*3+f_dir[i]+1, f_val[i], 1)
	}
	for i = 0; i < n_d; i++ {
		if d_dir[i] > 2 {

			pos = d_n[i]*3 + d_dir[i] - 2
			femMatPutAdd((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos, pos, d_val[i], 1)
		} else {

			pos = d_n[i]*3 + d_dir[i] + 1
			if math.Abs(d_val[i]) <= 1e-07 {
				femMatSetZeroCol((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos)
				femMatSetZeroRow((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos)
				femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&u))[:], pos, 0, 0)

				femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], pos, 0, 0)
				femMatPutAdd((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos, pos, 1, 0)
			} else {
				for j = 1; j <= n_n*3; j++ {
					femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], j, -1*femMatGet((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], j, pos)*d_val[i], 1)
				}
				femMatSetZeroCol((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos)
				femMatSetZeroRow((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos)
				femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&u))[:], pos, d_val[i], 0)
				femMatPutAdd((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos, pos, femVecGet((*[1000000]tVector)(unsafe.Pointer(&F))[:], pos)/d_val[i], 0)
			}
		}
	}
	return 0
}

func get_int_forces(el int32, N1 []float64, N2 []float64, M1 []float64, M2 []float64, Q []float64) {

	var t float64
	var L float64
	var R float64
	var j int32
	var posj int32
	femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&D))[:])
	femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&B))[:])
	femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&DB))[:])
	femVecSetZero((*[1000000]tVector)(unsafe.Pointer(&ue))[:])
	femVecSetZero((*[1000000]tVector)(unsafe.Pointer(&Fe))[:])
	{

		for j = 1; j <= 6; j++ {
			if j < 4 {
				posj = e_n1[el]*3 + j
			} else {
				posj = e_n2[el]*3 + j - 3
			}
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&ue))[:], j, femVecGet((*[1000000]tVector)(unsafe.Pointer(&u))[:], posj), 0)
		}
	}

	t = e_t[el]
	get_D_matrix(el, t, (*[1000000]tMatrix)(unsafe.Pointer(&D))[:])
	get_B_matrix(el, (*[1000000]tMatrix)(unsafe.Pointer(&B))[:], c4goUnsafeConvert_float64(&L), c4goUnsafeConvert_float64(&R))
	femMatMatMult((*[1000000]tMatrix)(unsafe.Pointer(&D))[:], (*[1000000]tMatrix)(unsafe.Pointer(&B))[:], (*[1000000]tMatrix)(unsafe.Pointer(&DB))[:])

	femMatVecMult((*[1000000]tMatrix)(unsafe.Pointer(&DB))[:], (*[1000000]tVector)(unsafe.Pointer(&ue))[:], (*[1000000]tVector)(unsafe.Pointer(&Fe))[:])
	N1[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 1)
	N2[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 2)
	M1[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 3)
	M2[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 4)
	Q[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 5)
}

func print_result(fw *noarch.File) int32 {
	return 0
}

func generate_rand_out_file(fw *noarch.File) {

	var i int32
	noarch.Fprintf(fw, []byte("%li\n\x00"), n_n*8+1)
	noarch.Fprintf(fw, []byte("FAIL 3 2\n\x00"))
	for i = 0; i < n_n; i++ {
		noarch.Fprintf(fw, []byte("UY%li 2\n\x00"), i)
		noarch.Fprintf(fw, []byte("UX%li 2\n\x00"), i)
		noarch.Fprintf(fw, []byte("RT%li 2\n\x00"), i)
		noarch.Fprintf(fw, []byte("NX%li 2\n\x00"), i)
		noarch.Fprintf(fw, []byte("NY%li 2\n\x00"), i)
		noarch.Fprintf(fw, []byte("MX%li 2\n\x00"), i)
		noarch.Fprintf(fw, []byte("MY%li 2\n\x00"), i)
		noarch.Fprintf(fw, []byte("QQ%li 2\n\x00"), i)
	}

	noarch.Fprintf(fw, []byte("0\n\x00"))
}

func generate_d_type(type_ int32) []byte {
	switch type_ {
	case 0:

		return []byte("UY\x00")
	case 1:
		return []byte("UX\x00")
	case 2:
		return []byte("RT\x00")
	case 3:
		return []byte("EY\x00")
	case 4:
		return []byte("EX\x00")
	case 5:
		return []byte("ER\x00")
		break
	}
	return []byte("XX\x00")
}

func generate_f_type(type_ int32) []byte {
	switch type_ {
	case 0:

		return []byte("FY\x00")
	case 1:
		return []byte("FX\x00")
	case 2:
		return []byte("MT\x00")
		break
	}
	return []byte("XX\x00")
}

func generate_w_type(type_ int32) []byte {
	switch type_ {
	case 0:

		return []byte("TOP\x00")
	case 1:
		return []byte("BOT\x00")
	case 2:
		return []byte("SIZE\x00")
		break
	}
	return []byte("XX\x00")
}

func generate_fc_type(type_ int32) []byte {
	switch fail_type {
	case 1:
		switch type_ {
		case 0:

			return []byte("COMPR\x00")
		case 1:
			return []byte("TENS\x00")
		default:
			return []byte("UNKNOWN\x00")
			break
		}
	default:
		return []byte("XX\x00")
		break
	}
	return []byte("XX\x00")
}

func generate_rand_input_file(fw *noarch.File) {

	var i int32
	noarch.Fprintf(fw, []byte("%li\n\x00"), n_r_inp)
	for i = 0; i < n_r_inp; i++ {
		switch rand_type[i] {
		case 0:
			switch rand_indx[i] {
			case 0:

				noarch.Fprintf(fw, []byte("MAT%li_E1 %e 1 normal-1-02.dis\n\x00"), rand_pos[i], m_E1[rand_pos[i]])
			case 1:
				noarch.Fprintf(fw, []byte("MAT%li_E2 %e 1 normal-1-02.dis\n\x00"), rand_pos[i], m_E2[rand_pos[i]])
			case 2:
				noarch.Fprintf(fw, []byte("MAT%li_G %e 1 normal-1-02.dis\n\x00"), rand_pos[i], m_G[rand_pos[i]])
			case 3:
				noarch.Fprintf(fw, []byte("MAT%li_NU1 %e 1 normal-1-02.dis\n\x00"), rand_pos[i], m_nu1[rand_pos[i]])
			case 4:
				noarch.Fprintf(fw, []byte("MAT%li_NU2 %e 1 normal-1-02.dis\n\x00"), rand_pos[i], m_nu2[rand_pos[i]])
			case 5:
				noarch.Fprintf(fw, []byte("MAT%li_VF %e 1 normal-1-02.dis\n\x00"), rand_pos[i], m_vp[rand_pos[i]])
			case 6:
				noarch.Fprintf(fw, []byte("MAT%li_T %e 1 normal-1-02.dis\n\x00"), rand_pos[i], m_t[rand_pos[i]])
				break
			}
		case 1:
			switch rand_indx[i] {
			case 0:

				noarch.Fprintf(fw, []byte("N%li_X %e 1 normal-1-02.dis\n\x00"), rand_pos[i], n_x[rand_pos[i]])
			case 1:
				noarch.Fprintf(fw, []byte("N%li_Y %e 1 normal-1-02.dis\n\x00"), rand_pos[i], n_y[rand_pos[i]])
				break
			}
		case 2:

			noarch.Fprintf(fw, []byte("E%li_WIDTH %e 1 normal-1-02.dis\n\x00"), rand_pos[i], e_t[rand_pos[i]])
		case 3:

			noarch.Fprintf(fw, []byte("D%li_%s_SIZE %e 1 normal-1-02.dis\n\x00"), rand_pos[i], generate_d_type(rand_indx[i]), d_val[rand_pos[i]])
		case 4:

			noarch.Fprintf(fw, []byte("F%li_%s_SIZE %e 1 normal-1-02.dis\n\x00"), rand_pos[i], generate_f_type(rand_indx[i]), f_val[rand_pos[i]])
		case 5:
			switch rand_indx[i] {
			case 0:

				noarch.Fprintf(fw, []byte("W_%s %e 1 normal-1-02.dis\n\x00"), generate_w_type(rand_indx[i]), w_top)
			case 1:
				noarch.Fprintf(fw, []byte("W_%s %e 1 normal-1-02.dis\n\x00"), generate_w_type(rand_indx[i]), w_bot)
			case 2:
				noarch.Fprintf(fw, []byte("W_%s %e 1 normal-1-02.dis\n\x00"), generate_w_type(rand_indx[i]), w_val)
				break
			}
		case 6:

			noarch.Fprintf(fw, []byte("FC_%s_%li %e 1 normal-1-02.dis\n\x00"), generate_fc_type(rand_indx[i]), rand_indx[i], fail_data[rand_indx[i]])
		default:
			noarch.Fprintf(msgout, []byte("Unused input random variable %li!\n\x00"), i)
			break
		}
	}
}

func fail_test_concrete() int32 {

	var N1 float64
	var N2 float64
	var Q float64
	var M1 float64
	var M2 float64
	var h float64
	var I1 float64
	var J2 float64
	var J3 float64
	var alpha float64
	var beta float64
	var lambda float64
	var k float64
	var cos3f float64
	var c1 float64
	var c2 float64
	var fc float64
	var s1 float64
	var s2 float64
	var sm float64
	var tmp float64
	var i int32
	k = fail_data[1] / fail_data[0]
	for i = 0; i < n_e; i++ {

		get_int_forces(i, c4goUnsafeConvert_float64(&N1), c4goUnsafeConvert_float64(&N2), c4goUnsafeConvert_float64(&M1), c4goUnsafeConvert_float64(&M2), c4goUnsafeConvert_float64(&Q))
		h = e_t[i]
		s1 = 6*M1/h + N1/h
		s2 = 6*M2/h + N2/h
		if s1 < s2 {
			tmp = s1
			s1 = s2
			s2 = tmp
		}
		I1 = s1 + s2
		sm = I1 / 3
		J3 = (s1 - sm) * (s2 - sm)
		J2 = 1 / 6 * (math.Pow(s1-s2, 2) + s1*s1 + s2*s2)
		alpha = 1 / (9 * math.Pow(k, 1.4))
		beta = 1 / (3.7 * math.Pow(k, 1.1))
		cos3f = 3 * math.Pow(3, 0.5) / 2 * (J3 / math.Pow(J2, 1.5))
		c1 = 1 / (0.7 * math.Pow(k, 1.1))
		c2 = 1 - 6.8*math.Pow(k-0.07, 2)
		if cos3f < 0 {
			lambda = c1 * math.Cos(3.141592653589793/3-1/3*math.Acos(0-c2*cos3f))
		} else {
			lambda = c1 * math.Cos(1/3*math.Acos(0-c2*cos3f))
		}
		fc = alpha*(J2/math.Pow(fail_data[0], 2)) + lambda*(math.Sqrt(J2)/fail_data[0]) + beta*(I1/fail_data[0])
		if fc > 1 {

			return 1
		}
	}
	return 0
}

func fail_test() int32 {
	switch fail_type {
	case 1:

		return fail_test_concrete()
	case 0:
		fallthrough
	default:

		return 0
		break
	}
	return 0
}

func compute_price() float64 {

	var price float64
	var volume float64
	var dx float64
	var dpx float64
	var dy float64
	var i int32
	price = 0
	for i = 0; i < n_e; i++ {

		dx = math.Abs(n_x[e_n2[i]] - n_x[e_n1[i]])

		dpx = n_x[e_n2[i]] + n_x[e_n1[i]]
		dy = math.Abs(n_y[e_n2[i]] - n_y[e_n1[i]])
		if dx <= 1e-07 {

			volume = dy * (2 * 3.141592653589793 * n_x[e_n2[i]])
		} else {
			if dy <= 1e-07 {

				volume = 3.141592653589793 * math.Abs(math.Pow(n_x[e_n2[i]], 2)-math.Pow(n_x[e_n1[i]], 2))
			} else {

				volume = 3.141592653589793 * dpx * math.Sqrt(dy*dy+dx*dx)
			}
		}
		price += e_t[i] * volume * m_vp[e_mat[i]]
	}
	return price
}

func optim_replace_data(ifld []float64) int32 {

	var i int32
	if len(ifld) == 0 || n_r_opt < 1 {
		return 0
	}
	for i = 0; i < n_r_opt; i++ {
		switch opt_type[i] {
		case 0:
			switch opt_indx[i] {
			case 0:

				m_E1[opt_pos[i]] = ifld[i]
			case 1:
				m_E2[opt_pos[i]] = ifld[i]
			case 2:
				m_G[opt_pos[i]] = ifld[i]
			case 3:
				m_nu1[opt_pos[i]] = ifld[i]
			case 4:
				m_nu2[opt_pos[i]] = ifld[i]
			case 5:
				m_q[opt_pos[i]] = ifld[i]
			case 6:
				m_t[opt_pos[i]] = ifld[i]
				break
			}
		case 1:
			switch opt_indx[i] {
			case 0:

				n_x[opt_pos[i]] = ifld[i]
			case 1:
				n_y[opt_pos[i]] = ifld[i]
				break
			}
		case 2:

			e_t[opt_pos[i]] = ifld[i]
		case 3:

			d_val[opt_pos[i]] = ifld[i]
		case 4:

			f_val[opt_pos[i]] = ifld[i]
		case 5:
			switch opt_indx[i] {
			case 0:

				w_top = ifld[i]
			case 1:
				w_bot = ifld[i]
			case 2:
				w_val = ifld[i]
				break
			}
		case 6:
			if opt_indx[i] < n_fail {

				fail_data[opt_indx[i]] = ifld[i]
			}
		default:
			noarch.Fprintf(msgout, []byte("Unused input optim variable %li!\n\x00"), i)
			break
		}
	}
	return 0
}

func print_help(argc int32, argv [][]byte) {

	fmt.Printf("\neSHELL 1.0: axisymetric shells solver\n")
	fmt.Printf("(C) 2010 VSB-TU of Ostrava \n")
	fmt.Printf("(C) 2003-2010 Jiri Brozovsky (uFEM libraries)\n")
	fmt.Printf("\nThis is free software licensed under GNU GPL 2.0\n")
	noarch.Printf([]byte("\nUsage: %s [parameters] <input >output\n\x00"), argv[0])
	fmt.Printf("\nParameters:\n")
	fmt.Printf("   -s        ... force solution only output\n")
	fmt.Printf("   -g        ... generate random data only \n")
	fmt.Printf("   -p        ... compute price function only\n")
	fmt.Printf("   -w        ... write input data and finish\n")
	fmt.Printf("   -h        ... print this help\n")
}

func cmd_param(argc int32, argv [][]byte) int32 {

	var i int32
	for i = 1; i < argc; i++ {
		if noarch.Strcmp(argv[i], []byte("-h\x00")) == 0 || noarch.Strcmp(argv[i], []byte("--help\x00")) == 0 {
			print_help(argc, argv)
			noarch.Exit(0)
		}
		if noarch.Strcmp(argv[i], []byte("-s\x00")) == 0 || noarch.Strcmp(argv[i], []byte("--solution\x00")) == 0 {
			solution_only = 1
			price_only = 0
			random_only = 0
		}
		if noarch.Strcmp(argv[i], []byte("-g\x00")) == 0 || noarch.Strcmp(argv[i], []byte("-r\x00")) == 0 || noarch.Strcmp(argv[i], []byte("--random\x00")) == 0 {
			solution_only = 0
			price_only = 0
			random_only = 1
		}
		if noarch.Strcmp(argv[i], []byte("-p\x00")) == 0 || noarch.Strcmp(argv[i], []byte("--price\x00")) == 0 {
			solution_only = 0
			price_only = 1
			random_only = 0
		}
		if noarch.Strcmp(argv[i], []byte("-w\x00")) == 0 || noarch.Strcmp(argv[i], []byte("--price\x00")) == 0 {
			write_only = 1
		}
	}
	return 0
}

func main() {
	argc := int32(len(os.Args))
	argv := [][]byte{}
	for _, argvSingle := range os.Args {
		argv = append(argv, []byte(argvSingle))
	}
	defer noarch.AtexitRun()

	var stat int32
	msgout = noarch.Stderr
	cmd_param(argc, argv)
	stat += read_input_data(noarch.Stdin)
	stat += alloc_solver_data()
	stat += optim_replace_data(opt_data)
	if write_only == 1 {
		write_input_data(noarch.Stdout)
		return
	}
	if solution_only == 1 {
		stat += get_matrix()
		stat += generate_water_load_x()
		stat += get_loads_and_supports()
		stat = femEqsCGwJ((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], (*[1000000]tVector)(unsafe.Pointer(&F))[:], (*[1000000]tVector)(unsafe.Pointer(&u))[:], 1e-09, 6*3*n_n)
	}
	if n_r_inp > 0 && random_only == 1 {
		if solution_only != 0 {
			print_result(noarch.Stderr)
		}
		generate_rand_input_file(noarch.Stdout)
		generate_rand_out_file(noarch.Stdout)
	} else {
		if solution_only == 1 {
			print_result(noarch.Stdout)
		}
	}
	if solution_only == 1 {
		if fail_test() != 0 {
			noarch.Fprintf(noarch.Stderr, []byte("# Structure FAILED\n\x00"))
		}
	}
	if price_only == 1 {
		if solution_only == 1 {
			noarch.Fprintf(msgout, []byte("# Price is %f\n\x00\x00"), compute_price())
		} else {
			noarch.Fprintf(noarch.Stdout, []byte("%e\n\x00"), compute_price())
		}
	}
	return
}

func femMatNull(mat []tMatrix) {

	mat[0].type_ = 0
	mat[0].rows = 0
	mat[0].cols = 0
	mat[0].len_ = 0
	mat[0].pos = nil
	mat[0].data = nil
	mat[0].frompos = nil
	mat[0].defpos = nil
}

func femMatFree(mat []tMatrix) {
	mat[0].type_ = 0
	mat[0].rows = 0
	mat[0].cols = 0
	mat[0].len_ = 0
	femIntFree(mat[0].pos)
	femDblFree(mat[0].data)
	femIntFree(mat[0].frompos)
	femIntFree(mat[0].defpos)
}

func femMatAlloc(mat []tMatrix, type_ int32, rows int32, cols int32, bandwidth int32, rowdesc []int32) int32 {
	var sum int32
	var i int32
	femMatNull(mat)
	if type_ >= 0 && type_ <= 1 {
		mat[0].type_ = type_
		switch type_ {
		case 0:
			mat[0].rows = rows
			mat[0].cols = cols
			mat[0].len_ = cols * rows
			if len((func() []float64 {
				mat[0].data = femDblAlloc(mat[0].len_)
				return mat[0].data
			}())) == 0 {
				goto memFree
			}
			mat[0].pos = nil
			mat[0].frompos = nil
			mat[0].defpos = nil
		case 1:
			mat[0].rows = rows
			mat[0].cols = cols
			if len((func() []int32 {
				mat[0].defpos = femIntAlloc(mat[0].rows)
				return mat[0].defpos
			}())) == 0 {
				goto memFree
			}
			if len((func() []int32 {
				mat[0].frompos = femIntAlloc(mat[0].rows)
				return mat[0].frompos
			}())) == 0 {
				goto memFree
			}
			if bandwidth > 0 && len(rowdesc) == 0 {
				mat[0].len_ = rows * bandwidth
				if len((func() []float64 {
					mat[0].data = femDblAlloc(mat[0].len_)
					return mat[0].data
				}())) == 0 {
					goto memFree
				}
				if len((func() []int32 {
					mat[0].pos = femIntAlloc(mat[0].len_)
					return mat[0].pos
				}())) == 0 {
					goto memFree
				}
				for i = 0; i < rows; i++ {
					mat[0].frompos[i] = bandwidth * i
				}
			} else {
				sum = 0
				for i = 0; i < rows; i++ {
					sum += rowdesc[i]
					mat[0].defpos[i] = rowdesc[i]
					mat[0].frompos[i] = sum - rowdesc[i]
				}
				mat[0].len_ = sum
				if len((func() []float64 {
					mat[0].data = femDblAlloc(mat[0].len_)
					return mat[0].data
				}())) == 0 {
					goto memFree
				}
				if len((func() []int32 {
					mat[0].pos = femIntAlloc(sum)
					return mat[0].pos
				}())) == 0 {
					goto memFree
				}
			}
			break
		}
		return 0
	} else {
		return -3
	}
memFree:
	;
	femMatFree(mat)
	return -4
}

func femMatGet(mat []tMatrix, row int32, col int32) float64 {

	var pos int32
	var i int32
	switch mat[0].type_ {
	case 0:
		pos = (row-1)*mat[0].cols + (col - 1)
		return mat[0].data[pos]
	case 1:
		for i = mat[0].frompos[row-1]; i < mat[0].frompos[row-1]+mat[0].defpos[row-1]; i++ {
			if mat[0].pos[i] == 0 {
				break
			}
			if mat[0].pos[i] == col {
				return mat[0].data[i]
				break
			}
		}
	default:
		return 0
		break
	}
	return 0
}

func femMatPutAdd(mat []tMatrix, row int32, col int32, val float64, mode int32) (c4goDefaultReturn int32) {

	var pos int32
	var i int32
	switch mat[0].type_ {
	case 0:
		pos = (row-1)*mat[0].cols + (col - 1)
		if mode == 1 {
			mat[0].data[pos] += val
		} else {
			mat[0].data[pos] = val
		}
		return 0
	case 1:
		{

			for i = mat[0].frompos[row-1]; i < mat[0].frompos[row-1]+mat[0].defpos[row-1]; i++ {
				if mat[0].pos[i] == col {
					if mode == 1 {
						mat[0].data[i] += val
					} else {
						mat[0].data[i] = val
					}
					return 0
				}
				if mat[0].pos[i] == 0 {

					mat[0].pos[i] = col
					if mode == 1 {
						mat[0].data[i] += val
					} else {
						mat[0].data[i] = val
					}
					return 0
				}
			}
		}

		return -11
	default:
		return -3
		break
	}
	return
}

func femMatPrn(mat []tMatrix, name []byte) {

	{
	}
}

func femMatPrnF(fname []byte, mat []tMatrix) int32 {

	var fw *noarch.File
	var rv int32
	var i int32
	var j int32
	if (func() *noarch.File {
		fw = noarch.Fopen(fname, []byte("w\x00"))
		return fw
	}()) == nil {
		return -2
	}
	for i = 1; i <= mat[0].rows; i++ {
		for j = 1; j <= mat[0].cols; j++ {
			noarch.Fprintf(fw, []byte(" %e \x00"), femMatGet(mat, i, j))
		}
		noarch.Fprintf(fw, []byte("\n\x00"))
	}
	if noarch.Fclose(fw) != 0 {
		rv = -2
	}
	return rv
}

func femSparseMatPrnF(fname []byte, mat []tMatrix) int32 {

	var fw *noarch.File
	var rv int32
	var i int32
	var j int32
	var sum int32
	if mat[0].type_ != 1 {
		return -3
	}
	if (func() *noarch.File {
		fw = noarch.Fopen(fname, []byte("w\x00"))
		return fw
	}()) == nil {
		return -2
	}
	noarch.Fprintf(fw, []byte("%li %li\n\x00"), mat[0].rows, mat[0].cols)
	for i = 0; i < mat[0].rows; i++ {
		sum = 0
		for j = mat[0].frompos[i]; j < mat[0].frompos[i]+mat[0].defpos[i]; j++ {
			if mat[0].pos[j] >= 0 {
				sum++
			} else {
				break
			}
		}
		noarch.Fprintf(fw, []byte("%li %li \x00"), i+1, sum)
		for j = mat[0].frompos[i]; j < mat[0].frompos[i]+sum; j++ {
			noarch.Fprintf(fw, []byte("%li %e \x00"), mat[0].pos[j], mat[0].data[j])
		}
		noarch.Fprintf(fw, []byte("\n\x00"))
	}
	if noarch.Fclose(fw) != 0 {
		rv = -2
	}
	return rv
}

func femSparseMarketMatPrnF(fname []byte, mat []tMatrix) int32 {

	var fw *noarch.File
	var rv int32
	var i int32
	var j int32
	var sum int32
	if mat[0].type_ != 1 {
		return -3
	}
	if (func() *noarch.File {
		fw = noarch.Fopen(fname, []byte("w\x00"))
		return fw
	}()) == nil {
		return -2
	}
	noarch.Fprintf(fw, []byte("%%%%MatrixMarket matrix coordinate real general\n\x00"))
	noarch.Fprintf(fw, []byte("%li %li %li\n\x00"), mat[0].rows, mat[0].cols, mat[0].len_)
	for i = 0; i < mat[0].rows; i++ {
		sum = 0
		for j = mat[0].frompos[i]; j < mat[0].frompos[i]+mat[0].defpos[i]; j++ {
			if mat[0].pos[j] >= 0 {
				sum++
			} else {
				break
			}
		}
		for j = mat[0].frompos[i]; j < mat[0].frompos[i]+sum; j++ {
			noarch.Fprintf(fw, []byte("%li %li %e\n\x00"), i+1, mat[0].pos[j], mat[0].data[j])
		}
	}
	if noarch.Fclose(fw) != 0 {
		rv = -2
	}
	return rv
}

func femSparseMatReadF(fname []byte, mat []tMatrix) int32 {

	var fw *noarch.File
	var rv int32
	var i int32
	var j int32
	var k int32
	var tmp int32
	var sum int32
	var size int32
	var ensize int32
	var pos0 []int32
	var data0 []float64
	if (func() *noarch.File {
		fw = noarch.Fopen(fname, []byte("r\x00"))
		return fw
	}()) == nil {
		return -2
	}
	noarch.Fscanf(fw, []byte("%li %li\n\x00"), (*[1000000]int32)(unsafe.Pointer(&mat[0].rows))[:], (*[1000000]int32)(unsafe.Pointer(&mat[0].cols))[:])
	if mat[0].rows <= 0 || mat[0].cols <= 0 {
		return -2
	}
	if len((func() []int32 {
		mat[0].frompos = femIntAlloc(mat[0].rows)
		return mat[0].frompos
	}())) == 0 {
		rv = -4
		goto memFree
	}
	if len((func() []int32 {
		mat[0].defpos = femIntAlloc(mat[0].rows)
		return mat[0].defpos
	}())) == 0 {
		rv = -4
		goto memFree
	}
	size = mat[0].rows * 300
	if len((func() []int32 {
		mat[0].pos = femIntAlloc(size)
		return mat[0].pos
	}())) == 0 {
		rv = -4
		goto memFree
	}
	if len((func() []float64 {
		mat[0].data = femDblAlloc(size)
		return mat[0].data
	}())) == 0 {
		rv = -4
		goto memFree
	}
	mat[0].type_ = 1
	sum = 0
	for i = 0; i < mat[0].rows; i++ {
		noarch.Fscanf(fw, []byte("%li %li \x00"), c4goUnsafeConvert_int32(&tmp), mat[0].defpos[i:])
		if i > 0 {
			mat[0].frompos[i] = mat[0].frompos[i-1] + mat[0].defpos[i-1]
		} else {

			mat[0].frompos[i] = 0
		}
		for j = 0; j < mat[0].defpos[i]; j++ {
			if sum >= size {

				ensize = size + 2*size*(i/mat[0].rows)
				if len((func() []int32 {
					pos0 = femIntAlloc(ensize)
					return pos0
				}())) == 0 {
					rv = -4
					goto memFree
				}
				if len((func() []float64 {
					data0 = femDblAlloc(ensize)
					return data0
				}())) == 0 {
					rv = -4
					goto memFree
				}
				for k = 0; k < sum; k++ {
					pos0[k] = mat[0].pos[k]
					data0[k] = mat[0].data[k]
				}
				_ = mat[0].pos
				_ = mat[0].data
				mat[0].pos = pos0
				mat[0].data = data0
				pos0 = nil
				data0 = nil
			}
			noarch.Fscanf(fw, []byte("%li %f \x00\x00"), mat[0].pos[sum:], mat[0].data[sum:])
			sum++
		}
	}
	if noarch.Fclose(fw) != 0 {
		rv = -2
	}
	return rv
memFree:
	;
	femMatFree(mat)
	return rv
}

func femMatOut(a []tMatrix, fw *noarch.File) int32 {

	var rv int32
	var i int32
	var j int32
	noarch.Fprintf(fw, []byte(" %li %li\n\x00"), a[0].rows, a[0].cols)
	for i = 1; i <= a[0].rows; i++ {
		for j = 1; j <= a[0].cols; j++ {
			noarch.Fprintf(fw, []byte(" %e \n\x00"), femMatGet(a, i, j))
		}
	}
	return rv
}

func femMatSetZeroBig(a []tMatrix) {

	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] = 0
	}
}

func femMatSetZero(a []tMatrix) {

	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] = 0
	}
}

func femMatSetZeroRow(a []tMatrix, row int32) {

	var i int32
	if a[0].type_ == 1 {
		for i = a[0].frompos[row-1]; i < a[0].frompos[row-1]+a[0].defpos[row-1]; i++ {
			if a[0].pos[i] == 0 {
				break
			}
			a[0].data[i] = 0
		}
	} else {

		for i = 1; i <= a[0].cols; i++ {
			femMatPutAdd(a, row, i, 0, 0)
		}
	}
}

func femMatSetZeroCol(a []tMatrix, Col int32) {

	var i int32
	var j int32
	var ifrom int32
	var ito int32
	var ipos int32
	if a[0].type_ == 1 {
		ifrom = a[0].pos[a[0].frompos[Col-1]] - 1
		ito = a[0].pos[a[0].frompos[Col-1]+a[0].defpos[Col-1]-1] - 1
		for i = ifrom; i < ito; i++ {
			for j = a[0].frompos[i]; j < a[0].frompos[i]+a[0].defpos[i]; j++ {
				if a[0].pos[j] == Col {
					a[0].data[j] = 0
				}
			}
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			femMatPutAdd(a, i, Col, 0, 0)
		}
	}
}

func femVecNull(mat []tVector) {

	mat[0].type_ = 0
	mat[0].rows = 0
	mat[0].len_ = 0
	mat[0].pos = nil
	mat[0].data = nil
}

func femVecFree(mat []tVector) {
	mat[0].type_ = 0
	mat[0].rows = 0
	mat[0].len_ = 0
	femIntFree(mat[0].pos)
	femDblFree(mat[0].data)
}

func femVecAlloc(mat []tVector, type_ int32, rows int32, items int32) int32 {
	femVecNull(mat)
	if type_ >= 0 && type_ <= 1 {
		mat[0].type_ = type_
		switch type_ {
		case 0:
			mat[0].rows = rows
			mat[0].len_ = rows
			if len((func() []float64 {
				mat[0].data = femDblAlloc(mat[0].len_)
				return mat[0].data
			}())) == 0 {
				goto memFree
			}
			mat[0].pos = nil
		case 1:

			noarch.Exit(-3)
			mat[0].rows = rows
			if items > 0 {
				mat[0].len_ = items
				if len((func() []float64 {
					mat[0].data = femDblAlloc(mat[0].len_)
					return mat[0].data
				}())) == 0 {
					goto memFree
				}
				if len((func() []int32 {
					mat[0].pos = femIntAlloc(mat[0].len_)
					return mat[0].pos
				}())) == 0 {
					goto memFree
				}
			} else {
				goto memFree
			}
			break
		}
		return 0
	} else {
		return -3
	}
memFree:
	;
	femVecFree(mat)
	return -4
}

func femVecPutAdd(vec []tVector, pos int32, val float64, mode int32) int32 {
	if pos > vec[0].rows {

		return -11
	}
	switch vec[0].type_ {
	case 0:
		if mode == 0 {

			vec[0].data[pos-1] = val
		} else {

			vec[0].data[pos-1] += val
		}
	case 1:

		noarch.Exit(-3)
	default:
		return -5
		break
	}
	return 0
}

func femVecGet(vec []tVector, pos int32) float64 {
	if pos > vec[0].rows {

		return float64(0)
	}
	switch vec[0].type_ {
	case 0:
		return vec[0].data[pos-1]
	case 1:

		noarch.Exit(0)
	default:
		return float64(0)
		break
	}
	return float64(0)
}

func femVecPrn(mat []tVector, name []byte) {

	{
	}
}

func femVecPrnF(fname []byte, mat []tVector) int32 {

	var fw *noarch.File
	var rv int32
	var i int32
	if (func() *noarch.File {
		fw = noarch.Fopen(fname, []byte("w\x00"))
		return fw
	}()) == nil {
		return -2
	}
	for i = 1; i <= mat[0].rows; i++ {
		noarch.Fprintf(fw, []byte(" %e \x00"), femVecGet(mat, i))
	}
	noarch.Fprintf(fw, []byte("\n\x00"))
	if noarch.Fclose(fw) != 0 {
		rv = -2
	}
	return rv
}

func femVecOut(a []tVector, fw *noarch.File) int32 {

	var rv int32
	var i int32
	noarch.Fprintf(fw, []byte(" %li\n\x00"), a[0].rows)
	for i = 1; i <= a[0].rows; i++ {
		noarch.Fprintf(fw, []byte(" %e \n\x00"), femVecGet(a, i))
	}
	return rv
}

func femVecSetZeroBig(a []tVector) {

	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] = 0
	}
}

func femVecSetZero(a []tVector) {

	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] = 0
	}
}

func femVecClone(src []tVector, dest []tVector) int32 {

	var i int32
	if src[0].type_ != 0 || dest[0].type_ != 0 {
		return -5
	}
	if src[0].len_ != dest[0].len_ {
		return -9
	}
	for i = 0; i < src[0].len_; i++ {
		dest[0].data[i] = src[0].data[i]
	}
	return 0
}

func femVecVecMultBig(a []tVector, b []tVector) float64 {

	var i int32
	var mult float64
	mult = 0
	if a[0].type_ == 0 && b[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			mult += a[0].data[i] * b[0].data[i]
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			mult += femVecGet(a, i) * femVecGet(b, i)
		}
	}
	return mult
}

func femVecVecMult(a []tVector, b []tVector) float64 {

	var i int32
	var mult float64
	mult = 0
	if a[0].type_ == 0 && b[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			mult += a[0].data[i] * b[0].data[i]
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			mult += femVecGet(a, i) * femVecGet(b, i)
		}
	}
	return mult
}

func femVecVecMulttoMat(a []tVector, b []tVector, c []tMatrix) int32 {

	var i int32
	var j int32
	if a[0].type_ == 0 && b[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			for j = 0; j < a[0].rows; j++ {
				c[0].data[i*c[0].cols+j] = a[0].data[i] * b[0].data[j]
			}
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			for j = 1; j <= a[0].rows; j++ {
				femMatPutAdd(c, i, j, femVecGet(a, i)*femVecGet(b, j), 0)
			}
		}
	}
	return 0
}

func femValVecMult(val float64, a []tVector, b []tVector) int32 {

	var i int32
	for i = 0; i < a[0].len_; i++ {
		b[0].data[i] = a[0].data[i] * val
	}
	return 0
}

func femValVecMultSelf(val float64, a []tVector) int32 {

	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] *= val
	}
	return 0
}

func femValMatMultSelf(val float64, a []tMatrix) int32 {

	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] *= val
	}
	return 0
}

func femVecMatMult(a []tVector, b []tMatrix, c []tVector) int32 {

	var i int32
	var j int32
	var val float64
	if a[0].rows != b[0].rows || b[0].cols != c[0].rows {
		return -9
	}
	if c[0].type_ != 0 {
		return -3
	}
	if a[0].type_ == 0 && b[0].type_ == 0 && c[0].type_ == 0 {
		for i = 0; i < b[0].cols; i++ {
			val = 0
			for j = 0; j < a[0].rows; j++ {
				val += a[0].data[j] * b[0].data[i+b[0].cols*j]
			}
			c[0].data[i] = val
		}
	} else {
		for i = 1; i <= b[0].cols; i++ {
			val = 0
			for j = 1; j <= a[0].rows; j++ {
				val += femVecGet(a, j) * femMatGet(b, j, i)
			}
			femVecPutAdd(c, i, val, 0)
		}
	}
	return 0
}

func femVecMatVecMult(a []tVector, b []tMatrix, c []tVector) float64 {

	var i int32
	var j int32
	var val float64
	var sum_tot float64
	sum_tot = 0
	if a[0].rows != b[0].rows || b[0].cols != c[0].rows {
		return float64(-9)
	}
	if c[0].type_ != 0 {
		return float64(-3)
	}
	if a[0].type_ == 0 && b[0].type_ == 0 && c[0].type_ == 0 {
		for i = 0; i < b[0].cols; i++ {
			val = 0
			for j = 0; j < a[0].rows; j++ {
				val += a[0].data[j] * b[0].data[i+b[0].cols*j]
			}
			sum_tot += c[0].data[i] * val
		}
	} else {
		for i = 1; i <= b[0].cols; i++ {
			val = 0
			for j = 1; j <= a[0].rows; j++ {
				val += femVecGet(a, j) * femMatGet(b, j, i)
			}
			sum_tot += femVecGet(c, i) * val
		}
	}
	return sum_tot
}

func femMatVecMultBig(a []tMatrix, b []tVector, c []tVector) int32 {

	var i int32
	var j int32
	var val float64
	if a[0].cols != b[0].rows || c[0].rows != a[0].rows {
		return -9
	}
	if c[0].type_ != 0 {
		return -3
	}
	if a[0].type_ == 0 && b[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			val = 0
			for j = 0; j < a[0].cols; j++ {
				val += b[0].data[j] * a[0].data[j+i*a[0].cols]
			}
			c[0].data[i] = val
		}
	} else {
		if a[0].type_ == 1 && b[0].type_ == 0 {
			femVecSetZero(c)
			for i = 0; i < a[0].rows; i++ {
				val = 0
				for j = a[0].frompos[i]; j < a[0].frompos[i]+a[0].defpos[i]; j++ {
					if a[0].pos[j] <= 0 {
						break
					}
					val += a[0].data[j] * b[0].data[a[0].pos[j]-1]
				}
				c[0].data[i] = val
			}
		} else {
			for i = 1; i <= a[0].rows; i++ {
				val = 0
				for j = 1; j <= a[0].cols; j++ {
					val += femMatGet(a, i, j) * femVecGet(b, j)
				}
				femVecPutAdd(c, i, val, 0)
			}
		}
	}
	return 0
}

func femMatVecMult(a []tMatrix, b []tVector, c []tVector) int32 {

	var i int32
	var j int32
	var val float64
	if a[0].cols != b[0].rows || c[0].rows != a[0].rows {
		return -9
	}
	if c[0].type_ != 0 {
		return -3
	}
	if a[0].type_ == 0 && b[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			val = 0
			for j = 0; j < a[0].cols; j++ {
				val += b[0].data[j] * a[0].data[j+i*a[0].cols]
			}
			c[0].data[i] = val
		}
	} else {
		if a[0].type_ == 1 && b[0].type_ == 0 {
			femVecSetZero(c)
			for i = 0; i < a[0].rows; i++ {
				val = 0
				for j = a[0].frompos[i]; j < a[0].frompos[i]+a[0].defpos[i]; j++ {
					if a[0].pos[j] <= 0 {
						break
					}
					val += a[0].data[j] * b[0].data[a[0].pos[j]-1]
				}
				c[0].data[i] = val
			}
		} else {
			for i = 1; i <= a[0].rows; i++ {
				val = 0
				for j = 1; j <= a[0].cols; j++ {
					val += femMatGet(a, i, j) * femVecGet(b, j)
				}
				femVecPutAdd(c, i, val, 0)
			}
		}
	}
	return 0
}

func femVecLinCombBig(amult float64, a []tVector, bmult float64, b []tVector, c []tVector) int32 {

	var i int32
	if a[0].type_ == 0 && b[0].type_ == 0 && c[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			c[0].data[i] = amult*a[0].data[i] + bmult*b[0].data[i]
		}
	} else {

		for i = 1; i <= a[0].rows; i++ {
			femVecPutAdd(c, i, femVecGet(a, i)*amult+femVecGet(b, i)*bmult, 0)
		}
	}
	return 0
}

func femVecLinComb(amult float64, a []tVector, bmult float64, b []tVector, c []tVector) int32 {

	var i int32
	if a[0].type_ == 0 && b[0].type_ == 0 && c[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			c[0].data[i] = amult*a[0].data[i] + bmult*b[0].data[i]
		}
	} else {

		for i = 1; i <= a[0].rows; i++ {
			femVecPutAdd(c, i, femVecGet(a, i)*amult+femVecGet(b, i)*bmult, 0)
		}
	}
	return 0
}

func femMatMatMult(a []tMatrix, b []tMatrix, c []tMatrix) int32 {

	var i int32
	var j int32
	var k int32
	var val float64
	if a[0].cols != b[0].rows || b[0].cols != c[0].cols || a[0].rows != c[0].rows {
		return -9
	}
	if c[0].type_ != 0 {
		return -3
	}
	if a[0].type_ == 0 && b[0].type_ == 0 && c[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			for j = 0; j < b[0].cols; j++ {
				val = 0
				for k = 0; k < a[0].cols; k++ {

					val += a[0].data[i*a[0].cols+k] * b[0].data[k*b[0].cols+j]
				}

				c[0].data[i*c[0].cols+j] = val
			}
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			for j = 1; j <= b[0].cols; j++ {
				val = 0
				for k = 1; k <= a[0].cols; k++ {
					val += femMatGet(a, i, k) * femMatGet(b, k, j)
				}
				femMatPutAdd(c, i, j, val, 0)
			}
		}
	}
	return 0
}

func femMatLinComb(am float64, a []tMatrix, bm float64, b []tMatrix, c []tMatrix) int32 {

	var i int32
	var j int32
	var val float64
	if a[0].cols != b[0].cols || a[0].rows != b[0].rows || a[0].rows != c[0].rows || a[0].cols != c[0].cols {
		return -9
	}
	if c[0].type_ != 0 {
		return -3
	}
	if a[0].type_ == 0 && b[0].type_ == 0 && c[0].type_ == 0 {
		for i = 0; i < a[0].rows*a[0].cols; i++ {
			c[0].data[i] = am*a[0].data[i] + bm*b[0].data[i]
		}
	} else {
		for i = 1; i <= c[0].rows; i++ {
			for j = 1; j <= c[0].cols; j++ {
				val = am*femMatGet(a, i, j) + bm*femMatGet(b, i, j)
				femMatPutAdd(c, i, j, val, 0)
			}
		}
	}
	return 0
}

func femMatTran(a []tMatrix, b []tMatrix) int32 {

	var i int32
	var j int32
	if a[0].cols != b[0].rows || b[0].cols != a[0].rows {
		return -9
	}
	if a[0].type_ != 0 || b[0].type_ != 0 {
		return -9
	}
	for i = 0; i < a[0].rows; i++ {
		for j = 0; j < a[0].cols; j++ {
			if a[0].cols == a[0].rows {
				b[0].data[j*a[0].cols+i] = a[0].data[i*a[0].cols+j]
			} else {
				femMatPutAdd(b, j+1, i+1, femMatGet(a, i+1, j+1), 0)
			}
		}
	}
	return 0
}

func femMatNormBig(a []tMatrix) float64 {

	var Norm float64
	var MaxNorm float64
	var val float64
	var i int32
	var j int32
	MaxNorm = 0
	if a[0].type_ == 1 {
		for i = 0; i < a[0].rows; i++ {
			Norm = 0
			for j = a[0].frompos[i]; j < a[0].frompos[i]+a[0].defpos[i]; j++ {
				if a[0].pos[j] <= 0 {
					break
				}
				Norm += a[0].data[j] * a[0].data[j]
			}
			Norm = math.Sqrt(Norm)
			if Norm > MaxNorm {
				MaxNorm = Norm
			}
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			Norm = 0
			for j = 1; j <= a[0].cols; j++ {
				val = femMatGet(a, i, j)
				Norm += val * val
			}
			Norm = math.Sqrt(Norm)
			if Norm > MaxNorm {
				MaxNorm = Norm
			}
		}
	}
	return MaxNorm
}

func femMatNorm(a []tMatrix) float64 {

	var Norm float64
	var MaxNorm float64
	var val float64
	var i int32
	var j int32
	MaxNorm = 0
	if a[0].type_ == 1 {
		for i = 0; i < a[0].rows; i++ {
			Norm = 0
			for j = a[0].frompos[i]; j < a[0].frompos[i]+a[0].defpos[i]; j++ {
				if a[0].pos[j] <= 0 {
					break
				}
				Norm += a[0].data[j] * a[0].data[j]
			}
			Norm = math.Sqrt(Norm)
			if Norm > MaxNorm {
				MaxNorm = Norm
			}
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			Norm = 0
			for j = 1; j <= a[0].cols; j++ {
				val = femMatGet(a, i, j)
				Norm += val * val
			}
			Norm = math.Sqrt(Norm)
			if Norm > MaxNorm {
				MaxNorm = Norm
			}
		}
	}
	return MaxNorm
}

func femVecNormBig(a []tVector) float64 {

	var Norm float64
	var val float64
	var i int32
	Norm = 0
	if a[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			Norm += a[0].data[i] * a[0].data[i]
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			val = femVecGet(a, i)
			Norm += val * val
		}
	}
	return math.Sqrt(Norm)
}

func femVecNorm(a []tVector) float64 {

	var Norm float64
	var val float64
	var i int32
	Norm = 0
	if a[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			Norm += a[0].data[i] * a[0].data[i]
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			val = femVecGet(a, i)
			Norm += val * val
		}
	}
	return math.Sqrt(Norm)
}

func femVecAddVec(orig []tVector, mult float64, addt []tVector) int32 {

	var i int32
	if orig[0].rows != addt[0].rows {
		return -9
	}
	if orig[0].type_ == 0 && addt[0].type_ == 0 {
		for i = 0; i < orig[0].len_; i++ {
			orig[0].data[i] += mult * addt[0].data[i]
		}
	} else {
		for i = 1; i <= orig[0].len_; i++ {
			femVecPutAdd(orig, i, mult*femVecGet(addt, i), 1)
		}
	}
	return 0
}

func femMatInv(a []tMatrix) int32 {

	var m int32
	var n int32
	var i int32
	var j int32
	var k int32
	var f float64
	var f2 float64
	var val float64
	var f1 tVector
	n = a[0].cols
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&f1))[:])
	if femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&f1))[:], 0, n, n) != 0 {
		return -4
	}
	m = n - 1
	val = femMatGet(a, 1, 1)
	femMatPutAdd(a, 1, 1, 1/val, 0)
	for i = 1; i <= m; i++ {
		for j = 1; j <= i; j++ {
			f = 0
			for k = 1; k <= i; k++ {
				f += femMatGet(a, j, k) * femMatGet(a, k, i+1)
			}
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&f1))[:], j, -f, 0)
		}
		f2 = femMatGet(a, i+1, i+1)
		for j = 1; j <= i; j++ {
			f2 += femMatGet(a, j, i+1) * femVecGet((*[1000000]tVector)(unsafe.Pointer(&f1))[:], j)
		}
		if math.Abs(f2/femMatGet(a, i+1, i+1)) < 1e-07 {
			return -3
		}
		f2 = 1 / f2
		femMatPutAdd(a, i+1, i+1, f2, 0)
		for j = 1; j <= i; j++ {
			for k = 1; k <= i; k++ {
				femMatPutAdd(a, j, k, femVecGet((*[1000000]tVector)(unsafe.Pointer(&f1))[:], j)*femVecGet((*[1000000]tVector)(unsafe.Pointer(&f1))[:], k)*f2+femMatGet(a, j, k), 0)
			}
		}
		for j = 1; j <= i; j++ {
			femMatPutAdd(a, j, i+1, femVecGet((*[1000000]tVector)(unsafe.Pointer(&f1))[:], j)*f2, 0)
			femMatPutAdd(a, i+1, j, femMatGet(a, j, i+1), 0)
		}
	}
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&f1))[:])
	return 0
}

func femLUdecomp(a []tMatrix, index []tVector) int32 {

	var rv int32
	var i int32
	var j int32
	var k int32
	var imax int32
	var n int32
	var big float64
	var dum float64
	var sum float64
	var temp float64
	var vv tVector
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&vv))[:])
	if (func() int32 {
		n = a[0].rows
		return n
	}()) <= 0 {
		return -9
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&vv))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	for i = 1; i <= n; i++ {
		big = 0
		for j = 1; j <= n; j++ {
			if (func() float64 {
				temp = math.Abs(femMatGet(a, i, j))
				return temp
			}()) > big {
				big = temp
			}
		}
		if big == 0 {

			return -3
		}
		femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&vv))[:], i, 1/big, 0)
	}
	for j = 1; j <= n; j++ {
		for i = 1; i < j; i++ {
			sum = femMatGet(a, i, j)
			for k = 1; k < i; k++ {
				sum -= femMatGet(a, i, k) * femMatGet(a, k, j)
			}
			femMatPutAdd(a, i, j, sum, 0)
		}
		big = 0
		for i = j; i <= n; i++ {
			sum = femMatGet(a, i, j)
			for k = 1; k < j; k++ {
				sum -= femMatGet(a, i, k) * femMatGet(a, k, j)
			}
			femMatPutAdd(a, i, j, sum, 0)
			if (func() float64 {
				dum = femVecGet((*[1000000]tVector)(unsafe.Pointer(&vv))[:], i) * math.Abs(sum)
				return dum
			}()) >= big {
				big = dum
				imax = i
			}
		}
		if j != imax {
			for k = 1; k <= n; k++ {
				dum = femMatGet(a, imax, k)
				femMatPutAdd(a, imax, k, femMatGet(a, j, k), 0)
				femMatPutAdd(a, j, k, dum, 0)
			}
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&vv))[:], imax, femVecGet((*[1000000]tVector)(unsafe.Pointer(&vv))[:], j), 0)
		}
		femVecPutAdd(index, j, float64(imax), 0)
		if femMatGet(a, j, j) == 0 {
			femMatPutAdd(a, j, j, 1e-20, 0)
		}
		if j != n {
			dum = 1 / femMatGet(a, j, j)
			for i = j + 1; i <= n; i++ {
				femMatPutAdd(a, i, j, dum*femMatGet(a, i, j), 0)
			}
		}
	}
memFree:
	;
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&vv))[:])
	return rv
}

func femLUback(a []tMatrix, index []tVector, b []tVector) int32 {

	var rv int32
	var i int32
	var ii int32
	var ip int32
	var j int32
	var n int32
	var sum float64
	ii = 0
	if (func() int32 {
		n = a[0].rows
		return n
	}()) <= 0 {
		return -9
	}
	for i = 1; i <= n; i++ {
		ip = int32(femVecGet(index, i))
		sum = femVecGet(b, ip)
		femVecPutAdd(b, ip, femVecGet(b, i), 0)
		if ii != 0 {

			for j = ii; j <= i-1; j++ {
				sum -= femMatGet(a, i, j) * femVecGet(b, j)
			}
		} else {
			if sum != 0 {
				ii = i
			}
		}
		femVecPutAdd(b, i, sum, 0)
	}
	for i = n; i >= 1; i-- {
		sum = femVecGet(b, i)
		for j = i + 1; j <= n; j++ {
			sum -= femMatGet(a, i, j) * femVecGet(b, j)
		}
		femVecPutAdd(b, i, sum/femMatGet(a, i, i), 0)
	}
	return rv
}

func femLUinverse(a []tMatrix) int32 {

	var rv int32
	var i int32
	var j int32
	var n int32
	var col tVector
	var index tVector
	var b tMatrix
	if (func() int32 {
		n = a[0].rows
		return n
	}()) <= 0 {
		return -9
	}
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&col))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&index))[:])
	femMatNull((*[1000000]tMatrix)(unsafe.Pointer(&b))[:])
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&col))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&index))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&b))[:], 0, n, n, 0, nil)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femLUdecomp(a, (*[1000000]tVector)(unsafe.Pointer(&index))[:])
		return rv
	}()) != 0 {
		goto memFree
	}
	for j = 1; j <= n; j++ {
		for i = 1; i <= n; i++ {
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&col))[:], i, 0, 0)
		}
		femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&col))[:], j, 1, 0)
		if (func() int32 {
			rv = femLUback(a, (*[1000000]tVector)(unsafe.Pointer(&index))[:], (*[1000000]tVector)(unsafe.Pointer(&col))[:])
			return rv
		}()) != 0 {
			goto memFree
		}
		for i = 1; i <= n; i++ {
			femMatPutAdd((*[1000000]tMatrix)(unsafe.Pointer(&b))[:], i, j, femVecGet((*[1000000]tVector)(unsafe.Pointer(&col))[:], i), 0)
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			femMatPutAdd(a, i, j, femMatGet((*[1000000]tMatrix)(unsafe.Pointer(&b))[:], i, j), 0)
		}
	}
memFree:
	;
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&col))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&index))[:])
	femMatFree((*[1000000]tMatrix)(unsafe.Pointer(&b))[:])
	return rv
}

func femVecSwitch(a []tVector, b []tVector) int32 {

	var val float64
	var i int32
	if a[0].rows != b[0].rows || a[0].type_ != 0 || b[0].type_ != 0 {
		return -9
	}
	for i = 0; i < a[0].len_; i++ {
		val = a[0].data[i]
		a[0].data[i] = b[0].data[i]
		b[0].data[i] = val
	}
	return 0
}

func femVecCloneDiff(orig []tVector, clone []tVector) int32 {

	var i int32
	var len_ int32
	if orig[0].type_ != 0 || clone[0].type_ != 0 {
		return -5
	}
	if clone[0].rows < 1 || orig[0].rows < 1 {
		return -9
	}
	if orig[0].rows > clone[0].rows {
		len_ = clone[0].rows
	} else {
		len_ = orig[0].rows
	}
	for i = 0; i < len_; i++ {
		clone[0].data[i] = orig[0].data[i]
	}
	return 0
}

func femMatCloneDiffToSame(orig []tMatrix, clone []tMatrix) int32 {

	var i int32
	var j int32
	var k int32
	var ko int32
	var kc int32
	if orig[0].type_ != 1 || clone[0].type_ != 1 {
		return -5
	}
	if orig[0].rows > clone[0].rows || orig[0].rows < 1 {
		return -9
	}
	if orig[0].cols > clone[0].cols || orig[0].cols < 1 {
		return -9
	}
	for i = 0; i < orig[0].rows; i++ {
		k = 0
		for j = orig[0].frompos[i]; j < orig[0].frompos[i]+orig[0].defpos[i]; j++ {
			ko = k + orig[0].frompos[i]
			kc = k + clone[0].frompos[i]
			clone[0].data[kc] = orig[0].data[ko]
			k++
		}
	}
	return 0
}

func femMatCloneDiffToEmpty(orig []tMatrix, clone []tMatrix) int32 {

	var i int32
	var j int32
	if orig[0].type_ != 1 || clone[0].type_ != 1 {
		return -5
	}
	if orig[0].rows > clone[0].rows || orig[0].rows < 1 {
		return -9
	}
	if orig[0].cols > clone[0].cols || orig[0].cols < 1 {
		return -9
	}
	for i = 0; i < orig[0].rows; i++ {
		for j = orig[0].frompos[i]; j < orig[0].frompos[i]+orig[0].defpos[i]; j++ {
			femMatPutAdd(clone, i+1, orig[0].pos[j], orig[0].data[j], 0)
		}
	}
	return 0
}

func eqsCompResid(a []tMatrix, x []tVector, b []tVector, r []tVector) int32 {

	var i int32
	var j int32
	if a[0].type_ == 1 {
		for i = 0; i < a[0].rows; i++ {
			r[0].data[i] = 0 - b[0].data[i]
			for j = a[0].frompos[i]; j < a[0].frompos[i]+a[0].defpos[i]; j++ {
				if a[0].pos[j] <= 0 {
					break
				}
				r[0].data[i] += a[0].data[j] * x[0].data[a[0].pos[j]-1]
			}
		}
	} else {
		for i = 1; i <= a[0].rows; i++ {
			femVecPutAdd(r, i, 0-femVecGet(b, i), 0)
			for j = 1; j < a[0].cols; j++ {
				femVecPutAdd(r, i, femMatGet(a, i, j)*femVecGet(x, j), 1)
			}
		}
	}
	return 0
}

func femEqsCGwJ(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {

	var M tVector
	var r tVector
	var z tVector
	var p tVector
	var q tVector
	var ro float64
	var alpha float64
	var beta float64
	var roro float64
	var n int32
	var i int32
	var j int32
	var rv int32
	var converged int32
	var normRes float64
	var normX float64
	var normA float64
	var normB float64
	if a[0].cols != x[0].rows || x[0].rows != b[0].rows {
		return -9
	}
	n = a[0].rows
	normA = femMatNormBig(a)
	normB = femVecNormBig(b)
	if normB <= 0 {
		femVecSetZeroBig(x)
		return 0
	}

	femVecNull((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&z))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&q))[:])
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&M))[:], 0, n, n)
		return rv
	}()) != 0 {

		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&r))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&z))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&p))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&q))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	{

		for i = 1; i <= n; i++ {
			M.data[i-1] = femMatGet(a, i, i)
			if math.Abs(M.data[i-1]) < 1e-07 {
				rv = -13
				goto memFree
			}
		}
	}

	femMatVecMultBig(a, x, (*[1000000]tVector)(unsafe.Pointer(&r))[:])
	for i = 0; i < n; i++ {
		r.data[i] = b[0].data[i] - r.data[i]
	}
	{

		for i = 1; i <= maxIt; i++ {
			{

				for j = 0; j < n; j++ {
					z.data[j] = r.data[j] / M.data[j]
				}
			}
			ro = femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&r))[:], (*[1000000]tVector)(unsafe.Pointer(&z))[:])
			if i == 1 {
				for j = 0; j < n; j++ {
					p.data[j] = z.data[j]
				}
			} else {
				beta = ro / roro
				for j = 0; j < n; j++ {
					p.data[j] = z.data[j] + beta*p.data[j]
				}
			}
			femMatVecMultBig(a, (*[1000000]tVector)(unsafe.Pointer(&p))[:], (*[1000000]tVector)(unsafe.Pointer(&q))[:])
			alpha = ro / femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&p))[:], (*[1000000]tVector)(unsafe.Pointer(&q))[:])
			for j = 0; j < n; j++ {
				x[0].data[j] = x[0].data[j] + alpha*p.data[j]
				r.data[j] = r.data[j] - alpha*q.data[j]
			}

			normRes = femVecNormBig((*[1000000]tVector)(unsafe.Pointer(&r))[:])
			normX = femVecNormBig(x)
			if normRes <= eps*(normA*normX+normB) {

				converged = 1
				break
			}
			roro = ro
		}
	}
	if converged != 1 {

		rv = -1
	}
memFree:
	;

	femVecFree((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&z))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&q))[:])
	return rv
}

func femEqsBiCCSwJ(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {

	var M tVector
	var r tVector
	var rr tVector
	var p tVector
	var pp tVector
	var s tVector
	var ss tVector
	var t tVector
	var v tVector
	var ro float64
	var beta float64
	var roro float64
	var alpha float64
	var omega float64
	var i int32
	var j int32

	var n int32
	var converged int32

	var res tVector

	var normRes float64
	var normX float64
	var normA float64
	var normB float64
	var rv int32
	n = a[0].rows
	normA = femMatNormBig(a)
	normX = femVecNormBig(x)
	normB = femVecNormBig(b)

	femVecNull((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&rr))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&pp))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&s))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&ss))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&t))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&v))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&res))[:])
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&M))[:], 0, n, n)
		return rv
	}()) != 0 {

		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&r))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&rr))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&p))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&pp))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&s))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&ss))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&t))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&v))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&res))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	{

		for i = 1; i <= n; i++ {
			M.data[i-1] = femMatGet(a, i, i)
			if math.Abs(M.data[i-1]) < 1e-07 {
				rv = -13
				goto memFree
			}
		}
	}

	femMatVecMultBig(a, x, (*[1000000]tVector)(unsafe.Pointer(&r))[:])
	for i = 0; i < n; i++ {
		r.data[i] = b[0].data[i] - r.data[i]
		rr.data[i] = r.data[i]
	}
	if femVecNormBig((*[1000000]tVector)(unsafe.Pointer(&r))[:]) <= 1e-07 {

		converged = 1
		goto memFree
	}
	{

		for i = 1; i <= maxIt; i++ {
			ro = femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&rr))[:], (*[1000000]tVector)(unsafe.Pointer(&r))[:])
			if math.Abs(ro) <= 0 {
				goto memFree
			}
			if i == 1 {

				for j = 0; j < n; j++ {
					p.data[j] = r.data[j]
				}
			} else {

				beta = ro / roro * (alpha / omega)
				for j = 0; j < n; j++ {
					p.data[j] = r.data[j] + beta*(p.data[j]-omega*v.data[j])
				}
			}
			{

				for j = 0; j < n; j++ {
					pp.data[j] = p.data[j] / M.data[j]
				}
			}
			femMatVecMultBig(a, (*[1000000]tVector)(unsafe.Pointer(&pp))[:], (*[1000000]tVector)(unsafe.Pointer(&v))[:])
			alpha = ro / femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&rr))[:], (*[1000000]tVector)(unsafe.Pointer(&v))[:])
			for j = 0; j < n; j++ {
				s.data[j] = r.data[j] - alpha*v.data[j]
			}
			if femVecNormBig((*[1000000]tVector)(unsafe.Pointer(&s))[:]) <= 1e-07 {
				{

					for j = 0; j < n; j++ {
						x[0].data[j] += alpha * pp.data[j]
					}
				}
				converged = 1
				break
			}
			{

				for j = 0; j < n; j++ {
					ss.data[j] = s.data[j] / M.data[j]
				}
			}
			femMatVecMultBig(a, (*[1000000]tVector)(unsafe.Pointer(&ss))[:], (*[1000000]tVector)(unsafe.Pointer(&t))[:])
			omega = femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&t))[:], (*[1000000]tVector)(unsafe.Pointer(&s))[:]) / femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&t))[:], (*[1000000]tVector)(unsafe.Pointer(&t))[:])
			for j = 0; j < n; j++ {
				x[0].data[j] += alpha*pp.data[j] + omega*ss.data[j]
				r.data[j] = s.data[j] - omega*t.data[j]
			}
			roro = ro

			eqsCompResid(a, b, x, (*[1000000]tVector)(unsafe.Pointer(&res))[:])
			normRes = femVecNormBig((*[1000000]tVector)(unsafe.Pointer(&res))[:])
			normX = femVecNormBig(x)
			if normRes < eps*(normA*normX+normB) {
				converged = 1
				break
			}
		}
	}
	if converged != 1 {

		{
		}
	}
memFree:
	;

	femVecFree((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&rr))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&pp))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&s))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&ss))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&t))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&v))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&res))[:])
	return 0
}

func femEqsLU(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {

	var rv int32
	var n int32
	var indx tVector
	if (func() int32 {
		n = a[0].rows
		return n
	}()) <= 0 {
		return -9
	}
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&indx))[:])
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&indx))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femLUdecomp(a, (*[1000000]tVector)(unsafe.Pointer(&indx))[:])
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femLUback(a, (*[1000000]tVector)(unsafe.Pointer(&indx))[:], b)
		return rv
	}()) != 0 {
		goto memFree
	}
memFree:
	;
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&indx))[:])
	return rv
}

func femEqsPCGwJ(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {

	var rv int32
	var converged int32
	var nui float64
	var dei float64
	var lambda float64
	var alpha float64

	var normRes float64
	var normX float64
	var normA float64
	var normB float64
	var p tVector
	var r tVector
	var d tVector

	var M tVector

	var ap tVector

	var n int32
	var i int32
	var j int32
	if a[0].rows != x[0].rows || x[0].rows != b[0].rows {
		return -9
	}
	n = a[0].rows
	normA = femMatNormBig(a)
	normB = femVecNormBig(b)
	normX = femVecNormBig(x)
	if normB <= 0 {
		femVecSetZeroBig(x)
		return 0
	}

	femVecNull((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&d))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&ap))[:])
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&p))[:], 0, n, n)
		return rv
	}()) != 0 {

		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&r))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&d))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&M))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&ap))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	{

		for i = 1; i <= n; i++ {
			M.data[i-1] = femMatGet(a, i, i)
			if math.Abs(M.data[i-1]) < 1e-07 {
				rv = -13
				goto memFree
			}
		}
	}

	femMatVecMultBig(a, x, (*[1000000]tVector)(unsafe.Pointer(&r))[:])
	for i = 0; i < n; i++ {
		r.data[i] = b[0].data[i] - r.data[i]
	}
	{

		for j = 0; j < n; j++ {
			d.data[j] = r.data[j] / M.data[j]
			p.data[j] = d.data[j]
		}
	}
	for i = 1; i <= maxIt; i++ {

		femMatVecMultBig(a, (*[1000000]tVector)(unsafe.Pointer(&p))[:], (*[1000000]tVector)(unsafe.Pointer(&ap))[:])
		nui = femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&r))[:], (*[1000000]tVector)(unsafe.Pointer(&d))[:])
		dei = femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&p))[:], (*[1000000]tVector)(unsafe.Pointer(&ap))[:])
		lambda = nui / dei
		for j = 0; j < n; j++ {
			x[0].data[j] += lambda * p.data[j]
			r.data[j] = r.data[j] - lambda*ap.data[j]
			d.data[j] = r.data[j] / M.data[j]
		}
		normRes = femVecNormBig((*[1000000]tVector)(unsafe.Pointer(&r))[:])
		normX = femVecNormBig(x)
		nui = femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&r))[:], (*[1000000]tVector)(unsafe.Pointer(&d))[:])
		dei = femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&p))[:], (*[1000000]tVector)(unsafe.Pointer(&ap))[:])
		if normRes < eps*(normA*normX+normB) {

			converged = 1
			break
		}
		alpha = nui / dei
		for j = 0; j < n; j++ {
			p.data[j] = d.data[j] + alpha*p.data[j]
		}
	}

	femVecPrn(x, []byte("X\x00"))
memFree:
	;
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&d))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&ap))[:])
	return rv
}

func femMatCholFact(a []tMatrix, C []tVector) int32 {

	var rv int32
	var sum float64
	var n int32
	var i int32
	var j int32
	var k int32
	var have_C int32
	n = a[0].rows
	if len(C) != 0 {
		if C[0].rows != a[0].rows {
			return -3
		} else {
			have_C = 1
		}
	}
	if have_C == 0 {
		femVecNull(C)
		if femVecAlloc(C, 0, n, n) != 0 {
			goto memFree
		}
	}
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			sum = femMatGet(a, i, j)
			for k = i - 1; k >= 1; k-- {
				sum -= femMatGet(a, i, k) * femMatGet(a, j, k)
			}
			if i == j {
				if sum <= 0 {
					rv = -3
					goto memFree
				}
				femVecPutAdd(C, i, math.Sqrt(sum), 0)
			} else {
				femMatPutAdd(a, j, i, sum/femVecGet(C, i), 0)
			}
		}
	}
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			if i != j {
				femMatPutAdd(a, i, j, femMatGet(a, j, i), 0)
				femMatPutAdd(a, j, i, 0, 0)
			} else {
				femMatPutAdd(a, j, i, femVecGet(C, i), 0)
			}
		}
	}
	femVecPrn(C, []byte("C\x00"))
memFree:
	;
	if have_C == 0 {

		femVecFree(C)
	}
	return rv
}

func femEqsChol(a []tMatrix, b []tVector, x []tVector) int32 {

	var rv int32
	var sum float64
	var n int32
	var i int32
	var j int32
	var k int32
	var C tVector
	n = a[0].rows
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&C))[:])
	if femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&C))[:], 0, n, n) != 0 {
		goto memFree
	}
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			sum = femMatGet(a, i, j)
			for k = i - 1; k >= 1; k-- {
				sum -= femMatGet(a, i, k) * femMatGet(a, j, k)
			}
			if i == j {
				if sum <= 0 {
					rv = -3
					goto memFree
				}
				femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&C))[:], i, math.Sqrt(sum), 0)
			} else {
				femMatPutAdd(a, j, i, sum/femVecGet((*[1000000]tVector)(unsafe.Pointer(&C))[:], i), 0)
			}
		}
	}
	{

		for i = 1; i <= n; i++ {
			sum = femVecGet(b, i)
			for k = i - 1; k >= 1; k-- {
				sum -= femMatGet(a, i, k) * femVecGet(x, k)
			}
			femVecPutAdd(x, i, sum/femVecGet((*[1000000]tVector)(unsafe.Pointer(&C))[:], i), 0)
		}
	}
	for i = n; i >= 1; i-- {
		sum = femVecGet(x, i)
		for k = i + 1; k <= n; k++ {
			sum -= femMatGet(a, k, i) * femVecGet(x, k)
		}
		femVecPutAdd(x, i, sum/femVecGet((*[1000000]tVector)(unsafe.Pointer(&C))[:], i), 0)
	}
memFree:
	;

	femVecFree((*[1000000]tVector)(unsafe.Pointer(&C))[:])
	return rv
}

func femMatJacRotate(a []tMatrix, i int32, j int32, k int32, l int32, g float64, h float64, s float64, tau float64) {

	g = femMatGet(a, i, j)
	h = femMatGet(a, k, l)
	femMatPutAdd(a, i, j, g-s*(h+g*tau), 0)
	femMatPutAdd(a, k, l, h+s*(g-h*tau), 0)
}

func femMatEigenJacobi(a []tMatrix, d []tVector, v []tMatrix, nrot []int32) int32 {

	var iters int32 = 100
	var i int32
	var iq int32
	var ip int32
	var j int32
	var n int32
	var sm float64
	var tresh float64
	var g float64
	var h float64
	var t float64
	var c float64
	var theta float64
	var s float64
	var tau float64
	var checkp float64
	var checkq float64
	var checkh float64
	var b tVector
	var z tVector
	nrot[0] = 0
	n = a[0].rows
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&b))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&z))[:])
	femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&b))[:], 0, n, n)
	femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&z))[:], 0, n, n)
	for i = 1; i <= n; i++ {
		femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&b))[:], i, femMatGet(a, i, i), 0)
		femVecPutAdd(d, i, femMatGet(a, i, i), 0)
		femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&z))[:], i, 0, 0)
		femMatPutAdd(v, i, i, 1, 0)
	}
	for i = 1; i <= iters; i++ {
		sm = 0
		for ip = 0; ip <= n-1; ip++ {
			for iq = ip + 1; iq <= n; iq++ {
				sm += math.Abs(femMatGet(a, ip, iq))
			}
		}
		if sm <= 1e-07 {

			femVecFree((*[1000000]tVector)(unsafe.Pointer(&b))[:])
			femVecFree((*[1000000]tVector)(unsafe.Pointer(&z))[:])
			return 0
		}
		if i < 4 {
			tresh = 0.2 * sm / float64(n*n)
		} else {
			tresh = 0
		}
		for ip = 1; ip <= n-1; ip++ {
			for iq = ip + 1; iq <= n; iq++ {
				g = 100 * math.Abs(femMatGet(a, ip, iq))
				checkp = math.Abs(g*math.Abs(femVecGet(d, ip)) - math.Abs(femVecGet(d, ip)))
				checkq = math.Abs(g*math.Abs(femVecGet(d, iq)) - math.Abs(femVecGet(d, iq)))
				if i > 4 && checkp <= 1e-07 && checkq <= 1e-07 {

					femMatPutAdd(a, ip, iq, 0, 0)
				} else {

					h = femVecGet(d, iq) - femVecGet(d, ip)
					checkh = math.Abs(math.Abs(h) + g - math.Abs(h))
					if checkh < 1e-07 {
						if h != 0 {
							t = femMatGet(a, ip, iq) / h
						} else {
							t = 0
						}
					} else {
						theta = 0.5 * h / femMatGet(a, ip, iq)
						t = 1 / (math.Abs(theta) + math.Sqrt(1+math.Pow(theta, 2)))
						if theta < 0 {
							t = -1 * t
						}
					}
					c = 1 / math.Sqrt(1+math.Pow(t, 2))
					s = t * c
					tau = s / (1 + c)
					h = t * femMatGet(a, ip, iq)
					femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&z))[:], ip, femVecGet((*[1000000]tVector)(unsafe.Pointer(&z))[:], ip)-h, 0)
					femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&z))[:], iq, femVecGet((*[1000000]tVector)(unsafe.Pointer(&z))[:], iq)+h, 0)
					femVecPutAdd(d, ip, femVecGet(d, ip)-h, 0)
					femVecPutAdd(d, iq, femVecGet(d, iq)+h, 0)
					femMatPutAdd(a, ip, iq, 0, 0)
					for j = 1; j <= ip-1; j++ {
						femMatJacRotate(a, j, ip, j, iq, g, h, s, tau)
					}
					for j = ip + 1; j <= iq-1; j++ {
						femMatJacRotate(a, ip, j, j, iq, g, h, s, tau)
					}
					for j = iq + 1; j <= n; j++ {
						femMatJacRotate(a, ip, j, iq, j, g, h, s, tau)
					}
					for j = 1; j <= n; j++ {
						femMatJacRotate(v, j, ip, j, iq, g, h, s, tau)
					}
					nrot[0] = nrot[0] + 1
				}
			}
		}
		for ip = 1; ip <= n; ip++ {
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&b))[:], ip, femVecGet((*[1000000]tVector)(unsafe.Pointer(&z))[:], ip), 1)
			femVecPutAdd(d, ip, femVecGet((*[1000000]tVector)(unsafe.Pointer(&b))[:], ip), 0)
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&z))[:], ip, 0, 0)
		}
	}
	return -1
}

func femEqsCGwSSOR(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {

	var M tVector
	var r tVector
	var z tVector
	var zz tVector
	var p tVector
	var q tVector
	var ro float64
	var alpha float64
	var beta float64
	var roro float64
	var n int32
	var i int32
	var ii int32
	var j int32
	var ipos int32
	var rv int32
	var converged int32
	var normRes float64
	var normX float64
	var normA float64
	var normB float64
	var val float64
	if a[0].cols != x[0].rows || x[0].rows != b[0].rows {
		return -9
	}
	n = a[0].rows
	normA = femMatNormBig(a)
	normB = femVecNormBig(b)
	if normB <= 0 {
		femVecSetZeroBig(x)
		return 0
	}

	femVecNull((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&z))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&zz))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&q))[:])
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&M))[:], 0, n, n)
		return rv
	}()) != 0 {

		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&r))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&z))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&zz))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&p))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int32 {
		rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&q))[:], 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	{

		for i = 1; i <= n; i++ {
			val = femMatGet(a, i, i)
			if math.Abs(val) < 1e-07 {
				rv = -13
				goto memFree
			}

			M.data[i-1] = 1 / val
		}
	}

	femMatVecMultBig(a, x, (*[1000000]tVector)(unsafe.Pointer(&r))[:])
	for i = 0; i < n; i++ {
		r.data[i] = b[0].data[i] - r.data[i]
	}
	{

		for i = 1; i <= maxIt; i++ {
			if a[0].type_ != 1 {
				{

					for ii = 1; ii <= n; ii++ {
						val = 0
						for j = 1; j < ii; j++ {
							val += femMatGet(a, ii, j) * femVecGet((*[1000000]tVector)(unsafe.Pointer(&zz))[:], j)
						}
						femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&zz))[:], ii, femVecGet((*[1000000]tVector)(unsafe.Pointer(&M))[:], ii)*(femVecGet((*[1000000]tVector)(unsafe.Pointer(&r))[:], ii)-val), 0)
					}
				}
				for ii = n; ii >= 1; ii-- {
					val = 0
					for j = ii + 1; j <= n; j++ {
						val += femMatGet(a, ii, j) * femVecGet((*[1000000]tVector)(unsafe.Pointer(&z))[:], j)
					}
					femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&z))[:], ii, femVecGet((*[1000000]tVector)(unsafe.Pointer(&zz))[:], ii)-femVecGet((*[1000000]tVector)(unsafe.Pointer(&M))[:], ii)*val, 0)
				}
			} else {
				{

					for ii = 1; ii <= n; ii++ {
						val = 0
						for j = a[0].frompos[ii-1]; j < a[0].frompos[ii-1]+a[0].defpos[ii-1]; j++ {
							ipos = a[0].pos[j]
							if ipos >= ii || ipos < 1 {
								continue
							}
							val += a[0].data[j] * zz.data[ipos-1]
						}
						femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&zz))[:], ii, femVecGet((*[1000000]tVector)(unsafe.Pointer(&M))[:], ii)*(femVecGet((*[1000000]tVector)(unsafe.Pointer(&r))[:], ii)-val), 0)
					}
				}
				for ii = n; ii >= 1; ii-- {
					val = 0
					for j = a[0].frompos[ii-1]; j < a[0].frompos[ii-1]+a[0].defpos[ii-1]; j++ {
						ipos = a[0].pos[j]
						if ipos > ii {
							val += a[0].data[j] * z.data[ipos-1]
						}
					}
					femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&z))[:], ii, femVecGet((*[1000000]tVector)(unsafe.Pointer(&zz))[:], ii)-femVecGet((*[1000000]tVector)(unsafe.Pointer(&M))[:], ii)*val, 0)
				}
			}

			ro = femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&r))[:], (*[1000000]tVector)(unsafe.Pointer(&z))[:])
			if i == 1 {
				for j = 0; j < n; j++ {
					p.data[j] = z.data[j]
				}
			} else {
				beta = ro / roro
				for j = 0; j < n; j++ {
					p.data[j] = z.data[j] + beta*p.data[j]
				}
			}
			femMatVecMultBig(a, (*[1000000]tVector)(unsafe.Pointer(&p))[:], (*[1000000]tVector)(unsafe.Pointer(&q))[:])
			alpha = ro / femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&p))[:], (*[1000000]tVector)(unsafe.Pointer(&q))[:])
			for j = 0; j < n; j++ {
				x[0].data[j] = x[0].data[j] + alpha*p.data[j]
				r.data[j] = r.data[j] - alpha*q.data[j]
			}

			normRes = femVecNormBig((*[1000000]tVector)(unsafe.Pointer(&r))[:])
			normX = femVecNormBig(x)
			if normRes <= eps*(normA*normX+normB) {

				converged = 1
				break
			}
			roro = ro
		}
	}
	if converged != 1 {

		rv = -1
	}
memFree:
	;

	femVecFree((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&z))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&zz))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&q))[:])
	return rv
}

func c4goUnsafeConvert_float64(c4go_name *float64) []float64 {
	return (*[1000000]float64)(unsafe.Pointer(c4go_name))[:]
}

func c4goUnsafeConvert_int32(c4go_name *int32) []int32 {
	return (*[1000000]int32)(unsafe.Pointer(c4go_name))[:]
}
