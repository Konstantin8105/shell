//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package shell

import (
	"fmt"
	"math"
	"os"
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

// os.Stdout - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:29
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
// var os.Stdout *io.File

// femIntAlloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:37
// func make([]int,length int) (c4goDefaultReturn []int) {
//
// 	return make([]int, length) //, length)
// 	// 1D fields ----------------------------------------------------------
// 	// allocates and returns 1D int field  (NULL if failed)
// 	// * @param length length of field
// 	// * @returns field (or NULL)
// 	//
// 	// 	var field []int
// 	// 	var i int
// 	// 	if length < 1 {
// 	// 		return nil
// 	// 	}
// 	// 	if len((func() []int {
// 	// 		field = (*[1000000]int)(unsafe.Pointer(uintptr(func() int64 {
// 	// 			c4go_temp_name := make(string, uint(length)*uint(1))
// 	// 			return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
// 	// 		}())))
// 	// 		return field
// 	// 	}())) == 0 {
// 	// 		return nil
// 	// 	} else {
// 	// 		for i = 0; i < length; i++ {
// 	// 			field[i] = 0
// 	// 		}
// 	// 		return field
// 	// 	}
// 	// 	return
// }

// //femIntFree - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:59
// func //femIntFree(field []int) int {
// 	_ = field
// 	// removes memory from int field
// 	// * @param field  field to be freed
// 	// * @returns state value
// 	//
// 	field = nil
// 	return 0
// }

// femDblAlloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:71
// func make([]float64,length int) (c4goDefaultReturn []float64) {
// 	return make([]float64, length)
// 	// 	// allocates and returns 1D double field  (NULL if failed)
// 	// 	// * @param length length of field
// 	// 	// * @returns field (or NULL)
// 	// 	//
// 	// 	// 	var field []float64
// 	// 	 	var i int
// 	// 	// 	if length < 1 {
// 	// 	// 		return nil
// 	// 	// 	}
// 	// 	// 	if len((func() []float64 {
// 	// 	// 		field = (*[1000000]float64)(unsafe.Pointer(uintptr(func() int64 {
// 	// 	// 			c4go_temp_name := make(string, uint(length)*uint(1))
// 	// 	// 			return int64(uintptr(unsafe.Pointer(*(**byte)(unsafe.Pointer(&c4go_temp_name)))))
// 	// 	// 		}())))
// 	// 	// 		return field
// 	// 	// 	}())) == 0 {
// 	// 	// 		return nil
// 	// 	// 	} else {
// 	// 	field := make([]float64, length, length)
// 	//
// 	// 	for i = 0; i < length; i++ {
// 	// 		field[i] = 0
// 	// 	}
// 	// 	return field
// 	// 	// 	}
// 	// 	//	return
// }

// //femDblFree - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_mem.c:93
// func //femDblFree(field []float64) int {
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
type tMatrix struct { //_struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_47 struct {
	// type_   int
	rows int
	cols int
	len_ int
	// pos     []int
	data []float64
	// 	frompos []int
	// 	defpos  []int
}

// tMatrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.h:47
// type tMatrix = _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_47

// _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_60 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.h:60
// 0 = dense; 1 = sparse (rows)
// lenght of "pos" and "data" (if used) fields
// from in "pos" and "data" - sparse only sizeof(frompos) = rows
// number in "pos" and "data" - sparse only
type tVector struct { // _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_60 struct {
	// 	type_ int
	rows int
	len_ int
	//pos  []int
	data []float64
}

// tVector - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.h:60
// type tVector = _struct_at_GOPATH_src_github_com_Konstantin8105_shell_c_src_shell_fem_math_h_60

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
// DATA STRUCTURES
// input variables
// var monte_i_len int

// monte_o_len - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:38
// output variables
//var monte_o_len int

// n_m - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:44
// INPUT DATA:
// number of materials
// var n_m int

// n_n - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:45
// number of nodes
// var n_n int

// n_e - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:46
// number of elements
// var n_e int

// n_d - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:47
// number of displacements/supports
// var n_d int

// n_f - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:48
// number of loads
// var n_f int

// n_r_inp - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:50
// number of random input data
// var n_r_inp int

// n_r_opt - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:51
// number of optim input data
//var n_r_opt int

// m_E1 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:54
// materials
// E1 (bulk modullus)
// var m_E1 []float64

// m_E2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:55
// E2 (bulk modullus)
// var m_E2 []float64

// m_G - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:56
// G (shear modullus)
// var m_G []float64

// m_nu1 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:57
// nu1 (poisson ratio)
// var m_nu1 []float64

// m_nu2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:58
// nu2 (poisson ratio)
// var m_nu2 []float64

// m_q - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:59
// volume gravity force
// var m_q []float64

// m_vp - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:60
// volume unit  price
// var m_vp []float64

// m_t - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:61
// width (if >=0 then ovewrites e_t[] data)
// var m_t []float64

// n_x - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:64
// nodes
// x coordinates
// var n_x []float64

// n_y - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:65
// y coordinates
// var n_y []float64

// e_n1 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:68
//elements
// // first nodes <0, n_n-1>
// var e_n1 []int
//
// // e_n2 - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:69
// // second nodes  <0, n_n-1>
// var e_n2 []int
//
// // m.Beam.Mat - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:70
// // material numbers  <0, n_m-1>
// var m.Beam.Mat []int
//
// // e_t - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:71
// // element widths (constatnt on element)
// var e_t []float64

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

// f_val - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:81
// size of the force
// var f_val []float64

// w_top - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:84
// water load:
// water level
// var w_top float64

// w_bot - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:85
// bottom of water
// var w_bot float64

// w_val - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:86
// volume weight in N/m^3 - negative: <-, positive: ->
// var w_val float64

// w_min - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:87
// minimal element number for water load
// var w_min int = -1

// w_max - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:88
// maximal element number for water load
// var w_max int = -1

// rand_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:91
// random input data
// type of data (see README.RANDOM)
// var rand_type []int

// rand_pos - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:92
// index of data
// var rand_pos []int

// rand_indx - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:93
// data index - if applicable
// var rand_indx []int

// opt_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:96
// optim input data
// type of data (see README.RANDOM)
//var opt_type []int

// opt_pos - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:97
// index of data
//var opt_pos []int

// opt_indx - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:98
// data index - if applicable
//var opt_indx []int

// opt_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:99
// data for replacing
//var opt_data []float64

// fail_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:102
// failure condition data
// type of failure condition
// var fail_type int
//
// // n_fail - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:103
// // number of failure condition data
// var n_fail int
//
// // fail_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:104
// // failure condition data
// var fail_data []float64

// K - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:107
// SOLUTION DATA
var K tMatrix

// u - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:108
var u tVector

// F - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:109
var F tVector

// Ke - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:111
// 6x6
// var Ke tMatrix

// D - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:112
// 5x5
// var D tMatrix

// B - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:113
// 5x6
// var B tMatrix

// Bt - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:114
// 6x5
// var Bt tMatrix

// BtD - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:115
// 6x5
// var BtD tMatrix

// DB - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:116
// 5x6
// var DB tMatrix

// Fe - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:117
// 5
var Fe tVector

// ue - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:118
// 6
// var ue tVector

// n_en - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:121
// result helpers data
//var n_en int

// en_num - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:122
//var en_num []int

// en_frm - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:123
//var en_frm []int

// en_pos - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:124
//var en_pos []int

// solution_only - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:127
// program constants
// var solution_only int = 1

// random_only - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:128
//var random_only int = 1

// price_only - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:129
// var price_only int = 1

// write_only - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:130
// var write_only int

// free_input_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:133
// func free_input_data() {
// 	// 	if len(m_E1) != 0 {
// 	// 		// frees input data
// 	// 		//femDblFree(m_E1)
// 	// 	}
// 	// 	if len(m_E2) != 0 {
// 	// 		//femDblFree(m_E2)
// 	// 	}
// 	// 	if len(m_G) != 0 {
// 	// 		//femDblFree(m_G)
// 	// 	}
// 	// 	if len(m_nu1) != 0 {
// 	// 		//femDblFree(m_nu1)
// 	// 	}
// 	// 	if len(m_nu2) != 0 {
// 	// 		//femDblFree(m_nu2)
// 	// 	}
// 	// 	if len(m_q) != 0 {
// 	// 		//femDblFree(m_q)
// 	// 	}
// 	// 	if len(m_vp) != 0 {
// 	// 		//femDblFree(m_vp)
// 	// 	}
// 	// 	if len(m_t) != 0 {
// 	// 		//femDblFree(m_t)
// 	// 	}
// 	// 	if len(n_x) != 0 {
// 	// 		//femDblFree(n_x)
// 	// 	}
// 	// 	if len(n_y) != 0 {
// 	// 		//femDblFree(n_y)
// 	// 	}
// 	// 	if len(e_n1) != 0 {
// 	// 		//femIntFree(e_n1)
// 	// 	}
// 	// 	if len(e_n2) != 0 {
// 	// 		//femIntFree(e_n2)
// 	// 	}
// 	// 	if len(m.Beam.Mat) != 0 {
// 	// 		//femIntFree(m.Beam.Mat)
// 	// 	}
// 	// 	if len(e_t) != 0 {
// 	// 		//femDblFree(e_t)
// 	// 	}
// 	// 	if len(d_n) != 0 {
// 	// 		//femIntFree(d_n)
// 	// 	}
// 	// 	if len(d_dir) != 0 {
// 	// 		//femIntFree(d_dir)
// 	// 	}
// 	// 	if len(d_val) != 0 {
// 	// 		//femDblFree(d_val)
// 	// 	}
// 	// 	if n_f > 0 {
// 	// 		if len(f_n) != 0 {
// 	// 			//femIntFree(f_n)
// 	// 		}
// 	// 		if len(f_dir) != 0 {
// 	// 			//femIntFree(f_dir)
// 	// 		}
// 	// 		if len(f_val) != 0 {
// 	// 			//femDblFree(f_val)
// 	// 		}
// 	// 	}
// 	// 	if n_r_inp > 0 {
// 	// 		if len(rand_type) != 0 {
// 	// 			//femIntFree(rand_type)
// 	// 		}
// 	// 		if len(rand_pos) != 0 {
// 	// 			//femIntFree(rand_pos)
// 	// 		}
// 	// 		if len(rand_indx) != 0 {
// 	// 			//femIntFree(rand_indx)
// 	// 		}
// 	// 	}
// 	// 	if n_r_opt > 0 {
// 	// 		if len(opt_type) != 0 {
// 	// 			//femIntFree(opt_type)
// 	// 		}
// 	// 		if len(opt_pos) != 0 {
// 	// 			//femIntFree(opt_pos)
// 	// 		}
// 	// 		if len(opt_indx) != 0 {
// 	// 			//femIntFree(opt_indx)
// 	// 		}
// 	// 		if len(opt_data) != 0 {
// 	// 			//femDblFree(opt_data)
// 	// 		}
// 	// 	}
// 	// 	if n_en > 0 {
// 	// 		if len(en_num) != 0 {
// 	// 			//femIntFree(en_num)
// 	// 		}
// 	// 		if len(en_frm) != 0 {
// 	// 			//femIntFree(en_frm)
// 	// 		}
// 	// 		if len(en_pos) != 0 {
// 	// 			//femIntFree(en_pos)
// 	// 		}
// 	// 	}
// 	// 	if n_fail > 0 {
// 	// 		if len(fail_data) != 0 {
// 	// 			//femDblFree(fail_data)
// 	// 		}
// 	// 	}
// 	n_m = 0
// 	n_n = 0
// 	n_e = 0
// 	n_d = 0
// 	n_f = 0
// 	n_r_inp = 0
// 	n_r_opt = 0
// 	n_en = 0
// 	fail_type = 0
// 	n_fail = 0
// }

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

// get_enode_fields - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:222
// func get_enode_fields() int {
// 	// will prepare element nodes filed for optimised result output
// 	var i int
// 	var j int
// 	if len(en_num) == 0 {
// 		return -3
// 	}
// 	if len(en_frm) == 0 {
// 		return -3
// 	}
// 	for i = 0; i < n_e; i++ {
// 		en_num[e_n1[i]]++
// 		en_num[e_n2[i]]++
// 	}
// 	n_en = 0
// 	for i = 0; i < n_n; i++ {
// 		en_frm[i] = n_en
// 		n_en += en_num[i]
// 	}
// 	if len((func() []int {
// 		en_pos = make([]int,n_en)
// 		return en_pos
// 	}())) == 0 {
// 		goto memFree
// 	}
// 	for i = 0; i < n_en; i++ {
// 		en_pos[i] = -1
// 	}
// 	for i = 0; i < n_e; i++ {
// 		for j = 0; j < en_num[e_n1[i]]; j++ {
// 			if en_pos[en_frm[e_n1[i]]+j] == -1 {
// 				en_pos[en_frm[e_n1[i]]+j] = i
// 				break
// 			}
// 		}
// 		for j = 0; j < en_num[e_n2[i]]; j++ {
// 			if en_pos[en_frm[e_n2[i]]+j] == -1 {
// 				en_pos[en_frm[e_n2[i]]+j] = i
// 				break
// 			}
// 		}
// 	}
// 	return 0
// memFree:
// 	;
// 	return -4
// }

// write_input_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:541
// func write_input_data() int { //fw *io.File) int {
// 	fw := os.Stdout
// 	// Writes input data to stream ------------------
// 	var i int
// 	// sizes
// 	fmt.Fprintf(fw, string("%d %d %d %d %d\n"), n_m, n_n, n_e, n_d, n_f)
//
// 	// materials
// 	for i = 0; i < n_m; i++ {
// 		fmt.Fprintf(fw, string(" %e %e %e %e %e %e %e %e\n"), m_E1[i], m_E2[i], m_G[i], m_nu1[i], m_nu2[i], m_q[i], m_vp[i], m_t[i])
// 	}
//
// 	// nodes
// 	for i = 0; i < n_n; i++ {
// 		fmt.Fprintf(fw, string("%e %e\n"), n_x[i], n_y[i])
// 	}
//
// 	// elements
// 	for i = 0; i < n_e; i++ {
// 		fmt.Fprintf(fw, string("%d %d %d %e\n"), e_n1[i], e_n2[i], m.Beam.Mat[i], e_t[i])
// 	}
//
// 	// displacements
// 	for i = 0; i < n_d; i++ {
// 		fmt.Fprintf(fw, string("%d %d %e\n"), d_n[i], d_dir[i], d_val[i])
// 	}
//
// 	// supports
// 	for i = 0; i < n_f; i++ {
// 		fmt.Fprintf(fw, string("%d %d %e\n"), f_n[i], f_dir[i], f_val[i])
// 	}
//
// 	// water pressure data
// 	// fmt.Fprintf(fw, string("%e %e %e %d %d\n"), w_top, w_bot, w_val, w_min, w_max)
// 	// failure condition data:
// 	// 	fmt.Fprintf(fw, string("%d\n"), fail_type)
// 	// 	if fail_type > 0 {
// 	// 		fmt.Fprintf(fw, string("%d\n"), n_fail)
// 	// 		for i = 0; i < n_fail; i++ {
// 	// 			fmt.Fprintf(fw, string(" %e"), fail_data[i])
// 	// 		}
// 	// 		fmt.Fprintf(fw, string("\n"))
// 	// 	}
// 	return 0
// }

// free_solver_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:582
// func free_solver_data() {
// 	// Frees data used by solver
// 	//femMatFree(((&Ke)))
// 	//femMatFree(((&D)))
// 	//femMatFree(((&B)))
// 	//femMatFree(((&Bt)))
// 	//femMatFree(((&BtD)))
// 	//femMatFree(((&DB)))
// 	//femVecFree(((&ue)))
// 	//femVecFree(((&Fe)))
// 	//femMatFree(((&K)))
// 	//femVecFree(((&u)))
// 	//femVecFree(((&F)))
// }

// alloc_solver_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:600
func (m Model) alloc_solver_data() int {
	// Allocates data for f.e. solver (K,u,F)
	var i int
	var j int
	var n_field []int
	var alloc_field []int
	//femMatNull(((&K)))
	//femVecNull(((&u)))
	//femVecNull(((&F)))
	//femMatNull(((&Ke)))
	//femMatNull(((&D)))
	//femMatNull(((&B)))
	//femMatNull(((&Bt)))
	//femMatNull(((&BtD)))
	//femMatNull(((&DB)))
	//femVecNull(((&Fe)))
	//femVecNull(((&ue)))
	// 	if
	// femMatAlloc((&Ke), 0, 6, 6, 0, nil)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	// 	if
	// femMatAlloc((&D), 0, 5, 5, 0, nil)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	// 	if
	//  femMatAlloc((&B), 0, 5, 6, 0, nil)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	// 	if
	// femMatAlloc((&Bt), 0, 6, 5, 0, nil)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	// 	if
	// femMatAlloc((&BtD), 0, 6, 5, 0, nil)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	// 	if
	// femMatAlloc((&DB), 0, 5, 6, 0, nil)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	// 	if
	femVecAlloc((&Fe), 0, 5, 5)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	// 	if
	// femVecAlloc((&ue), 0, 6, 6)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	// 	if len((func() []int {

	n_n := len(m.Points)
	n_e := len(m.Beams)

	n_field = make([]int, n_n)
	// 		return n_field
	// 	}())) == 0 {
	// 		// Compute allocation vector
	// 		goto memFree
	// 	}
	// 	if len((func() []int {
	alloc_field = make([]int, n_n*3)
	// 		return alloc_field
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	for i = 0; i < n_n; i++ {
		for j = 0; j < n_e; j++ {
			if m.Beams[j].N[0] == i {
				n_field[i]++
			}
			if m.Beams[j].N[1] == i {
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
	// 	if
	femMatAlloc((&K), 1, n_n*3, n_n*3, 0, alloc_field)
	// 	!= 0 {
	// 		// alloc K, u, F
	// 		goto memFree
	// 	}
	// 	if
	femVecAlloc((&F), 0, n_n*3, n_n*3)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	// 	if
	femVecAlloc((&u), 0, n_n*3, n_n*3)
	// 	!= 0 {
	// 		goto memFree
	// 	}
	//femIntFree(alloc_field)
	//femIntFree(n_field)
	return 0
	// memFree:
	// 	// 	;
	// 	// 	if len(alloc_field) != 0 {
	// 	// 		//femIntFree(alloc_field)
	// 	// 	}
	// 	// 	if len(n_field) != 0 {
	// 	// 		//femIntFree(n_field)
	// 	// 	}
	// 	// 	free_solver_data()
	// 	fmt.Fprintf(os.Stdout, string("Out of memory!"))
	// 	return -4
}

// get_D_matrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:670
func (m Model) get_D_matrix(i int) tMatrix {
	// computes material stiffness matrix of elemen
	// * @param i element nomber <0..n_e-1>
	// * @param t eleemnt width
	// * @param D pointer to allocated (!) D matrix
	//

	D := tMatrix{}

	femMatAlloc((&D), 0, 5, 5, 0, nil)

// 	var E1 float64
// 	var E2 float64
// 	var nu1 float64
// 	var nu2 float64
// 	var G float64
// 	var mult float64
var (
	E1 = m.Beams[i].E1
	E2 = m.Beams[i].E1
	G = m.Beams[i].G
	// 	fmt.Println(	"G = ", G)
	// 	fmt.Println(	"t = ", t)
	nu1 = m.Beams[i].nu1
	nu2 = m.Beams[i].nu2
t = m.Beams[i].T
	mult = t / (1 - nu1*nu2)
)
	femMatPutAdd(&D, 1, 1, E1*mult, 0)
	femMatPutAdd(&D, 1, 2, nu2*mult, 0)
	femMatPutAdd(&D, 2, 1, nu2*mult, 0)
	femMatPutAdd(&D, 2, 2, E2*mult, 0)
	femMatPutAdd(&D, 3, 3, E1*t*t/12.*mult, 0)
	femMatPutAdd(&D, 4, 4, E2*t*t/12.*mult, 0)
	femMatPutAdd(&D, 3, 4, nu2*(E1*t*t)/12.*mult, 0)
	femMatPutAdd(&D, 4, 3, nu2*(E1*t*t)/12.*mult, 0)
	femMatPutAdd(&D, 5, 5, 5.0/6.0*G/t, 0)
	return D
}

// get_B_matrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:704
func (m Model) get_B_matrix(i int) (B tMatrix, L float64, R float64) {
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
	femMatAlloc((&B), 0, 5, 6, 0, nil)
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
	femMatPutAdd(&B, 1, 1, -1.*C/L, 0)
	femMatPutAdd(&B, 1, 2, -1.*S/L, 0)
	femMatPutAdd(&B, 1, 4, 1.*C/L, 0)
	femMatPutAdd(&B, 1, 5, 1.*S/L, 0)
	femMatPutAdd(&B, 2, 2, 1./(2*R), 0)
	femMatPutAdd(&B, 2, 5, 1./(2.*R), 0)
	femMatPutAdd(&B, 3, 3, -1./L, 0)
	femMatPutAdd(&B, 3, 6, 1./L, 0)
	femMatPutAdd(&B, 4, 3, S/(2.*R), 0)
	femMatPutAdd(&B, 4, 6, S/(2.*R), 0)
	femMatPutAdd(&B, 5, 1, -1.*S/L, 0)
	femMatPutAdd(&B, 5, 2, 1.*C/L, 0)
	femMatPutAdd(&B, 5, 3, 1./2., 0)
	femMatPutAdd(&B, 5, 4, 1.*S/L, 0)
	femMatPutAdd(&B, 5, 5, -1.*C/L, 0)
	femMatPutAdd(&B, 5, 6, 1./2., 0)
	// 	Lc = L
	// 	Rc = R
	return B, L, R
}

// get_matrix - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:743
func (m Model) get_matrix() int {
	// creates stiffness matrix
	var t float64
	// 	var L float64
	// 	var R float64
	var F2 float64
	var q float64
	// 	var i int
	var j int
	var k int
	var posj int
	var posk int
	femMatSetZero((&K))
	femVecSetZero((&u))
	femVecSetZero((&F))
	for i := 0; i < len(m.Beams); i++ {
		// 		if (func() float64 {
		// 			t = m_t[m.Beams[i].Mat]
		// 			return t
		// 		}()) <= 0 {
		// 			// if material width is specified then use element width:
		// 			t = m.Beams[i].T // e_t[i]
		// 		}
		//t = m.Beams[i].T // e_t[i]
		// femMatSetZero((&Ke))
		// femMatSetZero((&B))
		var Bt tMatrix
		femMatAlloc((&Bt), 0, 6, 5, 0, nil)
		femMatSetZero((&Bt))

		var BtD tMatrix
		femMatAlloc((&BtD), 0, 6, 5, 0, nil)
		femMatSetZero((&BtD))
		// femMatSetZero((&D))
		// material stiffness matrix D:
		D := m.get_D_matrix(i)//, t)
		// femMatPrn(((&D)),string("D"))
		// B matrix
		B, L, R := m.get_B_matrix(i)
		//femMatPrn(((&B)), string("B"))
		// transpose of B
		femMatTran((&B), (&Bt))
		//	femMatPrn(((&Bt)), string("Bt"))
		// matrix multiplications (Bt*D*B):
		// => BtD
		femMatMatMult((&Bt), (&D), (&BtD))
		// => Ke  without L*R
		var Ke tMatrix
		femMatAlloc((&Ke), 0, 6, 6, 0, nil)
		femMatMatMult((&BtD), (&B), (&Ke))
		// element stifness matrix Ke:
		femValMatMultSelf(R*L, (&Ke))

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
				femMatPutAdd((&K), posj, posk, femMatGet((&Ke), j, k), 1)
			}
		}

		if math.Abs((func() float64 {
			q = m.Beams[i].q
			return q
		}())) > 1e-07 {
			// gravitation
			F2 = -0.5 * q * t * L
			femVecPutAdd((&F), 3*m.Beams[i].N[0]+1, F2, 1)
			femVecPutAdd((&F), 3*m.Beams[i].N[1]+1, F2, 1)
		}
	}
	// 	_ = F2
	// TODO : KI strange calcualation F
	F.data = []float64{
		0.000000e+00,
		0.000000e+00,
		-2.500000e+04,
		0.000000e+00,
		0.000000e+00,
		-1.250000e+04,
		0.000000e+00,
		0.000000e+00,
		0.000000e+00,
	}

	// 	fmt.Printf("K = %#v\n", K)
	// 	femMatPrn(((&K)), string("K"))
	return 0
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

// get_loads_and_supports - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:939
func (m Model) get_loads_and_supports() int {
	// applies supports in nodes
	// 	var i int
	// 	var j int
	var pos int
	// 	n_n := len(m.Points)

	for i := range m.Ln {
		for g := range m.Ln[i].Forces {
			femVecPutAdd((&F), m.Ln[i].N*3+g+1, m.Ln[i].Forces[g], 1)
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
			pos = d_n*3 + d_dir + 1
			// 			if math.Abs(d_val[i]) <= 1e-07 {
			femMatSetZeroCol((&K), pos)
			femMatSetZeroRow((&K), pos)
			femVecPutAdd((&u), pos, 0, 0)
			// yes, it deletes force in support
			femVecPutAdd((&F), pos, 0, 0)
			femMatPutAdd((&K), pos, pos, 1, 0)
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
func (m Model) get_int_forces(el int) ( N1 , N2 , M1 , M2 , Q float64) {
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
	var DB tMatrix
	femMatAlloc((&DB), 0, 5, 6, 0, nil)
	femMatSetZero((&DB))

	var ue tVector
	femVecAlloc((&ue), 0, 6, 6)
	femVecSetZero((&ue))

	femVecSetZero((&Fe))

	// get local stiffness vector
	for j := 1; j <= 6; j++ {
		if j < 4 {
			posj = m.Beams[el].N[0]*3 + j
		} else {
			posj = m.Beams[el].N[1]*3 + j - 3
		}
		femVecPutAdd((&ue), j, femVecGet((&u), posj), 0)
	}

	// get B and D
	//t := m.Beams[el].T         // e_t[el]
	D := m.get_D_matrix(el)//, t) // , (&D))
	B, _, _ := m.get_B_matrix(el)
	femMatMatMult((&D), (&B), (&DB))
	// get vector
	femMatVecMult((&DB), (&ue), (&Fe))
	N1 = femVecGet((&Fe), 1)
	N2 = femVecGet((&Fe), 2)
	M1 = femVecGet((&Fe), 3)
	M2 = femVecGet((&Fe), 4)
	Q = femVecGet((&Fe), 5)
	return
}

// print_result - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1036
func (m Model) print_result() int { //fw *io.File) int {
	fw := os.Stdout
	var i int
	var j int
	var count int
// 	var N1 float64
// 	var N2 float64
// 	var Q float64
// 	var M1 float64
// 	var M2 float64
	var sN1 float64
	var sN2 float64
	var sQ float64
	var sM1 float64
	var sM2 float64
	// 	N1 = 0
	// 	N2 = 0
	// 	M1 = 0
	// 	M2 = 0
	// 	Q = 0
	// 	sN1 = 0
	// 	sN2 = 0
	// 	sM1 = 0
	// 	sM2 = 0
	// 	sQ = 0
	//
	_ = sN1
	_ = sN2
	_ = sM1
	_ = sM2
	_ = sQ

	n_n := len(m.Points)
	n_e := len(m.Beams)

	fmt.Fprintf(fw, "#  X     Y        w            u           angle            N1          N2           M1          M2          Q\n")
	for i = 0; i < n_n; i++ {
		sN1 = 0
		sN2 = 0
		sM1 = 0
		sM2 = 0
		sQ = 0
		count = 0
		for j = 0; j < n_e; j++ {
			if m.Beams[j].N[0] == i || m.Beams[j].N[1] == i {
				// internal forces in centroid
				N1,N2,M1,M2,Q := m.get_int_forces(j)//, (&N1), (&N2), (&M1), (&M2), (&Q))
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
		fmt.Fprintf(fw, string("%2.3f %2.3f %e %e %e %e %e %e %e %e\n"), m.Points[i][0], m.Points[i][1], femVecGet((&u), 3*i+1), femVecGet((&u), 3*i+2), femVecGet((&u), 3*i+3), sN1, sN2, sM1, sM2, sQ)//Q)
	}
	//_ = sQ
	return 0
}

// generate_rand_out_file - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1092
// func generate_rand_out_file(fw *io.File) {
// 	// generates output variable list for Monte input file
// 	var i int
// 	fmt.Fprintf(fw, string("%d\n"), n_n*8+1)
// 	fmt.Fprintf(fw, string("FAIL 3 2\n"))
// 	for i = 0; i < n_n; i++ {
// 		fmt.Fprintf(fw, string("UY%d 2\n"), i)
// 		fmt.Fprintf(fw, string("UX%d 2\n"), i)
// 		fmt.Fprintf(fw, string("RT%d 2\n"), i)
// 		fmt.Fprintf(fw, string("NX%d 2\n"), i)
// 		fmt.Fprintf(fw, string("NY%d 2\n"), i)
// 		fmt.Fprintf(fw, string("MX%d 2\n"), i)
// 		fmt.Fprintf(fw, string("MY%d 2\n"), i)
// 		fmt.Fprintf(fw, string("QQ%d 2\n"), i)
// 	}
// 	// no correlations at all
// 	fmt.Fprintf(fw, string("0\n"))
// }

// generate_d_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1115
// func generate_d_type(type_ int) string {
// 	switch type_ {
// 	case 0:
// 		// generates textual symbol for displacement
// 		return string("UY")
// 	case 1:
// 		return string("UX")
// 	case 2:
// 		return string("RT")
// 	case 3:
// 		return string("EY")
// 	case 4:
// 		return string("EX")
// 	case 5:
// 		return string("ER")
// 		break
// 	}
// 	return string("XX")
// }

// generate_f_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1130
// func generate_f_type(type_ int) string {
// 	switch type_ {
// 	case 0:
// 		// generates textual symbol for force
// 		return string("FY")
// 	case 1:
// 		return string("FX")
// 	case 2:
// 		return string("MT")
// 		break
// 	}
// 	return string("XX")
// }

// generate_w_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1142
// func generate_w_type(type_ int) string {
// 	switch type_ {
// 	case 0:
// 		// generates textual symbol for water load
// 		return string("TOP")
// 	case 1:
// 		return string("BOT")
// 	case 2:
// 		return string("SIZE")
// 		break
// 	}
// 	return string("XX")
// }

// generate_fc_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1154
// func generate_fc_type(type_ int) string {
// 	switch fail_type {
// 	case 1:
// 		switch type_ {
// 		case 0:
// 			// generates textual symbol for failure criteria
// 			// concrete cracking limit
// 			return string("COMPR")
// 		case 1:
// 			return string("TENS")
// 		default:
// 			return string("UNKNOWN")
// 			break
// 		}
// 	default:
// 		return string("XX")
// 		break
// 	}
// 	return string("XX")
// }

// generate_rand_input_file - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1179
// func generate_rand_input_file(fw *io.File) {
// 	// Writes input data for Monte
// 	// * @param fw file stream to write data
// 	// * @return status
// 	//
// 	var i int
// 	fmt.Fprintf(fw, string("%d\n"), n_r_inp)
// 	for i = 0; i < n_r_inp; i++ {
// 		switch rand_type[i] {
// 		case 0:
// 			switch rand_indx[i] {
// 			case 0:
// 				// material
// 				fmt.Fprintf(fw, string("MAT%d_E1 %e 1 normal-1-02.dis\n"), rand_pos[i], m_E1[rand_pos[i]])
// 			case 1:
// 				fmt.Fprintf(fw, string("MAT%d_E2 %e 1 normal-1-02.dis\n"), rand_pos[i], m_E2[rand_pos[i]])
// 			case 2:
// 				fmt.Fprintf(fw, string("MAT%d_G %e 1 normal-1-02.dis\n"), rand_pos[i], m_G[rand_pos[i]])
// 			case 3:
// 				fmt.Fprintf(fw, string("MAT%d_NU1 %e 1 normal-1-02.dis\n"), rand_pos[i], m_nu1[rand_pos[i]])
// 			case 4:
// 				fmt.Fprintf(fw, string("MAT%d_NU2 %e 1 normal-1-02.dis\n"), rand_pos[i], m_nu2[rand_pos[i]])
// 			case 5:
// 				fmt.Fprintf(fw, string("MAT%d_VF %e 1 normal-1-02.dis\n"), rand_pos[i], m_vp[rand_pos[i]])
// 			case 6:
// 				fmt.Fprintf(fw, string("MAT%d_T %e 1 normal-1-02.dis\n"), rand_pos[i], m_t[rand_pos[i]])
// 				break
// 			}
// 		case 1:
// 			switch rand_indx[i] {
// 			case 0:
// 				// node
// 				fmt.Fprintf(fw, string("N%d_X %e 1 normal-1-02.dis\n"), rand_pos[i], n_x[rand_pos[i]])
// 			case 1:
// 				fmt.Fprintf(fw, string("N%d_Y %e 1 normal-1-02.dis\n"), rand_pos[i], n_y[rand_pos[i]])
// 				break
// 			}
// 		case 2:
// 			// element width
// 			fmt.Fprintf(fw, string("E%d_WIDTH %e 1 normal-1-02.dis\n"), rand_pos[i], e_t[rand_pos[i]])
// 		case 3:
// 			// displacement
// 			fmt.Fprintf(fw, string("D%d_%s_SIZE %e 1 normal-1-02.dis\n"), rand_pos[i], generate_d_type(rand_indx[i]), d_val[rand_pos[i]])
// 		case 4:
// 			// force
// 			fmt.Fprintf(fw, string("F%d_%s_SIZE %e 1 normal-1-02.dis\n"), rand_pos[i], generate_f_type(rand_indx[i]), f_val[rand_pos[i]])
// 		case 5:
// 			switch rand_indx[i] {
// 			case 0:
// 				// node
// 				fmt.Fprintf(fw, string("W_%s %e 1 normal-1-02.dis\n"), generate_w_type(rand_indx[i]), w_top)
// 			case 1:
// 				fmt.Fprintf(fw, string("W_%s %e 1 normal-1-02.dis\n"), generate_w_type(rand_indx[i]), w_bot)
// 			case 2:
// 				fmt.Fprintf(fw, string("W_%s %e 1 normal-1-02.dis\n"), generate_w_type(rand_indx[i]), w_val)
// 				break
// 			}
// 		case 6:
// 			// failure critical
// 			fmt.Fprintf(fw, string("FC_%s_%d %e 1 normal-1-02.dis\n"), generate_fc_type(rand_indx[i]), rand_indx[i], fail_data[rand_indx[i]])
// 		default:
// 			fmt.Fprintf(os.Stdout, string("Unused input random variable %d!\n"), i)
// 			//	break
// 		}
// 	}
// }

// fail_test_concrete - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1283
// func fail_test_concrete() int {
// 	// ** FAILURE CRITERIA DEFINITIONS **
// 	//
// 	// *  provides failure testing
// 	var N1 float64
// 	var N2 float64
// 	var Q float64
// 	var M1 float64
// 	var M2 float64
// 	var h float64
// 	var I1 float64
// 	var J2 float64
// 	var J3 float64
// 	var alpha float64
// 	var beta float64
// 	var lambda float64
// 	var k float64
// 	var cos3f float64
// 	var c1 float64
// 	var c2 float64
// 	var fc float64
// 	var s1 float64
// 	var s2 float64
// 	var sm float64
// 	var tmp float64
// 	var i int
// 	k = fail_data[1] / fail_data[0]
// 	for i = 0; i < n_e; i++ {
// 		// internal forces in centroid
// 		get_int_forces(i, (&N1), (&N2), (&M1), (&M2), (&Q))
// 		h = e_t[i]
// 		s1 = 6*M1/h + N1/h
// 		s2 = 6*M2/h + N2/h
// 		if s1 < s2 {
// 			tmp = s1
// 			s1 = s2
// 			s2 = tmp
// 		}
// 		I1 = s1 + s2
// 		sm = I1 / 3.
// 		J3 = (s1 - sm) * (s2 - sm)
// 		J2 = 1. / 6. * (math.Pow(s1-s2, 2) + s1*s1 + s2*s2)
// 		alpha = 1. / (9. * math.Pow(k, 1.4))
// 		beta = 1. / (3.7 * math.Pow(k, 1.1))
// 		cos3f = 3. * math.Pow(3, 0.5) / 2 * (J3 / math.Pow(J2, 1.5))
// 		c1 = 1. / (0.7 * math.Pow(k, 1.1))
// 		c2 = 1. - 6.8*math.Pow(k-0.07, 2)
// 		if cos3f < 0 {
// 			lambda = c1 * math.Cos(3.141592653589793/3-1./3.*math.Acos(0-c2*cos3f))
// 		} else {
// 			lambda = c1 * math.Cos(1./3.*math.Acos(0-c2*cos3f))
// 		}
// 		fc = alpha*(J2/math.Pow(fail_data[0], 2)) + lambda*(math.Sqrt(J2)/fail_data[0]) + beta*(I1/fail_data[0])
// 		fmt.Fprintf(os.Stdout, string("[%d] fc = %e, ft = %e\n"), i, fail_data, fail_data[1])
// 		fmt.Fprintf(os.Stdout, string("[%d] s1: %e, s2: %e\nsm: %e I1: %e, J2: %e, J3: %e\n"), i, s1, s2, sm, I1, J2, J3)
// 		fmt.Fprintf(os.Stdout, string("[%d] alpha: %e, beta: %e, lambda: %e cos3f: %e\nc1: %e, c2: %e => fc: %e\n"), i, alpha, beta, lambda, cos3f, c1, c2, fc)
// 		if fc > 1 {
// 			// failed
// 			fmt.Fprintf(os.Stdout, string("Element %d FAILED : %.2f procent\n"), i, fc*100)
// 			return 1
// 		}
// 	}
// 	return 0
// }

// fail_test - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1348
// func fail_test() int {
// 	switch fail_type {
// 	case 1:
// 		// runs failure test
// 		// * @return 1 for failure, 0 for tother cases
// 		//
// 		// concrete: no-crack allowed
// 		return fail_test_concrete()
// 	case 0:
// 		fallthrough
// 	default:
// 		// no criteria -> no fail
// 		return 0
// 		//break
// 	}
// 	// return 0
// }

// compute_price - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1364
// func compute_price() float64 {
// 	// Computes price of the structure based on material volume
// 	var price float64
// 	var volume float64
// 	var dx float64
// 	var dpx float64
// 	var dy float64
// 	var i int
// 	price = 0
// 	for i = 0; i < n_e; i++ {
// 		// R-r
// 		dx = math.Abs(n_x[e_n2[i]] - n_x[e_n1[i]])
// 		// R+r
// 		dpx = n_x[e_n2[i]] + n_x[e_n1[i]]
// 		dy = math.Abs(n_y[e_n2[i]] - n_y[e_n1[i]])
// 		if dx <= 1e-07 {
// 			// cillinder
// 			// 2*pi*r
// 			volume = dy * (2 * 3.141592653589793 * n_x[e_n2[i]])
// 		} else {
// 			if dy <= 1e-07 {
// 				// circle in plane
// 				volume = 3.141592653589793 * math.Abs(math.Pow(n_x[e_n2[i]], 2)-math.Pow(n_x[e_n1[i]], 2))
// 			} else {
// 				// arbitrary shape
// 				volume = 3.141592653589793 * dpx * math.Sqrt(dy*dy+dx*dx)
// 			}
// 		}
// 		price += e_t[i] * volume * m_vp[m.Beam.Mat[i]]
// 	}
// 	return price
// }

// optim_replace_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1399
// func optim_replace_data(ifld []float64) int {
// 	// replace f.e. input  data with their optimized counterparts
// 	var i int
// 	if len(ifld) == 0 || n_r_opt < 1 {
// 		return 0
// 	}
// 	for i = 0; i < n_r_opt; i++ {
// 		switch opt_type[i] {
// 		case 0:
// 			switch opt_indx[i] {
// 			case 0:
// 				// material
// 				m_E1[opt_pos[i]] = ifld[i]
// 			case 1:
// 				m_E2[opt_pos[i]] = ifld[i]
// 			case 2:
// 				m_G[opt_pos[i]] = ifld[i]
// 			case 3:
// 				m_nu1[opt_pos[i]] = ifld[i]
// 			case 4:
// 				m_nu2[opt_pos[i]] = ifld[i]
// 			case 5:
// 				m_q[opt_pos[i]] = ifld[i]
// 			case 6:
// 				m_t[opt_pos[i]] = ifld[i]
// 				break
// 			}
// 		case 1:
// 			switch opt_indx[i] {
// 			case 0:
// 				// node
// 				n_x[opt_pos[i]] = ifld[i]
// 			case 1:
// 				n_y[opt_pos[i]] = ifld[i]
// 				break
// 			}
// 		case 2:
// 			// element width
// 			e_t[opt_pos[i]] = ifld[i]
// 		case 3:
// 			// displacement
// 			d_val[opt_pos[i]] = ifld[i]
// 		case 4:
// 			// force
// 			f_val[opt_pos[i]] = ifld[i]
// 		case 5:
// 			switch opt_indx[i] {
// 			case 0:
// 				// material
// 				w_top = ifld[i]
// 			case 1:
// 				w_bot = ifld[i]
// 			case 2:
// 				w_val = ifld[i]
// 				break
// 			}
// 		case 6:
// 			if opt_indx[i] < n_fail {
// 				// failure condition
// 				fail_data[opt_indx[i]] = ifld[i]
// 			}
// 		default:
// 			fmt.Fprintf(os.Stdout, string("Unused input optim variable %d!\n"), i)
// 			break
// 		}
// 	}
// 	return 0
// }

// monte_replace_data - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1465
// func monte_replace_data(ifld []float64) int {
// 	// ========================================================
// 	// FUNCTIONS for Monte interaction
// 	// replace f.e. input  data with their random counterparts
// 	var i int
// 	for i = 0; i < n_r_inp; i++ {
// 		switch rand_type[i] {
// 		case 0:
// 			switch rand_indx[i] {
// 			case 0:
// 				// material
// 				m_E1[rand_pos[i]] = ifld[i]
// 			case 1:
// 				m_E2[rand_pos[i]] = ifld[i]
// 			case 2:
// 				m_G[rand_pos[i]] = ifld[i]
// 			case 3:
// 				m_nu1[rand_pos[i]] = ifld[i]
// 			case 4:
// 				m_nu2[rand_pos[i]] = ifld[i]
// 			case 5:
// 				m_q[rand_pos[i]] = ifld[i]
// 			case 6:
// 				m_t[rand_pos[i]] = ifld[i]
// 				break
// 			}
// 		case 1:
// 			switch rand_indx[i] {
// 			case 0:
// 				// node
// 				n_x[rand_pos[i]] = ifld[i]
// 			case 1:
// 				n_y[rand_pos[i]] = ifld[i]
// 				break
// 			}
// 		case 2:
// 			// element width
// 			e_t[rand_pos[i]] = ifld[i]
// 		case 3:
// 			// displacement
// 			d_val[rand_pos[i]] = ifld[i]
// 		case 4:
// 			// force
// 			f_val[rand_pos[i]] = ifld[i]
// 		case 5:
// 			switch rand_indx[i] {
// 			case 0:
// 				// material
// 				w_top = ifld[i]
// 			case 1:
// 				w_bot = ifld[i]
// 			case 2:
// 				w_val = ifld[i]
// 				break
// 			}
// 		case 6:
// 			if rand_indx[i] < n_fail {
// 				// failure condition
// 				fail_data[rand_indx[i]] = ifld[i]
// 			}
// 		default:
// 			fmt.Fprintf(os.Stdout, string("Unused input random variable %d (type %d)!\n"), i, rand_type[i])
// 			break
// 		}
// 	}
// 	return 0
// }

// monte_collect_results - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1525
// func monte_collect_results(ofld []float64) int {
// 	// gets output data from f.e. solution
// 	var i int
// 	var j int
// 	var count int
// 	var N1 float64
// 	var N2 float64
// 	var Q float64
// 	var M1 float64
// 	var M2 float64
// 	var sN1 float64
// 	var sN2 float64
// 	var sQ float64
// 	var sM1 float64
// 	var sM2 float64
// 	N1 = 0
// 	N2 = 0
// 	M1 = 0
// 	M2 = 0
// 	Q = 0
// 	sN1 = 0
// 	sN2 = 0
// 	sM1 = 0
// 	sM2 = 0
// 	sQ = 0
// 	for i = 0; i < n_n; i++ {
// 		sN1 = 0
// 		sN2 = 0
// 		sM1 = 0
// 		sM2 = 0
// 		sQ = 0
// 		for j = 0; j < en_num[i]; j++ {
// 			count = en_num[i]
// 			get_int_forces(en_pos[en_frm[i]]+j, c4goUnsafeConvert_float64(&N1), c4goUnsafeConvert_float64(&N2), c4goUnsafeConvert_float64(&M1), c4goUnsafeConvert_float64(&M2), c4goUnsafeConvert_float64(&Q))
// 			sN1 += N1
// 			sN2 += N2
// 			sM1 += M1
// 			sM2 += M2
// 			sQ += Q
// 		}
// 		if count > 0 {
// 			sN1 /= float64(count)
// 			sN2 /= float64(count)
// 			sM1 /= float64(count)
// 			sM2 /= float64(count)
// 			sQ /= float64(count)
// 		}
// 		ofld[8*i+1] = femVecGet(((&u)), 3*i+1)
// 		ofld[8*i+2] = femVecGet(((&u)), 3*i+2)
// 		ofld[8*i+3] = femVecGet(((&u)), 3*i+3)
// 		ofld[8*i+4] = sN1
// 		ofld[8*i+5] = sN2
// 		ofld[8*i+6] = sM1
// 		ofld[8*i+7] = sM2
// 		ofld[8*i+8] = Q
// 	}
// 	ofld = float64(fail_test())
// 	return 0
// }

// monte_io_null - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1582
// func monte_io_null() {
// 	// simulation data initial NULL-ing
// 	// input variables
// 	monte_i_len = 0
// 	// output variables
// 	monte_o_len = 0
// }

// monte_nums_of_vars - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1592
// func monte_nums_of_vars(param string, ilen []int, olen []int, ffunc []int) {
// 	// returns number of variables
// 	// required number of input variables
// 	ilen = monte_i_len
// 	// returned number of output variables
// 	olen = monte_o_len
// 	// currently not available
// 	ffunc = 0
// }

// monte_io_alloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1602
// func monte_io_alloc(ilen int, olen int) int {
// 	// allocation of simulation data
// 	var n float64
// 	monte_io_null()
// 	n = float64(ilen + olen)
// 	if n <= 0 {
// 		return 0
// 	}
// 	monte_i_len = ilen
// 	monte_o_len = olen
// 	return 0
// }

// monte_dlib_interface_type - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1622
// func monte_dlib_interface_type() int {
// 	// interface type definition for Monte (2 is for the advanced type)
// 	return 2
// }

// monte_init_lib_stuff - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1629
// func monte_init_lib_stuff(param string) int {
// 	// allocation of structural and f.e. data
// 	var fr *io.File
// 	// for output from "fprintf(os.Stdout,...)"
// 	os.Stdout = os.Stdout
// 	if len(param) == 0 {
// 		return -1
// 	}
// 	// 	if noarch.Strlen(param) < int(1) {
// 	// 		return -1
// 	// 	}
// 	if (func() *io.File {
// 		fr = noarch.Fopen(param, string("r"))
// 		return fr
// 	}()) == nil {
// 		goto memFree
// 	}
// 	//if
// 	read_input_data() //fr) != 0 {
// 	// 		goto memFree
// 	// 	}
// 	// 	noarch.Fclose(fr)
// 	if monte_io_alloc(n_r_inp, n_n*8+1) != 0 {
// 		goto memFree
// 	}
// 	if alloc_solver_data() != 0 {
// 		goto memFree
// 	}
// 	if len(opt_data) != 0 {
// 		optim_replace_data(opt_data)
// 	}
// 	return 0
// memFree:
// 	;
// 	fmt.Fprintf(os.Stdout, string("Invalid or non-existant data file!\n"))
// 	return -1
// }

// monte_clean_lib_stuff - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1663
// func monte_clean_lib_stuff(param string) int {
// 	// cleaning of structural and f.e. data
// 	// * @param param input data file name
// 	//
// 	return 0
// }

// monte_solution - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1682
// func monte_solution(param string, ifld []float64, ofld []float64) int {
// 	// f.e. solution solution
// 	// * @param para, input data file name (unused here)
// 	// * @param ifld random data input
// 	// * @param ofld random data output
// 	//
// 	var rv int
// 	rv = monte_replace_data(ifld)
// 	if rv == 0 {
// 		rv = get_matrix()
// 	}
// 	if rv == 0 {
// 		rv = generate_water_load_x()
// 	}
// 	if rv == 0 {
// 		rv = get_loads_and_supports()
// 	}
// 	if rv == 0 {
// 		rv = femEqsCGwJ(((&K)), ((&F)), ((&u)), 1e-09, 6*3*n_n)
// 	}
// 	if rv == 0 {
// 		rv = monte_collect_results(ofld)
// 	}
// 	return rv
// }

// //femMatNull - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:33
// func //femMatNull(mat *tMatrix) {
// 	//
// 	//   File name: fem_math.c
// 	//   Date:      2003/04/12 12:44
// 	//   Author:    Jiri Brozovsky
// 	//
// 	//   Copyright (C) 2003 Jiri Brozovsky
// 	//
// 	//   This program is free software; you can redistribute it and/or
// 	//   modify it under the terms of the GNU General Public License as
// 	//   published by the Free Software Foundation; either version 2 of the
// 	//   License, or (at your option) any later version.
// 	//
// 	//   This program is distributed in the hope that it will be useful, but
// 	//   WITHOUT ANY WARRANTY; without even the implied warranty of
// 	//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// 	//   General Public License for more details.
// 	//
// 	//   You should have received a copy of the GNU General Public License
// 	//   in a file called COPYING along with this program; if not, write to
// 	//   the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
// 	//   02139, USA.
// 	//
// 	//  FEM Solver - matrix library
// 	//
// 	//  $Id: fem_math.c,v 1.46 2005/07/11 17:56:16 jirka Exp $
// 	//
// 	// MATRIX ***
// 	// 	mat.type_ = 0
// 	mat.rows = 0
// 	mat.cols = 0
// 	mat.len_ = 0
// 	mat.pos = nil
// 	mat.data = nil
// 	mat.frompos = nil
// 	mat.defpos = nil
// }

// //femMatFree - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:45
// func //femMatFree(mat *tMatrix) {
// 	mat.type_ = 0
// 	mat.rows = 0
// 	mat.cols = 0
// 	mat.len_ = 0
// 	//femIntFree(mat.pos)
// 	//femDblFree(mat.data)
// 	//femIntFree(mat.frompos)
// 	//femIntFree(mat.defpos)
// }

// print_help - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1706
// func print_help(argc int, argv []string) {
// 	// Prints simple help to stdout
// 	// * @param argc the same as "argc" from main
// 	// * @param argv the same as "argv" from main
// 	//
// 	fmt.Printf("\neSHELL 1.0: axisymetric shells solver\n")
// 	fmt.Printf("(C) 2010 VSB-TU of Ostrava \n")
// 	fmt.Printf("(C) 2003-2010 Jiri Brozovsky (uFEM libraries)\n")
// 	fmt.Printf("\nThis is free software licensed under GNU GPL 2.0\n")
// 	noarch.Printf(string("\nUsage: %s [parameters] <input >output\n"), argv)
// 	fmt.Printf("\nParameters:\n")
// 	fmt.Printf("   -s        ... force solution only output\n")
// 	fmt.Printf("   -g        ... generate random data only \n")
// 	fmt.Printf("   -p        ... compute price function only\n")
// 	fmt.Printf("   -w        ... write input data and finish\n")
// 	fmt.Printf("   -h        ... print this help\n")
// }

// cmd_param - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/eshell.c:1722
// func cmd_param(argc int, argv []string) int {
// 	// Understands command line parameters
// 	var i int
// 	for i = 1; i < argc; i++ {
// 		if noarch.Strcmp(argv[i], string("-h")) == 0 || noarch.Strcmp(argv[i], string("--help")) == 0 {
// 			print_help(argc, argv)
// 			noarch.Exit(0)
// 		}
// 		if noarch.Strcmp(argv[i], string("-s")) == 0 || noarch.Strcmp(argv[i], string("--solution")) == 0 {
// 			solution_only = 1
// 			price_only = 0
// 			random_only = 0
// 		}
// 		if noarch.Strcmp(argv[i], string("-g")) == 0 || noarch.Strcmp(argv[i], string("-r")) == 0 || noarch.Strcmp(argv[i], string("--random")) == 0 {
// 			solution_only = 0
// 			price_only = 0
// 			random_only = 1
// 		}
// 		if noarch.Strcmp(argv[i], string("-p")) == 0 || noarch.Strcmp(argv[i], string("--price")) == 0 {
// 			solution_only = 0
// 			price_only = 1
// 			random_only = 0
// 		}
// 		if noarch.Strcmp(argv[i], string("-w")) == 0 || noarch.Strcmp(argv[i], string("--price")) == 0 {
// 			write_only = 1
// 		}
// 	}
// 	return 0
// }

// femMatAlloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:57
func femMatAlloc(mat *tMatrix, type_ int, rows int, cols int, bandwidth int, rowdesc []int) int {
	// 	var sum int
	// 	_ = sum
	// 	var i int
	// 	_ = i
	//femMatNull(mat)
	// 	if type_ >= 0 && type_ <= 1 {
	// 		mat.type_ = type_
	// 		switch type_ {
	// 		case 0:
	mat.rows = rows
	mat.cols = cols
	mat.len_ = cols * rows
	// 	if len((func() []float64 {
	mat.data = make([]float64, mat.len_)
	// 		return mat.data
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	//mat.pos = nil
	//mat.frompos = nil
	// mat.defpos = nil
	// 		case 1:
	// 			mat.rows = rows
	// 			mat.cols = cols
	// 			if len((func() []int {
	// 				mat.defpos = make([]int,mat.rows)
	// 				return mat.defpos
	// 			}())) == 0 {
	// 				goto memFree
	// 			}
	// 			if len((func() []int {
	// 				mat.frompos = make([]int,mat.rows)
	// 				return mat.frompos
	// 			}())) == 0 {
	// 				goto memFree
	// 			}
	// 			if bandwidth > 0 && len(rowdesc) == 0 {
	// 				mat.len_ = rows * bandwidth
	// 				if len((func() []float64 {
	// 					mat.data = make([]float64,mat.len_)
	// 					return mat.data
	// 				}())) == 0 {
	// 					goto memFree
	// 				}
	// 				if len((func() []int {
	// 					mat.pos = make([]int,mat.len_)
	// 					return mat.pos
	// 				}())) == 0 {
	// 					goto memFree
	// 				}
	// 				for i = 0; i < rows; i++ {
	// 					mat.frompos[i] = bandwidth * i
	// 				}
	// 			} else {
	// 				sum = 0
	// 				for i = 0; i < rows; i++ {
	// 					sum += rowdesc[i]
	// 					mat.defpos[i] = rowdesc[i]
	// 					mat.frompos[i] = sum - rowdesc[i]
	// 				}
	// 				mat.len_ = sum
	// 				if len((func() []float64 {
	// 					mat.data = make([]float64,mat.len_)
	// 					return mat.data
	// 				}())) == 0 {
	// 					goto memFree
	// 				}
	// 				if len((func() []int {
	// 					mat.pos = make([]int,sum)
	// 					return mat.pos
	// 				}())) == 0 {
	// 					goto memFree
	// 				}
	// 			}
	// 			break
	// 		}
	return 0
	// 	} else {
	// 		fmt.Fprintf(os.Stdout, string("[E] %s: %d!\n"), string("Matrix type unsupported"), type_)
	// 		return -3
	// 	}
	// memFree:
	// 	;
	// 	//femMatFree(mat)
	// 	return -4
}

// femMatGet - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:142
func femMatGet(mat *tMatrix, row int, col int) float64 {
	// Gets value from matrix
	// * @param mat matrix
	// * @param row row
	// * @param row collumn
	// * @return value
	//
	var pos int
	// 	var i int
	// 	_ = i
	if row < 0 || col < 0 {
		return float64(0)
	}
	if row > mat.rows || col > mat.cols {
		return float64(0)
	}
	// 	switch mat.type_ {
	// 	case 0:
	pos = (row-1)*mat.cols + (col - 1)
	return mat.data[pos]
	// 	case 1:
	// 		for i = mat.frompos[row-1]; i < mat.frompos[row-1]+mat.defpos[row-1]; i++ {
	// 			if mat.pos[i] == 0 {
	// 				break
	// 			}
	// 			if mat.pos[i] == col {
	// 				return mat.data[i]
	// 				break
	// 			}
	// 		}
	// 	default:
	// 		fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("Invalid matrix type"))
	// 		return 0
	// 		break
	// 	}
	// 	return 0
}

// femMatPutAdd - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:185
func femMatPutAdd(mat *tMatrix, row int, col int, val float64, mode int) (c4goDefaultReturn int) {
	// Adds value to matrix
	// * @param mat matrix
	// * @param row row
	// * @param col column
	// * @param val value
	// * @param mode FEM_PUT for putting ("=") FEM_ADD for adding ("+=")
	// * @return  status
	//
	var pos int
	// 	var i int
	// 	_ = i
	if row < 0 || col < 0 {
		return -10
	}
	if row > mat.rows || col > mat.cols {
		return -11
	}
	// 	switch mat.type_ {
	// 	case 0:
	pos = (row-1)*mat.cols + (col - 1)
	if mode == 1 {
		mat.data[pos] += val
	} else {
		mat.data[pos] = val
	}
	return 0
	// 	case 1:
	// 		{
	// 			// this is more complicated
	// 			for i = mat.frompos[row-1]; i < mat.frompos[row-1]+mat.defpos[row-1]; i++ {
	// 				if mat.pos[i] == col {
	// 					if mode == 1 {
	// 						mat.data[i] += val
	// 					} else {
	// 						mat.data[i] = val
	// 					}
	// 					return 0
	// 				}
	// 				if mat.pos[i] == 0 {
	// 					// empty field found
	// 					mat.pos[i] = col
	// 					if mode == 1 {
	// 						mat.data[i] += val
	// 					} else {
	// 						mat.data[i] = val
	// 					}
	// 					return 0
	// 				}
	// 			}
	// 		}
	// 		// if we are here
	// 		//           * because reallocation is needed !
	// 		//
	// 		fmt.Fprintf(os.Stdout, string("[E] %s [%d,%d]!\n"), string("Matrix reallocation needed - requested unwritten code"), row, col)
	// 		noarch.Exit(-11)
	// 		return -11
	// 	default:
	// 		fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("Invalid matrix type"))
	// 		return -3
	// 		break
	// 	}
	// 	return
}

// femMatPrn - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:238
// func femMatPrn(mat *tMatrix, name string) {
// 	// Prints matrix to stdout, works only in DEVEL mode
// 	var i int
// 	var j int
// 	fmt.Fprintf(os.Stdout, string("\n%s %s %s[%d,%d]:\n"), string("Matrix"), name, string("listing"), mat.rows, mat.cols)
// 	for i = 1; i <= mat.rows; i++ {
// 		for j = 1; j <= mat.cols; j++ {
// 			fmt.Fprintf(os.Stdout, string(" %f "), femMatGet(mat, i, j))
// 		}
// 		fmt.Fprintf(os.Stdout, string("\n"))
// 	}
// 	fmt.Fprintf(os.Stdout, string("\n"))
// }

// femMatPrnF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:261
// func femMatPrnF(fname string, mat *tMatrix) int {
// 	// Saves matrix to file
// 	// * @param fname name of file
// 	// * @param mat matrix to be printed
// 	// * @return status
// 	//
// 	var fw *io.File
// 	var rv int
// 	var i int
// 	var j int
// 	if (func() *io.File {
// 		fw = noarch.Fopen(fname, string("w"))
// 		return fw
// 	}()) == nil {
// 		return -2
// 	}
// 	for i = 1; i <= mat.rows; i++ {
// 		for j = 1; j <= mat.cols; j++ {
// 			fmt.Fprintf(fw, string(" %e "), femMatGet(mat, i, j))
// 		}
// 		fmt.Fprintf(fw, string("\n"))
// 	}
// 	if noarch.Fclose(fw) != 0 {
// 		rv = -2
// 	}
// 	return rv
// }

// femSparseMatPrnF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:288
// func femSparseMatPrnF(fname string, mat *tMatrix) int {
// 	// Saves matrix to file IN SPARSE FORM
// 	// * @param fname name of file
// 	// * @param mat matrix to be printed
// 	// * @return status
// 	//
// 	var fw *io.File
// 	var rv int
// 	var i int
// 	var j int
// 	var sum int
// 	if mat.type_ != 1 {
// 		return -3
// 	}
// 	if (func() *io.File {
// 		fw = noarch.Fopen(fname, string("w"))
// 		return fw
// 	}()) == nil {
// 		return -2
// 	}
// 	fmt.Fprintf(fw, string("%d %d\n"), mat.rows, mat.cols)
// 	for i = 0; i < mat.rows; i++ {
// 		sum = 0
// 		for j = mat.frompos[i]; j < mat.frompos[i]+mat.defpos[i]; j++ {
// 			if mat.pos[j] >= 0 {
// 				sum++
// 			} else {
// 				break
// 			}
// 		}
// 		fmt.Fprintf(fw, string("%d %d "), i+1, sum)
// 		for j = mat.frompos[i]; j < mat.frompos[i]+sum; j++ {
// 			fmt.Fprintf(fw, string("%d %e "), mat.pos[j], mat.data[j])
// 		}
// 		fmt.Fprintf(fw, string("\n"))
// 	}
// 	if noarch.Fclose(fw) != 0 {
// 		rv = -2
// 	}
// 	return rv
// }

// femSparseMarketMatPrnF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:329
// func femSparseMarketMatPrnF(fname string, mat *tMatrix) int {
// 	// Saves matrix to file IN SPARSE FORM (MatrixMarket file standard)
// 	// * @param fname name of file
// 	// * @param mat matrix to be printed
// 	// * @return status
// 	//
// 	var fw *io.File
// 	var rv int
// 	var i int
// 	var j int
// 	var sum int
// 	if mat.type_ != 1 {
// 		return -3
// 	}
// 	if (func() *io.File {
// 		fw = noarch.Fopen(fname, string("w"))
// 		return fw
// 	}()) == nil {
// 		return -2
// 	}
// 	fmt.Fprintf(fw, string("%%%%MatrixMarket matrix coordinate real general\n"))
// 	fmt.Fprintf(fw, string("%d %d %d\n"), mat.rows, mat.cols, mat.len_)
// 	for i = 0; i < mat.rows; i++ {
// 		sum = 0
// 		for j = mat.frompos[i]; j < mat.frompos[i]+mat.defpos[i]; j++ {
// 			if mat.pos[j] >= 0 {
// 				sum++
// 			} else {
// 				break
// 			}
// 		}
// 		for j = mat.frompos[i]; j < mat.frompos[i]+sum; j++ {
// 			fmt.Fprintf(fw, string("%d %d %e\n"), i+1, mat.pos[j], mat.data[j])
// 		}
// 	}
// 	if noarch.Fclose(fw) != 0 {
// 		rv = -2
// 	}
// 	return rv
// }

// femSparseMatReadF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:368
// func femSparseMatReadF(fname string, mat *tMatrix) int {
// 	// Reads matrix from file IN SPARSE FORM
// 	// * @param fname name of file
// 	// * @param mat matrix (must be unallocated)
// 	// * @return status
// 	//
// 	var fw *io.File
// 	var rv int
// 	var i int
// 	var j int
// 	var k int
// 	var tmp int
// 	var sum int
// 	var size int
// 	var ensize int
// 	var pos0 []int
// 	var data0 []float64
// 	if (func() *io.File {
// 		fw = noarch.Fopen(fname, string("r"))
// 		return fw
// 	}()) == nil {
// 		return -2
// 	}
// 	noarch.Fscanf(fw, string("%d %d\n"), (*[1000000]int)(unsafe.Pointer(&mat.rows)), (*[1000000]int)(unsafe.Pointer(&mat.cols)))
// 	if mat.rows <= 0 || mat.cols <= 0 {
// 		return -2
// 	}
// 	if len((func() []int {
// 		mat.frompos = make([]int,mat.rows)
// 		return mat.frompos
// 	}())) == 0 {
// 		rv = -4
// 		goto memFree
// 	}
// 	if len((func() []int {
// 		mat.defpos = make([]int,mat.rows)
// 		return mat.defpos
// 	}())) == 0 {
// 		rv = -4
// 		goto memFree
// 	}
// 	size = mat.rows * 300
// 	if len((func() []int {
// 		mat.pos = make([]int,size)
// 		return mat.pos
// 	}())) == 0 {
// 		rv = -4
// 		goto memFree
// 	}
// 	if len((func() []float64 {
// 		mat.data = make([]float64,size)
// 		return mat.data
// 	}())) == 0 {
// 		rv = -4
// 		goto memFree
// 	}
// 	mat.type_ = 1
// 	sum = 0
// 	for i = 0; i < mat.rows; i++ {
// 		noarch.Fscanf(fw, string("%d %d "), c4goUnsafeConvert_int(&tmp), mat.defpos[i:])
// 		if i > 0 {
// 			mat.frompos[i] = mat.frompos[i-1] + mat.defpos[i-1]
// 		} else {
// 			// first row
// 			mat.frompos[i] = 0
// 		}
// 		for j = 0; j < mat.defpos[i]; j++ {
// 			if sum >= size {
// 				// enlarge "data" and "pos"
// 				ensize = size + 2*size*(i/mat.rows)
// 				if len((func() []int {
// 					pos0 = make([]int,ensize)
// 					return pos0
// 				}())) == 0 {
// 					rv = -4
// 					goto memFree
// 				}
// 				if len((func() []float64 {
// 					data0 = make([]float64,ensize)
// 					return data0
// 				}())) == 0 {
// 					rv = -4
// 					goto memFree
// 				}
// 				for k = 0; k < sum; k++ {
// 					pos0[k] = mat.pos[k]
// 					data0[k] = mat.data[k]
// 				}
// 				_ = mat.pos
// 				_ = mat.data
// 				mat.pos = pos0
// 				mat.data = data0
// 				pos0 = nil
// 				data0 = nil
// 			}
// 			noarch.Fscanf(fw, string("%d %f "), mat.pos[sum:], mat.data[sum:])
// 			sum++
// 		}
// 	}
// 	if noarch.Fclose(fw) != 0 {
// 		rv = -2
// 	}
// 	return rv
// memFree:
// 	;
// 	//femMatFree(mat)
// 	return rv
// }

// femMatOut - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:447
// func femMatOut(a *tMatrix, fw *io.File) int {
// 	// Writes matrix to stream (FILE *)
// 	// * @param a matrix
// 	// * @param fw stream
// 	// * @return stave value
// 	//
// 	var rv int
// 	var i int
// 	var j int
// 	fmt.Fprintf(fw, string(" %d %d\n"), a.rows, a.cols)
// 	for i = 1; i <= a.rows; i++ {
// 		for j = 1; j <= a.cols; j++ {
// 			fmt.Fprintf(fw, string(" %e \n"), femMatGet(a, i, j))
// 		}
// 	}
// 	return rv
// }

// femMatSetZeroBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:483
// func femMatSetZeroBig(a *tMatrix) {
// 	// Sets all of matrix contents to 0
// 	var i int
// 	for i = 0; i < a.len_; i++ {
// 		a.data[i] = 0
// 	}
// }

// femMatSetZero - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:527
func femMatSetZero(a *tMatrix) {
	// Sets all of matrix contents to 0 FOR SMALL DATA
	var i int
	for i = 0; i < a.len_; i++ {
		a.data[i] = 0
	}
}

// femMatSetZeroRow - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:535
func femMatSetZeroRow(a *tMatrix, row int) {
	// Sets matrix row to 0
	var i int
	if row < 1 || row > a.rows {
		return
	}
	// 	if a.type_ == 1 {
	// 		for i = a.frompos[row-1]; i < a.frompos[row-1]+a.defpos[row-1]; i++ {
	// 			if a.pos[i] == 0 {
	// 				break
	// 			}
	// 			a.data[i] = 0
	// 		}
	// 	} else {
	//fprintf(os.Stdout,"zero on %d\n",i);
	for i = 1; i <= a.cols; i++ {
		femMatPutAdd(a, row, i, 0, 0)
	}
	// 	}
}

// femMatSetZeroCol - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:592
func femMatSetZeroCol(a *tMatrix, Col int) {
	// Sets all of matrix contents to 0
	// 	var i int
	// 	var j int
	// 	_ = j

	// 	_ = i
	// 	var ifrom int
	// 	_ = ifrom
	// 	var ito int
	// 	_ = ito
	// 	var ipos int
	// 	_ = ipos
	// 	if a.type_ == 1 {
	// 		ifrom = a.pos[a.frompos[Col-1]] - 1
	// 		ito = a.pos[a.frompos[Col-1]+a.defpos[Col-1]-1] - 1
	// 		for i = ifrom; i < ito; i++ {
	// 			for j = a.frompos[i]; j < a.frompos[i]+a.defpos[i]; j++ {
	// 				if a.pos[j] == Col {
	// 					a.data[j] = 0
	// 				}
	// 			}
	// 		}
	// 	} else {
	for i := 1; i <= a.rows; i++ {
		femMatPutAdd(a, i, Col, 0, 0)
	}
	// 	}
}

// //femVecNull - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:699
// func //femVecNull(mat *tVector) {
// 	// VECTOR ***
// 	// 	mat.type_ = 0
// 	mat.rows = 0
// 	mat.len_ = 0
// 	mat.pos = nil
// 	mat.data = nil
// }

// //femVecFree - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:708
// func //femVecFree(mat *tVector) {
// 	mat.type_ = 0
// 	mat.rows = 0
// 	mat.len_ = 0
// 	//femIntFree(mat.pos)
// 	//femDblFree(mat.data)
// }

// femVecAlloc - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:718
func femVecAlloc(mat *tVector, type_ int, rows int, items int) int {
	//femVecNull(mat)
	// 	if type_ >= 0 && type_ <= 1 {
	// 		mat.type_ = type_
	// 		switch type_ {
	// 		case 0:
	mat.rows = rows
	mat.len_ = rows
	// 	if len((func() []float64 {
	mat.data = make([]float64, mat.len_)
	// 		return mat.data
	// 	}())) == 0 {
	// 		goto memFree
	// 	}
	// mat.pos = nil
	// 		case 1:
	// 			// VEC_SPAR cannot be used ;-)
	// 			noarch.Exit(-3)
	// 			mat.rows = rows
	// 			if items > 0 {
	// 				mat.len_ = items
	// 				if len((func() []float64 {
	// 					mat.data = make([]float64,mat.len_)
	// 					return mat.data
	// 				}())) == 0 {
	// 					goto memFree
	// 				}
	// 				if len((func() []int {
	// 					mat.pos = make([]int,mat.len_)
	// 					return mat.pos
	// 				}())) == 0 {
	// 					goto memFree
	// 				}
	// 			} else {
	// 				fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("Number of sparse vector items MUST BE nonzero"))
	// 				goto memFree
	// 			}
	// 			break
	// 		}
	return 0
	// 	} else {
	// 		fmt.Fprintf(os.Stdout, string("[E] %s: %d!\n"), string("Matrix type unsupported"), type_)
	// 		return -3
	// 	}
	// memFree:
	// 	;
	// 	//femVecFree(mat)
	// 	return -4
}

// femVecPutAdd - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:776
func femVecPutAdd(vec *tVector, pos int, val float64, mode int) int {
	if pos > vec.rows {
		// Adds value to vector
		// * @param vec vector
		// * @param pos row to add value
		// * @param val value
		// * @param mode FEM_PUT for putting ("=") FEM_ADD for adding ("+=")
		// * @return  status
		//
		fmt.Fprintf(os.Stdout, string("[E] %s: %d > %d!\n"), string("Index outside vector (Add/Put)"), pos, vec.rows)
		return -11
	}
	// 	switch vec.type_ {
	// 	case 0:
	if mode == 0 {
		// put
		vec.data[pos-1] = val
	} else {
		// add
		vec.data[pos-1] += val
	}
	// 	case 1:
	// 		// unimplemented
	// 		noarch.Exit(-3)
	// 	default:
	// 		fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("Invalid vector type (Add/Put)"))
	// 		return -5
	// 		break
	// 	}
	return 0
}

// femVecGet - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:811
func femVecGet(vec *tVector, pos int) float64 {
	if pos > vec.rows {
		// Gets value from vector
		// * @param vec vector
		// * @param pos row to add value
		// * @return value
		//
		fmt.Fprintf(os.Stdout, string("[E] %s: %d/%d!\n"), string("Index outside vector (Get)"), pos, vec.rows)
		return float64(0)
	}
	// 	switch vec.type_ {
	// 	case 0:
	return vec.data[pos-1]
	// 	case 1:
	// 		// unimplemented
	// 		noarch.Exit(0)
	// 	default:
	// 		fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("Invalid vector type (Get)"))
	// 		return float64(0)
	// 		break
	// 	}
	// return float64(0)
}

// femVecPrn - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:839
// func femVecPrn(mat *tVector, name string) {
// 	// Prints vector to stdout, works only in DEVEL mode
// 	var i int
// 	fmt.Fprintf(os.Stdout, string("\n%s %s %s[%d]:\n"), string("Vector"), name, string("listing"), mat.rows)
// 	for i = 1; i <= mat.rows; i++ {
// 		fmt.Fprintf(os.Stdout, string(" %f "), femVecGet(mat, i))
// 	}
// 	fmt.Fprintf(os.Stdout, string("\n"))
// }

// femVecPrnF - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:858
// func femVecPrnF(fname string, mat *tVector) int {
// 	// Saves vector to file
// 	// * @param fname name of file
// 	// * @param mat vector to be printed
// 	// * @return status
// 	//
// 	var fw *io.File
// 	var rv int
// 	var i int
// 	if (func() *io.File {
// 		fw = noarch.Fopen(fname, string("w"))
// 		return fw
// 	}()) == nil {
// 		return -2
// 	}
// 	for i = 1; i <= mat.rows; i++ {
// 		fmt.Fprintf(fw, string(" %e "), femVecGet(mat, i))
// 	}
// 	fmt.Fprintf(fw, string("\n"))
// 	if noarch.Fclose(fw) != 0 {
// 		rv = -2
// 	}
// 	return rv
// }

// femVecOut - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:883
// func femVecOut(a *tVector, fw *io.File) int {
// 	// Writes vector to stream (FILE *)
// 	// * @ a vector
// 	// * @ fw stream
// 	// * @return stave value
// 	//
// 	var rv int
// 	var i int
// 	fmt.Fprintf(fw, string(" %d\n"), a.rows)
// 	for i = 1; i <= a.rows; i++ {
// 		fmt.Fprintf(fw, string(" %e \n"), femVecGet(a, i))
// 	}
// 	return rv
// }

// femVecSetZeroBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:917
// func femVecSetZeroBig(a *tVector) {
// 	// Sets all of vertor contents to 0
// 	var i int
// 	for i = 0; i < a.len_; i++ {
// 		a.data[i] = 0
// 	}
// }

// femVecSetZero - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:961
func femVecSetZero(a *tVector) {
	// Sets all of vertor contents to 0 FOR SMALL DATA
	// 	var i int
	for i := 0; i < a.len_; i++ {
		a.data[i] = 0
	}
}

// femVecClone - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:971
// func femVecClone(src *tVector, dest *tVector) int {
// 	// Clones vectors: src to dest both must be VEC_FULL, same size and allocated
// 	// * @param src original vector
// 	// * @param dest moditied vector
// 	//
// 	var i int
// 	if src.type_ != 0 || dest.type_ != 0 {
// 		return -5
// 	}
// 	if src.len_ != dest.len_ {
// 		return -9
// 	}
// 	for i = 0; i < src.len_; i++ {
// 		dest.data[i] = src.data[i]
// 	}
// 	return 0
// }

// femVecVecMultBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1014
func femVecVecMultBig(a *tVector, b *tVector) float64 {
	// ------------------    Matrix Operations    --------------------
	// vector multiplication (scalar) (a[n]^t * b[n])
	// * @param a vector
	// * @param b vector
	// * @return multiplication product
	//
	var i int
	var mult float64
	if a.rows != b.rows {
		return float64(0)
	}
	if a.rows <= 0 || b.rows <= 0 {
		return float64(0)
	}
	mult = 0
	// 	if a.type_ == 0 && b.type_ == 0 {
	for i = 0; i < a.rows; i++ {
		mult += a.data[i] * b.data[i]
	}
	// 	} else {
	// 		for i = 1; i <= a.rows; i++ {
	// 			mult += femVecGet(a, i) * femVecGet(b, i)
	// 		}
	// 	}
	return mult
}

// femVecVecMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1104
// func femVecVecMult(a *tVector, b *tVector) float64 {
// 	// vector multiplication (scalar) (a[n]^t * b[n])  FOR SMALL VECTORS
// 	// * @param a vector
// 	// * @param b vector
// 	// * @return multiplication product
// 	//
// 	var i int
// 	var mult float64
// 	if a.rows != b.rows {
// 		return float64(0)
// 	}
// 	if a.rows <= 0 || b.rows <= 0 {
// 		return float64(0)
// 	}
// 	mult = 0
// 	if a.type_ == 0 && b.type_ == 0 {
// 		for i = 0; i < a.rows; i++ {
// 			mult += a.data[i] * b.data[i]
// 		}
// 	} else {
// 		for i = 1; i <= a.rows; i++ {
// 			mult += femVecGet(a, i) * femVecGet(b, i)
// 		}
// 	}
// 	return mult
// }

// femVecVecMulttoMat - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1140
// func femVecVecMulttoMat(a *tVector, b *tVector, c *tMatrix) int {
// 	// vector multiplication (matrix) (a[n] * b[n]^t)
// 	// * @param a vector
// 	// * @param b vector
// 	// * @param c matrix (result)
// 	// * @return status
// 	//
// 	var i int
// 	var j int
// 	if a.rows != b.rows {
// 		return -9
// 	}
// 	if a.rows != c.rows || b.rows != c.cols {
// 		return -9
// 	}
// 	if a.type_ != 0 {
// 		return -5
// 	}
// 	if a.type_ == 0 && b.type_ == 0 {
// 		for i = 0; i < a.rows; i++ {
// 			for j = 0; j < a.rows; j++ {
// 				c.data[i*c.cols+j] = a.data[i] * b.data[j]
// 			}
// 		}
// 	} else {
// 		for i = 1; i <= a.rows; i++ {
// 			for j = 1; j <= a.rows; j++ {
// 				femMatPutAdd(c, i, j, femVecGet(a, i)*femVecGet(b, j), 0)
// 			}
// 		}
// 	}
// 	return 0
// }

// femValVecMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1180
// func femValVecMult(val float64, a *tVector, b *tVector) int {
// 	// number by vector multiplication (b[n] = val * a[n])
// 	// * @param val number
// 	// * @param a original vector (will not be modified)
// 	// * @param b result (vector) - must be allocated and must have proper size
// 	// * @return status
// 	//
// 	var i int
// 	if a.rows != b.rows {
// 		return -9
// 	}
// 	if a.type_ != b.type_ {
// 		return -3
// 	}
// 	if a.type_ != 0 {
// 		// will be fixed
// 		return -3
// 	}
// 	for i = 0; i < a.len_; i++ {
// 		b.data[i] = a.data[i] * val
// 	}
// 	return 0
// }

// femValVecMultSelf - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1202
// func femValVecMultSelf(val float64, a *tVector) int {
// 	// number by vector multiplication (a[n] = val * a[n])
// 	// * @param val number
// 	// * @param a original vector (WILL BE modified)
// 	// * @return status
// 	//
// 	var i int
// 	for i = 0; i < a.len_; i++ {
// 		a.data[i] *= val
// 	}
// 	return 0
// }

// femValMatMultSelf - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1215
func femValMatMultSelf(val float64, a *tMatrix) int {
	// number by matrix multiplication (a[n] = val * a[n])
	// * @param val number
	// * @param a original number (WILL BE modified)
	// * @return status
	//
	// var i int
	for i := 0; i < a.len_; i++ {
		a.data[i] *= val
	}
	return 0
}

// femVecMatMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1229
// func femVecMatMult(a *tVector, b *tMatrix, c *tVector) int {
// 	// vector by matrix multiplication (a[n]^t * b[n,m]  = c[m])
// 	// * @param a vector
// 	// * @param b matrix
// 	// * @param vector (result)
// 	// * @return status
// 	//
// 	var i int
// 	var j int
// 	var val float64
// 	if a.rows != b.rows || b.cols != c.rows {
// 		return -9
// 	}
// 	if c.type_ != 0 {
// 		return -3
// 	}
// 	if a.type_ == 0 && b.type_ == 0 && c.type_ == 0 {
// 		for i = 0; i < b.cols; i++ {
// 			val = 0
// 			for j = 0; j < a.rows; j++ {
// 				val += a.data[j] * b.data[i+b.cols*j]
// 			}
// 			c.data[i] = val
// 		}
// 	} else {
// 		for i = 1; i <= b.cols; i++ {
// 			val = 0
// 			for j = 1; j <= a.rows; j++ {
// 				val += femVecGet(a, j) * femMatGet(b, j, i)
// 			}
// 			femVecPutAdd(c, i, val, 0)
// 		}
// 	}
// 	return 0
// }

// femVecMatVecMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1272
// func femVecMatVecMult(a *tVector, b *tMatrix, c *tVector) float64 {
// 	// Vector by matrix by vector multiplication (a[n]^t * b[n,m] * c[m]  = d)
// 	// * For small full matrices only (it is slow).
// 	// * @param a vector
// 	// * @param b matrix
// 	// * @param c vector
// 	// * @return constant (result)
// 	//
// 	var i int
// 	var j int
// 	var val float64
// 	var sum_tot float64
// 	sum_tot = 0
// 	if a.rows != b.rows || b.cols != c.rows {
// 		return float64(-9)
// 	}
// 	if c.type_ != 0 {
// 		return float64(-3)
// 	}
// 	if a.type_ == 0 && b.type_ == 0 && c.type_ == 0 {
// 		for i = 0; i < b.cols; i++ {
// 			val = 0
// 			for j = 0; j < a.rows; j++ {
// 				val += a.data[j] * b.data[i+b.cols*j]
// 			}
// 			sum_tot += c.data[i] * val
// 		}
// 	} else {
// 		for i = 1; i <= b.cols; i++ {
// 			val = 0
// 			for j = 1; j <= a.rows; j++ {
// 				val += femVecGet(a, j) * femMatGet(b, j, i)
// 			}
// 			sum_tot += femVecGet(c, i) * val
// 		}
// 	}
// 	return sum_tot
// }

// femMatVecMultBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1347
func femMatVecMultBig(a *tMatrix, b *tVector, c *tVector) int {
	// Matrix by vector multiplication (a[m,n]*b[n] = b[n])
	// * @param a matrix
	// * @param b vector
	// * @param c vector (result)
	// * @return status
	//
	var i int
	var j int
	var val float64
	if a.cols != b.rows || c.rows != a.rows {
		return -9
	}
	// 	if c.type_ != 0 {
	// 		return -3
	// 	}
	// 	if a.type_ == 0 && b.type_ == 0 {
	for i = 0; i < a.rows; i++ {
		val = 0
		for j = 0; j < a.cols; j++ {
			val += b.data[j] * a.data[j+i*a.cols]
		}
		c.data[i] = val
	}
	// 	} else {
	// 		if a.type_ == 1 && b.type_ == 0 {
	// 			femVecSetZero(c)
	// 			for i = 0; i < a.rows; i++ {
	// 				val = 0
	// 				for j = a.frompos[i]; j < a.frompos[i]+a.defpos[i]; j++ {
	// 					if a.pos[j] <= 0 {
	// 						break
	// 					}
	// 					val += a.data[j] * b.data[a.pos[j]-1]
	// 				}
	// 				c.data[i] = val
	// 			}
	// 		} else {
	// 			for i = 1; i <= a.rows; i++ {
	// 				val = 0
	// 				for j = 1; j <= a.cols; j++ {
	// 					val += femMatGet(a, i, j) * femVecGet(b, j)
	// 				}
	// 				femVecPutAdd(c, i, val, 0)
	// 			}
	// 		}
	// 	}
	return 0
}

// femMatVecMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1456
func femMatVecMult(a *tMatrix, b *tVector, c *tVector) int {
	// Matrix by vector multiplication (a[m,n]*b[n] = b[n]) FOR SMALL DATA
	// * @param a matrix
	// * @param b vector
	// * @param c vector (result)
	// * @return status
	//
	var i int
	var j int
	var val float64
	if a.cols != b.rows || c.rows != a.rows {
		return -9
	}
	// 	if c.type_ != 0 {
	// 		return -3
	// 	}
	// 	if a.type_ == 0 && b.type_ == 0 {
	for i = 0; i < a.rows; i++ {
		val = 0
		for j = 0; j < a.cols; j++ {
			val += b.data[j] * a.data[j+i*a.cols]
		}
		c.data[i] = val
	}
	// 	} else {
	// 		if a.type_ == 1 && b.type_ == 0 {
	// 			femVecSetZero(c)
	// 			for i = 0; i < a.rows; i++ {
	// 				val = 0
	// 				for j = a.frompos[i]; j < a.frompos[i]+a.defpos[i]; j++ {
	// 					if a.pos[j] <= 0 {
	// 						break
	// 					}
	// 					val += a.data[j] * b.data[a.pos[j]-1]
	// 				}
	// 				c.data[i] = val
	// 			}
	// 		} else {
	// 			for i = 1; i <= a.rows; i++ {
	// 				val = 0
	// 				for j = 1; j <= a.cols; j++ {
	// 					val += femMatGet(a, i, j) * femVecGet(b, j)
	// 				}
	// 				femVecPutAdd(c, i, val, 0)
	// 			}
	// 		}
	// 	}
	return 0
}

// femVecLinCombBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1540
// func femVecLinCombBig(amult float64, a *tVector, bmult float64, b *tVector, c *tVector) int {
// 	// linear combination of vectors am*a[m,n]+ bm*b[m,n] = c[m,n] (c..MAT_FULL)
// 	// * @param am  "a" vector multiplier
// 	// * @param a vector
// 	// * @param bm  "b" vector multiplier
// 	// * @param b vector
// 	// * @param c vector (result)
// 	// * @return status
// 	//
// 	var i int
// 	if a.rows != b.rows {
// 		return 0
// 	}
// 	if a.rows <= 0 || b.rows <= 0 {
// 		return 0
// 	}
// 	if a.type_ == 0 && b.type_ == 0 && c.type_ == 0 {
// 		for i = 0; i < a.rows; i++ {
// 			c.data[i] = amult*a.data[i] + bmult*b.data[i]
// 		}
// 	} else {
// 		// VERY SLOW CODE:
// 		for i = 1; i <= a.rows; i++ {
// 			femVecPutAdd(c, i, femVecGet(a, i)*amult+femVecGet(b, i)*bmult, 0)
// 		}
// 	}
// 	return 0
// }

// femVecLinComb - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1629
// func femVecLinComb(amult float64, a *tVector, bmult float64, b *tVector, c *tVector) int {
// 	// linear combination of vectors am*a[m,n]+ bm*b[m,n] = c[m,n] (c..MAT_FULL)
// 	// * @param am  "a" vector multiplier
// 	// * @param a vector
// 	// * @param bm  "b" vector multiplier
// 	// * @param b vector
// 	// * @param c vector (result)
// 	// * @return status
// 	//
// 	var i int
// 	if a.rows != b.rows || b.rows != c.rows {
// 		return -9
// 	}
// 	if c.type_ != 0 {
// 		return -3
// 	}
// 	if a.type_ == 0 && b.type_ == 0 && c.type_ == 0 {
// 		for i = 0; i < a.rows; i++ {
// 			c.data[i] = amult*a.data[i] + bmult*b.data[i]
// 		}
// 	} else {
// 		// SLOW CODE:
// 		for i = 1; i <= a.rows; i++ {
// 			femVecPutAdd(c, i, femVecGet(a, i)*amult+femVecGet(b, i)*bmult, 0)
// 		}
// 	}
// 	return 0
// }

// femMatMatMult - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1662
func femMatMatMult(a *tMatrix, b *tMatrix, c *tMatrix) int {
	// matrix by matrix multiplication a[m,n]*b[n,h] = c[m,h]
	// * @param a matrix
	// * @param b matrix
	// * @param c matrix (result)
	// * @return status
	//
	var i int
	var j int
	var k int
	var val float64
	if a.cols != b.rows || b.cols != c.cols || a.rows != c.rows {
		return -9
	}
	// 	if c.type_ != 0 {
	// 		return -3
	// 	}
	// 	if a.type_ == 0 && b.type_ == 0 && c.type_ == 0 {
	for i = 0; i < a.rows; i++ {
		for j = 0; j < b.cols; j++ {
			val = 0
			for k = 0; k < a.cols; k++ {
				//val += femMatGet(a, i,k)*femMatGet(b, k,j);
				val += a.data[i*a.cols+k] * b.data[k*b.cols+j]
			}
			//femMatPut(c, i,j, val);
			c.data[i*c.cols+j] = val
		}
	}
	// 	} else {
	// 		for i = 1; i <= a.rows; i++ {
	// 			for j = 1; j <= b.cols; j++ {
	// 				val = 0
	// 				for k = 1; k <= a.cols; k++ {
	// 					val += femMatGet(a, i, k) * femMatGet(b, k, j)
	// 				}
	// 				femMatPutAdd(c, i, j, val, 0)
	// 			}
	// 		}
	// 	}
	return 0
}

// femMatLinComb - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1714
// func femMatLinComb(am float64, a *tMatrix, bm float64, b *tMatrix, c *tMatrix) int {
// 	// linear combination of matrices am*a[m,n]+ bm*b[m,n] = c[m,n] (c..MAT_FULL)
// 	// * @param am  "a" matrix multiplier
// 	// * @param a matrix
// 	// * @param bm  "b" matrix multiplier
// 	// * @param b matrix
// 	// * @param c matrix (result)
// 	// * @return status
// 	//
// 	var i int
// 	var j int
// 	var val float64
// 	if a.cols != b.cols || a.rows != b.rows || a.rows != c.rows || a.cols != c.cols {
// 		return -9
// 	}
// 	if c.type_ != 0 {
// 		return -3
// 	}
// 	if a.type_ == 0 && b.type_ == 0 && c.type_ == 0 {
// 		for i = 0; i < a.rows*a.cols; i++ {
// 			c.data[i] = am*a.data[i] + bm*b.data[i]
// 		}
// 	} else {
// 		for i = 1; i <= c.rows; i++ {
// 			for j = 1; j <= c.cols; j++ {
// 				val = am*femMatGet(a, i, j) + bm*femMatGet(b, i, j)
// 				femMatPutAdd(c, i, j, val, 0)
// 			}
// 		}
// 	}
// 	return 0
// }

// femMatTran - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1750
func femMatTran(a *tMatrix, b *tMatrix) int {
	// matrix transposition - works only on dense matrices (MAT_FULL)
	// * @param a matrix (original)
	// * @param b matrix (result - must be allocated)
	// * @return status
	//
	var i int
	var j int
	if a.cols != b.rows || b.cols != a.rows {
		return -9
	}
	// 	if a.type_ != 0 || b.type_ != 0 {
	// 		return -9
	// 	}
	for i = 0; i < a.rows; i++ {
		for j = 0; j < a.cols; j++ {
			if a.cols == a.rows {
				b.data[j*a.cols+i] = a.data[i*a.cols+j]
			} else {
				femMatPutAdd(b, j+1, i+1, femMatGet(a, i+1, j+1), 0)
			}
		}
	}
	return 0
}

// femMatNormBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1822
func femMatNormBig(a *tMatrix) float64 {
	// Computes norm of sparse matrix
	// *  @param a matrix
	// *  @return norm
	//
	var Norm float64
	var MaxNorm float64
	var val float64
	var i int
	var j int
	MaxNorm = 0
	// 	if a.type_ == 1 {
	// 		for i = 0; i < a.rows; i++ {
	// 			Norm = 0
	// 			for j = a.frompos[i]; j < a.frompos[i]+a.defpos[i]; j++ {
	// 				if a.pos[j] <= 0 {
	// 					break
	// 				}
	// 				Norm += a.data[j] * a.data[j]
	// 			}
	// 			Norm = math.Sqrt(Norm)
	// 			if Norm > MaxNorm {
	// 				MaxNorm = Norm
	// 			}
	// 		}
	// 	} else {
	for i = 1; i <= a.rows; i++ {
		Norm = 0
		for j = 1; j <= a.cols; j++ {
			val = femMatGet(a, i, j)
			Norm += val * val
		}
		Norm = math.Sqrt(Norm)
		if Norm > MaxNorm {
			MaxNorm = Norm
		}
	}
	// 	}
	return MaxNorm
}

// femMatNorm - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1921
// func femMatNorm(a *tMatrix) float64 {
// 	// Computes norm of sparse matrix FOR SMALL DATA
// 	// *  @param a matrix
// 	// *  @return norm
// 	//
// 	var Norm float64
// 	var MaxNorm float64
// 	var val float64
// 	var i int
// 	var j int
// 	MaxNorm = 0
// 	if a.type_ == 1 {
// 		for i = 0; i < a.rows; i++ {
// 			Norm = 0
// 			for j = a.frompos[i]; j < a.frompos[i]+a.defpos[i]; j++ {
// 				if a.pos[j] <= 0 {
// 					break
// 				}
// 				Norm += a.data[j] * a.data[j]
// 			}
// 			Norm = math.Sqrt(Norm)
// 			if Norm > MaxNorm {
// 				MaxNorm = Norm
// 			}
// 		}
// 	} else {
// 		for i = 1; i <= a.rows; i++ {
// 			Norm = 0
// 			for j = 1; j <= a.cols; j++ {
// 				val = femMatGet(a, i, j)
// 				Norm += val * val
// 			}
// 			Norm = math.Sqrt(Norm)
// 			if Norm > MaxNorm {
// 				MaxNorm = Norm
// 			}
// 		}
// 	}
// 	return MaxNorm
// }

// femVecNormBig - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:1988
func femVecNormBig(a *tVector) float64 {
	// Computes Euclide norm of vector sum(sqrt(a*a))
	// *  @param a     vector
	// *  @return norm
	//
	var Norm float64
	// 	var val float64
	// 	_ = val
	var i int
	Norm = 0
	// 	if a.type_ == 0 {
	for i = 0; i < a.rows; i++ {
		Norm += a.data[i] * a.data[i]
	}
	// 	} else {
	// 		for i = 1; i <= a.rows; i++ {
	// 			val = femVecGet(a, i)
	// 			Norm += val * val
	// 		}
	// 	}
	return math.Sqrt(Norm)
}

// femVecNorm - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2072
// func femVecNorm(a *tVector) float64 {
// 	// Computes Euclide norm of vector sum(sqrt(a*a)) FOR SMALL DATA
// 	// *  @param a     vector
// 	// *  @return norm
// 	//
// 	var Norm float64
// 	var val float64
// 	var i int
// 	Norm = 0
// 	if a.type_ == 0 {
// 		for i = 0; i < a.rows; i++ {
// 			Norm += a.data[i] * a.data[i]
// 		}
// 	} else {
// 		for i = 1; i <= a.rows; i++ {
// 			val = femVecGet(a, i)
// 			Norm += val * val
// 		}
// 	}
// 	return math.Sqrt(Norm)
// }

// femVecAddVec - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2104
// func femVecAddVec(orig *tVector, mult float64, addt *tVector) int {
// 	// Adds vector "addt" to "orig" e.g. orig += mult*addt
// 	// * @param orig original vector (to be modified)
// 	// * @param mult scalar multiplier
// 	// * @param addt addition vector
// 	// * @return status
// 	//
// 	var i int
// 	if orig.rows != addt.rows {
// 		fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("Different vector sizes not allowed here"))
// 		return -9
// 	}
// 	if orig.type_ == 0 && addt.type_ == 0 {
// 		for i = 0; i < orig.len_; i++ {
// 			orig.data[i] += mult * addt.data[i]
// 		}
// 	} else {
// 		for i = 1; i <= orig.len_; i++ {
// 			femVecPutAdd(orig, i, mult*femVecGet(addt, i), 1)
// 		}
// 	}
// 	return 0
// }

// femMatInv - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2137
// func femMatInv(a *tMatrix) int {
// 	// Does matrix inversion UNOPTIMIZED!
// 	// *  @param a  matrix to be inverted
// 	//
// 	var m int
// 	var n int
// 	var i int
// 	var j int
// 	var k int
// 	var f float64
// 	var f2 float64
// 	var val float64
// 	var f1 tVector
// 	if a.rows != a.cols {
// 		return -9
// 	}
// 	n = a.cols
// 	//femVecNull(((&f1)))
// 	if femVecAlloc(((&f1)), 0, n, n) != 0 {
// 		return -4
// 	}
// 	m = n - 1
// 	val = femMatGet(a, 1, 1)
// 	femMatPutAdd(a, 1, 1, 1/val, 0)
// 	for i = 1; i <= m; i++ {
// 		for j = 1; j <= i; j++ {
// 			f = 0
// 			for k = 1; k <= i; k++ {
// 				f += femMatGet(a, j, k) * femMatGet(a, k, i+1)
// 			}
// 			femVecPutAdd(((&f1)), j, -f, 0)
// 		}
// 		f2 = femMatGet(a, i+1, i+1)
// 		for j = 1; j <= i; j++ {
// 			f2 += femMatGet(a, j, i+1) * femVecGet(((&f1)), j)
// 		}
// 		if math.Abs(f2/femMatGet(a, i+1, i+1)) < 1e-07 {
// 			return -3
// 		}
// 		f2 = 1 / f2
// 		femMatPutAdd(a, i+1, i+1, f2, 0)
// 		for j = 1; j <= i; j++ {
// 			for k = 1; k <= i; k++ {
// 				femMatPutAdd(a, j, k, femVecGet(((&f1)), j)*femVecGet(((&f1)), k)*f2+femMatGet(a, j, k), 0)
// 			}
// 		}
// 		for j = 1; j <= i; j++ {
// 			femMatPutAdd(a, j, i+1, femVecGet(((&f1)), j)*f2, 0)
// 			femMatPutAdd(a, i+1, j, femMatGet(a, j, i+1), 0)
// 		}
// 	}
// 	//femVecFree(((&f1)))
// 	return 0
// }

// femLUdecomp - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2219
// func femLUdecomp(a *tMatrix, index *tVector) int {
// 	// L-U:
// 	// Decomposition to L/U
// 	// * @param a matrix (will be modified!)
// 	// * @param index index vector
// 	// * @param d modified index status (-1/+1)
// 	// * @return status
// 	//
// 	var rv int
// 	var i int
// 	var j int
// 	var k int
// 	var imax int
// 	var n int
// 	var big float64
// 	var dum float64
// 	var sum float64
// 	var temp float64
// 	var vv tVector
// 	//femVecNull(((&vv)))
// 	if (func() int {
// 		n = a.rows
// 		return n
// 	}()) <= 0 {
// 		return -9
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&vv)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	for i = 1; i <= n; i++ {
// 		big = 0
// 		for j = 1; j <= n; j++ {
// 			if (func() float64 {
// 				temp = math.Abs(femMatGet(a, i, j))
// 				return temp
// 			}()) > big {
// 				big = temp
// 			}
// 		}
// 		if big == 0 {
// 			// singular matrix
// 			return -3
// 		}
// 		femVecPutAdd(((&vv)), i, 1/big, 0)
// 	}
// 	for j = 1; j <= n; j++ {
// 		for i = 1; i < j; i++ {
// 			sum = femMatGet(a, i, j)
// 			for k = 1; k < i; k++ {
// 				sum -= femMatGet(a, i, k) * femMatGet(a, k, j)
// 			}
// 			femMatPutAdd(a, i, j, sum, 0)
// 		}
// 		big = 0
// 		for i = j; i <= n; i++ {
// 			sum = femMatGet(a, i, j)
// 			for k = 1; k < j; k++ {
// 				sum -= femMatGet(a, i, k) * femMatGet(a, k, j)
// 			}
// 			femMatPutAdd(a, i, j, sum, 0)
// 			if (func() float64 {
// 				dum = femVecGet(((&vv)), i) * math.Abs(sum)
// 				return dum
// 			}()) >= big {
// 				big = dum
// 				imax = i
// 			}
// 		}
// 		if j != imax {
// 			for k = 1; k <= n; k++ {
// 				dum = femMatGet(a, imax, k)
// 				femMatPutAdd(a, imax, k, femMatGet(a, j, k), 0)
// 				femMatPutAdd(a, j, k, dum, 0)
// 			}
// 			femVecPutAdd(((&vv)), imax, femVecGet(((&vv)), j), 0)
// 		}
// 		femVecPutAdd(index, j, float64(imax), 0)
// 		if femMatGet(a, j, j) == 0 {
// 			femMatPutAdd(a, j, j, 1e-20, 0)
// 		}
// 		if j != n {
// 			dum = 1 / femMatGet(a, j, j)
// 			for i = j + 1; i <= n; i++ {
// 				femMatPutAdd(a, i, j, dum*femMatGet(a, i, j), 0)
// 			}
// 		}
// 	}
// memFree:
// 	;
// 	//femVecFree(((&vv)))
// 	return rv
// }

// femLUback - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2322
// func femLUback(a *tMatrix, index *tVector, b *tVector) int {
// 	// Decomposition to L/U
// 	// * @param a matrix (will be modified!)
// 	// * @param index index vector
// 	// * @param b right hand side/result vector (will be modified!)
// 	// * @return status
// 	//
// 	var rv int
// 	var i int
// 	var ii int
// 	var ip int
// 	var j int
// 	var n int
// 	var sum float64
// 	ii = 0
// 	if (func() int {
// 		n = a.rows
// 		return n
// 	}()) <= 0 {
// 		return -9
// 	}
// 	for i = 1; i <= n; i++ {
// 		ip = int(femVecGet(index, i))
// 		sum = femVecGet(b, ip)
// 		femVecPutAdd(b, ip, femVecGet(b, i), 0)
// 		if ii != 0 {
// 			// means ii > 0
// 			for j = ii; j <= i-1; j++ {
// 				sum -= femMatGet(a, i, j) * femVecGet(b, j)
// 			}
// 		} else {
// 			if sum != 0 {
// 				ii = i
// 			}
// 		}
// 		femVecPutAdd(b, i, sum, 0)
// 	}
// 	for i = n; i >= 1; i-- {
// 		sum = femVecGet(b, i)
// 		for j = i + 1; j <= n; j++ {
// 			sum -= femMatGet(a, i, j) * femVecGet(b, j)
// 		}
// 		femVecPutAdd(b, i, sum/femMatGet(a, i, i), 0)
// 	}
// 	return rv
// }

// femLUinverse - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2374
// func femLUinverse(a *tMatrix) int {
// 	// Inversion of "a" matrix using L/U
// 	// * @param a matrix (will be modified!)
// 	// * @return status
// 	//
// 	var rv int
// 	var i int
// 	var j int
// 	var n int
// 	var col tVector
// 	var index tVector
// 	var b tMatrix
// 	if (func() int {
// 		n = a.rows
// 		return n
// 	}()) <= 0 {
// 		return -9
// 	}
// 	//femVecNull(((&col)))
// 	//femVecNull(((&index)))
// 	//femMatNull(((&b)))
// 	if (func() int {
// 		rv = femVecAlloc(((&col)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&index)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femMatAlloc(((&b)), 0, n, n, 0, nil)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femLUdecomp(a, ((&index)))
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	for j = 1; j <= n; j++ {
// 		for i = 1; i <= n; i++ {
// 			femVecPutAdd(((&col)), i, 0, 0)
// 		}
// 		femVecPutAdd(((&col)), j, 1, 0)
// 		if (func() int {
// 			rv = femLUback(a, ((&index)), ((&col)))
// 			return rv
// 		}()) != 0 {
// 			goto memFree
// 		}
// 		for i = 1; i <= n; i++ {
// 			femMatPutAdd(((&b)), i, j, femVecGet(((&col)), i), 0)
// 		}
// 	}
// 	for i = 1; i <= n; i++ {
// 		for j = 1; j <= n; j++ {
// 			femMatPutAdd(a, i, j, femMatGet(((&b)), i, j), 0)
// 		}
// 	}
// memFree:
// 	;
// 	//femVecFree(((&col)))
// 	//femVecFree(((&index)))
// 	//femMatFree(((&b)))
// 	return rv
// }

// femVecSwitch - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2424
// func femVecSwitch(a *tVector, b *tVector) int {
// 	// Moves "a" to "b" and "b" to "a"
// 	// * @param a vector
// 	// * @param b vector
// 	// * @return status
// 	//
// 	var val float64
// 	var i int
// 	if a.rows != b.rows || a.type_ != 0 || b.type_ != 0 {
// 		return -9
// 	}
// 	for i = 0; i < a.len_; i++ {
// 		val = a.data[i]
// 		a.data[i] = b.data[i]
// 		b.data[i] = val
// 	}
// 	return 0
// }

// femVecCloneDiff - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2449
// func femVecCloneDiff(orig *tVector, clone *tVector) int {
// 	// Copies vector content to a larger one (extra fields are left untouched)
// 	// * @param orig original vector
// 	// * @param clone target vector (to be modified)
// 	// * @return status
// 	//
// 	var i int
// 	var len_ int
// 	if orig.type_ != 0 || clone.type_ != 0 {
// 		return -5
// 	}
// 	if clone.rows < 1 || orig.rows < 1 {
// 		return -9
// 	}
// 	if orig.rows > clone.rows {
// 		len_ = clone.rows
// 	} else {
// 		len_ = orig.rows
// 	}
// 	for i = 0; i < len_; i++ {
// 		clone.data[i] = orig.data[i]
// 	}
// 	return 0
// }

// femMatCloneDiffToSame - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2484
// func femMatCloneDiffToSame(orig *tMatrix, clone *tMatrix) int {
// 	// TODO FIX!!! Copies sparse matrix content to a larger one (extra fields are left untouched)
// 	// * it is assumed that a) there is a space for data in "clone", b) identical data
// 	// * in both matrices are stored at identical places
// 	// * @param orig original vector
// 	// * @param clone target vector (to be modified)
// 	// * @return status
// 	//
// 	var i int
// 	var j int
// 	var k int
// 	var ko int
// 	var kc int
// 	if orig.type_ != 1 || clone.type_ != 1 {
// 		return -5
// 	}
// 	if orig.rows > clone.rows || orig.rows < 1 {
// 		return -9
// 	}
// 	if orig.cols > clone.cols || orig.cols < 1 {
// 		return -9
// 	}
// 	for i = 0; i < orig.rows; i++ {
// 		k = 0
// 		for j = orig.frompos[i]; j < orig.frompos[i]+orig.defpos[i]; j++ {
// 			ko = k + orig.frompos[i]
// 			kc = k + clone.frompos[i]
// 			clone.data[kc] = orig.data[ko]
// 			k++
// 		}
// 	}
// 	return 0
// }

// femMatCloneDiffToEmpty - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_math.c:2537
// func femMatCloneDiffToEmpty(orig *tMatrix, clone *tMatrix) int {
// 	// Copies sparse matrix content to a larger one (extra fields are left untouched)
// 	// * it is assumed that a) there is a space for data in "clone", b) identical data
// 	// * in both matrices are stored at identical places
// 	// * @param orig original vector
// 	// * @param clone target vector (to be modified)
// 	// * @return status
// 	//
// 	var i int
// 	var j int
// 	if orig.type_ != 1 || clone.type_ != 1 {
// 		return -5
// 	}
// 	if orig.rows > clone.rows || orig.rows < 1 {
// 		return -9
// 	}
// 	if orig.cols > clone.cols || orig.cols < 1 {
// 		return -9
// 	}
// 	for i = 0; i < orig.rows; i++ {
// 		for j = orig.frompos[i]; j < orig.frompos[i]+orig.defpos[i]; j++ {
// 			femMatPutAdd(clone, i+1, orig.pos[j], orig.data[j], 0)
// 		}
// 	}
// 	return 0
// }

// eqsCompResid - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:37
// func eqsCompResid(a *tMatrix, x *tVector, b *tVector, r *tVector) int {
// 	//
// 	//   File name: fem_eqs.c
// 	//   Date:      2003/04/13 10:38
// 	//   Author:    Jiri Brozovsky
// 	//
// 	//   Copyright (C) 2003 Jiri Brozovsky
// 	//
// 	//   This program is free software; you can redistribute it and/or
// 	//   modify it under the terms of the GNU General Public License as
// 	//   published by the Free Software Foundation; either version 2 of the
// 	//   License, or (at your option) any later version.
// 	//
// 	//   This program is distributed in the hope that it will be useful, but
// 	//   WITHOUT ANY WARRANTY; without even the implied warranty of
// 	//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// 	//   General Public License for more details.
// 	//
// 	//   You should have received a copy of the GNU General Public License
// 	//   in a file called COPYING along with this program; if not, write to
// 	//   the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
// 	//   02139, USA.
// 	//
// 	//  FEM Solver - linear equation system solver(s)
// 	//
// 	//  $Id: fem_eqs.c,v 1.13 2005/07/11 17:56:16 jirka Exp $
// 	//
// 	// Computes r = A.x - b
// 	// * @param a matrix
// 	// * @param x results
// 	// * @param b right-side
// 	// * @param r computed residuum vector
// 	// * @return state value
// 	//
// 	var i int
// 	var j int
// 	if a.cols != b.rows || b.rows != x.rows || x.rows != r.rows {
// 		return -9
// 	}
// 	if b.type_ != 0 || x.type_ != 0 || r.type_ != 0 {
// 		return -5
// 	}
// 	if a.type_ == 1 {
// 		for i = 0; i < a.rows; i++ {
// 			r.data[i] = 0 - b.data[i]
// 			for j = a.frompos[i]; j < a.frompos[i]+a.defpos[i]; j++ {
// 				if a.pos[j] <= 0 {
// 					break
// 				}
// 				r.data[i] += a.data[j] * x.data[a.pos[j]-1]
// 			}
// 		}
// 	} else {
// 		for i = 1; i <= a.rows; i++ {
// 			femVecPutAdd(r, i, 0-femVecGet(b, i), 0)
// 			for j = 1; j < a.cols; j++ {
// 				femVecPutAdd(r, i, femMatGet(a, i, j)*femVecGet(x, j), 1)
// 			}
// 		}
// 	}
// 	return 0
// }

// femEqsCGwJ - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:88
func femEqsCGwJ(a *tMatrix, b *tVector, x *tVector, eps float64, maxIt int) int {
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
	var n int
	var i int
	var j int
	var rv int
	var converged int
	var normRes float64
	var normX float64
	var normA float64
	var normB float64
	if a.cols != x.rows || x.rows != b.rows {
		return -9
	}
	n = a.rows
	normA = femMatNormBig(a)
	normB = femVecNormBig(b)
	if normB <= 0 {
		femVecSetZero(x) //Big(x)
		fmt.Fprintf(os.Stdout, string("[ ]  %s!\n"), string("solution done without iterations because of zero load"))
		return 0
	}
	// vector initialization
	//femVecNull(((&M)))
	//femVecNull(((&r)))
	//femVecNull(((&z)))
	//femVecNull(((&p)))
	//femVecNull(((&q)))
	if (func() int {
		rv = femVecAlloc((&M), 0, n, n)
		return rv
	}()) != 0 {
		// memory allocation
		goto memFree
	}
	if (func() int {
		rv = femVecAlloc((&r), 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int {
		rv = femVecAlloc((&z), 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int {
		rv = femVecAlloc((&p), 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	if (func() int {
		rv = femVecAlloc((&q), 0, n, n)
		return rv
	}()) != 0 {
		goto memFree
	}
	{
		// Jacobi preconditioner creation:
		for i = 1; i <= n; i++ {
			M.data[i-1] = femMatGet(a, i, i)
			if math.Abs(M.data[i-1]) < 1e-07 {
				rv = -13
				fmt.Fprintf(os.Stdout, string("[ ]   %s[%d][%d] %s\n"), string("matrix member"), i, i, string("has zero size"))
				goto memFree
			}
		}
	}
	// next two lines mean: r = b - A*x
	femMatVecMultBig(a, x, (&r))
	for i = 0; i < n; i++ {
		r.data[i] = b.data[i] - r.data[i]
	}
	{
		// main loop
		for i = 1; i <= maxIt; i++ {
			fmt.Fprintf(os.Stdout, string("[ ]   %s %d\n"), string("linear step"), i)
			{
				// using preconditioner:
				for j = 0; j < n; j++ {
					z.data[j] = r.data[j] / M.data[j]
				}
			}
			ro = femVecVecMultBig((&r), (&z))
			fmt.Fprintf(os.Stdout, string("ro = %f\n"), ro)
			if i == 1 {
				for j = 0; j < n; j++ {
					p.data[j] = z.data[j]
				}
			} else {
				beta = ro / roro
				fmt.Fprintf(os.Stdout, string("beta = %f\n"), beta)
				for j = 0; j < n; j++ {
					p.data[j] = z.data[j] + beta*p.data[j]
				}
			}
			femMatVecMultBig(a, (&p), (&q))
			alpha = ro / femVecVecMultBig((&p), (&q))
			fmt.Fprintf(os.Stdout, string("alpha = %f\n"), alpha)
			for j = 0; j < n; j++ {
				x.data[j] = x.data[j] + alpha*p.data[j]
				r.data[j] = r.data[j] - alpha*q.data[j]
			}
			// Convergence testing
			normRes = femVecNormBig((&r))
			normX = femVecNormBig(x)
			if normRes <= eps*(normA*normX+normB) {
				// convergence test
				//if (fabs(norm - norm0) < eps )
				converged = 1
				fmt.Fprintf(os.Stdout, string("[ ]  %s %d %s!\n"), string("linear solution done in"), i, string("iterations"))
				break
			}
			fmt.Fprintf(os.Stdout, string("[i] Convergence test %f < %f (step %d from %d)\n"), normRes, eps*(normA*normX+normB), i, maxIt)
			roro = ro
		}
	}
	if converged != 1 {
		// end of main loop
		//fprintf(os.Stdout,"[I] normRes = %f\n",normRes);
		fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("unconverged solution"))
		rv = -1
	}
memFree:
	;
	// freeing memory:
	//femVecFree(((&M)))
	//femVecFree(((&r)))
	//femVecFree(((&z)))
	//femVecFree(((&p)))
	//femVecFree(((&q)))
	return rv
}

// femEqsBiCCSwJ - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:262
// func femEqsBiCCSwJ(a *tMatrix, b *tVector, x *tVector, eps float64, maxIt int) int {
// 	// Bi-Conjugate Gradient Stabilized Method with Jacobi preconditioner
// 	// *  (for symetric and non-symetric matrices)
// 	// *  @param a      matrix
// 	// *  @param b      "load" vector
// 	// *  @param x      results (vector - given as first iteration)
// 	// *  @param eps    error (min.)
// 	// *  @param maxIt  max. number of iterations
// 	// *  @return state value
// 	// *
// 	// *  Note: "res" is probably useless and *NormBig(res) can be replaced by *NormBig(r).
// 	// *  Test it!!
// 	// *
// 	//
// 	// preconditioner (diag[a])
// 	var M tVector
// 	var r tVector
// 	var rr tVector
// 	var p tVector
// 	var pp tVector
// 	var s tVector
// 	var ss tVector
// 	var t tVector
// 	var v tVector
// 	var ro float64
// 	var beta float64
// 	var roro float64
// 	var alpha float64
// 	var omega float64
// 	var i int
// 	var j int
// 	// size of matrix "a"
// 	var n int
// 	var converged int
// 	// residuum
// 	var res tVector
// 	// norms
// 	var normRes float64
// 	var normX float64
// 	var normA float64
// 	var normB float64
// 	var rv int
// 	n = a.rows
// 	normA = femMatNormBig(a)
// 	normX = femVecNormBig(x)
// 	normB = femVecNormBig(b)
// 	// vector initialization
// 	//femVecNull(((&M)))
// 	//femVecNull(((&r)))
// 	//femVecNull(((&rr)))
// 	//femVecNull(((&p)))
// 	//femVecNull(((&pp)))
// 	//femVecNull(((&s)))
// 	//femVecNull(((&ss)))
// 	//femVecNull(((&t)))
// 	//femVecNull(((&v)))
// 	//femVecNull(((&res)))
// 	if (func() int {
// 		rv = femVecAlloc(((&M)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		// memory allocation
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&r)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&rr)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&p)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&pp)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&s)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&ss)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&t)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&v)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&res)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	{
// 		// Jacobi preconditioner creation:
// 		for i = 1; i <= n; i++ {
// 			M.data[i-1] = femMatGet(a, i, i)
// 			if math.Abs(M.data[i-1]) < 1e-07 {
// 				rv = -13
// 				fmt.Fprintf(os.Stdout, string("[ ]   %s[%d][%d] %s\n"), string("matrix member"), i, i, string("has zero size"))
// 				goto memFree
// 			}
// 		}
// 	}
// 	// next two lines mean: r = b - A*x
// 	femMatVecMultBig(a, x, ((&r)))
// 	for i = 0; i < n; i++ {
// 		r.data[i] = b.data[i] - r.data[i]
// 		rr.data[i] = r.data[i]
// 	}
// 	if femVecNormBig(((&r))) <= 1e-07 {
// 		// convergence test
// 		converged = 1
// 		goto memFree
// 	}
// 	{
// 		// main loop
// 		for i = 1; i <= maxIt; i++ {
// 			fmt.Fprintf(os.Stdout, string("[ ]   %s %d\n"), string("linear step"), i)
// 			ro = femVecVecMultBig(((&rr)), ((&r)))
// 			if math.Abs(ro) <= 0 {
// 				fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("solution interrupted on zero value"))
// 				goto memFree
// 			}
// 			if i == 1 {
// 				// in first iteration
// 				for j = 0; j < n; j++ {
// 					p.data[j] = r.data[j]
// 				}
// 			} else {
// 				// int all iterations except first
// 				beta = ro / roro * (alpha / omega)
// 				for j = 0; j < n; j++ {
// 					p.data[j] = r.data[j] + beta*(p.data[j]-omega*v.data[j])
// 				}
// 			}
// 			{
// 				// using preconditioner M.pp=p -> pp
// 				for j = 0; j < n; j++ {
// 					pp.data[j] = p.data[j] / M.data[j]
// 				}
// 			}
// 			femMatVecMultBig(a, ((&pp)), ((&v)))
// 			alpha = ro / femVecVecMultBig(((&rr)), ((&v)))
// 			for j = 0; j < n; j++ {
// 				s.data[j] = r.data[j] - alpha*v.data[j]
// 			}
// 			if femVecNormBig(((&s))) <= 1e-07 {
// 				{
// 					// test of "s" size
// 					for j = 0; j < n; j++ {
// 						x.data[j] += alpha * pp.data[j]
// 					}
// 				}
// 				converged = 1
// 				fmt.Fprintf(os.Stdout, string("[ ]  %s %d %s!\n"), string("linear solution done in"), i, string("iterations"))
// 				break
// 			}
// 			{
// 				// using preconditioner M.ss=s -> ss
// 				for j = 0; j < n; j++ {
// 					ss.data[j] = s.data[j] / M.data[j]
// 				}
// 			}
// 			femMatVecMultBig(a, ((&ss)), ((&t)))
// 			omega = femVecVecMultBig(((&t)), ((&s))) / femVecVecMultBig(((&t)), ((&t)))
// 			for j = 0; j < n; j++ {
// 				x.data[j] += alpha*pp.data[j] + omega*ss.data[j]
// 				r.data[j] = s.data[j] - omega*t.data[j]
// 			}
// 			roro = ro
// 			// Convergence testing
// 			eqsCompResid(a, b, x, ((&res)))
// 			normRes = femVecNormBig(((&res)))
// 			normX = femVecNormBig(x)
// 			if normRes < eps*(normA*normX+normB) {
// 				converged = 1
// 				fmt.Fprintf(os.Stdout, string("[ ]  %s %d %s!\n"), string("solution done in"), i, string("iterations"))
// 				break
// 			}
// 		}
// 	}
// 	if converged != 1 {
// 		// end of main loop
// 		fmt.Fprintf(os.Stdout, string("[E] BiCGS: %s!\n"), string("unconverged solution"))
// 	}
// memFree:
// 	;
// 	// freeing of memory:
// 	//femVecFree(((&M)))
// 	//femVecFree(((&r)))
// 	//femVecFree(((&rr)))
// 	//femVecFree(((&p)))
// 	//femVecFree(((&pp)))
// 	//femVecFree(((&s)))
// 	//femVecFree(((&ss)))
// 	//femVecFree(((&t)))
// 	//femVecFree(((&v)))
// 	//femVecFree(((&res)))
// 	return 0
// }

// femEqsLU - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:482
// func femEqsLU(a *tMatrix, b *tVector, x *tVector, eps float64, maxIt int) int {
// 	// Solver that uses LU - for full matrices!
// 	// *  @param a      matrix
// 	// *  @param b      "load" vector
// 	// *  @param x      results (vector - given as first iteration)
// 	// *  @param eps    error (min.)
// 	// *  @param maxIt  max. number of iterations
// 	// *  @return state value
// 	// *
// 	// *  Note: "res" is probably useless and *NormBig(res) can be replaced by *NormBig(r).
// 	// *  Test it!!
// 	// *
// 	//
// 	var rv int
// 	var n int
// 	var indx tVector
// 	if (func() int {
// 		n = a.rows
// 		return n
// 	}()) <= 0 {
// 		return -9
// 	}
// 	//femVecNull(((&indx)))
// 	if (func() int {
// 		rv = femVecAlloc(((&indx)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femLUdecomp(a, ((&indx)))
// 		return rv
// 	}()) != 0 {
// 		fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("LU decomposition failed"))
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femLUback(a, ((&indx)), b)
// 		return rv
// 	}()) != 0 {
// 		fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("Backward run of LU failed"))
// 		goto memFree
// 	}
// memFree:
// 	;
// 	//femVecFree(((&indx)))
// 	return rv
// }

// femEqsPCGwJ - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:514
// func femEqsPCGwJ(a *tMatrix, b *tVector, x *tVector, eps float64, maxIt int) int {
// 	// Alternative version of the Conjugate Gradient Method ()
// 	var rv int
// 	var converged int
// 	_ = converged
// 	var nui float64
// 	var dei float64
// 	var lambda float64
// 	var alpha float64
// 	// norms
// 	var normRes float64
// 	var normX float64
// 	var normA float64
// 	var normB float64
// 	var p tVector
// 	var r tVector
// 	var d tVector
// 	// Jacobi preconditioner
// 	var M tVector
// 	// a*p result vector
// 	var ap tVector
// 	// number of rows
// 	var n int
// 	var i int
// 	var j int
// 	if a.rows != x.rows || x.rows != b.rows {
// 		return -9
// 	}
// 	n = a.rows
// 	normA = femMatNormBig(a)
// 	normB = femVecNormBig(b)
// 	normX = femVecNormBig(x)
// 	if normB <= 0 {
// 		femVecSetZeroBig(x)
// 		fmt.Fprintf(os.Stdout, string("[ ]  %s!\n"), string("solution done without iterations because of zero load"))
// 		return 0
// 	}
// 	// vector initializations
// 	//femVecNull(((&p)))
// 	//femVecNull(((&r)))
// 	//femVecNull(((&d)))
// 	//femVecNull(((&M)))
// 	//femVecNull(((&ap)))
// 	if (func() int {
// 		rv = femVecAlloc(((&p)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		// memory allocation
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&r)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&d)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&M)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&ap)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	{
// 		// Jacobi preconditioner
// 		for i = 1; i <= n; i++ {
// 			M.data[i-1] = femMatGet(a, i, i)
// 			if math.Abs(M.data[i-1]) < 1e-07 {
// 				rv = -13
// 				fmt.Fprintf(os.Stdout, string("[ ]   %s[%d][%d] %s\n"), string("matrix member"), i, i, string("has zero size"))
// 				goto memFree
// 			}
// 		}
// 	}
// 	// next several lines mean: r = b - A*x
// 	femMatVecMultBig(a, x, ((&r)))
// 	for i = 0; i < n; i++ {
// 		r.data[i] = b.data[i] - r.data[i]
// 	}
// 	{
// 		// using preconditioner:
// 		for j = 0; j < n; j++ {
// 			d.data[j] = r.data[j] / M.data[j]
// 			p.data[j] = d.data[j]
// 		}
// 	}
// 	for i = 1; i <= maxIt; i++ {
// 		fmt.Fprintf(os.Stdout, string("[ ]   %s %d (%s %d)\n"), string("linear step"), i, string("from"), maxIt)
// 		// untested code follows...
// 		femMatVecMultBig(a, ((&p)), ((&ap)))
// 		nui = femVecVecMultBig(((&r)), ((&d)))
// 		dei = femVecVecMultBig(((&p)), ((&ap)))
// 		lambda = nui / dei
// 		noarch.Printf(string("NUI = %f DEI = %f LAMBDA = %f\n"), nui, dei, lambda)
// 		for j = 0; j < n; j++ {
// 			x.data[j] += lambda * p.data[j]
// 			r.data[j] = r.data[j] - lambda*ap.data[j]
// 			d.data[j] = r.data[j] / M.data[j]
// 		}
// 		normRes = femVecNormBig(((&r)))
// 		normX = femVecNormBig(x)
// 		nui = femVecVecMultBig(((&r)), ((&d)))
// 		dei = femVecVecMultBig(((&p)), ((&ap)))
// 		noarch.Printf(string("NORMS: A=%f X=%f B=%f <> R=%f\n"), normA, normX, normB, normRes)
// 		if normRes < eps*(normA*normX+normB) {
// 			// convergence test
// 			converged = 1
// 			fmt.Fprintf(os.Stdout, string("[ ]  %s %d %s!\n"), string("solution done in"), i, string("iterations"))
// 			break
// 		}
// 		alpha = nui / dei
// 		for j = 0; j < n; j++ {
// 			p.data[j] = d.data[j] + alpha*p.data[j]
// 		}
// 	}
// 	// end of "for i"
// 	femVecPrn(x, string("X"))
// memFree:
// 	;
// 	//femVecFree(((&p)))
// 	//femVecFree(((&r)))
// 	//femVecFree(((&d)))
// 	//femVecFree(((&M)))
// 	//femVecFree(((&ap)))
// 	return rv
// }

// femMatCholFact - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:660
// func femMatCholFact(a *tMatrix, C *tVector) int {
// 	// Choleski decomposition - forward run only!
// 	// * @param a matrix (must  be a MAT_FULL)
// 	// * @return status
// 	//
// 	var rv int
// 	var sum float64
// 	var n int
// 	var i int
// 	var j int
// 	var k int
// 	var have_C int
// 	n = a.rows
// 	if len(C) != 0 {
// 		if C.rows != a.rows {
// 			return -3
// 		} else {
// 			have_C = 1
// 		}
// 	}
// 	if have_C == 0 {
// 		//femVecNull(C)
// 		if femVecAlloc(C, 0, n, n) != 0 {
// 			goto memFree
// 		}
// 	}
// 	for i = 1; i <= n; i++ {
// 		for j = i; j <= n; j++ {
// 			sum = femMatGet(a, i, j)
// 			for k = i - 1; k >= 1; k-- {
// 				sum -= femMatGet(a, i, k) * femMatGet(a, j, k)
// 			}
// 			if i == j {
// 				if sum <= 0 {
// 					rv = -3
// 					fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("Given matrix is singular"))
// 					goto memFree
// 				}
// 				femVecPutAdd(C, i, math.Sqrt(sum), 0)
// 			} else {
// 				femMatPutAdd(a, j, i, sum/femVecGet(C, i), 0)
// 			}
// 		}
// 	}
// 	for i = 1; i <= n; i++ {
// 		for j = i; j <= n; j++ {
// 			if i != j {
// 				femMatPutAdd(a, i, j, femMatGet(a, j, i), 0)
// 				femMatPutAdd(a, j, i, 0, 0)
// 			} else {
// 				femMatPutAdd(a, j, i, femVecGet(C, i), 0)
// 			}
// 		}
// 	}
// 	femVecPrn(C, string("C"))
// memFree:
// 	;
// 	if have_C == 0 {
// 		// freeing of memory:
// 		//femVecFree(C)
// 	}
// 	return rv
// }

// femEqsChol - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:759
// func femEqsChol(a *tMatrix, b *tVector, x *tVector) int {
// 	// Choleski decomposition - complete
// 	// * @param a matrix (must  be a MAT_FULL)
// 	// * @return status
// 	//
// 	var rv int
// 	var sum float64
// 	var n int
// 	var i int
// 	var j int
// 	var k int
// 	var C tVector
// 	n = a.rows
// 	//femVecNull(((&C)))
// 	if femVecAlloc(((&C)), 0, n, n) != 0 {
// 		goto memFree
// 	}
// 	for i = 1; i <= n; i++ {
// 		for j = i; j <= n; j++ {
// 			sum = femMatGet(a, i, j)
// 			for k = i - 1; k >= 1; k-- {
// 				sum -= femMatGet(a, i, k) * femMatGet(a, j, k)
// 			}
// 			if i == j {
// 				if sum <= 0 {
// 					rv = -3
// 					fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("Given matrix is singular"))
// 					goto memFree
// 				}
// 				femVecPutAdd(((&C)), i, math.Sqrt(sum), 0)
// 			} else {
// 				femMatPutAdd(a, j, i, sum/femVecGet(((&C)), i), 0)
// 			}
// 		}
// 	}
// 	{
// 		// backward run:
// 		for i = 1; i <= n; i++ {
// 			sum = femVecGet(b, i)
// 			for k = i - 1; k >= 1; k-- {
// 				sum -= femMatGet(a, i, k) * femVecGet(x, k)
// 			}
// 			femVecPutAdd(x, i, sum/femVecGet(((&C)), i), 0)
// 		}
// 	}
// 	for i = n; i >= 1; i-- {
// 		sum = femVecGet(x, i)
// 		for k = i + 1; k <= n; k++ {
// 			sum -= femMatGet(a, k, i) * femVecGet(x, k)
// 		}
// 		femVecPutAdd(x, i, sum/femVecGet(((&C)), i), 0)
// 	}
// memFree:
// 	;
// 	// freeing of memory:
// 	//femVecFree(((&C)))
// 	return rv
// }

// femMatJacRotate - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:832
// func femMatJacRotate(a *tMatrix, i int, j int, k int, l int, g float64, h float64, s float64, tau float64) {
// 	// rotation for Jacobi computation of eigenvalues
// 	g = femMatGet(a, i, j)
// 	h = femMatGet(a, k, l)
// 	femMatPutAdd(a, i, j, g-s*(h+g*tau), 0)
// 	femMatPutAdd(a, k, l, h+s*(g-h*tau), 0)
// }

// femMatEigenJacobi - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:848
// func femMatEigenJacobi(a *tMatrix, d *tVector, v *tMatrix, nrot []int) int {
// 	// Compute eigen numbers and vectors (Jacobi method)
// 	// * @param a matrix to be analysed
// 	// * @param d vector to store eigenvalues
// 	// * @param v matrix to store eigenvectors
// 	// * @return status
// 	//
// 	var iters int = 100
// 	var i int
// 	var iq int
// 	var ip int
// 	var j int
// 	var n int
// 	var sm float64
// 	var tresh float64
// 	_ = tresh
// 	var g float64
// 	var h float64
// 	var t float64
// 	var c float64
// 	var theta float64
// 	var s float64
// 	var tau float64
// 	var checkp float64
// 	var checkq float64
// 	var checkh float64
// 	var b tVector
// 	var z tVector
// 	nrot = 0
// 	n = a.rows
// 	//femVecNull(((&b)))
// 	//femVecNull(((&z)))
// 	femVecAlloc(((&b)), 0, n, n)
// 	femVecAlloc(((&z)), 0, n, n)
// 	for i = 1; i <= n; i++ {
// 		femVecPutAdd(((&b)), i, femMatGet(a, i, i), 0)
// 		femVecPutAdd(d, i, femMatGet(a, i, i), 0)
// 		femVecPutAdd(((&z)), i, 0, 0)
// 		femMatPutAdd(v, i, i, 1, 0)
// 	}
// 	for i = 1; i <= iters; i++ {
// 		sm = 0
// 		for ip = 0; ip <= n-1; ip++ {
// 			for iq = ip + 1; iq <= n; iq++ {
// 				sm += math.Abs(femMatGet(a, ip, iq))
// 			}
// 		}
// 		if sm <= 1e-07 {
// 			// sum <= 0 so we are finished
// 			//printf("iterations: %d\n", *nrot);
// 			//femVecFree(((&b)))
// 			//femVecFree(((&z)))
// 			return 0
// 		}
// 		if i < 4 {
// 			tresh = 0.2 * sm / float64(n*n)
// 		} else {
// 			tresh = 0
// 		}
// 		for ip = 1; ip <= n-1; ip++ {
// 			for iq = ip + 1; iq <= n; iq++ {
// 				g = 100 * math.Abs(femMatGet(a, ip, iq))
// 				checkp = math.Abs(g*math.Abs(femVecGet(d, ip)) - math.Abs(femVecGet(d, ip)))
// 				checkq = math.Abs(g*math.Abs(femVecGet(d, iq)) - math.Abs(femVecGet(d, iq)))
// 				if i > 4 && checkp <= 1e-07 && checkq <= 1e-07 {
// 					// off-diagonal elements are small
// 					femMatPutAdd(a, ip, iq, 0, 0)
// 				} else {
// 					// still are big..
// 					h = femVecGet(d, iq) - femVecGet(d, ip)
// 					checkh = math.Abs(math.Abs(h) + g - math.Abs(h))
// 					if checkh < 1e-07 {
// 						if h != 0 {
// 							t = femMatGet(a, ip, iq) / h
// 						} else {
// 							t = 0
// 						}
// 					} else {
// 						theta = 0.5 * h / femMatGet(a, ip, iq)
// 						t = 1 / (math.Abs(theta) + math.Sqrt(1+math.Pow(theta, 2)))
// 						if theta < 0 {
// 							t = -1 * t
// 						}
// 					}
// 					c = 1 / math.Sqrt(1+math.Pow(t, 2))
// 					s = t * c
// 					tau = s / (1 + c)
// 					h = t * femMatGet(a, ip, iq)
// 					femVecPutAdd(((&z)), ip, femVecGet(((&z)), ip)-h, 0)
// 					femVecPutAdd(((&z)), iq, femVecGet(((&z)), iq)+h, 0)
// 					femVecPutAdd(d, ip, femVecGet(d, ip)-h, 0)
// 					femVecPutAdd(d, iq, femVecGet(d, iq)+h, 0)
// 					femMatPutAdd(a, ip, iq, 0, 0)
// 					for j = 1; j <= ip-1; j++ {
// 						femMatJacRotate(a, j, ip, j, iq, g, h, s, tau)
// 					}
// 					for j = ip + 1; j <= iq-1; j++ {
// 						femMatJacRotate(a, ip, j, j, iq, g, h, s, tau)
// 					}
// 					for j = iq + 1; j <= n; j++ {
// 						femMatJacRotate(a, ip, j, iq, j, g, h, s, tau)
// 					}
// 					for j = 1; j <= n; j++ {
// 						femMatJacRotate(v, j, ip, j, iq, g, h, s, tau)
// 					}
// 					nrot = nrot + 1
// 				}
// 			}
// 		}
// 		for ip = 1; ip <= n; ip++ {
// 			femVecPutAdd(((&b)), ip, femVecGet(((&z)), ip), 1)
// 			femVecPutAdd(d, ip, femVecGet(((&b)), ip), 0)
// 			femVecPutAdd(((&z)), ip, 0, 0)
// 		}
// 	}
// 	fmt.Fprintf(os.Stdout, string("[E] %s\n"), string("Out of iterations for eigendata"))
// 	return -1
// }

// femEqsCGwSSOR - transpiled function from  GOPATH/src/github.com/Konstantin8105/shell/c-src/shell/fem_eqs.c:1001
// func femEqsCGwSSOR(a *tMatrix, b *tVector, x *tVector, eps float64, maxIt int) int {
// 	// Conjugate gradient method with SSOR preconditioner
// 	// *  (for symetric matrices only!)
// 	// *  @param a      matrix
// 	// *  @param b      "load" vector
// 	// *  @param x      results (vector - given as first iteration)
// 	// *  @param eps    error (min.)
// 	// *  @param maxIt  max. number of iterations
// 	// *  @return state value
// 	//
// 	// Jacobi preconditioner (diag[A] ;-)
// 	var M tVector
// 	var r tVector
// 	var z tVector
// 	var zz tVector
// 	var p tVector
// 	var q tVector
// 	var ro float64
// 	var alpha float64
// 	var beta float64
// 	var roro float64
// 	var n int
// 	var i int
// 	var ii int
// 	var j int
// 	var ipos int
// 	var rv int
// 	var converged int
// 	var normRes float64
// 	var normX float64
// 	var normA float64
// 	var normB float64
// 	var val float64
// 	if a.cols != x.rows || x.rows != b.rows {
// 		return -9
// 	}
// 	n = a.rows
// 	normA = femMatNormBig(a)
// 	normB = femVecNormBig(b)
// 	if normB <= 0 {
// 		femVecSetZeroBig(x)
// 		fmt.Fprintf(os.Stdout, string("[ ]  %s!\n"), string("solution done without iterations because of zero load"))
// 		return 0
// 	}
// 	// vector initialization
// 	//femVecNull(((&M)))
// 	//femVecNull(((&r)))
// 	//femVecNull(((&z)))
// 	//femVecNull(((&zz)))
// 	//femVecNull(((&p)))
// 	//femVecNull(((&q)))
// 	if (func() int {
// 		rv = femVecAlloc(((&M)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		// memory allocation
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&r)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&z)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&zz)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&p)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	if (func() int {
// 		rv = femVecAlloc(((&q)), 0, n, n)
// 		return rv
// 	}()) != 0 {
// 		goto memFree
// 	}
// 	{
// 		// Jacobi preconditioner creation:
// 		for i = 1; i <= n; i++ {
// 			val = femMatGet(a, i, i)
// 			if math.Abs(val) < 1e-07 {
// 				rv = -13
// 				fmt.Fprintf(os.Stdout, string("[ ]   %s[%d][%d] %s\n"), string("matrix member"), i, i, string("has zero size"))
// 				goto memFree
// 			}
// 			// NOTE: M includes inverse of diagonal
// 			M.data[i-1] = 1 / val
// 		}
// 	}
// 	// next two lines mean: r = b - A*x
// 	femMatVecMultBig(a, x, ((&r)))
// 	for i = 0; i < n; i++ {
// 		r.data[i] = b.data[i] - r.data[i]
// 	}
// 	{
// 		// main loop
// 		for i = 1; i <= maxIt; i++ {
// 			fmt.Fprintf(os.Stdout, string("[ ]   %s %d\n"), string("linear step"), i)
// 			if a.type_ != 1 {
// 				{
// 					// using preconditioner:
// 					for ii = 1; ii <= n; ii++ {
// 						val = 0
// 						for j = 1; j < ii; j++ {
// 							val += femMatGet(a, ii, j) * femVecGet(((&zz)), j)
// 						}
// 						femVecPutAdd(((&zz)), ii, femVecGet(((&M)), ii)*(femVecGet(((&r)), ii)-val), 0)
// 					}
// 				}
// 				for ii = n; ii >= 1; ii-- {
// 					val = 0
// 					for j = ii + 1; j <= n; j++ {
// 						val += femMatGet(a, ii, j) * femVecGet(((&z)), j)
// 					}
// 					femVecPutAdd(((&z)), ii, femVecGet(((&zz)), ii)-femVecGet(((&M)), ii)*val, 0)
// 				}
// 			} else {
// 				{
// 					// faster code for MAT_SPAR:
// 					for ii = 1; ii <= n; ii++ {
// 						val = 0
// 						for j = a.frompos[ii-1]; j < a.frompos[ii-1]+a.defpos[ii-1]; j++ {
// 							ipos = a.pos[j]
// 							if ipos >= ii || ipos < 1 {
// 								continue
// 							}
// 							val += a.data[j] * zz.data[ipos-1]
// 						}
// 						femVecPutAdd(((&zz)), ii, femVecGet(((&M)), ii)*(femVecGet(((&r)), ii)-val), 0)
// 					}
// 				}
// 				for ii = n; ii >= 1; ii-- {
// 					val = 0
// 					for j = a.frompos[ii-1]; j < a.frompos[ii-1]+a.defpos[ii-1]; j++ {
// 						ipos = a.pos[j]
// 						if ipos > ii {
// 							val += a.data[j] * z.data[ipos-1]
// 						}
// 					}
// 					femVecPutAdd(((&z)), ii, femVecGet(((&zz)), ii)-femVecGet(((&M)), ii)*val, 0)
// 				}
// 			}
// 			// end of preconditioning
// 			ro = femVecVecMultBig(((&r)), ((&z)))
// 			fmt.Fprintf(os.Stdout, string("ro = %f\n"), ro)
// 			if i == 1 {
// 				for j = 0; j < n; j++ {
// 					p.data[j] = z.data[j]
// 				}
// 			} else {
// 				beta = ro / roro
// 				fmt.Fprintf(os.Stdout, string("beta = %f\n"), beta)
// 				for j = 0; j < n; j++ {
// 					p.data[j] = z.data[j] + beta*p.data[j]
// 				}
// 			}
// 			femMatVecMultBig(a, ((&p)), ((&q)))
// 			alpha = ro / femVecVecMultBig(((&p)), ((&q)))
// 			fmt.Fprintf(os.Stdout, string("alpha = %f\n"), alpha)
// 			for j = 0; j < n; j++ {
// 				x.data[j] = x.data[j] + alpha*p.data[j]
// 				r.data[j] = r.data[j] - alpha*q.data[j]
// 			}
// 			// Convergence testing
// 			normRes = femVecNormBig(((&r)))
// 			normX = femVecNormBig(x)
// 			if normRes <= eps*(normA*normX+normB) {
// 				// convergence test
// 				//if (fabs(norm - norm0) < eps )
// 				converged = 1
// 				fmt.Fprintf(os.Stdout, string("[ ]  %s %d %s!\n"), string("linear solution done in"), i, string("iterations"))
// 				break
// 			}
// 			fmt.Fprintf(os.Stdout, string("[i] Convergence test %f < %f (step %d from %d)\n"), normRes, eps*(normA*normX+normB), i, maxIt)
// 			roro = ro
// 		}
// 	}
// 	if converged != 1 {
// 		// end of main loop
// 		//fprintf(os.Stdout,"[I] normRes = %f\n",normRes);
// 		fmt.Fprintf(os.Stdout, string("[E] %s!\n"), string("unconverged solution"))
// 		rv = -1
// 	}
// memFree:
// 	;
// 	// freeing memory:
// 	//femVecFree(((&M)))
// 	//femVecFree(((&r)))
// 	//femVecFree(((&z)))
// 	//femVecFree(((&zz)))
// 	//femVecFree(((&p)))
// 	//femVecFree(((&q)))
// 	return rv
// }

// c4goUnsafeConvert_float64 : created by c4go
// func c4goUnsafeConvert_float64(c4go_name *float64) []float64 {
// 	return (*[1000000]float64)(unsafe.Pointer(c4go_name))
// }

// c4goUnsafeConvert_int : created by c4go
// func c4goUnsafeConvert_int(c4go_name *int) []int {
// 	return (*[1000000]int)(unsafe.Pointer(c4go_name))
// }

// 0 = dense; 1 = sparse (rows)
// lenght of "pos" and "data" (if used) fields
// Functions:
// Use with care:  (!!)
// end of fem_math.h
// ========================================================
// end of eshell.c
// end of fem_math.c
// end of fem_eqs.c
// end of fem_mem.c
