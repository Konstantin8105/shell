//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"github.com/Konstantin8105/c4go/noarch"
)

// msgout - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:29
//
//   File name: fem_mem.c
//   Date:      2003/04/07 10:16
//   Author:    Jiri Brozovsky
//
//   Copyright (C) 2003 Jiri Brozovsky
//
//   This program is free software; you can redistribute it and/or
//   modify it under the terms of the GNU General Public License as
//   published by the Free Software Foundation; either version 2 of the
//   License, or (at your option) any later version.
//
//   This program is distributed in the hope that it will be useful, but
//   WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   in a file called COPYING along with this program; if not, write to
//   the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
//   02139, USA.
//
//	FEM Solver - memory handling
//  $Id: fem_mem.c,v 1.5 2004/07/06 21:03:44 jirka Exp $
//
var msgout = os.Stdout // *fmt.File

// femIntAlloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:37
func femIntAlloc(length int32) (c4goDefaultReturn []int32) {
	// 1D fields ----------------------------------------------------------
	// allocates and returns 1D int field  (NULL if failed)
	// * @param length length of field
	// * @returns field (or NULL)
	//
	// 	var field []int32
	// 	var i int32
	// 	if length < 1 {
	// 		return nil
	// 	}
	// 	if len((func() []int32 {
	// 		field = (*[1000000]int32)(unsafe.Pointer(uintptr(func() int64 {
	// 			c4go_temp_name := make(, uint32(length)*uint32(1))
	// 			return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	// 		}())))[:]
	// 		return field
	// 	}())) == 0 {
	// 		return nil
	// 	} else {
	// 		for i = 0; i < length; i++ {
	// 			field[i] = 0
	// 		}
	// 		return field
	// 	}
	// 	return
	return make([]int32, length)
}

// femIntFree - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:59
// func femIntFree(field []int32) int32 {
// 	_ = field
// 	// removes memory from int field
// 	// * @param field  field to be freed
// 	// * @returns state value
// 	//
// 	field = nil
// 	return 0
// }

// femDblAlloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:71
func femDblAlloc(length int32) (c4goDefaultReturn []float64) {
	// allocates and returns 1D double field  (NULL if failed)
	// * @param length length of field
	// * @returns field (or NULL)
	//
	// 	var field []float64
	// 	var i int32
	// 	if length < 1 {
	// 		return nil
	// 	}
	// 	if len((func() []float64 {
	// 		field = (*[1000000]float64)(unsafe.Pointer(uintptr(func() int64 {
	// 			c4go_temp_name := make(, uint32(length)*uint32(1))
	// 			return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
	// 		}())))[:]
	// 		return field
	// 	}())) == 0 {
	// 		return nil
	// 	} else {
	// 		for i = 0; i < length; i++ {
	// 			field[i] = 0
	// 		}
	// 		return field
	// 	}

	return make([]float64, length, length)
}

// femDblFree - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:93
// func femDblFree(field []float64) int32 {
// 	_ = field
// 	// removes memory from dbl field
// 	// * @param field  field to be freed
// 	// * @returns state value
// 	//
// 	field = nil
// 	return 0
// }

// _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_47 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.h:47
//
//   File name: fem_math.h
//   Date:      2003/04/12 12:45
//   Author:    Jiri Brozovsky
//
//   Copyright (C) 2003 Jiri Brozovsky
//
//   This program is free software; you can redistribute it and/or
//   modify it under the terms of the GNU General Public License as
//   published by the Free Software Foundation; either version 2 of the
//   License, or (at your option) any later version.
//
//   This program is distributed in the hope that it will be useful, but
//   WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   in a file called COPYING along with this program; if not, write to
//   the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
//   02139, USA.
//
//	 FEM Software - matrix library - header file
//
//	 $Id: fem_math.h,v 1.21 2005/07/11 17:56:16 jirka Exp $
//
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

// tMatrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.h:47
type tMatrix = _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_47

// _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_60 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.h:60
// 0 = dense; 1 = sparse (rows)
// lenght of "pos" and "data" (if used) fields
// from in "pos" and "data" - sparse only sizeof(frompos) = rows
// number in "pos" and "data" - sparse only
type _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_60 struct {
	type_ int32
	rows  int32
	len_  int32
	pos   []int32
	data  []float64
}

// tVector - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.h:60
type tVector = _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_60

// n_m - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:44
//
//   File name: eshell.c
//   Date:      2010/07/21 17:18
//   Author:    Jiri Brozovsky
//
//   Copyright (C) 2010 VSB-TU of Ostrava
//
//   This program is free software; you can redistribute it and/or
//   modify it under the terms of the GNU General Public License as
//   published by the Free Software Foundation; either version 2 of the
//   License.
//
//   This program is distributed in the hope that it will be useful, but
//   WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//   General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   in a file called COPYING along with this program; if not, write to
//   the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
//   02139, USA.
//
//
//   Axisymetric shell solver. Use: eshell <input >output
//
//   See for details:
//   viz Schneider, Vykutil: Stavba chemickych zarizeni II.a
//       Mikropocitacove aplikace MKP ve statice rotacnich
//       skorepin, ES VUT Brno, Brno, Czechoslovakia, 1986
//
// INPUT DATA:
// number of materials
var n_m int32

// n_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:45
// number of nodes
var n_n int32

// n_e - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:46
// number of elements
var n_e int32

// n_d - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:47
// number of displacements/supports
var n_d int32

// n_f - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:48
// number of loads
var n_f int32

// n_r_inp - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:50
// number of random input data
var n_r_inp int32

// n_r_opt - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:51
// number of optim input data
var n_r_opt int32

// m_E1 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:54
// materials
// E1 (bulk modullus)
var m_E1 []float64

// m_E2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:55
// E2 (bulk modullus)
var m_E2 []float64

// m_G - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:56
// G (shear modullus)
var m_G []float64

// m_nu1 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:57
// nu1 (poisson ratio)
var m_nu1 []float64

// m_nu2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:58
// nu2 (poisson ratio)
var m_nu2 []float64

// m_q - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:59
// volume gravity force
var m_q []float64

// m_vp - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:60
// volume unit  price
var m_vp []float64

// m_t - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:61
// width (if >=0 then ovewrites e_t[] data)
var m_t []float64

// n_x - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:64
// nodes
// x coordinates
var n_x []float64

// n_y - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:65
// y coordinates
var n_y []float64

// e_n1 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:68
//elements
// first nodes <0, n_n-1>
var e_n1 []int32

// e_n2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:69
// second nodes  <0, n_n-1>
var e_n2 []int32

// e_mat - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:70
// material numbers  <0, n_m-1>
var e_mat []int32

// e_t - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:71
// element widths (constatnt on element)
var e_t []float64

// d_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:74
// displacements
// nodes <0, n_n-1>
var d_n []int32

// d_dir - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:75
// orientation w=0, u=1, pho=2, Ez=3, Ex=4, Erot=5
var d_dir []int32

// d_val - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:76
// size of displacement or stiffness
var d_val []float64

// f_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:79
// forces in nodes
// nodes <0, n_n-1>
var f_n []int32

// f_dir - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:80
// orientation Fw=0, Fu=1, Mpho=2
var f_dir []int32

// f_val - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:81
// size of the force
var f_val []float64

// w_top - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:84
// water load:
// water level
var w_top float64

// w_bot - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:85
// bottom of water
var w_bot float64

// w_val - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:86
// volume weight in N/m^3 - negative: <-, positive: ->
var w_val float64

// w_min - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:87
// minimal element number for water load
var w_min int32 = -1

// w_max - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:88
// maximal element number for water load
var w_max int32 = -1

// rand_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:91
// random input data
// type of data (see README.RANDOM)
var rand_type []int32

// rand_pos - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:92
// index of data
var rand_pos []int32

// rand_indx - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:93
// data index - if applicable
var rand_indx []int32

// opt_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:96
// optim input data
// type of data (see README.RANDOM)
var opt_type []int32

// opt_pos - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:97
// index of data
var opt_pos []int32

// opt_indx - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:98
// data index - if applicable
var opt_indx []int32

// opt_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:99
// data for replacing
var opt_data []float64

// fail_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:102
// failure condition data
// type of failure condition
var fail_type int32

// n_fail - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:103
// number of failure condition data
var n_fail int32

// fail_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:104
// failure condition data
var fail_data []float64

// K - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:107
// SOLUTION DATA
var K tMatrix

// u - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:108
var u tVector

// F - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:109
var F tVector

// Ke - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:111
// 6x6
var Ke tMatrix

// D - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:112
// 5x5
var D tMatrix

// B - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:113
// 5x6
var B tMatrix

// Bt - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:114
// 6x5
var Bt tMatrix

// BtD - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:115
// 6x5
var BtD tMatrix

// DB - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:116
// 5x6
var DB tMatrix

// Fe - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:117
// 5
var Fe tVector

// ue - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:118
// 6
var ue tVector

// n_en - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:121
// result helpers data
var n_en int32

// en_num - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:122
var en_num []int32

// en_frm - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:123
var en_frm []int32

// en_pos - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:124
var en_pos []int32

// solution_only - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:127
// program constants
var solution_only int32 = 1

// random_only - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:128
var random_only int32 = 1

// price_only - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:129
var price_only int32 = 1

// write_only - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:130
var write_only int32

// free_input_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:133
func free_input_data() {
	// 	if len(m_E1) != 0 {
	// 		// frees input data
	// 		femDblFree(m_E1)
	// 	}
	// 	if len(m_E2) != 0 {
	// 		femDblFree(m_E2)
	// 	}
	// 	if len(m_G) != 0 {
	// 		femDblFree(m_G)
	// 	}
	// 	if len(m_nu1) != 0 {
	// 		femDblFree(m_nu1)
	// 	}
	// 	if len(m_nu2) != 0 {
	// 		femDblFree(m_nu2)
	// 	}
	// 	if len(m_q) != 0 {
	// 		femDblFree(m_q)
	// 	}
	// 	if len(m_vp) != 0 {
	// 		femDblFree(m_vp)
	// 	}
	// 	if len(m_t) != 0 {
	// 		femDblFree(m_t)
	// 	}
	// 	if len(n_x) != 0 {
	// 		femDblFree(n_x)
	// 	}
	// 	if len(n_y) != 0 {
	// 		femDblFree(n_y)
	// 	}
	// 	if len(e_n1) != 0 {
	// 		femIntFree(e_n1)
	// 	}
	// 	if len(e_n2) != 0 {
	// 		femIntFree(e_n2)
	// 	}
	// 	if len(e_mat) != 0 {
	// 		femIntFree(e_mat)
	// 	}
	// 	if len(e_t) != 0 {
	// 		femDblFree(e_t)
	// 	}
	// 	if len(d_n) != 0 {
	// 		femIntFree(d_n)
	// 	}
	// 	if len(d_dir) != 0 {
	// 		femIntFree(d_dir)
	// 	}
	// 	if len(d_val) != 0 {
	// 		femDblFree(d_val)
	// 	}
	// 	if n_f > 0 {
	// 		if len(f_n) != 0 {
	// 			femIntFree(f_n)
	// 		}
	// 		if len(f_dir) != 0 {
	// 			femIntFree(f_dir)
	// 		}
	// 		if len(f_val) != 0 {
	// 			femDblFree(f_val)
	// 		}
	// 	}
	// 	if n_r_inp > 0 {
	// 		if len(rand_type) != 0 {
	// 			femIntFree(rand_type)
	// 		}
	// 		if len(rand_pos) != 0 {
	// 			femIntFree(rand_pos)
	// 		}
	// 		if len(rand_indx) != 0 {
	// 			femIntFree(rand_indx)
	// 		}
	// 	}
	// 	if n_r_opt > 0 {
	// 		if len(opt_type) != 0 {
	// 			femIntFree(opt_type)
	// 		}
	// 		if len(opt_pos) != 0 {
	// 			femIntFree(opt_pos)
	// 		}
	// 		if len(opt_indx) != 0 {
	// 			femIntFree(opt_indx)
	// 		}
	// 		if len(opt_data) != 0 {
	// 			femDblFree(opt_data)
	// 		}
	// 	}
	// 	if n_en > 0 {
	// 		if len(en_num) != 0 {
	// 			femIntFree(en_num)
	// 		}
	// 		if len(en_frm) != 0 {
	// 			femIntFree(en_frm)
	// 		}
	// 		if len(en_pos) != 0 {
	// 			femIntFree(en_pos)
	// 		}
	// 	}
	// 	if n_fail > 0 {
	// 		if len(fail_data) != 0 {
	// 			femDblFree(fail_data)
	// 		}
	// 	}
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

// check_elem_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:205
func check_elem_data() {
	// first node must be always under the second - it exchanges them
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

// get_enode_fields - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:222
func get_enode_fields() int32 {
	// will prepare element nodes filed for optimised result output
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
	// 	if len((func() []int32 {
	en_pos = femIntAlloc(n_en)
	// 		return en_pos
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
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
	// memFree:
	// 	;
	// 	return -4
}

// read_input_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:276
func read_input_data() { //fw *fmt.File) int32 {
	// reads data from stream
	// * @param fw stream for reading
	//
	var i int32
	n_m = 1
	// 	if fmt.Fscanf(fw, ("%v"), c4goUnsafeConvert_int32(&n_m)) <= 0 {
	// 		goto memFree
	// 	}
	// 	if n_m < 1 {
	// 		fmt.Fprintf(msgout, ("Invalid number of materials!\n"))
	// 		goto memFree
	// 	}
	n_n = 3
	// 	if fmt.Fscanf(fw, ("%v"), c4goUnsafeConvert_int32(&n_n)) <= 0 {
	// 		goto memFree
	// 	}
	// 	if n_n < 2 {
	// 		fmt.Fprintf(msgout, ("Invalid number of nodes!\n"))
	// 		goto memFree
	// 	}
	n_e = 2
	// 	if fmt.Fscanf(fw, ("%v"), c4goUnsafeConvert_int32(&n_e)) <= 0 {
	// 		goto memFree
	// 	}
	// 	if n_e < 2 {
	// 		fmt.Fprintf(msgout, ("Invalid number of elements!\n"))
	// 		goto memFree
	// 	}
	n_d = 3
	// 	if fmt.Fscanf(fw, ("%v"), c4goUnsafeConvert_int32(&n_d)) <= 0 {
	// 		goto memFree
	// 	}
	// 	if n_d < 3 {
	// 		fmt.Fprintf(msgout, ("Invalid number of supports!\n"))
	// 		goto memFree
	// 	}
	n_f = 1
	// 	if fmt.Fscanf(fw, ("%v"), c4goUnsafeConvert_int32(&n_f)) < 0 {
	// 		goto memFree
	// 	}
	// 	if n_f < 0 {
	// 		fmt.Fprintf(msgout, ("Invalid number of forces!\n"))
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	m_E1 = femDblAlloc(n_m)
	// 		return m_E1
	// 	}())) == 0 {
	// 		// data allocations
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	m_E2 = femDblAlloc(n_m)
	// 		return m_E2
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	m_G = femDblAlloc(n_m)
	// 		return m_G
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	m_nu1 = femDblAlloc(n_m)
	// 		return m_nu1
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	m_nu2 = femDblAlloc(n_m)
	// 		return m_nu2
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	m_q = femDblAlloc(n_m)
	// 		return m_q
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	m_vp = femDblAlloc(n_m)
	// 		return m_vp
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	m_t = femDblAlloc(n_m)
	// 		return m_t
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	n_x = femDblAlloc(n_n)
	// 		return n_x
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	n_y = femDblAlloc(n_n)
	// 		return n_y
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	e_n1 = femIntAlloc(n_e)
	// 		return e_n1
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	e_n2 = femIntAlloc(n_e)
	// 		return e_n2
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	e_mat = femIntAlloc(n_e)
	// 		return e_mat
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	e_t = femDblAlloc(n_e)
	// 		return e_t
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	d_n = femIntAlloc(n_d)
	// 		return d_n
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	d_dir = femIntAlloc(n_d)
	// 		return d_dir
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	d_val = femDblAlloc(n_d)
	// 		return d_val
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if n_f > 0 {
	// 		if len((func() []int32 {
	f_n = femIntAlloc(n_f)
	// 			return f_n
	// 		}())) == 0 {
	// 			goto memFree
	// 		}
	// 		if len((func() []int32 {
	f_dir = femIntAlloc(n_f)
	// 			return f_dir
	// 		}())) == 0 {
	// 			goto memFree
	// 		}
	// 		if len((func() []float64 {
	f_val = femDblAlloc(n_f)
	// 			return f_val
	// 		}())) == 0 {
	// 			goto memFree
	// 		}
	// 	}
	// 	if len((func() []int32 {
	en_num = femIntAlloc(n_n)
	// 		return en_num
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	en_frm = femIntAlloc(n_n)
	// 		return en_frm
	// 	}())) == 0 {
	// 		goto memFree
	// 	}

	m_E1[0], m_E2[0], m_G[0], m_nu1[0], m_nu2[0], m_q[0], m_vp[0], m_t[0] = 20e9, 0, 0, 0.2, 0, 25000, 1000, 0

	// 	{
	// 		// reading of data:
	// 		// materials
	for i = 0; i < n_m; i++ {
		// 	 			if fmt.Fscanf(fw, (" %f %f %f %f %f %f %f %f"), m_E1[i:], m_E2[i:], m_G[i:], m_nu1[i:], m_nu2[i:], m_q[i:], m_vp[i:], m_t[i:]) <= 0 {
		// 	 				goto memFree
		// 	 			}
		if m_E1[i] == m_E2[i] || m_E2[i] <= 0 {
			// isotropic
			m_E2[i] = m_E1[i]
			m_nu2[i] = m_nu1[i]
			if m_G[i] <= 0 {
				m_G[i] = m_E1[i] / (2 * (1 + m_nu1[i]))
			}
			// 	 			} else {
			// 	 				if m_E1[i] <= 0 || m_E2[i] <= 0 || m_G[i] <= 0 || m_nu1[i] <= 0 || m_nu2[i] <= 0 {
			// 	 					fmt.Fprintf(msgout, ("Invalid or incomplete data for material %v\n"), i)
			// 	 					goto memFree
			// 	 				}
		}
	}
	// 	}
	n_x[0], n_y[0] = 10, 0
	n_x[1], n_y[1] = 10, 5
	n_x[2], n_y[2] = 10, 10

	// 	{
	// 		// nodes
	// 		for i = 0; i < n_n; i++ {
	// 			if fmt.Fscanf(fw, ("%f %f"), n_x[i:], n_y[i:]) <= 0 {
	// 				goto memFree
	// 			}
	// 		}
	// 	}
	e_n1[0], e_n2[0], e_mat[0], e_t[0] = 0, 1, 0, 0.2
	e_n1[1], e_n2[1], e_mat[1], e_t[1] = 1, 2, 0, 0.2

	// 	{
	// 		// elements
	// 		for i = 0; i < n_e; i++ {
	// 			if fmt.Fscanf(fw, ("%v %v %v %f"), e_n1[i:], e_n2[i:], e_mat[i:], e_t[i:]) <= 0 {
	// 				goto memFree
	// 			}
	// 			if e_n1[i] < 0 || e_n1[i] >= n_n {
	// 				fmt.Fprintf(msgout, ("Invalid n1 in element %v\n"), i)
	// 				goto memFree
	// 			}
	// 			if e_n2[i] < 0 || e_n2[i] >= n_n {
	// 				fmt.Fprintf(msgout, ("Invalid n2 in element %v\n"), i)
	// 				goto memFree
	// 			}
	// 			if e_n1[i] == e_n2[i] {
	// 				fmt.Fprintf(msgout, ("Invalid nodes in element %v\n"), i)
	// 				goto memFree
	// 			}
	// 			if e_mat[i] < 0 || e_mat[i] >= n_m {
	// 				fmt.Fprintf(msgout, ("Invalid material in element %v\n"), i)
	// 				goto memFree
	// 			}
	// 			if e_t[i] <= 0 {
	// 				fmt.Fprintf(msgout, ("Invalid width in element %v\n"), i)
	// 				goto memFree
	// 			}
	// 		}
	// 	}
	d_n[0], d_dir[0], d_val[0] = 0, 0, 0
	d_n[1], d_dir[1], d_val[1] = 0, 1, 0
	d_n[2], d_dir[2], d_val[2] = 2, 1, 0
	// 	{
	// 		// displacements
	// 		for i = 0; i < n_d; i++ {
	// 			if fmt.Fscanf(fw, ("%v %v %f"), d_n[i:], d_dir[i:], d_val[i:]) <= 0 {
	// 				goto memFree
	// 			}
	// 			if d_n[i] < 0 || d_n[i] >= n_n {
	// 				fmt.Fprintf(msgout, ("Invalid node in support %v\n"), i)
	// 				goto memFree
	// 			}
	// 			if d_dir[i] < 0 || d_dir[i] >= 6 {
	// 				fmt.Fprintf(msgout, ("Invalid direction in support %v\n"), i)
	// 				goto memFree
	// 			}
	// 			if d_dir[i] > 2 && d_val[i] < 0 {
	// 				fmt.Fprintf(msgout, ("Invalid stiffness in support %v\n"), i)
	// 				goto memFree
	// 			}
	// 		}
	// 	}
	f_n[0], f_dir[0], f_val[0] = 1, 1, 11.899e6

	// 	{
	// 		// forces
	// 		for i = 0; i < n_f; i++ {
	// 			if fmt.Fscanf(fw, ("%v %v %f"), f_n[i:], f_dir[i:], f_val[i:]) <= 0 {
	// 				goto memFree
	// 			}
	// 			if f_n[i] < 0 || f_n[i] >= n_n {
	// 				fmt.Fprintf(msgout, ("Invalid node for force %v\n"), i)
	// 				goto memFree
	// 			}
	// 			if f_dir[i] < 0 || f_dir[i] >= 3 {
	// 				fmt.Fprintf(msgout, ("Invalid direction for force %v\n"), i)
	// 				goto memFree
	// 			}
	// 		}
	// 	}
	w_top, w_bot, w_val, w_min, w_max = 0, 0, 0, 0, 0

	// 	if fmt.Fscanf(fw, ("%f %f %f %v %v"), c4goUnsafeConvert_float64(&w_top), c4goUnsafeConvert_float64(&w_bot), c4goUnsafeConvert_float64(&w_val), c4goUnsafeConvert_int32(&w_min), c4goUnsafeConvert_int32(&w_max)) <= 0 {
	// 		// water pressure data
	// 		goto memFree
	// 	}
	// 	// check of element nodes
	check_elem_data()
	// 	if get_enode_fields() != 0 {
	// 		goto memFree
	// 	}
	fail_type = 1
	// 	if fmt.Fscanf(fw, ("%v"), c4goUnsafeConvert_int32(&fail_type)) <= 0 {
	// 		// failure condition data:
	// 		// that's great, no failure is needed
	// 		fail_type = 0
	// 		n_fail = 0
	// 	} else {
	// 		if fail_type > 0 {
	n_fail = 2
	// 			if fmt.Fscanf(fw, ("%v"), c4goUnsafeConvert_int32(&n_fail)) <= 0 {
	// 				fail_type = 0
	// 			} else {
	// 				if len((func() []float64 {
	fail_data = femDblAlloc(n_fail)
	// 					return fail_data
	// 				}())) == 0 {
	// 					fmt.Fprintf(msgout, ("Cannot allocate memory for failure data!\n"))
	// 					goto memFree
	// 				}
	fail_data[0] = 20e6
	fail_data[1] = 1e6
	// 				for i = 0; i < n_fail; i++ {
	// 					if fmt.Fscanf(fw, ("%f"), fail_data[i:]) <= 0 {
	// 						fmt.Fprintf(msgout, ("Invalid failure data!\n"))
	// 						goto memFree
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	n_r_inp = 1
	// 	if fmt.Fscanf(fw, ("%v"), c4goUnsafeConvert_int32(&n_r_inp)) <= 0 {
	// 		// random variables:
	// 		n_r_inp = 0
	// 		// fprintf(msgout, "No random data found.\n");
	// 		return 0
	// 	}
	// 	if n_r_inp < 1 {
	// 		return 0
	// 	}
	// 	if len((func() []int32 {
	rand_type = femIntAlloc(n_r_inp)
	// 		return rand_type
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	rand_pos = femIntAlloc(n_r_inp)
	// 		return rand_pos
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	rand_indx = femIntAlloc(n_r_inp)
	// 		return rand_indx
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	rand_type[0], rand_pos[0], rand_indx[0] = 4, 0, 0

	// 	for i = 0; i < n_r_inp; i++ {
	// 		if fmt.Fscanf(fw, ("%v %v %v"), rand_type[i:], rand_pos[i:], rand_indx[i:]) <= 0 {
	// 			goto memFree
	// 		}
	// 	}
	// 	if fmt.Fscanf(fw, ("%v"), c4goUnsafeConvert_int32(&n_r_opt)) <= 0 {
	// 		// optimized variables: -------------------------------------
	n_r_opt = 0
	// 		//fprintf(msgout, "No optim. data found.\n");
	//return 0
	// 	}
	// 	if n_r_opt < 1 {
	// 		fmt.Fprintf(msgout, ("Invalid number of optim. inputs!\n"))
	// 		return 0
	// 	}
	// 	if len((func() []int32 {
	// 		opt_type = femIntAlloc(n_r_opt)
	// 		return opt_type
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	// 		opt_pos = femIntAlloc(n_r_opt)
	// 		return opt_pos
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int32 {
	// 		opt_indx = femIntAlloc(n_r_opt)
	// 		return opt_indx
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []float64 {
	// 		opt_data = femDblAlloc(n_r_opt)
	// 		return opt_data
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// 	for i = 0; i < n_r_opt; i++ {
	// 		if fmt.Fscanf(fw, ("%v %v %v"), opt_type[i:], opt_pos[i:], opt_indx[i:]) <= 0 {
	// 			goto memFree
	// 		}
	// 	}
	// 	for i = 0; i < n_r_opt; i++ {
	// 		if fmt.Fscanf(fw, ("%f"), opt_data[i:]) <= 0 {
	// 			femDblFree(opt_data)
	// 			femIntFree(opt_indx)
	// 			femIntFree(opt_pos)
	// 			femIntFree(opt_type)
	// 			n_r_opt = 0
	// 		}
	// 	}
	// 	return 0
	// memFree:
	// 	;
	// 	free_input_data()
	// 	fmt.Fprintf(msgout, ("Error when reading input!\n"))
	// 	return -2
}

// write_input_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:541
func write_input_data() int32 { // fw *fmt.File
	fw := os.Stdout
	// Writes input data to stream ------------------
	var i int32
	// sizes
	fmt.Fprintf(fw, ("%v %v %v %v %v\n"), n_m, n_n, n_e, n_d, n_f)
	{
		// materials
		for i = 0; i < n_m; i++ {
			fmt.Fprintf(fw, (" %v %v %v %v %v %v %v %v\n"), m_E1[i], m_E2[i], m_G[i], m_nu1[i], m_nu2[i], m_q[i], m_vp[i], m_t[i])
		}
	}
	{
		// nodes
		for i = 0; i < n_n; i++ {
			fmt.Fprintf(fw, ("%v %v\n"), n_x[i], n_y[i])
		}
	}
	{
		// elements
		for i = 0; i < n_e; i++ {
			fmt.Fprintf(fw, ("%v %v %v %v\n"), e_n1[i], e_n2[i], e_mat[i], e_t[i])
		}
	}
	{
		// displacements
		for i = 0; i < n_d; i++ {
			fmt.Fprintf(fw, ("%v %v %v\n"), d_n[i], d_dir[i], d_val[i])
		}
	}
	{
		// supports
		for i = 0; i < n_f; i++ {
			fmt.Fprintf(fw, ("%v %v %v\n"), f_n[i], f_dir[i], f_val[i])
		}
	}
	// water pressure data
	fmt.Fprintf(fw, ("%v %v %v %v %v\n"), w_top, w_bot, w_val, w_min, w_max)
	// failure condition data:
	fmt.Fprintf(fw, ("%v\n"), fail_type)
	if fail_type > 0 {
		fmt.Fprintf(fw, ("%v\n"), n_fail)
		for i = 0; i < n_fail; i++ {
			fmt.Fprintf(fw, ("%v "), fail_data[i])
		}
		fmt.Fprintf(fw, ("\n"))
	}
	return 0
}

// free_solver_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:582
func free_solver_data() {
	// Frees data used by solver
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

// alloc_solver_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:600
func alloc_solver_data() int32 {
	// Allocates data for f.e. solver (K,u,F)
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
		// Compute allocation vector
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
			// is "6" enough?
			alloc_field[3*i+j] = 3 * 6 * n_field[i]
		}
	}
	if femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], 1, n_n*3, n_n*3, 0, alloc_field) != 0 {
		// alloc K, u, F
		goto memFree
	}
	if femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&F))[:], 0, n_n*3, n_n*3) != 0 {
		goto memFree
	}
	if femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&u))[:], 0, n_n*3, n_n*3) != 0 {
		goto memFree
	}
	// 	femIntFree(alloc_field)
	// 	femIntFree(n_field)
	return 0
memFree:
	;
	// 	if len(alloc_field) != 0 {
	// 		femIntFree(alloc_field)
	// 	}
	// 	if len(n_field) != 0 {
	// 		femIntFree(n_field)
	// 	}
	free_solver_data()
	fmt.Fprintf(msgout, ("Out of memory!"))
	return -4
}

// get_D_matrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:670
func get_D_matrix(i int32, t float64, D []tMatrix) {
	// computes material stiffness matrix of elemen
	// * @param i element nomber <0..n_e-1>
	// * @param t eleemnt width
	// * @param D pointer to allocated (!) D matrix
	//
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

// get_B_matrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:704
func get_B_matrix(i int32, B []tMatrix, Lc []float64, Rc []float64) {
	// computes B matrix
	// * @param i element number
	// * @param B pointer to allocated (!) B matrix
	// * @param Lc element length (result)
	// * @param Rc average distance from axis or revolution
	//
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
	// B matrix:
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

// get_matrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:743
func get_matrix() int32 {
	// creates stiffness matrix
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
	// fmt.Println(":::: n_e", n_e, " e_n1 == ", e_n1)
	for i = 0; i < n_e; i++ {
		if (func() float64 {
			t = m_t[e_mat[i]]
			return t
		}()) <= 0 {
			// if material width is specified then use element width:
			t = e_t[i]
		}
		t = e_t[i]
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:])
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&B))[:])
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&Bt))[:])
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&BtD))[:])
		femMatSetZero((*[1000000]tMatrix)(unsafe.Pointer(&D))[:])
		// material stiffness matrix D:
		get_D_matrix(i, t, (*[1000000]tMatrix)(unsafe.Pointer(&D))[:])
		// B matrix
		get_B_matrix(i, (*[1000000]tMatrix)(unsafe.Pointer(&B))[:], c4goUnsafeConvert_float64(&L), c4goUnsafeConvert_float64(&R))
		// transpose of B
		femMatTran((*[1000000]tMatrix)(unsafe.Pointer(&B))[:], (*[1000000]tMatrix)(unsafe.Pointer(&Bt))[:])
		// matrix multiplications (Bt*D*B):
		// => BtD
		femMatMatMult((*[1000000]tMatrix)(unsafe.Pointer(&Bt))[:], (*[1000000]tMatrix)(unsafe.Pointer(&D))[:], (*[1000000]tMatrix)(unsafe.Pointer(&BtD))[:])
		// => Ke  without L*R
		femMatMatMult((*[1000000]tMatrix)(unsafe.Pointer(&BtD))[:], (*[1000000]tMatrix)(unsafe.Pointer(&B))[:], (*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:])
		// element stifness matrix Ke:
		femValMatMultSelf(R*L, (*[1000000]tMatrix)(unsafe.Pointer(&Ke))[:])

		// localisation to "K":
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

		if math.Abs((func() float64 {
			q = m_q[e_mat[i]]
			return q
		}())) > 1e-07 {
			// gravitation
			F2 = -0.5 * q * t * L
			// fmt.Println("i = ", i, e_n1, e_n2, F)
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], 3*e_n1[i], F2, 1)
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], 3*e_n2[i], F2, 1)
		}
	}
	return 0
}

// generate_water_load_x - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:809
func generate_water_load_x() int32 {
	// generates water pressure load
	// it goes through elements and decides if they are under the
	//   * water level (or over the bottom) then it computes horizontal
	//   * pressure on the element nodes
	//
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
	// don't ignore this node
	var use_1 int32 = 1
	// don't ignore this node
	var use_2 int32 = 1
	var pos1 int32
	var pos2 int32
	// real limits of water position
	var y_max float64
	var y_min float64
	// hydrostatic pressures on element - top, bot
	var a float64
	var b float64
	if math.Abs(w_val) > 100*1e-07 {
		if w_max-w_min == 0 {
			// limits for element testing (probably unused):
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
		// setting of unreachable limits for water
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
			// adjusting limits:
			y_max = w_top
		}
		if w_bot > y_min {
			y_min = w_bot
		}
		for i = from; i < to; i++ {
			y1 = n_y[e_n1[i]]
			y2 = n_y[e_n2[i]]
			if y1 > y_max || y1 < y_min {
				// geometric features:
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
				// nothing to do
				continue
			}
			// TODO: compute limit values
			b = (y_max - y1) * w_val
			a = (y_max - y2) * w_val
			if use_1 == 0 {
				// set values in nodes:
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
				// positions of loads
				// val1 (lower) is at n1
				pos1 = e_n1[i]*3 + 1
				pos2 = e_n2[i]*3 + 1
			} else {
				// val1 is at n2
				pos1 = e_n2[i]*3 + 1
				pos2 = e_n1[i]*3 + 1
			}
			// adding of loads:
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], pos1, val1, 1)
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], pos2, val2, 1)
		}
	}
	return 0
}

// get_loads_and_supports - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:939
func get_loads_and_supports() int32 {
	// applies supports in nodes
	var i int32
	var j int32
	var pos int32
	for i = 0; i < n_f; i++ {
		femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&F))[:], f_n[i]*3+f_dir[i]+1, f_val[i], 1)
	}
	for i = 0; i < n_d; i++ {
		if d_dir[i] > 2 {
			// stifnesses
			pos = d_n[i]*3 + d_dir[i] - 2
			femMatPutAdd((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos, pos, d_val[i], 1)
		} else {
			// displacements
			pos = d_n[i]*3 + d_dir[i] + 1
			if math.Abs(d_val[i]) <= 1e-07 {
				femMatSetZeroCol((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos)
				femMatSetZeroRow((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], pos)
				femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&u))[:], pos, 0, 0)
				// yes, it deletes force in support
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

// get_int_forces - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:994
func get_int_forces(el int32, N1 []float64, N2 []float64, M1 []float64, M2 []float64, Q []float64) {
	// computes internal force is nodes
	// * @param el element number <0..n_e-1>
	// * @param N1 meridian force
	// * @param N2 perpendicular force
	// * @param M1 meridian moment
	// * @param M2 perpendicular force
	// * @param Q tangent force
	// * @return status
	//
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
		// get local stiffness vector
		for j = 1; j <= 6; j++ {
			if j < 4 {
				posj = e_n1[el]*3 + j
			} else {
				posj = e_n2[el]*3 + j - 3
			}
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&ue))[:], j, femVecGet((*[1000000]tVector)(unsafe.Pointer(&u))[:], posj), 0)
		}
	}
	// get B and D
	t = e_t[el]
	get_D_matrix(el, t, (*[1000000]tMatrix)(unsafe.Pointer(&D))[:])
	get_B_matrix(el, (*[1000000]tMatrix)(unsafe.Pointer(&B))[:], c4goUnsafeConvert_float64(&L), c4goUnsafeConvert_float64(&R))
	femMatMatMult((*[1000000]tMatrix)(unsafe.Pointer(&D))[:], (*[1000000]tMatrix)(unsafe.Pointer(&B))[:], (*[1000000]tMatrix)(unsafe.Pointer(&DB))[:])
	// get vector
	femMatVecMult((*[1000000]tMatrix)(unsafe.Pointer(&DB))[:], (*[1000000]tVector)(unsafe.Pointer(&ue))[:], (*[1000000]tVector)(unsafe.Pointer(&Fe))[:])
	N1[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 1)
	N2[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 2)
	M1[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 3)
	M2[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 4)
	Q[0] = femVecGet((*[1000000]tVector)(unsafe.Pointer(&Fe))[:], 5)
}

// print_result - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1036
// func print_result(fw *fmt.File) int32 {
// 	return 0
// }

// generate_rand_out_file - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1092
func generate_rand_out_file() { //fw *fmt.File) {
	fw := os.Stdout
	// generates output variable list for Monte input file
	var i int32
	fmt.Fprintf(fw, ("%v\n"), n_n*8+1)
	fmt.Fprintf(fw, ("FAIL 3 2\n"))
	for i = 0; i < n_n; i++ {
		fmt.Fprintf(fw, ("UY%v 2\n"), i)
		fmt.Fprintf(fw, ("UX%v 2\n"), i)
		fmt.Fprintf(fw, ("RT%v 2\n"), i)
		fmt.Fprintf(fw, ("NX%v 2\n"), i)
		fmt.Fprintf(fw, ("NY%v 2\n"), i)
		fmt.Fprintf(fw, ("MX%v 2\n"), i)
		fmt.Fprintf(fw, ("MY%v 2\n"), i)
		fmt.Fprintf(fw, ("QQ%v 2\n"), i)
	}
	// no correlations at all
	fmt.Fprintf(fw, ("0\n"))
}

// generate_d_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1115
func generate_d_type(type_ int32) string {
	switch type_ {
	case 0:
		// generates textual symbol for displacement
		return ("UY")
	case 1:
		return ("UX")
	case 2:
		return ("RT")
	case 3:
		return ("EY")
	case 4:
		return ("EX")
	case 5:
		return ("ER")
		break
	}
	return ("XX")
}

// generate_f_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1130
func generate_f_type(type_ int32) string {
	switch type_ {
	case 0:
		// generates textual symbol for force
		return ("FY")
	case 1:
		return ("FX")
	case 2:
		return ("MT")
		break
	}
	return ("XX")
}

// generate_w_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1142
func generate_w_type(type_ int32) string {
	switch type_ {
	case 0:
		// generates textual symbol for water load
		return ("TOP")
	case 1:
		return ("BOT")
	case 2:
		return ("SIZE")
		break
	}
	return ("XX")
}

// generate_fc_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1154
func generate_fc_type(type_ int32) string {
	switch fail_type {
	case 1:
		switch type_ {
		case 0:
			// generates textual symbol for failure criteria
			// concrete cracking limit
			return ("COMPR")
		case 1:
			return ("TENS")
		default:
			return ("UNKNOWN")
			break
		}
	default:
		return ("XX")
		break
	}
	return ("XX")
}

// generate_rand_input_file - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1179
func generate_rand_input_file() { //fw *fmt.File) {
	//
	// 	// Writes input data for Monte
	// 	// * @param fw file stream to write data
	// 	// * @return status
	// 	//
	// 	var i int32
	// 	fmt.Fprintf(fw, ("%v\n"), n_r_inp)
	// 	for i = 0; i < n_r_inp; i++ {
	// 		switch rand_type[i] {
	// 		case 0:
	// 			switch rand_indx[i] {
	// 			case 0:
	// 				// material
	// 				fmt.Fprintf(fw, ("MAT%v_E1 %v 1 normal-1-02.dis\n"), rand_pos[i], m_E1[rand_pos[i]])
	// 			case 1:
	// 				fmt.Fprintf(fw, ("MAT%v_E2 %v 1 normal-1-02.dis\n"), rand_pos[i], m_E2[rand_pos[i]])
	// 			case 2:
	// 				fmt.Fprintf(fw, ("MAT%v_G %v 1 normal-1-02.dis\n"), rand_pos[i], m_G[rand_pos[i]])
	// 			case 3:
	// 				fmt.Fprintf(fw, ("MAT%v_NU1 %v 1 normal-1-02.dis\n"), rand_pos[i], m_nu1[rand_pos[i]])
	// 			case 4:
	// 				fmt.Fprintf(fw, ("MAT%v_NU2 %v 1 normal-1-02.dis\n"), rand_pos[i], m_nu2[rand_pos[i]])
	// 			case 5:
	// 				fmt.Fprintf(fw, ("MAT%v_VF %v 1 normal-1-02.dis\n"), rand_pos[i], m_vp[rand_pos[i]])
	// 			case 6:
	// 				fmt.Fprintf(fw, ("MAT%v_T %v 1 normal-1-02.dis\n"), rand_pos[i], m_t[rand_pos[i]])
	// 				break
	// 			}
	// 		case 1:
	// 			switch rand_indx[i] {
	// 			case 0:
	// 				// node
	// 				fmt.Fprintf(fw, ("N%v_X %v 1 normal-1-02.dis\n"), rand_pos[i], n_x[rand_pos[i]])
	// 			case 1:
	// 				fmt.Fprintf(fw, ("N%v_Y %v 1 normal-1-02.dis\n"), rand_pos[i], n_y[rand_pos[i]])
	// 				break
	// 			}
	// 		case 2:
	// 			// element width
	// 			fmt.Fprintf(fw, ("E%v_WIDTH %v 1 normal-1-02.dis\n"), rand_pos[i], e_t[rand_pos[i]])
	// 		case 3:
	// 			// displacement
	// 			fmt.Fprintf(fw, ("D%v_%s_SIZE %v 1 normal-1-02.dis\n"), rand_pos[i], generate_d_type(rand_indx[i]), d_val[rand_pos[i]])
	// 		case 4:
	// 			// force
	// 			fmt.Fprintf(fw, ("F%v_%s_SIZE %v 1 normal-1-02.dis\n"), rand_pos[i], generate_f_type(rand_indx[i]), f_val[rand_pos[i]])
	// 		case 5:
	// 			switch rand_indx[i] {
	// 			case 0:
	// 				// node
	// 				fmt.Fprintf(fw, ("W_%s %v 1 normal-1-02.dis\n"), generate_w_type(rand_indx[i]), w_top)
	// 			case 1:
	// 				fmt.Fprintf(fw, ("W_%s %v 1 normal-1-02.dis\n"), generate_w_type(rand_indx[i]), w_bot)
	// 			case 2:
	// 				fmt.Fprintf(fw, ("W_%s %v 1 normal-1-02.dis\n"), generate_w_type(rand_indx[i]), w_val)
	// 				break
	// 			}
	// 		case 6:
	// 			// failure critical
	// 			fmt.Fprintf(fw, ("FC_%s_%v %v 1 normal-1-02.dis\n"), generate_fc_type(rand_indx[i]), rand_indx[i], fail_data[rand_indx[i]])
	// 		default:
	// 			fmt.Fprintf(msgout, ("Unused input random variable %v!\n"), i)
	// 			break
	// 		}
	// 	}
}

// fail_test_concrete - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1283
func fail_test_concrete() int32 {
	// ** FAILURE CRITERIA DEFINITIONS **
	//
	// *  provides failure testing
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
		// internal forces in centroid
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
			// failed
			return 1
		}
	}
	return 0
}

// fail_test - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1348
func fail_test() int32 {
	switch fail_type {
	case 1:
		// runs failure test
		// * @return 1 for failure, 0 for tother cases
		//
		// concrete: no-crack allowed
		return fail_test_concrete()
	case 0:
		fallthrough
	default:
		// no criteria -> no fail
		return 0
		break
	}
	return 0
}

// compute_price - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1364
func compute_price() float64 {
	// Computes price of the structure based on material volume
	var price float64
	var volume float64
	var dx float64
	var dpx float64
	var dy float64
	var i int32
	price = 0
	for i = 0; i < n_e; i++ {
		// R-r
		dx = math.Abs(n_x[e_n2[i]] - n_x[e_n1[i]])
		// R+r
		dpx = n_x[e_n2[i]] + n_x[e_n1[i]]
		dy = math.Abs(n_y[e_n2[i]] - n_y[e_n1[i]])
		if dx <= 1e-07 {
			// cillinder
			// 2*pi*r
			volume = dy * (2 * 3.141592653589793 * n_x[e_n2[i]])
		} else {
			if dy <= 1e-07 {
				// circle in plane
				volume = 3.141592653589793 * math.Abs(math.Pow(n_x[e_n2[i]], 2)-math.Pow(n_x[e_n1[i]], 2))
			} else {
				// arbitrary shape
				volume = 3.141592653589793 * dpx * math.Sqrt(dy*dy+dx*dx)
			}
		}
		price += e_t[i] * volume * m_vp[e_mat[i]]
	}
	return price
}

// optim_replace_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1399
func optim_replace_data(ifld []float64) int32 {
	// replace f.e. input  data with their optimized counterparts
	var i int32
	if len(ifld) == 0 || n_r_opt < 1 {
		return 0
	}
	for i = 0; i < n_r_opt; i++ {
		switch opt_type[i] {
		case 0:
			switch opt_indx[i] {
			case 0:
				// material
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
				// node
				n_x[opt_pos[i]] = ifld[i]
			case 1:
				n_y[opt_pos[i]] = ifld[i]
				break
			}
		case 2:
			// element width
			e_t[opt_pos[i]] = ifld[i]
		case 3:
			// displacement
			d_val[opt_pos[i]] = ifld[i]
		case 4:
			// force
			f_val[opt_pos[i]] = ifld[i]
		case 5:
			switch opt_indx[i] {
			case 0:
				// material
				w_top = ifld[i]
			case 1:
				w_bot = ifld[i]
			case 2:
				w_val = ifld[i]
				break
			}
		case 6:
			if opt_indx[i] < n_fail {
				// failure condition
				fail_data[opt_indx[i]] = ifld[i]
			}
		default:
			fmt.Fprintf(msgout, ("Unused input optim variable %v!\n"), i)
			break
		}
	}
	return 0
}

// print_help - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1706
// func print_help(argc int32, argv []) {
// 	// Prints simple help to stdout
// 	// * @param argc the same as "argc" from main
// 	// * @param argv the same as "argv" from main
// 	//
// 	fmt.Printf("\neSHELL 1.0: axisymetric shells solver\n")
// 	fmt.Printf("(C) 2010 VSB-TU of Ostrava \n")
// 	fmt.Printf("(C) 2003-2010 Jiri Brozovsky (uFEM libraries)\n")
// 	fmt.Printf("\nThis is free software licensed under GNU GPL 2.0\n")
// 	noarch.Printf(("\nUsage: %s [parameters] <input >output\n"), argv[0])
// 	fmt.Printf("\nParameters:\n")
// 	fmt.Printf("   -s        ... force solution only output\n")
// 	fmt.Printf("   -g        ... generate random data only \n")
// 	fmt.Printf("   -p        ... compute price function only\n")
// 	fmt.Printf("   -w        ... write input data and finish\n")
// 	fmt.Printf("   -h        ... print this help\n")
// }

// cmd_param - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1722
func cmd_param() { //argc int32, argv []) int32 {

	solution_only = 1
	price_only = 1
	write_only = 1

	// Understands command line parameters
	// 	var i int32
	// 	for i = 1; i < argc; i++ {
	// 		if noarch.Strcmp(argv[i], ("-h")) == 0 || noarch.Strcmp(argv[i], ("--help")) == 0 {
	// 			print_help(argc, argv)
	// 			noarch.Exit(0)
	// 		}
	// 		if noarch.Strcmp(argv[i], ("-s")) == 0 || noarch.Strcmp(argv[i], ("--solution")) == 0 {
	// 			solution_only = 1
	// 			price_only = 0
	// 			random_only = 0
	// 		}
	// 		if noarch.Strcmp(argv[i], ("-g")) == 0 || noarch.Strcmp(argv[i], ("-r")) == 0 || noarch.Strcmp(argv[i], ("--random")) == 0 {
	// 			solution_only = 0
	// 			price_only = 0
	// 			random_only = 1
	// 		}
	// 		if noarch.Strcmp(argv[i], ("-p")) == 0 || noarch.Strcmp(argv[i], ("--price")) == 0 {
	// 			solution_only = 0
	// 			price_only = 1
	// 			random_only = 0
	// 		}
	// 		if noarch.Strcmp(argv[i], ("-w")) == 0 || noarch.Strcmp(argv[i], ("--price")) == 0 {
	// 			write_only = 1
	// 		}
	// 	}
	// return 0
}

// main - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1744
func main() {
	// 	argc := int32(len(os.Args))
	// 	argv := []{}
	// 	for _, argvSingle := range os.Args {
	// 		argv = append(argv, (argvSingle))
	// 	}
	// 	defer noarch.AtexitRun()
	// main() routine for standalone program only.
	var stat int32
	//msgout = noarch.Stderr
	cmd_param() //argc, argv)
	//stat +=
	read_input_data() //noarch.Stdin)
	stat += alloc_solver_data()
	stat += optim_replace_data(opt_data)
	//if write_only == 1 {
	write_input_data() //noarch.Stdout)
	//	return
	//}
	//if solution_only == 1 {
	stat += get_matrix()
	stat += generate_water_load_x()
	stat += get_loads_and_supports()
	stat = femEqsCGwJ((*[1000000]tMatrix)(unsafe.Pointer(&K))[:], (*[1000000]tVector)(unsafe.Pointer(&F))[:], (*[1000000]tVector)(unsafe.Pointer(&u))[:], 1e-09, 6*3*n_n)
	//	}
	// if n_r_inp > 0 && random_only == 1 {
	//	if solution_only != 0 {
	//print_result(noarch.Stderr)
	//	}
	//	generate_rand_input_file(noarch.Stdout)
	//	generate_rand_out_file(noarch.Stdout)
	//} else {
	//	if solution_only == 1 {
	//print_result(noarch.Stdout)
	//	}
	//	}
	// 	if solution_only == 1 {
	// 		if fail_test() != 0 {
	// 			fmt.Fprintf(noarch.Stderr, ("# Structure FAILED\n"))
	// 		}
	// 	}
	// 	if price_only == 1 {
	// 		if solution_only == 1 {
	// 			fmt.Fprintf(msgout, ("# Price is %f\n"), compute_price())
	// 		} else {
	fmt.Fprintf(os.Stdout, ("%v\n"), compute_price())
	// 		}
	// 	}
	return
}

// femMatNull - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:33
func femMatNull(mat []tMatrix) {
	//
	//   File name: fem_math.c
	//   Date:      2003/04/12 12:44
	//   Author:    Jiri Brozovsky
	//
	//   Copyright (C) 2003 Jiri Brozovsky
	//
	//   This program is free software; you can redistribute it and/or
	//   modify it under the terms of the GNU General Public License as
	//   published by the Free Software Foundation; either version 2 of the
	//   License, or (at your option) any later version.
	//
	//   This program is distributed in the hope that it will be useful, but
	//   WITHOUT ANY WARRANTY; without even the implied warranty of
	//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
	//   General Public License for more details.
	//
	//   You should have received a copy of the GNU General Public License
	//   in a file called COPYING along with this program; if not, write to
	//   the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
	//   02139, USA.
	//
	//  FEM Solver - matrix library
	//
	//  $Id: fem_math.c,v 1.46 2005/07/11 17:56:16 jirka Exp $
	//
	// MATRIX ***
	mat[0].type_ = 0
	mat[0].rows = 0
	mat[0].cols = 0
	mat[0].len_ = 0
	mat[0].pos = nil
	mat[0].data = nil
	mat[0].frompos = nil
	mat[0].defpos = nil
}

// femMatFree - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:45
func femMatFree(mat []tMatrix) {
	mat[0].type_ = 0
	mat[0].rows = 0
	mat[0].cols = 0
	mat[0].len_ = 0
	// 	femIntFree(mat[0].pos)
	// 	femDblFree(mat[0].data)
	// 	femIntFree(mat[0].frompos)
	// 	femIntFree(mat[0].defpos)
}

// femMatAlloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:57
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
			//if len((func() []float64 {
			mat[0].data = femDblAlloc(mat[0].len_)
			//	return mat[0].data
			//}())) == 0 {
			//	goto memFree
			//	}
			mat[0].pos = nil
			mat[0].frompos = nil
			mat[0].defpos = nil
		case 1:
			mat[0].rows = rows
			mat[0].cols = cols
			//if len((func() []int32 {
			mat[0].defpos = femIntAlloc(mat[0].rows)
			//	return mat[0].defpos
			//}())) == 0 {
			//	goto memFree
			//}
			//if len((func() []int32 {
			mat[0].frompos = femIntAlloc(mat[0].rows)
			//	return mat[0].frompos
			//}())) == 0 {
			//	goto memFree
			//}
			if bandwidth > 0 && len(rowdesc) == 0 {
				mat[0].len_ = rows * bandwidth
				//if len((func() []float64 {
				mat[0].data = femDblAlloc(mat[0].len_)
				//	return mat[0].data
				//	}())) == 0 {
				//		goto memFree
				//	}
				//	if len((func() []int32 {
				mat[0].pos = femIntAlloc(mat[0].len_)
				//		return mat[0].pos
				//	}())) == 0 {
				//		goto memFree
				//	}
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
				//	if len((func() []float64 {
				mat[0].data = femDblAlloc(mat[0].len_)
				//		return mat[0].data
				//	}())) == 0 {
				//		goto memFree
				//	}
				//	if len((func() []int32 {
				mat[0].pos = femIntAlloc(sum)
				//		return mat[0].pos
				//	}())) == 0 {
				//		goto memFree
				//	}
			}
			break
		}
		return 0
	} else {
		return -3
	}
	// memFree:
	// 	;
	// 	femMatFree(mat)
	// 	return -4
}

// femMatGet - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:142
func femMatGet(mat []tMatrix, row int32, col int32) float64 {
	// Gets value from matrix
	// * @param mat matrix
	// * @param row row
	// * @param row collumn
	// * @return value
	//
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

// femMatPutAdd - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:185
func femMatPutAdd(mat []tMatrix, row int32, col int32, val float64, mode int32) (c4goDefaultReturn int32) {
	// Adds value to matrix
	// * @param mat matrix
	// * @param row row
	// * @param col column
	// * @param val value
	// * @param mode FEM_PUT for putting ("=") FEM_ADD for adding ("+=")
	// * @return  status
	//
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
			// this is more complicated
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
					// empty field found
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
		// if we are here
		//           * because reallocation is needed !
		//
		return -11
	default:
		return -3
		break
	}
	return
}

// femMatPrn - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:238
// func femMatPrn(mat []tMatrix, name ) {
// 	// Prints matrix to stdout, works only in DEVEL mode
// 	{
// 	}
// }

// femMatPrnF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:261
func femMatPrnF(fname, mat []tMatrix) int32 {
	// Saves matrix to file
	// * @param fname name of file
	// * @param mat matrix to be printed
	// * @return status
	//
	// var fw *fmt.File
	fw := os.Stdout
	var rv int32
	var i int32
	var j int32
	// 	if (func() *fmt.File {
	// 		fw = fmt.Fopen(fname, ("w"))
	// 		return fw
	// 	}()) == nil {
	// 		return -2
	// 	}
	for i = 1; i <= mat[0].rows; i++ {
		for j = 1; j <= mat[0].cols; j++ {
			fmt.Fprintf(fw, (" %v "), femMatGet(mat, i, j))
		}
		fmt.Fprintf(fw, ("\n"))
	}
	// 	if fmt.Fclose(fw) != 0 {
	// 		rv = -2
	// 	}
	return rv
}

// femSparseMatPrnF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:288
func femSparseMatPrnF(fname, mat []tMatrix) int32 {
	// Saves matrix to file IN SPARSE FORM
	// * @param fname name of file
	// * @param mat matrix to be printed
	// * @return status
	//
	// var fw *fmt.File
	fw := os.Stdout
	var rv int32
	var i int32
	var j int32
	var sum int32
	if mat[0].type_ != 1 {
		return -3
	}
	// 	if (func() *fmt.File {
	// 		fw = fmt.Fopen(fname, ("w"))
	// 		return fw
	// 	}()) == nil {
	// 		return -2
	// 	}
	fmt.Fprintf(fw, ("%v %v\n"), mat[0].rows, mat[0].cols)
	for i = 0; i < mat[0].rows; i++ {
		sum = 0
		for j = mat[0].frompos[i]; j < mat[0].frompos[i]+mat[0].defpos[i]; j++ {
			if mat[0].pos[j] >= 0 {
				sum++
			} else {
				break
			}
		}
		fmt.Fprintf(fw, ("%v %v "), i+1, sum)
		for j = mat[0].frompos[i]; j < mat[0].frompos[i]+sum; j++ {
			fmt.Fprintf(fw, ("%v %v "), mat[0].pos[j], mat[0].data[j])
		}
		fmt.Fprintf(fw, ("\n"))
	}
	// 	if fmt.Fclose(fw) != 0 {
	// 		rv = -2
	// 	}
	return rv
}

// femSparseMarketMatPrnF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:329
func femSparseMarketMatPrnF(fname, mat []tMatrix) int32 {
	// Saves matrix to file IN SPARSE FORM (MatrixMarket file standard)
	// * @param fname name of file
	// * @param mat matrix to be printed
	// * @return status
	//
	//var fw *fmt.File
	fw := os.Stdout
	var rv int32
	var i int32
	var j int32
	var sum int32
	if mat[0].type_ != 1 {
		return -3
	}
	// 	if (func() *fmt.File {
	// 		fw = fmt.Fopen(fname, ("w"))
	// 		return fw
	// 	}()) == nil {
	// 		return -2
	// 	}
	fmt.Fprintf(fw, ("%%%%MatrixMarket matrix coordinate real general\n"))
	fmt.Fprintf(fw, ("%v %v %v\n"), mat[0].rows, mat[0].cols, mat[0].len_)
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
			fmt.Fprintf(fw, ("%v %v %v\n"), i+1, mat[0].pos[j], mat[0].data[j])
		}
	}
	// 	if fmt.Fclose(fw) != 0 {
	// 		rv = -2
	// 	}
	return rv
}

// femSparseMatReadF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:368
func femSparseMatReadF(fname, mat []tMatrix) int32 {
	// Reads matrix from file IN SPARSE FORM
	// * @param fname name of file
	// * @param mat matrix (must be unallocated)
	// * @return status
	//
	// var fw *fmt.File
	fw := os.Stdout
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
	// 	if (func() *fmt.File {
	// 		fw = fmt.Fopen(fname, ("r"))
	// 		return fw
	// 	}()) == nil {
	// 		return -2
	// 	}
	fmt.Fscanf(fw, ("%v %v\n"), (*[1000000]int32)(unsafe.Pointer(&mat[0].rows))[:], (*[1000000]int32)(unsafe.Pointer(&mat[0].cols))[:])
	if mat[0].rows <= 0 || mat[0].cols <= 0 {
		return -2
	}
	//if len((func() []int32 {
	mat[0].frompos = femIntAlloc(mat[0].rows)
	//	return mat[0].frompos
	//}())) == 0 {
	//	rv = -4
	//	goto memFree
	//}
	//if len((func() []int32 {
	mat[0].defpos = femIntAlloc(mat[0].rows)
	//	return mat[0].defpos
	//	}())) == 0 {
	//	rv = -4
	//	goto memFree
	//}
	size = mat[0].rows * 300
	//	if len((func() []int32 {
	mat[0].pos = femIntAlloc(size)
	//		return mat[0].pos
	//	}())) == 0 {
	//		rv = -4
	//		goto memFree
	//	}
	//	if len((func() []float64 {
	mat[0].data = femDblAlloc(size)
	//		return mat[0].data
	//	}())) == 0 {
	//		rv = -4
	//		goto memFree
	//	}
	mat[0].type_ = 1
	sum = 0
	for i = 0; i < mat[0].rows; i++ {
		fmt.Fscanf(fw, ("%v %v "), c4goUnsafeConvert_int32(&tmp), mat[0].defpos[i:])
		if i > 0 {
			mat[0].frompos[i] = mat[0].frompos[i-1] + mat[0].defpos[i-1]
		} else {
			// first row
			mat[0].frompos[i] = 0
		}
		for j = 0; j < mat[0].defpos[i]; j++ {
			if sum >= size {
				// enlarge "data" and "pos"
				ensize = size + 2*size*(i/mat[0].rows)
				//	if len((func() []int32 {
				pos0 = femIntAlloc(ensize)
				//		return pos0
				//	}())) == 0 {
				//		rv = -4
				//		goto memFree
				//	}
				//	if len((func() []float64 {
				data0 = femDblAlloc(ensize)
				//		return data0
				//	}())) == 0 {
				//		rv = -4
				//		goto memFree
				//	}
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
			fmt.Fscanf(fw, ("%v %f "), mat[0].pos[sum:], mat[0].data[sum:])
			sum++
		}
	}
	// 	if fmt.Fclose(fw) != 0 {
	// 		rv = -2
	// 	}
	return rv
	// memFree:
	// 	;
	// 	femMatFree(mat)
	// 	return rv
}

// femMatOut - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:447
func femMatOut(a []tMatrix) int32 { //, fw *fmt.File) int32 {
	// Writes matrix to stream (FILE *)
	// * @param a matrix
	// * @param fw stream
	// * @return stave value
	//
	fw := os.Stdout
	var rv int32
	var i int32
	var j int32
	fmt.Fprintf(fw, (" %v %v\n"), a[0].rows, a[0].cols)
	for i = 1; i <= a[0].rows; i++ {
		for j = 1; j <= a[0].cols; j++ {
			fmt.Fprintf(fw, (" %v \n"), femMatGet(a, i, j))
		}
	}
	return rv
}

// femMatSetZeroBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:483
func femMatSetZeroBig(a []tMatrix) {
	// Sets all of matrix contents to 0
	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] = 0
	}
}

// femMatSetZero - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:527
func femMatSetZero(a []tMatrix) {
	// Sets all of matrix contents to 0 FOR SMALL DATA
	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] = 0
	}
}

// femMatSetZeroRow - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:535
func femMatSetZeroRow(a []tMatrix, row int32) {
	// Sets matrix row to 0
	var i int32
	if a[0].type_ == 1 {
		for i = a[0].frompos[row-1]; i < a[0].frompos[row-1]+a[0].defpos[row-1]; i++ {
			if a[0].pos[i] == 0 {
				break
			}
			a[0].data[i] = 0
		}
	} else {
		//fprintf(msgout,"zero on %v\n",i);
		for i = 1; i <= a[0].cols; i++ {
			femMatPutAdd(a, row, i, 0, 0)
		}
	}
}

// femMatSetZeroCol - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:592
func femMatSetZeroCol(a []tMatrix, Col int32) {
	// Sets all of matrix contents to 0
	var i int32
	var j int32
	var ifrom int32
	var ito int32
	var ipos int32
	_ = ipos
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

// femVecNull - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:699
func femVecNull(mat []tVector) {
	// VECTOR ***
	mat[0].type_ = 0
	mat[0].rows = 0
	mat[0].len_ = 0
	mat[0].pos = nil
	mat[0].data = nil
}

// femVecFree - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:708
func femVecFree(mat []tVector) {
	mat[0].type_ = 0
	mat[0].rows = 0
	mat[0].len_ = 0
	// 	femIntFree(mat[0].pos)
	// 	femDblFree(mat[0].data)
}

// femVecAlloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:718
func femVecAlloc(mat []tVector, type_ int32, rows int32, items int32) int32 {
	femVecNull(mat)
	if type_ >= 0 && type_ <= 1 {
		mat[0].type_ = type_
		switch type_ {
		case 0:
			mat[0].rows = rows
			mat[0].len_ = rows
			//	if len((func() []float64 {
			mat[0].data = femDblAlloc(mat[0].len_)
			//		return mat[0].data
			//	}())) == 0 {
			//		goto memFree
			//	}
			mat[0].pos = nil
		case 1:
			// VEC_SPAR cannot be used ;-)
			noarch.Exit(-3)
			mat[0].rows = rows
			// if items > 0 {
			mat[0].len_ = items
			//	if len((func() []float64 {
			mat[0].data = femDblAlloc(mat[0].len_)
			//		return mat[0].data
			//	}())) == 0 {
			//		goto memFree
			//	}
			//	if len((func() []int32 {
			mat[0].pos = femIntAlloc(mat[0].len_)
			//		return mat[0].pos
			//	}())) == 0 {
			//		goto memFree
			//	}
			// 			} else {
			// 				goto memFree
			// 			}
			break
		}
		return 0
	} else {
		return -3
	}
	// memFree:
	// 	;
	// 	femVecFree(mat)
	// 	return -4
}

// femVecPutAdd - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:776
func femVecPutAdd(vec []tVector, pos int32, val float64, mode int32) int32 {
	if pos > vec[0].rows {
		// Adds value to vector
		// * @param vec vector
		// * @param pos row to add value
		// * @param val value
		// * @param mode FEM_PUT for putting ("=") FEM_ADD for adding ("+=")
		// * @return  status
		//
		return -11
	}

	switch vec[0].type_ {
	case 0:
		if mode == 0 {
			// put
			// vec[0].data[pos-1] = val
			vec[0].data[pos] = val
		} else {
			// add
			// fmt.Println(">", pos, vec[0].type_, ":::", mode, vec[0].data)
			// vec[0].data[pos-1] += val
			vec[0].data[pos] += val
		}
	case 1:
		// unimplemented
		noarch.Exit(-3)
	default:
		return -5
		break
	}
	return 0
}

// femVecGet - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:811
func femVecGet(vec []tVector, pos int32) float64 {
	if pos > vec[0].rows {
		// Gets value from vector
		// * @param vec vector
		// * @param pos row to add value
		// * @return value
		//
		return float64(0)
	}
	switch vec[0].type_ {
	case 0:
		return vec[0].data[pos-1]
	case 1:
		// unimplemented
		noarch.Exit(0)
	default:
		return float64(0)
		break
	}
	return float64(0)
}

// femVecPrn - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:839
func femVecPrn(mat []tVector, name string) {
	// Prints vector to stdout, works only in DEVEL mode
	{
		fmt.Println(mat, name)
	}
}

// femVecPrnF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:858
func femVecPrnF(fname, mat []tVector) int32 {
	// Saves vector to file
	// * @param fname name of file
	// * @param mat vector to be printed
	// * @return status
	//
	// var fw *fmt.File
	fw := os.Stdout
	var rv int32
	var i int32
	// 	if (func() *fmt.File {
	// 		fw = fmt.Fopen(fname, ("w"))
	// 		return fw
	// 	}()) == nil {
	// 		return -2
	// 	}
	for i = 1; i <= mat[0].rows; i++ {
		fmt.Fprintf(fw, (" %v "), femVecGet(mat, i))
	}
	fmt.Fprintf(fw, ("\n"))
	// 	if fmt.Fclose(fw) != 0 {
	// 		rv = -2
	// 	}
	return rv
}

// femVecOut - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:883
func femVecOut(a []tVector) int32 {
	// Writes vector to stream (FILE *)
	// * @ a vector
	// * @ fw stream
	// * @return stave value
	//
	fw := os.Stdout
	var rv int32
	var i int32
	fmt.Fprintf(fw, (" %v\n"), a[0].rows)
	for i = 1; i <= a[0].rows; i++ {
		fmt.Fprintf(fw, (" %v \n"), femVecGet(a, i))
	}
	return rv
}

// femVecSetZeroBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:917
func femVecSetZeroBig(a []tVector) {
	// Sets all of vertor contents to 0
	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] = 0
	}
}

// femVecSetZero - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:961
func femVecSetZero(a []tVector) {
	// Sets all of vertor contents to 0 FOR SMALL DATA
	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] = 0
	}
}

// femVecClone - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:971
// func femVecClone(src []tVector, dest []tVector) int32 {
// 	// Clones vectors: src to dest both must be VEC_FULL, same size and allocated
// 	// * @param src original vector
// 	// * @param dest moditied vector
// 	//
// 	var i int32
// 	if src[0].type_ != 0 || dest[0].type_ != 0 {
// 		return -5
// 	}
// 	if src[0].len_ != dest[0].len_ {
// 		return -9
// 	}
// 	for i = 0; i < src[0].len_; i++ {
// 		dest[0].data[i] = src[0].data[i]
// 	}
// 	return 0
// }

// femVecVecMultBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1014
func femVecVecMultBig(a []tVector, b []tVector) float64 {
	// ------------------    Matrix Operations    --------------------
	// vector multiplication (scalar) (a[n]^t * b[n])
	// * @param a vector
	// * @param b vector
	// * @return multiplication product
	//
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

// femVecVecMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1104
func femVecVecMult(a []tVector, b []tVector) float64 {
	// vector multiplication (scalar) (a[n]^t * b[n])  FOR SMALL VECTORS
	// * @param a vector
	// * @param b vector
	// * @return multiplication product
	//
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

// femVecVecMulttoMat - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1140
func femVecVecMulttoMat(a []tVector, b []tVector, c []tMatrix) int32 {
	// vector multiplication (matrix) (a[n] * b[n]^t)
	// * @param a vector
	// * @param b vector
	// * @param c matrix (result)
	// * @return status
	//
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

// femValVecMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1180
func femValVecMult(val float64, a []tVector, b []tVector) int32 {
	// number by vector multiplication (b[n] = val * a[n])
	// * @param val number
	// * @param a original vector (will not be modified)
	// * @param b result (vector) - must be allocated and must have proper size
	// * @return status
	//
	var i int32
	for i = 0; i < a[0].len_; i++ {
		b[0].data[i] = a[0].data[i] * val
	}
	return 0
}

// femValVecMultSelf - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1202
func femValVecMultSelf(val float64, a []tVector) int32 {
	// number by vector multiplication (a[n] = val * a[n])
	// * @param val number
	// * @param a original vector (WILL BE modified)
	// * @return status
	//
	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] *= val
	}
	return 0
}

// femValMatMultSelf - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1215
func femValMatMultSelf(val float64, a []tMatrix) int32 {
	// number by matrix multiplication (a[n] = val * a[n])
	// * @param val number
	// * @param a original number (WILL BE modified)
	// * @return status
	//
	var i int32
	for i = 0; i < a[0].len_; i++ {
		a[0].data[i] *= val
	}
	return 0
}

// femVecMatMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1229
func femVecMatMult(a []tVector, b []tMatrix, c []tVector) int32 {
	// vector by matrix multiplication (a[n]^t * b[n,m]  = c[m])
	// * @param a vector
	// * @param b matrix
	// * @param vector (result)
	// * @return status
	//
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

// femVecMatVecMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1272
func femVecMatVecMult(a []tVector, b []tMatrix, c []tVector) float64 {
	// Vector by matrix by vector multiplication (a[n]^t * b[n,m] * c[m]  = d)
	// * For small full matrices only (it is slow).
	// * @param a vector
	// * @param b matrix
	// * @param c vector
	// * @return constant (result)
	//
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

// femMatVecMultBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1347
func femMatVecMultBig(a []tMatrix, b []tVector, c []tVector) int32 {
	// Matrix by vector multiplication (a[m,n]*b[n] = b[n])
	// * @param a matrix
	// * @param b vector
	// * @param c vector (result)
	// * @return status
	//
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

// femMatVecMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1456
func femMatVecMult(a []tMatrix, b []tVector, c []tVector) int32 {
	// Matrix by vector multiplication (a[m,n]*b[n] = b[n]) FOR SMALL DATA
	// * @param a matrix
	// * @param b vector
	// * @param c vector (result)
	// * @return status
	//
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

// femVecLinCombBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1540
func femVecLinCombBig(amult float64, a []tVector, bmult float64, b []tVector, c []tVector) int32 {
	// linear combination of vectors am*a[m,n]+ bm*b[m,n] = c[m,n] (c..MAT_FULL)
	// * @param am  "a" vector multiplier
	// * @param a vector
	// * @param bm  "b" vector multiplier
	// * @param b vector
	// * @param c vector (result)
	// * @return status
	//
	var i int32
	if a[0].type_ == 0 && b[0].type_ == 0 && c[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			c[0].data[i] = amult*a[0].data[i] + bmult*b[0].data[i]
		}
	} else {
		// VERY SLOW CODE:
		for i = 1; i <= a[0].rows; i++ {
			femVecPutAdd(c, i, femVecGet(a, i)*amult+femVecGet(b, i)*bmult, 0)
		}
	}
	return 0
}

// femVecLinComb - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1629
func femVecLinComb(amult float64, a []tVector, bmult float64, b []tVector, c []tVector) int32 {
	// linear combination of vectors am*a[m,n]+ bm*b[m,n] = c[m,n] (c..MAT_FULL)
	// * @param am  "a" vector multiplier
	// * @param a vector
	// * @param bm  "b" vector multiplier
	// * @param b vector
	// * @param c vector (result)
	// * @return status
	//
	var i int32
	if a[0].type_ == 0 && b[0].type_ == 0 && c[0].type_ == 0 {
		for i = 0; i < a[0].rows; i++ {
			c[0].data[i] = amult*a[0].data[i] + bmult*b[0].data[i]
		}
	} else {
		// SLOW CODE:
		for i = 1; i <= a[0].rows; i++ {
			femVecPutAdd(c, i, femVecGet(a, i)*amult+femVecGet(b, i)*bmult, 0)
		}
	}
	return 0
}

// femMatMatMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1662
func femMatMatMult(a []tMatrix, b []tMatrix, c []tMatrix) int32 {
	// matrix by matrix multiplication a[m,n]*b[n,h] = c[m,h]
	// * @param a matrix
	// * @param b matrix
	// * @param c matrix (result)
	// * @return status
	//
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
					//val += femMatGet(a, i,k)*femMatGet(b, k,j);
					val += a[0].data[i*a[0].cols+k] * b[0].data[k*b[0].cols+j]
				}
				//femMatPut(c, i,j, val);
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

// femMatLinComb - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1714
func femMatLinComb(am float64, a []tMatrix, bm float64, b []tMatrix, c []tMatrix) int32 {
	// linear combination of matrices am*a[m,n]+ bm*b[m,n] = c[m,n] (c..MAT_FULL)
	// * @param am  "a" matrix multiplier
	// * @param a matrix
	// * @param bm  "b" matrix multiplier
	// * @param b matrix
	// * @param c matrix (result)
	// * @return status
	//
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

// femMatTran - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1750
func femMatTran(a []tMatrix, b []tMatrix) int32 {
	// matrix transposition - works only on dense matrices (MAT_FULL)
	// * @param a matrix (original)
	// * @param b matrix (result - must be allocated)
	// * @return status
	//
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

// femMatNormBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1822
func femMatNormBig(a []tMatrix) float64 {
	// Computes norm of sparse matrix
	// *  @param a matrix
	// *  @return norm
	//
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

// femMatNorm - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1921
func femMatNorm(a []tMatrix) float64 {
	// Computes norm of sparse matrix FOR SMALL DATA
	// *  @param a matrix
	// *  @return norm
	//
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

// femVecNormBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1988
func femVecNormBig(a []tVector) float64 {
	// Computes Euclide norm of vector sum(sqrt(a*a))
	// *  @param a     vector
	// *  @return norm
	//
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

// femVecNorm - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2072
func femVecNorm(a []tVector) float64 {
	// Computes Euclide norm of vector sum(sqrt(a*a)) FOR SMALL DATA
	// *  @param a     vector
	// *  @return norm
	//
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

// femVecAddVec - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2104
func femVecAddVec(orig []tVector, mult float64, addt []tVector) int32 {
	// Adds vector "addt" to "orig" e.g. orig += mult*addt
	// * @param orig original vector (to be modified)
	// * @param mult scalar multiplier
	// * @param addt addition vector
	// * @return status
	//
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

// femMatInv - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2137
func femMatInv(a []tMatrix) int32 {
	// Does matrix inversion UNOPTIMIZED!
	// *  @param a  matrix to be inverted
	//
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
	//if
	femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&f1))[:], 0, n, n)
	//!= 0 {
	//	return -4
	//}
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

// femLUdecomp - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2219
func femLUdecomp(a []tMatrix, index []tVector) int32 {
	// L-U:
	// Decomposition to L/U
	// * @param a matrix (will be modified!)
	// * @param index index vector
	// * @param d modified index status (-1/+1)
	// * @return status
	//
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
	//if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&vv))[:], 0, n, n)
	//	return rv
	//}()) != 0 {
	//	goto memFree
	//}
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
			// singular matrix
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
	// memFree:
	// 	;
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&vv))[:])
	return rv
}

// femLUback - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2322
func femLUback(a []tMatrix, index []tVector, b []tVector) int32 {
	// Decomposition to L/U
	// * @param a matrix (will be modified!)
	// * @param index index vector
	// * @param b right hand side/result vector (will be modified!)
	// * @return status
	//
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
			// means ii > 0
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

// femLUinverse - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2374
func femLUinverse(a []tMatrix) int32 {
	// Inversion of "a" matrix using L/U
	// * @param a matrix (will be modified!)
	// * @return status
	//
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
	//if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&col))[:], 0, n, n)
	//		return rv
	//	}()) != 0 {
	//		goto memFree
	//	}
	//	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&index))[:], 0, n, n)
	//		return rv
	//	}()) != 0 {
	//		goto memFree
	//	}
	//	if (func() int32 {
	rv = femMatAlloc((*[1000000]tMatrix)(unsafe.Pointer(&b))[:], 0, n, n, 0, nil)
	//		return rv
	//	}()) != 0 {
	//		goto memFree
	//	}
	//	if (func() int32 {
	rv = femLUdecomp(a, (*[1000000]tVector)(unsafe.Pointer(&index))[:])
	//		return rv
	//	}()) != 0 {
	//		goto memFree
	//	}
	for j = 1; j <= n; j++ {
		for i = 1; i <= n; i++ {
			femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&col))[:], i, 0, 0)
		}
		femVecPutAdd((*[1000000]tVector)(unsafe.Pointer(&col))[:], j, 1, 0)
		//		if (func() int32 {
		rv = femLUback(a, (*[1000000]tVector)(unsafe.Pointer(&index))[:], (*[1000000]tVector)(unsafe.Pointer(&col))[:])
		//			return rv
		//		}()) != 0 {
		//			goto memFree
		//		}
		for i = 1; i <= n; i++ {
			femMatPutAdd((*[1000000]tMatrix)(unsafe.Pointer(&b))[:], i, j, femVecGet((*[1000000]tVector)(unsafe.Pointer(&col))[:], i), 0)
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			femMatPutAdd(a, i, j, femMatGet((*[1000000]tMatrix)(unsafe.Pointer(&b))[:], i, j), 0)
		}
	}
	// memFree:
	// 	;
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&col))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&index))[:])
	// 	femMatFree((*[1000000]tMatrix)(unsafe.Pointer(&b))[:])
	return rv
}

// femVecSwitch - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2424
func femVecSwitch(a []tVector, b []tVector) int32 {
	// Moves "a" to "b" and "b" to "a"
	// * @param a vector
	// * @param b vector
	// * @return status
	//
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

// femVecCloneDiff - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2449
func femVecCloneDiff(orig []tVector, clone []tVector) int32 {
	// Copies vector content to a larger one (extra fields are left untouched)
	// * @param orig original vector
	// * @param clone target vector (to be modified)
	// * @return status
	//
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

// femMatCloneDiffToSame - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2484
func femMatCloneDiffToSame(orig []tMatrix, clone []tMatrix) int32 {
	// TODO FIX!!! Copies sparse matrix content to a larger one (extra fields are left untouched)
	// * it is assumed that a) there is a space for data in "clone", b) identical data
	// * in both matrices are stored at identical places
	// * @param orig original vector
	// * @param clone target vector (to be modified)
	// * @return status
	//
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

// femMatCloneDiffToEmpty - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2537
func femMatCloneDiffToEmpty(orig []tMatrix, clone []tMatrix) int32 {
	// Copies sparse matrix content to a larger one (extra fields are left untouched)
	// * it is assumed that a) there is a space for data in "clone", b) identical data
	// * in both matrices are stored at identical places
	// * @param orig original vector
	// * @param clone target vector (to be modified)
	// * @return status
	//
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

// eqsCompResid - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:37
func eqsCompResid(a []tMatrix, x []tVector, b []tVector, r []tVector) int32 {
	//
	//   File name: fem_eqs.c
	//   Date:      2003/04/13 10:38
	//   Author:    Jiri Brozovsky
	//
	//   Copyright (C) 2003 Jiri Brozovsky
	//
	//   This program is free software; you can redistribute it and/or
	//   modify it under the terms of the GNU General Public License as
	//   published by the Free Software Foundation; either version 2 of the
	//   License, or (at your option) any later version.
	//
	//   This program is distributed in the hope that it will be useful, but
	//   WITHOUT ANY WARRANTY; without even the implied warranty of
	//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
	//   General Public License for more details.
	//
	//   You should have received a copy of the GNU General Public License
	//   in a file called COPYING along with this program; if not, write to
	//   the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
	//   02139, USA.
	//
	//  FEM Solver - linear equation system solver(s)
	//
	//  $Id: fem_eqs.c,v 1.13 2005/07/11 17:56:16 jirka Exp $
	//
	// Computes r = A.x - b
	// * @param a matrix
	// * @param x results
	// * @param b right-side
	// * @param r computed residuum vector
	// * @return state value
	//
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

// femEqsCGwJ - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:88
func femEqsCGwJ(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {
	// Conjugate gradient method with Jacobi preconditioner
	// *  (for symetric matrices only!)
	// *  @param a      matrix
	// *  @param b      "load" vector
	// *  @param x      results (vector - given as first iteration)
	// *  @param eps    error (min.)
	// *  @param maxIt  max. number of iterations
	// *  @return state value
	//
	// Jacobi preconditioner (diag[A] ;-)
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
	// vector initialization
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&z))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&q))[:])
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&M))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		// memory allocation
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&r))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&z))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&p))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&q))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	{
		// Jacobi preconditioner creation:
		for i = 1; i <= n; i++ {
			M.data[i-1] = femMatGet(a, i, i)
			if math.Abs(M.data[i-1]) < 1e-07 {
				rv = -13
				goto memFree
			}
		}
	}
	// next two lines mean: r = b - A*x
	femMatVecMultBig(a, x, (*[1000000]tVector)(unsafe.Pointer(&r))[:])
	for i = 0; i < n; i++ {
		r.data[i] = b[0].data[i] - r.data[i]
	}
	{
		// main loop
		for i = 1; i <= maxIt; i++ {
			{
				// using preconditioner:
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
			// Convergence testing
			normRes = femVecNormBig((*[1000000]tVector)(unsafe.Pointer(&r))[:])
			normX = femVecNormBig(x)
			if normRes <= eps*(normA*normX+normB) {
				// convergence test
				//if (fabs(norm - norm0) < eps )
				converged = 1
				break
			}
			roro = ro
		}
	}
	if converged != 1 {
		// end of main loop
		//fprintf(msgout,"[I] normRes = %f\n",normRes);
		rv = -1
	}
memFree:
	;
	// freeing memory:
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&z))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecFree((*[1000000]tVector)(unsafe.Pointer(&q))[:])
	return rv
}

// femEqsBiCCSwJ - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:262
func femEqsBiCCSwJ(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {
	// Bi-Conjugate Gradient Stabilized Method with Jacobi preconditioner
	// *  (for symetric and non-symetric matrices)
	// *  @param a      matrix
	// *  @param b      "load" vector
	// *  @param x      results (vector - given as first iteration)
	// *  @param eps    error (min.)
	// *  @param maxIt  max. number of iterations
	// *  @return state value
	// *
	// *  Note: "res" is probably useless and *NormBig(res) can be replaced by *NormBig(r).
	// *  Test it!!
	// *
	//
	// preconditioner (diag[a])
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
	// size of matrix "a"
	var n int32
	var converged int32
	// residuum
	var res tVector
	// norms
	var normRes float64
	var normX float64
	var normA float64
	var normB float64
	var rv int32
	_ = rv
	n = a[0].rows
	normA = femMatNormBig(a)
	normX = femVecNormBig(x)
	normB = femVecNormBig(b)
	// vector initialization
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
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&M))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		// memory allocation
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&r))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&rr))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&p))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&pp))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&s))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&ss))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&t))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&v))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&res))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	{
		// Jacobi preconditioner creation:
		for i = 1; i <= n; i++ {
			M.data[i-1] = femMatGet(a, i, i)
			if math.Abs(M.data[i-1]) < 1e-07 {
				rv = -13
				goto memFree
			}
		}
	}
	// next two lines mean: r = b - A*x
	femMatVecMultBig(a, x, (*[1000000]tVector)(unsafe.Pointer(&r))[:])
	for i = 0; i < n; i++ {
		r.data[i] = b[0].data[i] - r.data[i]
		rr.data[i] = r.data[i]
	}
	if femVecNormBig((*[1000000]tVector)(unsafe.Pointer(&r))[:]) <= 1e-07 {
		// convergence test
		converged = 1
		goto memFree
	}
	{
		// main loop
		for i = 1; i <= maxIt; i++ {
			ro = femVecVecMultBig((*[1000000]tVector)(unsafe.Pointer(&rr))[:], (*[1000000]tVector)(unsafe.Pointer(&r))[:])
			if math.Abs(ro) <= 0 {
				goto memFree
			}
			if i == 1 {
				// in first iteration
				for j = 0; j < n; j++ {
					p.data[j] = r.data[j]
				}
			} else {
				// int all iterations except first
				beta = ro / roro * (alpha / omega)
				for j = 0; j < n; j++ {
					p.data[j] = r.data[j] + beta*(p.data[j]-omega*v.data[j])
				}
			}
			{
				// using preconditioner M.pp=p -> pp
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
					// test of "s" size
					for j = 0; j < n; j++ {
						x[0].data[j] += alpha * pp.data[j]
					}
				}
				converged = 1
				break
			}
			{
				// using preconditioner M.ss=s -> ss
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
			// Convergence testing
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
		// end of main loop
		{
		}
	}
memFree:
	;
	// freeing of memory:
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&rr))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&pp))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&s))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&ss))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&t))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&v))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&res))[:])
	return 0
}

// femEqsLU - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:482
func femEqsLU(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {
	// Solver that uses LU - for full matrices!
	// *  @param a      matrix
	// *  @param b      "load" vector
	// *  @param x      results (vector - given as first iteration)
	// *  @param eps    error (min.)
	// *  @param maxIt  max. number of iterations
	// *  @return state value
	// *
	// *  Note: "res" is probably useless and *NormBig(res) can be replaced by *NormBig(r).
	// *  Test it!!
	// *
	//
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
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&indx))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femLUdecomp(a, (*[1000000]tVector)(unsafe.Pointer(&indx))[:])
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femLUback(a, (*[1000000]tVector)(unsafe.Pointer(&indx))[:], b)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// memFree:
	// 	;
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&indx))[:])
	return rv
}

// femEqsPCGwJ - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:514
func femEqsPCGwJ(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {
	// Alternative version of the Conjugate Gradient Method ()
	var rv int32
	var converged int32
	_ = converged
	var nui float64
	var dei float64
	var lambda float64
	var alpha float64
	// norms
	var normRes float64
	var normX float64
	var normA float64
	var normB float64
	var p tVector
	var r tVector
	var d tVector
	// Jacobi preconditioner
	var M tVector
	// a*p result vector
	var ap tVector
	// number of rows
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
	// vector initializations
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&d))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&ap))[:])
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&p))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		// memory allocation
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&r))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&d))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&M))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&ap))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	{
		// Jacobi preconditioner
		for i = 1; i <= n; i++ {
			M.data[i-1] = femMatGet(a, i, i)
			if math.Abs(M.data[i-1]) < 1e-07 {
				rv = -13
				goto memFree
			}
		}
	}
	// next several lines mean: r = b - A*x
	femMatVecMultBig(a, x, (*[1000000]tVector)(unsafe.Pointer(&r))[:])
	for i = 0; i < n; i++ {
		r.data[i] = b[0].data[i] - r.data[i]
	}
	{
		// using preconditioner:
		for j = 0; j < n; j++ {
			d.data[j] = r.data[j] / M.data[j]
			p.data[j] = d.data[j]
		}
	}
	for i = 1; i <= maxIt; i++ {
		// untested code follows...
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
			// convergence test
			converged = 1
			break
		}
		alpha = nui / dei
		for j = 0; j < n; j++ {
			p.data[j] = d.data[j] + alpha*p.data[j]
		}
	}
	// end of "for i"
	femVecPrn(x, ("X"))
memFree:
	;
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&d))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&ap))[:])
	return rv
}

// femMatCholFact - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:660
func femMatCholFact(a []tMatrix, C []tVector) int32 {
	// Choleski decomposition - forward run only!
	// * @param a matrix (must  be a MAT_FULL)
	// * @return status
	//
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
		//	if
		femVecAlloc(C, 0, n, n) //!= 0 {
		//		goto memFree
		//	}
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
	femVecPrn(C, ("C"))
memFree:
	;
	if have_C == 0 {
		// freeing of memory:
		femVecFree(C)
	}
	return rv
}

// femEqsChol - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:759
func femEqsChol(a []tMatrix, b []tVector, x []tVector) int32 {
	// Choleski decomposition - complete
	// * @param a matrix (must  be a MAT_FULL)
	// * @return status
	//
	var rv int32
	var sum float64
	var n int32
	var i int32
	var j int32
	var k int32
	var C tVector
	n = a[0].rows
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&C))[:])
	//	if
	femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&C))[:], 0, n, n)
	// 	!= 0 {
	// 		goto memFree
	// 	}
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
		// backward run:
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
	// freeing of memory:
	// femVecFree((*[1000000]tVector)(unsafe.Pointer(&C))[:])
	return rv
}

// femMatJacRotate - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:832
func femMatJacRotate(a []tMatrix, i int32, j int32, k int32, l int32, g float64, h float64, s float64, tau float64) {
	// rotation for Jacobi computation of eigenvalues
	g = femMatGet(a, i, j)
	h = femMatGet(a, k, l)
	femMatPutAdd(a, i, j, g-s*(h+g*tau), 0)
	femMatPutAdd(a, k, l, h+s*(g-h*tau), 0)
}

// femMatEigenJacobi - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:848
func femMatEigenJacobi(a []tMatrix, d []tVector, v []tMatrix, nrot []int32) int32 {
	// Compute eigen numbers and vectors (Jacobi method)
	// * @param a matrix to be analysed
	// * @param d vector to store eigenvalues
	// * @param v matrix to store eigenvectors
	// * @return status
	//
	var iters int32 = 100
	var i int32
	var iq int32
	var ip int32
	var j int32
	var n int32
	var sm float64
	var tresh float64
	_ = tresh
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
			// sum <= 0 so we are finished
			//printf("iterations: %v\n", *nrot);
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
					// off-diagonal elements are small
					femMatPutAdd(a, ip, iq, 0, 0)
				} else {
					// still are big..
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

// femEqsCGwSSOR - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:1001
func femEqsCGwSSOR(a []tMatrix, b []tVector, x []tVector, eps float64, maxIt int32) int32 {
	// Conjugate gradient method with SSOR preconditioner
	// *  (for symetric matrices only!)
	// *  @param a      matrix
	// *  @param b      "load" vector
	// *  @param x      results (vector - given as first iteration)
	// *  @param eps    error (min.)
	// *  @param maxIt  max. number of iterations
	// *  @return state value
	//
	// Jacobi preconditioner (diag[A] ;-)
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
	// vector initialization
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&z))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&zz))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	femVecNull((*[1000000]tVector)(unsafe.Pointer(&q))[:])
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&M))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		// memory allocation
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&r))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&z))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&zz))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&p))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	// 	if (func() int32 {
	rv = femVecAlloc((*[1000000]tVector)(unsafe.Pointer(&q))[:], 0, n, n)
	// 		return rv
	// 	}()) != 0 {
	// 		goto memFree
	// 	}
	{
		// Jacobi preconditioner creation:
		for i = 1; i <= n; i++ {
			val = femMatGet(a, i, i)
			if math.Abs(val) < 1e-07 {
				rv = -13
				goto memFree
			}
			// NOTE: M includes inverse of diagonal
			M.data[i-1] = 1 / val
		}
	}
	// next two lines mean: r = b - A*x
	femMatVecMultBig(a, x, (*[1000000]tVector)(unsafe.Pointer(&r))[:])
	for i = 0; i < n; i++ {
		r.data[i] = b[0].data[i] - r.data[i]
	}
	{
		// main loop
		for i = 1; i <= maxIt; i++ {
			if a[0].type_ != 1 {
				{
					// using preconditioner:
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
					// faster code for MAT_SPAR:
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
			// end of preconditioning
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
			// Convergence testing
			normRes = femVecNormBig((*[1000000]tVector)(unsafe.Pointer(&r))[:])
			normX = femVecNormBig(x)
			if normRes <= eps*(normA*normX+normB) {
				// convergence test
				//if (fabs(norm - norm0) < eps )
				converged = 1
				break
			}
			roro = ro
		}
	}
	if converged != 1 {
		// end of main loop
		//fprintf(msgout,"[I] normRes = %f\n",normRes);
		rv = -1
	}
memFree:
	;
	// freeing memory:
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&M))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&r))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&z))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&zz))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&p))[:])
	// 	femVecFree((*[1000000]tVector)(unsafe.Pointer(&q))[:])
	return rv
}

// c4goUnsafeConvert_float64 : created by c4go
func c4goUnsafeConvert_float64(c4go_name *float64) []float64 {
	return (*[1000000]float64)(unsafe.Pointer(c4go_name))[:]
}

// c4goUnsafeConvert_int32 : created by c4go
func c4goUnsafeConvert_int32(c4go_name *int32) []int32 {
	return (*[1000000]int32)(unsafe.Pointer(c4go_name))[:]
}

// 0 = dense; 1 = sparse (rows)
// lenght of "pos" and "data" (if used) fields
// Functions:
// Use with care:  (!!)
// end of fem_math.h
// end of eshell.c
// end of fem_math.c
// end of fem_eqs.c
// end of fem_mem.c
