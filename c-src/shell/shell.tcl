#! /usr/freeware/bin/wish -f
#
# TODO:
# - elastic supports
# - dialog for setting of canvas size and grid step
# - more intelligent replotting of stuff in "mouse_action"
# - results ;-)
# - saving of graphics
# - text output
# - color switching (black/white)
# - toolbar/gizmo
#

# Compatibility stuff #############################################
proc abs {a} {
  if {$a < 0.0} { return [expr -1 * $a] } else { return $a }
}

set PI 3.1415926535897931
set dialog  .d
set dialog1 .d1
set dialog2 .d1
set activesel 0 ;# to implement (broken?) modal behaviour
set statusbar "Ready..."
set pickmode 0 ;# pick items: 0=no, 1=nodes, 2=elements 3=disps..
set workmode 0 ;# work mode: 0=no 1=add 2=edit 3=delete
set enode_id 0 ;# 1st node of new element (id number)
set processing 0 ;# 0=preprocessing, 1=postprocessing

# Colors ##########################################################
set black "black"
set white "darkgray"

# Motion stuff:   ###############################################
set m_x1 0
set m_y1 0
set m_line 0
set m_button 0
	
# Data structures #################################################

set data_file ""

proc make_empty_data {} {
  global mlen mat_E1 mat_nu1 mat_grav mat_price mat_width
  global nlen n_id n_x n_y
  global elen e_id e_na e_nb e_h e_E
  global dlen d_id d_n d_dir d_val 
  global flen f_id f_n f_dir f_val 
  global w_top w_bot w_force w_min w_max

# Materials:
set mlen 0
set mat_E1    [list]
set mat_nu1   [list]
set mat_grav  [list]
set mat_price [list]
set mat_width [list]
  
# Nodes:
set nlen 0
set n_id [list] ;# if=0 the node doesn't exist
set n_x [list]
set n_y [list]

# Elements:
set elen 0
set e_id [list] ;# if=0 the element doesn't exist
set e_na [list]
set e_nb [list]
set e_h [list]
set e_E [list]

# Displacements
set dlen 0
set d_id [list]  
set d_n [list]  ;# node number
set d_dir [list] ;# x=0, y=1, rotz=2, stiffx=3, stiffy=4, stiffrz=5
set d_val [list]

# Forces:
set flen 0
set f_id [list]  
set f_n [list]  ;# node number
set f_dir [list] ;# fx=0, fy=1, mz=2
set f_val [list]

# Water:
set w_top 0
set w_bot 0
set w_force 0
set w_min 0
set w_max 0

}

# Default values:
set def_id     0
set def_x      0.5
set def_y      0.5
set def_ha     1
set def_hb     1
set def_E      0
set def_h      0.1
set def_dval   0
set def_n     -1
set def_id1   -1
set def_id2   -1
set def_id3   -1
set def_is1    0
set def_is2    0
set def_is3    0
set def_type1  0
set def_type2  0
set def_type3  0
set def_dir1   0
set def_dir2   0
set def_dir3   0
set def_dval1  0
set def_dval2  0
set def_dval3  0
set def_lc     1
set def_le     0
set def_ltype  0
set def_ldir   0
set def_lvala  0.0
set def_lvalb  0.0
set def_fld [list]
set def_maxlc  0
set def_fldlen 0

proc unsetData {} {
  global mlen mat_E1 mat_nu1 mat_grav mat_price mat_width
  global  nlen  n_id  n_x  n_y
  global  elen  e_id  e_na  e_nb  e_h  e_E
  global  dlen  d_id  d_n  d_dir  d_val  
  global  flen  f_id  f_n  f_dir  f_val  
  global w_top w_bot w_force w_min w_max

  ##############################

  # Materials:
  unset mlen 
  unset mat_E1    
  unset mat_nu1   
  unset mat_grav  
  unset mat_price 
  unset mat_width


  # Nodes:
  unset nlen
  unset n_id
  unset n_x
  unset n_y

  # Elements:
  unset elen
  unset e_id
  unset e_na
  unset e_nb
  unset e_h
  unset e_E

  # Displacements
  unset dlen
  unset d_id
  unset d_n
  unset d_dir
  unset d_val

  # Forces:
  unset flen
  unset f_id
  unset f_n
  unset f_dir
  unset f_val

  # Water
  unset w_top
  unset w_bot
  unset w_force
  unset w_min
  unset w_max 
}


# I/O operations ##################################################

proc read_data {fname} {
  global mlen mat_E1 mat_nu1 mat_grav mat_price mat_width
  global nlen n_id n_x n_y
  global elen e_id e_na e_nb e_h e_E
  global dlen d_id d_n d_dir d_val 
  global flen f_id f_n f_dir f_val 
  global w_top w_bot w_force w_min w_max

  if {[file exists $fname]} {
    if {[file readable $fname]} {
        # open file:
        if [ catch {set fr [open $fname "r"]} result] { puts $result ; return -1 }

        # data sizes
        scan [ gets $fr ] "%li %li %li %li %li" mlen nlen elen dlen flen

        # materials:
        if {$mlen <= 0} { close $fr ; puts "I/O error (materials)!" ; return -1 }
        if {$mlen > 0} {
          for {set i 0} {$i< $mlen} {incr i} {
            scan [gets $fr] "%e %e %e %e %e %e %e %e" E10 E20 G0 nu10 nu20 grav0 price0 width0
            lappend mat_E1 $E10
            lappend mat_nu1 $nu10
            lappend mat_grav $grav0
            lappend mat_price $price0
            lappend mat_width $width0
          }
        }

        # nodes:
        if {$nlen <= -1} { close $fr ; puts "I/O error (nodes)!" ; return -1 }
        if {$nlen > 0} {
          for {set i 0} {$i< $nlen} {incr i} {
            scan [gets $fr] " %e %e" nx0 ny0
            lappend n_id [expr $i+1]
            lappend n_x $nx0
            lappend n_y $ny0
          }
        }

        # elements:
        if {$elen <= -1} { close $fr ; puts "I/O error!" ; return -1 }
        if {$elen > 0} {
          for {set i 0} {$i< $elen} {incr i} {
            scan [gets $fr] "%li %li %li %e" e_na0 e_nb0 e_E0 e_h0
            lappend e_id    [expr $i+1]
            lappend e_na    [expr $e_na0 +1]
            lappend e_nb    [expr $e_nb0 +1]
            lappend e_E     $e_E0 ; # material type
            lappend e_h     $e_h0
          }
        }

        # displacements:
        if {$dlen <= -1} { close $fr ; puts "I/O error!" ; return -1 }
        if {$dlen > 0} {
          for {set i 0} {$i< $dlen} {incr i} {
            scan [gets $fr] " %li %li %e" d_n0 d_dir0 d_val0
            lappend d_id  [expr $i+1]
            lappend d_n   [expr $d_n0+1]
            lappend d_dir $d_dir0
            lappend d_val $d_val0
          }
        }

        # forces:
        if {$flen <= -1} { close $fr ; puts "I/O error!" ; return -1 }
        if {$flen > 0} {
          for {set i 0} {$i< $flen} {incr i} {
            scan [gets $fr] " %li %li %e" f_n0 f_dir0 f_val0
            lappend f_id  [expr $i+1]
            lappend f_n   [expr $f_n0+1]
            lappend f_dir $f_dir0
            lappend f_val $f_val0
          }
        }

        # water:
        scan [gets $fr] "%li %li %e %li %li" w_top0 w_bot0 w_force0 w_min0 w_max0
        lappend w_top0    $w_top
        lappend w_bot0    $w_bot
        lappend w_force0  $w_force
        lappend w_min0    $w_min
        lappend w_max0    $w_max

        # close file:
        if [ catch [close $fr] result] { puts $result ; return -1 }

      } else { 
        puts "Can not read file!" 
        return -1
      }
    } else { 
      puts "Requested file not found!" 
      return -1
  }
  return 0
}

proc write_data { fname } { ;# TODO
  global mlen mat_E1 mat_nu1 mat_grav mat_price mat_width
  global nlen n_id n_x n_y
  global elen e_id e_na e_nb e_h e_E
  global dlen d_id d_n d_dir d_val 
  global flen f_id f_n f_dir f_val 
  global w_top w_bot w_force w_min w_max
  
    #if {[file writable $fname]} {
        # open file:
        if [ catch {set fr [open $fname "w"]} result] { puts $result ; return -1 }

        puts -nonewline $fr $mlen ; puts -nonewline $fr " "
        puts -nonewline $fr $nlen  ; puts -nonewline $fr " "
        puts -nonewline $fr $elen  ; puts -nonewline $fr " "
        puts -nonewline $fr $dlen  ; puts -nonewline $fr " "
        puts $fr $flen


        # materials:
        if {$mlen > 0} {
          for {set i 0} {$i< $mlen} {incr i} {
            puts -nonewline $fr [lindex $mat_E1 $i] ; puts -nonewline $fr " "
            puts -nonewline $fr "0" ; puts -nonewline $fr " "
            puts -nonewline $fr "0" ; puts -nonewline $fr " "
            puts -nonewline $fr [lindex $mat_nu1 $i] ; puts -nonewline $fr " "
            puts -nonewline $fr "0" ; puts -nonewline $fr " "
            puts -nonewline $fr [lindex $mat_grav $i] ; puts -nonewline $fr " "
            puts -nonewline $fr [lindex $mat_price $i] ; puts -nonewline $fr " "
            puts $fr [lindex $mat_width $i]
          }
        }

        # nodes:
        if {$nlen > 0} {
          for {set i 0} {$i< $nlen} {incr i} {
            puts -nonewline $fr [lindex $n_x $i] ; puts -nonewline $fr " "
            puts $fr [lindex $n_y $i]
          }
        }

        # elements:
        if {$elen > 0} {
          for {set i 0} {$i< $elen} {incr i} {
            puts -nonewline $fr [expr [ lindex $e_na $i] -1] ; puts -nonewline $fr " "
            puts -nonewline $fr [expr [ lindex $e_nb $i] -1] ; puts -nonewline $fr " "
            puts -nonewline $fr [ lindex $e_E  $i] ; puts -nonewline $fr " "
            puts $fr [ lindex $e_h     $i]
          }
        }

        # displacements:
        if {$dlen > 0} {
          for {set i 0} {$i< $dlen} {incr i} {
            puts -nonewline $fr [expr [ lindex $d_n $i] -1] ; puts -nonewline $fr " "
            puts -nonewline $fr [ lindex $d_dir $i] ; puts -nonewline $fr " "
            puts $fr [ lindex $d_val $i]
          }
        }

        # forces:
        if {$flen > 0} {
          for {set i 0} {$i< $flen} {incr i} {
            puts -nonewline $fr [expr [ lindex $f_n $i ] -1] ; puts -nonewline $fr " "
            puts -nonewline $fr [ lindex $f_dir $i ] ; puts -nonewline $fr " "
            puts $fr [ lindex $f_val $i ]
          }
        }

        # water:
        puts -nonewline $fr $w_top   ; puts -nonewline $fr " "
        puts -nonewline $fr $w_bot   ; puts -nonewline $fr " "
        puts -nonewline $fr $w_force ; puts -nonewline $fr " "
        puts -nonewline $fr $w_min   ; puts -nonewline $fr " "
        puts $fr $w_max 

        # close file:
        if [ catch [close $fr] result] { puts $result ; return -1 }

      #} else { 
        #return -1 
        ##puts "Can not write file!" 
      #}
  return 0 
}

proc get_max_id {id} {
  set max 0

  set len [llength $id]

  if {$len > 0} {
    for {set i 0} {$i<$len} {incr i} {
      if {$max < [lindex $id $i]} {
        set max [lindex $id $i]
      }
    }
    incr max
    return $max
  } else {
    return 1
  }
}

# test if such element exists; if yes returnes its id 
proc check_elem_exists {na nb} {
  global elen e_id e_na e_nb e_h e_E

	for {set i 0} {$i < $elen} {incr i} {
		if {$na == [lindex $e_na $i]} {
			if {$nb == [lindex $e_nb $i]} {
				return [lindex $e_id $i]
			}
		}
		if {$na == [lindex $e_nb $i]} {
			if {$nb == [lindex $e_na $i]} {
				return [lindex $e_id $i]
			}
		}
	}
	return -1
}

# test if such displacements exists; if yes returnes its id 
proc check_disp_exists {n dir lc} {
  global dlen d_id d_n d_dir d_val d_lc

	for {set i 0} {$i < $dlen} {incr i} {
		if {$lc == [lindex $d_lc $i]} {
			if {$n == [lindex $d_n $i]} {
				if {$dir == [lindex $d_dir $i]} {
					return [lindex $d_id $i]
				}
				if {[expr $dir - 3 ] == [lindex $d_dir $i]} {
					return [lindex $d_id $i]
				}
				if {[ expr $dir + 3 ] == [lindex $d_dir $i]} {
					return [lindex $d_id $i]
				}

			}
		}
	}
	return 0
}


# test if such force exists; if yes returnes its id 
proc check_force_exists {n dir lc} {
  global flen f_id f_n f_dir f_val f_lc

	for {set i 0} {$i < $flen} {incr i} {
		if {$lc == [lindex $f_lc $i]} {
			if {$n == [lindex $f_n $i]} {
				if {$dir == [lindex $f_dir $i]} {
					return [lindex $f_id $i]
				}
			}
		}
	}
	return 0
}

# adds new node:
proc add_node { x y } {
  global nlen n_id n_x n_y

  if {$x <= 0.0} { return -1  ; }
  if {$y <= 0.0} { return -1  ; }
  
  if {$nlen > 0} {
    set id [get_max_id $n_id]
  } else {
    set id 1
  }

  lappend n_id $id
  lappend n_x $x
  lappend n_y $y
  incr nlen
}

# add new element
proc add_elements { na nb ha hb A I h E ro alpha } {
  global n_id
  global elen e_id e_na e_nb e_h e_E

  if {[lsearch $n_id $na] < 0} {
    return -1  ;
  }
  if {[lsearch $n_id $nb] < 0} {
    return -1  ;
  }

  if {$E < 0} { return -1  ; }

  if {$ha < 0} { return -1  ; }
  if {$hb < 0} { return -1  ; }
  
  if {$elen > 0} {
    set id [get_max_id $e_id]
  } else {
    set id 1
  }

  lappend e_id $id
  lappend e_na $na
  lappend e_nb $nb
  lappend e_h $h
  lappend e_E $E
  incr elen
}

# add new displacement
proc add_disp {n dir val lc} {
  global n_id
  global dlen d_id  d_n d_dir d_val d_lc

  if {[lsearch $n_id $n] < 0} {
    return -1  ;
  }

  set id [check_disp_exists $n $dir $lc]

  if {$id > 0} {
    change_disp $id $n $dir $val $lc
  } else {

    if {$dlen > 0} {
      set id [get_max_id $d_id]
    } else {
      set id 1
    }

    lappend d_id $id
    lappend d_n $n
    lappend d_dir $dir
    lappend d_val $val
    lappend d_lc $lc
    incr dlen
  }
}

# add new force
proc add_force {n dir val lc} {
  global n_id
  global flen f_id  f_n f_dir f_val f_lc
  global def_lc

  if {$def_lc <= 0} {
    tk_messageBox -type ok -title "Error!" -icon  "error" -message "No load allowed for load case 0"
    return -1
  }

  if {[lsearch $n_id $n] < 0} {
    return -1  ;
  }

  if {$val == 0.0} {
    return -1
  }

  set id [check_force_exists $n $dir $lc]

  if {$id > 0} {
    change_force $id $n $dir $val $lc
  } else {
    if {$flen > 0} {
      set id [get_max_id $f_id]
    } else {
      set id 1
    }


    lappend f_id $id
    lappend f_n $n
    lappend f_dir $dir
    lappend f_val $val
    lappend f_lc $lc
    incr flen
  }
}

# deletes a node: 
proc delete_node { id } {
  global nlen n_id n_x n_y
  global e_na e_nb d_n f_n

  set pos [lsearch $n_id $id]

  puts $n_id
  puts $id
  puts -nonewline "Deleteing node: "
  puts $pos

  if {$pos < 0} {
    return -1  ;
  } else {
    if {[lsearch $e_na $id] >= 0} {return -2 }
    if {[lsearch $e_nb $id] >= 0} {return -2 }
    if {[lsearch $d_n $id] >= 0} {return -2 }
    if {[lsearch $f_n $id] >= 0} {return -2 }

    set n_id0 [list]
    set n_x0 [list]
    set n_y0 [list]

    for {set i 0} {$i < [expr $nlen]} {incr i} {
      if {$i != $pos} {
        lappend n_id0 [lindex $n_id $i] 
        lappend n_x0 [lindex $n_x $i] 
        lappend n_y0 [lindex $n_y $i] 
      }
    }
    unset n_id
    unset n_x
    unset n_y
    
    incr nlen -1

    set n_id $n_id0
    set n_x $n_x0
    set n_y $n_y0
  }
  return 1
}

# deletes an element:
proc delete_elem { id } {
  global elen e_id e_na e_nb e_h e_E
  global l_e

  set pos [lsearch $e_id $id]

  if {$pos < 0} {
    return -1  ;
  } else {
    if {[lsearch $l_e $id] >= 0} {return -2 }

    set e_id0 [list]
    set e_na0 [list]
    set e_nb0 [list]
    set e_E0 [list]
    set e_h0 [list]


    for {set i 0} {$i < [expr $elen]} {incr i} {
      if {$i != $pos} {
        lappend e_id0 [lindex $e_id $i] 
        lappend e_na0 [lindex $e_na $i] 
        lappend e_nb0 [lindex $e_nb $i] 
        lappend e_E0 [lindex $e_E $i] 
        lappend e_h0 [lindex $e_h $i] 
      }
    }
    unset e_id    
    unset e_na    
    unset e_nb    
    unset e_E     
    unset e_h     
    
    incr elen -1

    set e_id    $e_id0
    set e_na    $e_na0
    set e_nb    $e_nb0
    set e_E     $e_E0
    set e_h     $e_h0

  }
  return 1
}

# deletes a displacement:
proc delete_disp { id } {
  global dlen d_id d_n d_dir d_val d_lc

  set pos [lsearch $d_id $id]

  if {$pos < 0} {
    return -1  ;
  } else {

    set d_id0 [list]
    set d_n0 [list]
    set d_dir0 [list]
    set d_val0 [list]
    set d_lc0 [list]


    for {set i 0} {$i < [expr $dlen]} {incr i} {
      if {$i != $pos} {
        lappend d_id0 [lindex $d_id $i] 
        lappend d_n0 [lindex $d_n $i] 
        lappend d_dir0 [lindex $d_dir $i] 
        lappend d_val0 [lindex $d_val $i] 
        lappend d_lc0 [lindex $d_lc $i] 
      }
    }
    
    incr dlen -1

    unset d_id   
    unset d_n    
    unset d_dir  
    unset d_val  
    unset d_lc   

    set d_id   $d_id0
    set d_n    $d_n0
    set d_dir  $d_dir0
    set d_val  $d_val0
    set d_lc   $d_lc0

  }
  return 1
}

# deletes a force:
proc delete_force { id } {
  global flen f_id f_n f_dir f_val f_lc

  set pos [lsearch $f_id $id]

  if {$pos < 0} {
    return -1  ;
  } else {

    set f_id0 [list]
    set f_n0 [list]
    set f_dir0 [list]
    set f_val0 [list]
    set f_lc0 [list]


    for {set i 0} {$i < [expr $flen]} {incr i} {
      if {$i != $pos} {
        lappend f_id0 [lindex $f_id $i] 
        lappend f_n0 [lindex $f_n $i] 
        lappend f_dir0 [lindex $f_dir $i] 
        lappend f_val0 [lindex $f_val $i] 
        lappend f_lc0 [lindex $f_lc $i] 
      }
    }

    unset f_id   
    unset f_n    
    unset f_dir  
    unset f_val  
    unset f_lc   
    
    incr flen -1

    set f_id   $f_id0
    set f_n    $f_n0
    set f_dir  $f_dir0
    set f_val  $f_val0
    set f_lc   $f_lc0

  }
  return 1
}


# changes a node: 
proc change_node { id new_x new_y } {
  global nlen n_id n_x n_y

  set pos [lsearch $n_id $id]

  if {$pos < 0} {
    return -1  ;
  } else {
    set n_x [lreplace $n_x $pos $pos $new_x]
    set n_y [lreplace $n_y $pos $pos $new_y]
  puts $new_x; puts $new_y
  }
  return 1
}

# changes an element:
proc change_elem { id new_na new_nb new_ha new_hb new_A new_I new_E new_ro new_alpha} {
  global elen e_id e_na e_nb e_h e_E

  set pos [lsearch $e_id $id]

  if {$pos < 0} {
    return -1  ;
  } else {
    set e_na [lreplace $e_na $pos $pos $new_na]
    set e_nb [lreplace $e_nb $pos $pos $new_nb]
    set e_E [lreplace $e_E $pos $pos $new_E]
    set e_h [lreplace $e_h $pos $pos $new_I]
  }
  return 1
}

# changes a displacement: 
proc change_disp { id new_n new_dir new_val new_lc } {
  global dlen d_id d_n d_dir d_val d_lc

  set pos [lsearch $d_id $id]

  if {$pos < 0} {
    return -1  ;
  } else {
    set d_n [lreplace $d_n $pos $pos $new_n]
    set d_dir [lreplace $d_dir $pos $pos $new_dir]
    set d_val [lreplace $d_val $pos $pos $new_val]
    set d_lc [lreplace $d_lc $pos $pos $new_lc]
  }
  return 1
}

# changes a force: 
proc change_force { id new_n new_dir new_val new_lc } {
  global flen f_id f_n f_dir f_val f_lc

  set pos [lsearch $f_id $id]

  if {$pos < 0} {
    return -1  ;
  } else {
    set f_n [lreplace $f_n $pos $pos $new_n]
    set f_dir [lreplace $f_dir $pos $pos $new_dir]
    set f_val [lreplace $f_val $pos $pos $new_val]
    set f_lc [lreplace $f_lc $pos $pos $new_lc]
  }
  return 1
}

# Geometry finding functions ######################################

# Find node from real coordinates:
proc what_node_id {x y} {
  global nlen n_id n_x n_y
  set dl0 0.0
  set dl  0.0
  set id  -1

  for {set i 0} {$i < $nlen} {incr i} {
    set xx [lindex $n_x $i]
    set yy [lindex $n_y $i]
    set dl [expr sqrt(($y-$yy)*($y-$yy) + ($x-$xx)*($x-$xx)) ]
    if {$i <= 0} {
      set dl0 $dl
      set id [lindex $n_id $i]
    }

    if {$dl < $dl0} {
      set id [lindex $n_id $i]
      set dl0 $dl
    }
  }
  return $id
}

# Tests if node exists on given location (returns its ID)
proc if_node_exists {x y} {
  global nlen n_id n_x n_y
  
  for {set i 0} {$i < $nlen} {incr i} {
    if {[lindex $n_x $i] == $x} {
      if {[lindex $n_y $i] == $y} {
        return [lindex $n_id $i]
      }
    }
  }

  return -1
}

# Find element from real coordinates:
proc what_elem_id {x y} {
  global nlen n_id n_x n_y
  global elen e_id e_na e_nb e_h e_E

  set dl0 0.0
  set dl  0.0
  set id  -1

  for {set i 0} {$i < $elen} {incr i} {
    set na  [lsearch $n_id [lindex $e_na $i]]
    set nb  [lsearch $n_id [lindex $e_nb $i]]
    set xx [expr 0.5*([lindex $n_x $na] + [lindex $n_x $nb])]
    set yy [expr 0.5*([lindex $n_y $na] + [lindex $n_y $nb])]
    set dl [expr sqrt(($y-$yy)*($y-$yy) + ($x-$xx)*($x-$xx)) ]
    if {$i <= 0} {
      set dl0 $dl
      set id [lindex $e_id $i]
    }

    if {$dl < $dl0} {
      set id [lindex $e_id $i]
      set dl0 $dl
    }
  }
  return $id
}


# Dialogs      ####################################################

proc check_set_node {id x y} {
  global dialog grstep activesel

  if {$x < $grstep} {return -1}
  if {$y < $grstep} {return -1}

  change_node $id $x $y

  destroy $dialog
  set activesel 0
  plot_stuff
}

# Node (changes only):
proc node_dialog {id} {
  global dialog activesel
  global nlen n_id n_x n_y
  global def_id def_x def_y

  if {$id == -1} {return -1}

  if {$activesel > 0} {return -1}
  set activesel 1

  toplevel $dialog -takefocus 1
  wm title $dialog [format "Node %i" $id]
  wm transient $dialog .

  set pos [lsearch $n_id $id]
  set def_x [lindex $n_x $pos]
  set def_y [lindex $n_y $pos]
  set def_id $id

  label $dialog.l1 -text "X: " ; grid $dialog.l1 -row 1 

  entry $dialog.e1 -textvariable def_x -width 10 -justify right
  grid $dialog.e1 -row 1 -column 2 

  label $dialog.l2 -text "Y: " ; grid $dialog.l2 -row 2 

  entry $dialog.e2 -textvariable def_y -width 10 -justify right
  grid $dialog.e2 -row 2 -column 2 

  button $dialog.ok -text "OK" -command {check_set_node $def_id $def_x $def_y} -width 8
  grid $dialog.ok -row 3 

  button $dialog.cancel -text "Cancel" -command {set activesel 0 ; destroy $dialog} -width 8
  grid $dialog.cancel -row 3 -column 2

  focus $dialog.e1

  # "True" modal behaviour (see "http://wiki.tcl.tk/3541"):
  tkwait visibility .
  grab $dialog
  wm transient $dialog .
  wm protocol $dialog WM_DELETE_WINDOW {grab release $dialog; destroy $dialog}
  raise $dialog
  tkwait window $dialog
}

proc check_set_elem {id} {
  global dialog grstep activesel
  global def_id def_E def_A def_I def_h def_ro def_alpha  def_ha def_hb
  global elen e_id e_na e_nb

  if {$def_E <= 0} {return -1}
  if {$def_A <= 0} {return -1}
  if {$def_I <= 0} {return -1}
  if {$def_h <= 0} {return -1}
  if {$def_ro <= 0} {return -1}
  if {$def_alpha <= 0} {return -1}

  if {$id > 0} {
    # modification of existing element
    set pos [lsearch $e_id $id]
    if [change_elem $id [lindex $e_na $pos] [lindex $e_nb $pos] $def_ha $def_hb $def_A $def_I $def_E $def_ro $def_alpha] {
    } else {
      return -1
    }
  } 

  destroy $dialog
  set activesel 0
}

# Elements:
proc elem_dialog {id } {
  global dialog activesel
  global nlen n_id n_x n_y
  global elen e_id e_na e_nb e_h e_E
  global def_id
  global def_E def_A def_I def_h def_ro def_alpha def_ha def_hb

  # Modal behaviour:
  if {$activesel > 0} {return -1}
  set activesel 1

  toplevel $dialog -takefocus 1
  if {$id > 0} {
    wm title $dialog [format "Element %i" $id]
    set $def_id $id
  } else {
    wm title $dialog "New Element" 
    set $def_id -1
  }
  wm transient $dialog .

  if {$id > 0} {
    # existing
    set pos [lsearch $e_id $id]
    set def_na [lindex $e_na $pos]
    set def_nb [lindex $e_nb $pos]
    set def_E [lindex $e_E $pos]
    set def_h [lindex $e_h $pos]
  } else {
    # new: set possible values of material type TODO
  }

  label $dialog.l1 -text "Material type: " ; grid $dialog.l1 -row 1 -sticky w
  entry $dialog.e1 -textvariable def_E -width 14 -justify right
  grid $dialog.e1 -row 1 -column 2 

  label $dialog.l2 -text "Width: " ; grid $dialog.l2 -row 2 -sticky w
  entry $dialog.e2 -textvariable def_h -width 14 -justify right
  grid $dialog.e2 -row 2 -column 2 


  button $dialog.ok -text "OK" -command {check_set_elem $def_id } -width 8
  grid $dialog.ok -row 3 

  button $dialog.cancel -text "Cancel" -command {set activesel 0 ; destroy $dialog ; return -1} -width 8
  grid $dialog.cancel -row 3 -column 2

  focus $dialog.e1

  # "True" modal behaviour (see "http://wiki.tcl.tk/3541"):
  tkwait visibility .
  grab $dialog
  wm transient $dialog .
  wm protocol $dialog WM_DELETE_WINDOW {grab release $dialog; destroy $dialog}
  raise $dialog
  tkwait window $dialog
}

proc check_set_disp { n lc } {
  global dialog activesel
  global dlen d_id d_n d_dir d_val d_lc
  global def_id1 def_id2 def_id3
  global def_is1 def_is2 def_is3
  global def_type1 def_type2 def_type3
  global def_dir1 def_dir2 def_dir3
  global def_val1 def_val2 def_val3
  global def_n def_lc

  destroy $dialog
  set activesel 0

  if {$def_type1 == 0} {
    set def_dir1 0
  } else {
    set def_dir1 4
  }

  if {$def_type2 == 0} {
    set def_dir2 1
  } else {
    set def_dir2 5
  }

  if {$def_type3 == 0} {
    set def_dir3 2
  } else {
    set def_dir3 6
  }

  if {$n > 0} {

  if {$def_id1 <= 0} {
    if {$def_is1 > 0} {
      # new
      add_disp $n $def_dir1 $def_val1 $lc
    } else {
      # nothing to do
    }
  } else { 
    if {$def_is1 > 0} {
      # modification of existing
      change_disp $def_id1 $n $def_dir1 $def_val1 $lc
    } else {
      # delete 
      delete_disp $def_id1
    }
  }

  if {$def_id2 <= 0} {
    if {$def_is2 > 0} {
      # new
      add_disp $n $def_dir2 $def_val2 $lc
    } else {
      # nothing to do
    }
  } else { 
    if {$def_is2 > 0} {
      # modification of existing
      change_disp $def_id2 $n $def_dir2 $def_val2 $lc
    } else {
      # delete 
      delete_disp $def_id2
    }
  }

  if {$def_id3 <= 0} {
    if {$def_is3 > 0} {
      # new
      add_disp $n $def_dir3 $def_val3 $lc
    } else {
      # nothing to do
    }
  } else { 
    if {$def_is3 > 0} {
      # modification of existing
      change_disp $def_id3 $n $def_dir3 $def_val3 $lc
    } else {
      # delete 
      delete_disp $def_id3
    }
  }

  }
}

# Displacements:
proc disp_dialog { n lc } {
  global dialog activesel
  global nlen n_id
  global dlen d_id d_n d_dir d_val d_lc
  global def_id1 def_id2 def_id3
  global def_is1 def_is2 def_is3
  global def_type1 def_type2 def_type3
  global def_dir1 def_dir2 def_dir3
  global def_val1 def_val2 def_val3
  global def_n def_lc

  # Modal behaviour:
  if {$activesel > 0} {return -1}
  set activesel 1

  toplevel $dialog -takefocus 1
  wm title $dialog "Support"
  wm transient $dialog .

  set def_n $n
  set def_lc $lc
  set def_dir1 0.0
  set def_type1 0.0
  set def_val1 0.0
  set def_dir2 0.0
  set def_type2 0.0
  set def_val2 0.0
  set def_dir3 0.0
  set def_type3 0.0
  set def_val3 0.0
  set def_is1 -1
  set def_is2 -1
  set def_is3 -1
  set def_id1  0
  set def_id2  0
  set def_id3  0

  set def_id1 [check_disp_exists $n 0 $lc]
  if {$def_id1 < 0} { set def_id1 [check_disp_exists $n 3 $lc] }

  set def_id2 [check_disp_exists $n 1 $lc]
  if {$def_id2 < 0} { set def_id2 [check_disp_exists $n 4 $lc] }

  set def_id3 [check_disp_exists $n 2 $lc]
  if {$def_id3 < 0} { set def_id3 [check_disp_exists $n 5 $lc] }


  if {$def_id1 > 0} {
    set def_is1 1
    set pos [lsearch $f_id $def_id1]
    set dir [lindex $f_dir $pos]
    if {$dir == 0} {set def_type1 0}
    if {$dir == 3} {set def_type1 1}
    set def_val1 [lindex $f_val $pos]
  }

  if {$def_id2 > 0} {
    set def_is2 1
    set pos [lsearch $f_id $def_id2]
    set dir [lindex $f_dir $pos]
    if {$dir == 1} {set def_type2 0}
    if {$dir == 4} {set def_type2 1}
    set def_val2 [lindex $f_val $pos]
  }

  if {$def_id3 > 0} {
    set def_is3 1
    set pos [lsearch $f_id $def_id3]
    set dir [lindex $f_dir $pos]
    if {$dir == 2} {set def_type3 0}
    if {$dir == 5} {set def_type3 1}
    set def_val3 [lindex $f_val $pos]
  }

  checkbutton $dialog.c1 -text "U Y: " -variable def_is1
  grid $dialog.c1 -row 1 -column 1 -sticky w
  entry $dialog.e1  -textvariable def_val1 -justify right -width 8
  grid $dialog.e1 -row 1 -column 2 
  label $dialog.l1 -text "   "
  grid $dialog.l1 -row 1 -column 3 
  checkbutton $dialog.t1 -variable def_type1 -text "stiffness"
  grid $dialog.t1 -row 1 -column 4 

  checkbutton $dialog.c2 -text "U X: " -variable def_is2
  grid $dialog.c2 -row 2 -column 1 -sticky w
  entry $dialog.e2  -textvariable def_val2 -justify right -width 8
  grid $dialog.e2 -row 2 -column 2 
  label $dialog.l2 -text "   "
  grid $dialog.l2 -row 2 -column 3 
  checkbutton $dialog.t2 -variable def_type2 -text "stiffness"
  grid $dialog.t2 -row 2 -column 4 

  checkbutton $dialog.c3 -text "ROT: " -variable def_is3
  grid $dialog.c3 -row 3 -column 1 -sticky w
  entry $dialog.e3  -textvariable def_val3 -justify right -width 8
  grid $dialog.e3 -row 3 -column 2 
  label $dialog.l3 -text "   "
  grid $dialog.l3 -row 3 -column 3 
  checkbutton $dialog.t3 -variable def_type3 -text "stiffness"
  grid $dialog.t3 -row 3 -column 4 


  button $dialog.ok -text "OK" -command {check_set_disp $def_n $def_lc} -width 8
  grid $dialog.ok -row 4  -column 1

  button $dialog.cancel -text "Cancel" -command {set activesel 0 ; destroy $dialog ; return -1} -width 8
  grid $dialog.cancel -row 4 -column 4

  focus $dialog.e1

  # "True" modal behaviour (see "http://wiki.tcl.tk/3541"):
  tkwait visibility .
  grab $dialog
  wm transient $dialog .
  wm protocol $dialog WM_DELETE_WINDOW {grab release $dialog; destroy $dialog}
  raise $dialog
  tkwait window $dialog
}

#############################

proc check_set_force { n lc } {
  global dialog activesel
  global flen f_id f_n f_dir f_val f_lc
  global def_id1 def_id2 def_id3
  global def_val1 def_val2 def_val3
  global def_n def_lc

  destroy $dialog
  set activesel 0

  if {$n > 0} {

  if {$def_id1 <= 0} {
      if {$def_val1 != 0.0} { add_force $n 0 $def_val1 $lc}
  } else { 
      # modification of existing
      if {$def_val1 != 0.0} {
        change_force $def_id1 $n 0 $def_val1 $lc
      } else {
        delete_force $def_id1
      }
  }

  if {$def_id2 <= 0} {
      if {$def_val2 != 0.0} {add_force $n 1 $def_val2 $lc}
  } else { 
      # modification of existing
      if {$def_val2 != 0.0} {
        change_force $def_id2 $n 1 $def_val2 $lc
      } else {
        delete_force $def_id2
      }
  }

  if {$def_id3 <= 0} {
      if {$def_val3 != 0.0} {add_force $n 2 $def_val3 $lc}
  } else { 
      # modification of existing
      if {$def_val3 != 0.0} {
        change_force $def_id3 $n 2 $def_val3 $lc
      } else {
        delete_force $def_id3
      }
  }

  }
}

# Forces:
proc force_dialog { n lc } {
  global dialog activesel
  global nlen n_id
  global flen f_id f_n f_dir f_val f_lc
  global def_id1 def_id2 def_id3
  global def_val1 def_val2 def_val3
  global def_n def_lc

  # Modal behaviour:
  if {$activesel > 0} {return -1}
  set activesel 1

  toplevel $dialog -takefocus 1
  wm title $dialog "Support"
  wm transient $dialog .

  set def_n $n
  set def_lc $lc

  set def_val1 0.0
  set def_val2 0.0
  set def_val3 0.0
  set def_id1  0
  set def_id2  0
  set def_id3  0

  set def_id1 [check_force_exists $n 0 $lc]
  set def_id2 [check_force_exists $n 1 $lc]
  set def_id3 [check_force_exists $n 2 $lc]


  if {$def_id1 > 0} {
    set pos [lsearch $f_id $def_id1]
    set def_val1 [lindex $f_val $pos]
  }

  if {$def_id2 > 0} {
    set pos [lsearch $f_id $def_id2]
    set def_val2 [lindex $f_val $pos]
  }

  if {$def_id3 > 0} {
    set pos [lsearch $f_id $def_id3]
    set def_val3 [lindex $f_val $pos]
  }

  label $dialog.c1 -text "FX: "
  grid $dialog.c1 -row 1 -column 1 -sticky w
  entry $dialog.e1  -textvariable def_val1 -justify right -width 8
  grid $dialog.e1 -row 1 -column 2 

  label $dialog.c2 -text "FY: "
  grid $dialog.c2 -row 2 -column 1 -sticky w
  entry $dialog.e2  -textvariable def_val2 -justify right -width 8
  grid $dialog.e2 -row 2 -column 2 

  label $dialog.c3 -text "MZ: "
  grid $dialog.c3 -row 3 -column 1 -sticky w
  entry $dialog.e3  -textvariable def_val3 -justify right -width 8
  grid $dialog.e3 -row 3 -column 2 


  button $dialog.ok -text "OK" -command {check_set_force $def_n $def_lc} -width 8
  grid $dialog.ok -row 4  -column 1

  button $dialog.cancel -text "Cancel" -command {set activesel 0 ; destroy $dialog ; return -1} -width 8
  grid $dialog.cancel -row 4 -column 2

  focus $dialog.e1

  # "True" modal behaviour (see "http://wiki.tcl.tk/3541"):
  tkwait visibility .
  grab $dialog
  wm transient $dialog .
  wm protocol $dialog WM_DELETE_WINDOW {grab release $dialog; destroy $dialog}
  raise $dialog
  tkwait window $dialog
}



proc init_data_error {} {
  tk_messageBox -type ok -title "Error!" -icon  "error" -message "Invalid data!\n\nAll values must be non-negative and grid step should not be too small!"
}

# Initial dialog for "new" data:
proc check_set_init_data {} {
  global dialog
  global max_x max_y grstep
  global def_lc

  if {$max_x <= 0} {init_data_error ; return -1}
  if {$max_y <= 0} {init_data_error ; return -1}
  if {$grstep <= 0} {init_data_error ; return -1}
  if {[expr $grstep] < [expr $max_x/50]} {init_data_error ; return -1}
  if {[expr $grstep] < [expr $max_y/50]} {init_data_error ; return -1}

  set max_x [expr $max_x + 2*$grstep]
  set max_y [expr $max_y + 2*$grstep]

  destroy $dialog

  fileNew

  set def_lc 1
  plot_stuff
  plot_grid
  make_frame
}

proc init_dialog {} {
  global dialog activesel
  global max_x max_y grstep

  toplevel $dialog -takefocus 1
  wm title $dialog "Initial workspace"
  wm transient $dialog .

  label $dialog.l1 -text "Width: " ; grid $dialog.l1 -row 1 

  entry $dialog.e1 -textvariable max_x -width 10 -justify right
  grid $dialog.e1 -row 1 -column 2 

  label $dialog.l2 -text "Height: " ; grid $dialog.l2 -row 2 

  entry $dialog.e2 -textvariable max_y -width 10 -justify right
  grid $dialog.e2 -row 2 -column 2 

  label $dialog.l3 -text "Grid step: " ; grid $dialog.l3 -row 3 

  entry $dialog.e3 -textvariable grstep -width 10 -justify right
  grid $dialog.e3 -row 3 -column 2 

  button $dialog.ok -text "OK" -command {
    check_set_init_data
  } -width 8
  grid $dialog.ok -row 4

  button $dialog.cancel -text "Cancel" -command {
    set max_x 6
    set max_y 5
    set grstep 0.5
    destroy $dialog
    return -1
  } -width 8
  grid $dialog.cancel -row 4 -column 2

  focus $dialog.e1

  # "True" modal behaviour (see "http://wiki.tcl.tk/3541"):
  tkwait visibility .
  grab $dialog
  wm transient $dialog .
  wm protocol $dialog WM_DELETE_WINDOW {grab release $dialog; destroy $dialog}
  raise $dialog
  tkwait window $dialog
}



proc get_max_xy {} {
  global nlen n_id n_x n_y
  global max_x max_y 
  global c_w c_h

  set max_x 0.0
  set max_y 0.0

  for {set i 0} {$i < $nlen} {incr i} {
    if {$max_x < [lindex $n_x $i]} {set max_x [lindex $n_x $i]}
    if {$max_y < [lindex $n_y $i]} {set max_y [lindex $n_y $i]}
  }

  set max_x [expr $max_x * 1.20]
  set max_y [expr $max_y * 1.20]

  if {$max_x <= 0.0} {
    if {$max_y > 0.0} {
      set max_x [expr $max_y*(1.0*$c_w/$c_h)] 
    } else {
      set max_x 6.0
    }
  }

  if {$max_y <= 0.0} {
    if {$max_x <= 0.0} {
      set max_x 6.0
    }
    set max_y [expr $max_x*(1.0*$c_h/$c_w)] 
  }
}

# fit canvas
proc fit_canvas { } {
  global max_x max_y scalemult

  set scalemult 1.0

  get_max_xy

  if {$max_x > 0.0} {
    if {$max_y > 0.0} {
      compute_scale $max_x $max_y
    }
  }
}

proc fit_replt_canvas { } {
  fit_canvas
  plot_stuff
  plot_grid
  make_frame
}

proc zoom_canvas { step } {
  global scalemult max_x max_y

  if {[expr 1.0+$step] <= 0.0} {
    set scalemult [expr $scalemult * 1.0]
  } else {
    set scalemult [expr $scalemult * (1.0+$step)]
  }

  get_max_xy
  compute_scale $max_x $max_y
  plot_stuff
  plot_grid
  make_frame
}

# Drawing area ####################################################
canvas .c -background $white -width 600 -height 400 -relief sunken -background $black
pack .c -fill both -anchor nw -expand yes

# width and height:
set c_w  1 
set c_h  1
set mov 5
set scale  1.0
set scalemult  1.0
set grstep 0.5
set max_x 0
set max_y 0

proc switch_colors {} { 
  global white black

  if {$white == "white"} {
    set white "black"
    set black "white"
  } else {
    set white "white"
    set black "black"
  }
}

proc make_frame {} {
  global c_w c_h mov

  update idletasks
  set c_w  [winfo width .c]
  set c_h  [winfo height .c]

  .c create line $mov $mov [expr $c_w-$mov] $mov -fill darkgray
  .c create line $mov [expr $c_h-$mov] [expr $c_w-$mov] [expr $c_h-$mov] -fill darkgray
  .c create line [expr $c_w-$mov] $mov [expr $c_w-$mov] [expr $c_h-$mov] -fill darkgray
  .c create line $mov $mov $mov [expr $c_h-$mov] -fill darkgray

  set u [expr $c_h/8]
  .c create line $mov [expr $c_h-$mov-2*$mov] [expr $mov+$u] [expr $c_h-$mov] -fill darkgray
  .c create line $mov [expr $c_h-$mov-$u] [expr $mov+2*$mov] [expr $c_h-$mov] -fill darkgray

  .c create text [expr $mov+$u] [expr $c_h-2*$mov] -fill darkgray -text "x" -anchor sw
  .c create text [expr 2*$mov] [expr $c_h-$mov-$u] -fill darkgray -text "y" -anchor sw
}

# Plots huge point:
proc plot_point {c x y} {
  .c create line [expr $x-2] $y [expr $x+3] $y -fill red
  .c create line $x [expr $y-2] $x [expr $y+3] -fill red
}

# Computes real to canvas scale:
proc compute_scale {mmax_x mmax_y} {
  global c_w c_h mov scale scalemult
  set scalea 1.0
  set scaleb 1.0

  set scalea [expr ($c_w-2*$mov)/$mmax_x]
  set scaleb [expr ($c_h-2*$mov)/$mmax_y]

  if {$scalea <= $scaleb} {
    set scale $scalea
  } else {
    set scale $scaleb
  }

  if {$scale <= 0} { set scale 32 }

  set scale [expr $scale * $scalemult]
  
  return scale
}

# Computes screen coordinates:
proc x_pix {x} {
  global scale mov
  return [expr $x*$scale + $mov]
}

proc y_pix {y} {
  global scale mov c_h
  return [expr $c_h - $mov - $y*$scale]
}

# computes real coordinates from screen:
proc x_real {x} {
  global scale mov
  return [expr ($x-$mov)/$scale]
}

proc y_real {y} {
  global scale mov c_h
  return [expr ($c_h - 2*$mov - $y)/$scale]
}

# Grid plotting:
proc plot_grid {} {
  global grstep c_w c_h scale
  
  set istop [expr 2*$c_w/$scale]
  set jstop [expr 2*$c_h/$scale]
  set k 0

  for {set i 1} {$i < $istop} {incr i} {
    for {set j 1} {$j < $jstop} {incr j} {
      set x [x_pix [expr $i*$grstep]] 
      set y [y_pix [expr $j*$grstep]]
      .c create line [expr $x] $y [expr $x+1] $y -fill red -tags [incr k]
    }
  }
}

# Obtaining from grid:
proc what_grid_point_x {px} {
  global grstep c_w scale

  set gx -1
  
  set istop [expr 2*$c_w/$scale]

  for {set i 0} {$i < $istop} {incr i} {
      set x0 [expr $i*$grstep] 

      set rx [expr [x_real $px] - 0.5*$grstep]
      set rx1 [expr [x_real $px] + 0.5*$grstep]

      if {$rx <= $x0} {
        if {$rx1 > $x0} {
           return $x0 
        }
      }
  }
  return -1
}

proc what_grid_point_y {py} {
  global grstep c_h scale

  set gy -1
  
  set jstop [expr 2*$c_h/$scale]

    for {set j 0} {$j < $jstop} {incr j} {
      set y0 [expr $j*$grstep]

      set ry [expr [y_real $py] - 0.5*$grstep]
      set ry1 [expr [y_real $py] + 0.5*$grstep]

      if {$ry <= $y0} {
        if {$ry1 > $y0} {
          return $y0
        }
      }
    }
  return -1
}


# Plots node:
proc plot_node {c id x y} {
  $c create line [expr $x-2] $y [expr $x+3] $y -fill cyan
  $c create line $x [expr $y-2] $x [expr $y+3] -fill cyan
  #TODO: number
}

# element lenght
proc get_elem_lenght { xa ya xb yb } {
  return [expr sqrt( ($xa-$xb)*($xa-$xb) + ($ya-$yb)*($ya-$yb) )]
}

# rotation angle
proc get_rot_angle { xa ya xb yb } {
  global PI

  if {[expr $xb-$xa] == 0.0 } {
    if {[expr $yb-$ya] > 0.0} {
      return [expr $PI / 2]
    } else {
      return [expr -$PI / 2]
    }
  } else {
    if {[expr $xb-$xa] < 0.0} {
      if {[expr $yb-$ya] >= 0} {
        return [expr $PI + atan(($yb-$ya)/($xb-$xa))]   
      } else {
        return [expr atan(($yb-$ya)/($xb-$xa))]   
      }
    } else {
      if {[expr $xb-$xa] < 0.0} {
        if {[expr $yb-$ya] < 0.0} {
          return [expr $PI + atan(($yb-$ya)/($xb-$xa))]   
        } else {
          return [expr atan(($yb-$ya)/($xb-$xa))]   
        }
      } else {
        return [expr atan(($yb-$ya)/($xb-$xa))]   
      }
    }
  }
}

# rotated x coordinate
proc x_rot_pix { x0 pho x y } {
  return [expr $x0 + $x * cos($pho) + $y * sin($pho)]
}

# rotated y coordinate
proc y_rot_pix { y0 pho x y } {
  return [expr $y0 + $x * sin($pho) - $y * cos($pho)]
}

# Plots element:
proc plot_element {c id xa ya xb yb } {
  global white
  global mov

  set angle [get_rot_angle $xa $ya $xb $yb] 

  # line:
  $c create line $xa $ya $xb $yb -fill $white -width 3

  #TODO: number
}

# Plots support (displacement):
proc plot_disp {c id x y dir val color} {
  global mov

  switch $dir {
    1 { ;# ux
        #set val -10
        $c create line $x $y [expr $x-2*$mov] $y -fill $color
        $c create line [expr $x-2*$mov] [expr $y-$mov/2] [expr $x-2*$mov] [expr $y+$mov/2+1] -fill $color 
        
        if {$val > 0.0} {
          $c create line [expr $x-2*$mov] [expr $y-$mov/2] [expr $x-2*$mov+$mov] [expr $y] -fill $color 
          $c create line [expr $x-2*$mov] [expr $y+$mov/2] [expr $x-2*$mov+$mov] [expr $y] -fill $color 
        }
        if {$val < 0.0} {
          $c create line [expr $x-2*$mov] [expr $y-$mov/2] [expr $x-2*$mov-$mov] [expr $y] -fill $color 
          $c create line [expr $x-2*$mov] [expr $y+$mov/2] [expr $x-2*$mov-$mov] [expr $y] -fill $color 
        }
        if {$val != 0.0} {
          $c create text [expr $x-4*$mov] [expr $y-$mov*1] -text [abs $val] -fill $color -anchor sw
        }
    }
    0 { ;# uy
        #set val -20
        $c create line $x $y $x [expr $y+2*$mov]  -fill $color
        $c create line [expr $x-$mov/2] [expr $y+2*$mov] [expr $x+$mov/2+1] [expr $y+2*$mov] -fill $color
        
        if {$val > 0.0} {
          $c create line [expr $x-$mov/2] [expr $y+2*$mov] [expr $x] [expr $y+2*$mov-$mov] -fill $color
          $c create line [expr $x+$mov/2] [expr $y+2*$mov] [expr $x] [expr $y+2*$mov-$mov] -fill $color
        }
        if {$val < 0.0} {
          $c create line [expr $x-$mov/2] [expr $y+2*$mov] [expr $x] [expr $y+2*$mov+$mov] -fill $color
          $c create line [expr $x+$mov/2] [expr $y+2*$mov] [expr $x] [expr $y+2*$mov+$mov] -fill $color
        }
        if {$val != 0.0} {
          $c create text [expr $x+$mov] [expr $y+$mov*4] -text [abs $val] -fill $color -anchor sw
        }
    }
    2 { ;# rotz
        #set val 30
        $c create line [expr $x-$mov] [expr $y-$mov] [expr $x+$mov] [expr $y-$mov]  -fill $color
        $c create line [expr $x-$mov] [expr $y+$mov] [expr $x+$mov] [expr $y+$mov]  -fill $color
        $c create line [expr $x+$mov] [expr $y-$mov] [expr $x+$mov] [expr $y+$mov]  -fill $color
        $c create line [expr $x-$mov] [expr $y-$mov] [expr $x-$mov] [expr $y+$mov]  -fill $color
 
        if {$val > 0.0} {
          $c create line [expr $x+$mov] [expr $y-$mov-$mov/2] [expr $x+$mov+$mov] [expr $y-$mov] -fill $color 
          $c create line [expr $x+$mov] [expr $y-$mov+$mov/2] [expr $x+$mov+$mov] [expr $y-$mov] -fill $color 
          $c create line [expr $x+$mov] [expr $y-$mov-$mov/2] [expr $x+$mov] [expr $y-$mov+$mov/2+1] -fill $color 
        }
        if {$val < 0.0} {
          $c create line [expr $x+2*$mov] [expr $y-$mov-$mov/2] [expr $x+$mov] [expr $y-$mov] -fill $color 
          $c create line [expr $x+2*$mov] [expr $y-$mov+$mov/2] [expr $x+$mov] [expr $y-$mov] -fill $color 
          $c create line [expr $x+2*$mov] [expr $y-$mov-$mov/2] [expr $x+2*$mov] [expr $y-$mov+$mov/2+1] -fill $color 
        }
        if {$val != 0.0} {
          $c create text [expr $x+2*$mov] [expr $y-1*$mov] -text [abs $val] -fill $color -anchor sw
        }
    }
    3 { ;# stiff y TODO
    }
    4 { ;# stiff x TODO
    }
    5 { ;# stiff mz TODO
    }
  }
}

# Plots forces:
proc plot_force {c id x y dir val} {
  global mov
  
  #set dir 2 ;# for TESTING only!

  switch $dir {
    1 { ;# fx
        #set val -20000
        if {$val > 0.0} {
          $c create line $x $y [expr $x-6*$mov] $y -width 1 -fill red
          $c create line [expr $x-2*$mov] [expr $y-$mov] $x  $y -width 1 -fill red
          $c create line [expr $x-2*$mov] [expr $y+$mov] $x  $y -width 1 -fill red
          $c create line [expr $x-2*$mov] [expr $y+$mov] [expr $x-2*$mov] [expr $y-$mov] -width 1 -fill red
          $c create text [expr $x-6*$mov] [expr $y-$mov] -text [abs $val] -fill red -anchor sw
        } else {
          $c create line $x $y [expr $x+6*$mov] $y -width 1 -fill red
          $c create line [expr $x+2*$mov] [expr $y-$mov] $x  $y -width 1 -fill red
          $c create line [expr $x+2*$mov] [expr $y+$mov] $x  $y -width 1 -fill red
          $c create line [expr $x+2*$mov] [expr $y+$mov] [expr $x+2*$mov] [expr $y-$mov] -width 1 -fill red
          $c create text [expr $x+6*$mov] [expr $y-$mov] -text [abs $val] -fill red -anchor sw
        }
    }
    0 { ;# fy
        #set val -30000
        if {$val > 0.0} {
          $c create line $x $y $x [expr $y+6*$mov] -width 1 -fill red
          $c create line [expr $x+$mov] [expr $y+2*$mov] $x  $y -width 1 -fill red
          $c create line [expr $x-$mov] [expr $y+2*$mov] $x  $y -width 1 -fill red
          $c create line [expr $x+$mov] [expr $y+2*$mov] [expr $x-$mov] [expr $y+2*$mov] -width 1 -fill red
          $c create text [expr $x+$mov] [expr $y+6*$mov]  -text [abs $val] -fill red -anchor sw
        } else {
          $c create line $x $y $x [expr $y-6*$mov] -width 1 -fill red
          $c create line [expr $x+$mov] [expr $y-2*$mov] $x  $y -width 1 -fill red
          $c create line [expr $x-$mov] [expr $y-2*$mov] $x  $y -width 1 -fill red
          $c create line [expr $x+$mov] [expr $y-2*$mov] [expr $x-$mov] [expr $y-2*$mov] -width 1 -fill red
          $c create text [expr $x+$mov] [expr $y-6*$mov]  -text [abs $val] -fill red -anchor sw
        }
    }
    2 { ;# mz
        #set val  -40000
        if {$val > 0.0} {
          $c create line [expr $x-3*$mov] $y [expr $x-3*$mov] [expr $y+3*$mov] -width 1 -fill red
          $c create line [expr $x-3*$mov] [expr $y+3*$mov] [expr $x+3*$mov] [expr $y+3*$mov]   -width 1 -fill red
          $c create line [expr $x+3*$mov] [expr $y+3*$mov] [expr $x+3*$mov] [expr $y-3*$mov]   -width 1 -fill red
          $c create line [expr $x-3*$mov] [expr $y-3*$mov] [expr $x+3*$mov] [expr $y-3*$mov]   -width 1 -fill red

          $c create line [expr $x-3*$mov] [expr $y-3*$mov] [expr $x-1*$mov] [expr $y-2*$mov]   -width 1 -fill red
          $c create line [expr $x-3*$mov] [expr $y-3*$mov] [expr $x-1*$mov] [expr $y-4*$mov+1]   -width 1 -fill red
          $c create line [expr $x-1*$mov] [expr $y-2*$mov] [expr $x-1*$mov] [expr $y-4*$mov+1]   -width 1 -fill red

          $c create text [expr $x+4*$mov] [expr $y-0*$mov]  -text [abs $val] -fill red -anchor sw
        } else {
          $c create line [expr $x+3*$mov] $y [expr $x+3*$mov] [expr $y+3*$mov] -width 1 -fill red
          $c create line [expr $x+3*$mov] [expr $y+3*$mov] [expr $x-3*$mov] [expr $y+3*$mov]   -width 1 -fill red
          $c create line [expr $x-3*$mov] [expr $y+3*$mov] [expr $x-3*$mov] [expr $y-3*$mov]   -width 1 -fill red
          $c create line [expr $x+3*$mov] [expr $y-3*$mov] [expr $x-3*$mov] [expr $y-3*$mov]   -width 1 -fill red

          $c create line [expr $x+3*$mov] [expr $y-3*$mov] [expr $x+2*$mov] [expr $y-2*$mov]   -width 1 -fill red
          $c create line [expr $x+3*$mov] [expr $y-3*$mov] [expr $x+2*$mov] [expr $y-4*$mov]   -width 1 -fill red
          $c create line [expr $x+2*$mov] [expr $y-2*$mov] [expr $x+2*$mov] [expr $y-4*$mov]   -width 1 -fill red

          $c create text [expr $x+4*$mov] [expr $y-0*$mov]  -text [abs $val] -fill red -anchor sw
        }

    }
  }
}


# Plots element loads:
proc plot_eload {c e na nb id type dir val1 val2} {
  global mov n_x n_y

  # "x" and "y" of both element's nodes:
  set xa [lindex $n_x $na]
  set ya [lindex $n_y $na]
  set xb [lindex $n_x $nb]
  set yb [lindex $n_y $nb]

  set angle [get_rot_angle [x_pix $xa] [y_pix $ya] [x_pix $xb] [y_pix $yb]]

  switch $type {
    0 { ;# force
      # TODO: compute x and y of force
      set rpos [expr $val1 / [get_elem_lenght $xa $ya $xb $yb]]
      set xpix  [get_elem_lenght [x_pix $xa] [y_pix $ya] [x_pix $xb] [y_pix $yb]]

      set x [x_rot_pix [x_pix $xa] $angle [expr $xpix*$rpos] 0]
      set y [y_rot_pix [y_pix $ya] $angle [expr $xpix*$rpos] 0]

      puts $xpix
      puts $x
      puts $y

      #set x [expr ([x_pix $xa] + [x_pix $xb]) / 2.0 ]
      #set y [expr ([y_pix $ya] + [y_pix $yb]) / 2.0 ]

      # TODO:
  switch $dir {
    0 { ;# fx
        if {$val2 > 0.0} {
          set x1 [x_rot_pix $x $angle [expr -6*$mov] 0 ]
          set y1 [y_rot_pix $y $angle [expr -6*$mov] 0 ]
            $c create line $x $y $x1 $y1 -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr -2*$mov] [expr -$mov]]
          set y1 [y_rot_pix $y $angle [expr -2*$mov] [expr -$mov]]
            $c create line $x1 $y1 $x  $y -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr -2*$mov] [expr $mov]]
          set y1 [y_rot_pix $y $angle [expr -2*$mov] [expr $mov]]
            $c create line $x1 $y1 $x  $y -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr -2*$mov] [expr $mov]]
          set y1 [y_rot_pix $y $angle [expr -2*$mov] [expr $mov]]
          set x2 [x_rot_pix $x $angle [expr -2*$mov] [expr -$mov]]
          set y2 [y_rot_pix $y $angle [expr -2*$mov] [expr -$mov]]
            $c create line $x1 $y1 $x2 $y2 -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr 1*$mov] [expr -$mov]]
          set y1 [y_rot_pix $y $angle [expr 1*$mov] [expr -$mov]]
            $c create text $x1 $y1 -text [abs $val2] -fill red -anchor sw
        } else {
          set x1 [x_rot_pix $x $angle [expr 6*$mov] 0 ]
          set y1 [y_rot_pix $y $angle [expr 6*$mov] 0 ]
            $c create line $x $y $x1 $y1 -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr 2*$mov] [expr -$mov] ]
          set y1 [y_rot_pix $y $angle [expr 2*$mov] [expr -$mov] ]
            $c create line $x1 $y1 $x  $y -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr 2*$mov] [expr $mov] ]
          set y1 [y_rot_pix $y $angle [expr 2*$mov] [expr $mov] ]
            $c create line $x1 $y1 $x  $y -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr +2*$mov] [expr $mov]]
          set y1 [y_rot_pix $y $angle [expr +2*$mov] [expr $mov]]
          set x2 [x_rot_pix $x $angle [expr +2*$mov] [expr -$mov]]
          set y2 [y_rot_pix $y $angle [expr +2*$mov] [expr -$mov]]
            $c create line $x1 $y1 $x2 $y2 -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr +6*$mov] [expr -$mov]]
          set y1 [y_rot_pix $y $angle [expr +6*$mov] [expr -$mov]]
            $c create text $x1 $y1 -text [abs $val2] -fill red -anchor sw
        }
    }
    1 { ;# fy
        #set val -30000
        if {$val2 > 0.0} {
          set x1 [x_rot_pix $x $angle 0 [expr +6*$mov] ]
          set y1 [y_rot_pix $y $angle 0 [expr +6*$mov] ]
            $c create line $x $y $x1 $y1 -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr $mov] [expr 2*$mov] ]
          set y1 [y_rot_pix $y $angle [expr $mov] [expr 2*$mov] ]
            $c create line $x1 $y1 $x  $y -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr -$mov] [expr 2*$mov] ]
          set y1 [y_rot_pix $y $angle [expr -$mov] [expr 2*$mov] ]
            $c create line $x1 $y1 $x  $y -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr $mov] [expr 2*$mov] ]
          set y1 [y_rot_pix $y $angle [expr $mov] [expr 2*$mov] ]
          set x2 [x_rot_pix $x $angle [expr -$mov] [expr 2*$mov] ]
          set y2 [y_rot_pix $y $angle [expr -$mov] [expr 2*$mov] ]
            $c create line $x1 $y1 $x2 $y2 -width 1 -fill red
          set x1 [x_rot_pix $x $angle 0 [expr +6*$mov] ]
          set y1 [y_rot_pix $y $angle 0 [expr +6*$mov] ]
            $c create text $x1 $y1  -text [abs $val2] -fill red -anchor sw
        } else {
          set x1 [x_rot_pix $x $angle 0 [expr -6*$mov] ]
          set y1 [y_rot_pix $y $angle 0 [expr -6*$mov] ]
            $c create line $x $y $x1 $y1 -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr $mov] [expr -2*$mov] ]
          set y1 [y_rot_pix $y $angle [expr $mov] [expr -2*$mov] ]
            $c create line $x1 $y1 $x  $y -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr -$mov] [expr -2*$mov] ]
          set y1 [y_rot_pix $y $angle [expr -$mov] [expr -2*$mov] ]
            $c create line $x1 $y1 $x  $y -width 1 -fill red
          set x1 [x_rot_pix $x $angle [expr $mov] [expr -2*$mov] ]
          set y1 [y_rot_pix $y $angle [expr $mov] [expr -2*$mov] ]
          set x2 [x_rot_pix $x $angle [expr -$mov] [expr -2*$mov] ]
          set y2 [y_rot_pix $y $angle [expr -$mov] [expr -2*$mov] ]
            $c create line $x1 $y1 $x2 $y2 -width 1 -fill red
          set x1 [x_rot_pix $x $angle 0 [expr -7*$mov] ]
          set y1 [y_rot_pix $y $angle 0 [expr -7*$mov] ]
            $c create text $x1 $y1  -text [abs $val2] -fill red -anchor sw
        }
    }
    2 { ;# mz
        if {$val2 > 0.0} {
          $c create line [expr $x-3*$mov] $y [expr $x-3*$mov] [expr $y+3*$mov] -width 1 -fill red
          $c create line [expr $x-3*$mov] [expr $y+3*$mov] [expr $x+3*$mov] [expr $y+3*$mov]   -width 1 -fill red
          $c create line [expr $x+3*$mov] [expr $y+3*$mov] [expr $x+3*$mov] [expr $y-3*$mov]   -width 1 -fill red
          $c create line [expr $x-3*$mov] [expr $y-3*$mov] [expr $x+3*$mov] [expr $y-3*$mov]   -width 1 -fill red

          $c create line [expr $x-3*$mov] [expr $y-3*$mov] [expr $x-1*$mov] [expr $y-2*$mov]   -width 1 -fill red
          $c create line [expr $x-3*$mov] [expr $y-3*$mov] [expr $x-1*$mov] [expr $y-4*$mov+1]   -width 1 -fill red
          $c create line [expr $x-1*$mov] [expr $y-2*$mov] [expr $x-1*$mov] [expr $y-4*$mov+1]   -width 1 -fill red

          $c create text [expr $x+4*$mov] [expr $y-0*$mov]  -text [abs $val2] -fill red -anchor sw
        } else {
          $c create line [expr $x+3*$mov] $y [expr $x+3*$mov] [expr $y+3*$mov] -width 1 -fill red
          $c create line [expr $x+3*$mov] [expr $y+3*$mov] [expr $x-3*$mov] [expr $y+3*$mov]   -width 1 -fill red
          $c create line [expr $x-3*$mov] [expr $y+3*$mov] [expr $x-3*$mov] [expr $y-3*$mov]   -width 1 -fill red
          $c create line [expr $x+3*$mov] [expr $y-3*$mov] [expr $x-3*$mov] [expr $y-3*$mov]   -width 1 -fill red

          $c create line [expr $x+3*$mov] [expr $y-3*$mov] [expr $x+2*$mov] [expr $y-2*$mov]   -width 1 -fill red
          $c create line [expr $x+3*$mov] [expr $y-3*$mov] [expr $x+2*$mov] [expr $y-4*$mov]   -width 1 -fill red
          $c create line [expr $x+2*$mov] [expr $y-2*$mov] [expr $x+2*$mov] [expr $y-4*$mov]   -width 1 -fill red

          plot_point $c $x $y
          $c create text [expr $x+4*$mov] [expr $y-0*$mov]  -text [abs $val2] -fill red -anchor sw
        }
    }
  }

    }
    1 { ;# continuous load
      # TODO

      if {$val1 > $val2 } {
        set mova [expr 4*$mov]
        set movb [expr 2*$mov]
      } else {
        if {$val1 == $val2} {
          set mova [expr 4*$mov]
          set movb [expr 4*$mov]
        } else {
          set mova [expr 2*$mov]
          set movb [expr 4*$mov]
        }
      }
      if {$val1 == 0.0} { set mova 0 }
      if {$val1 == 0.0} { set movb 0 }

      switch $dir {
        0 { ;# fx
          if {$val1 < 0} {
            set mova [expr -1*$mova]
          } else {
            if {$val1 == 0} {
              set mova 0
            }
          }

          if {$val2 < 0} {
            set movb [expr -1*$movb]
          } else {
            if {$val1 == 0} {
              set mova 0
            }
          }

          set x1 [x_pix $xa]
          set y1 [y_pix $ya]
          set x2 [x_rot_pix [x_pix $xa] $angle 0 [expr -$mova] ]
          set y2 [y_rot_pix [y_pix $ya] $angle 0 [expr -$mova] ]
            $c create line $x1 $y1 $x2  $y2 -width 1 -fill red
          set x1 [x_pix $xb]
          set y1 [y_pix $yb]
          set x2 [x_rot_pix [x_pix $xb] $angle 0 [expr -$movb] ]
          set y2 [y_rot_pix [y_pix $yb] $angle 0 [expr -$movb] ]
            $c create line $x1 $y1 $x2  $y2 -width 1 -fill red
          set x1 [x_rot_pix [x_pix $xa] $angle 0 [expr -$mova] ]
          set y1 [y_rot_pix [y_pix $ya] $angle 0 [expr -$mova] ]
          set x2 [x_rot_pix [x_pix $xb] $angle 0 [expr -$movb] ]
          set y2 [y_rot_pix [y_pix $yb] $angle 0 [expr -$movb] ]
            $c create line $x1 $y1 $x2  $y2 -width 1 -fill red

          set x1 [x_rot_pix [x_pix $xa] $angle 0 [expr -1.5*$mova] ]
          set y1 [y_rot_pix [y_pix $ya] $angle 0 [expr -1.5*$mova] ]
          set x2 [x_rot_pix [x_pix $xb] $angle 0 [expr -1.5*$movb] ]
          set y2 [y_rot_pix [y_pix $yb] $angle 0 [expr -1.5*$movb] ]
          $c create text $x1 $y1 -text [abs $val1] -fill red -anchor sw
          $c create text $x2 $y2 -text [abs $val2] -fill red -anchor sw

          # arrows
          if {$val1 > 0} {
            # TODO code here
              set x1 [x_rot_pix [x_pix $xa] $angle 0 [expr -$mova/2] ]
              set y1 [y_rot_pix [y_pix $ya] $angle 0 [expr -$mova/2] ]
              set x2 [x_rot_pix [x_pix $xa] $angle $mova 0  ]
              set y2 [y_rot_pix [y_pix $ya] $angle $mova 0 ]
              $c create line $x1 $y1 $x2  $y2 -width 1 -fill red
          } else {
            if {$val1 < 0} {
              # TODO code here
              set x1 [x_rot_pix [x_pix $xa] $angle 0 0 ]
              set y1 [y_rot_pix [y_pix $ya] $angle 0 0 ]
              set x2 [x_rot_pix [x_pix $xa] $angle [expr -$mova] [expr -$mova/2] ]
              set y2 [y_rot_pix [y_pix $ya] $angle [expr -$mova] [expr -$mova/2] ]
              $c create line $x1 $y1 $x2  $y2 -width 1 -fill red

              set x1 [x_rot_pix [x_pix $xa] $angle [expr -$mova] 0 ]
              set y1 [y_rot_pix [y_pix $ya] $angle [expr -$mova] 0 ]
              set x2 [x_rot_pix [x_pix $xa] $angle [expr -$mova] [expr -$mova/2] ]
              set y2 [y_rot_pix [y_pix $ya] $angle [expr -$mova] [expr -$mova/2] ]
              $c create line $x1 $y1 $x2  $y2 -width 1 -fill red
            }
          }
          
          if {$val2 > 0} {
            # TODO code here
              set x1 [x_rot_pix [x_pix $xb] $angle 0 0 ]
              set y1 [y_rot_pix [y_pix $yb] $angle 0 0 ]
              set x2 [x_rot_pix [x_pix $xb] $angle [expr -$movb] [expr -$movb/2] ]
              set y2 [y_rot_pix [y_pix $yb] $angle [expr -$movb] [expr -$movb/2] ]
              $c create line $x2 $y2 $x1  $y1 -width 1 -fill red

              set x1 [x_rot_pix [x_pix $xb] $angle [expr -$movb] 0 ]
              set y1 [y_rot_pix [y_pix $yb] $angle [expr -$movb] 0 ]
              set x2 [x_rot_pix [x_pix $xb] $angle [expr -$movb] [expr -$movb/2] ]
              set y2 [y_rot_pix [y_pix $yb] $angle [expr -$movb] [expr -$movb/2] ]
              $c create line $x2 $y2 $x1  $y1 -width 1 -fill red
          } else {
            if {$val2 < 0} {
              # TODO code here
              set x1 [x_rot_pix [x_pix $xb] $angle 0 [expr -$movb/2] ]
              set y1 [y_rot_pix [y_pix $yb] $angle 0 [expr -$movb/2] ]
              set x2 [x_rot_pix [x_pix $xb] $angle [expr $movb] 0 ]
              set y2 [y_rot_pix [y_pix $yb] $angle [expr $movb] 0 ]
              $c create line $x1 $y1 $x2  $y2 -width 1 -fill red
            }
          }

        }
        1 { ;# fy
          if {$val1 < 0} {
            set mova [expr -1*$mova]
          } else {
            if {$val1 == 0} {
              set mova 0
            }
          }
          if {$val2 < 0} {
            set movb [expr -1*$movb]
          } else {
            if {$val1 == 0} {
              set mova 0
            }
          }

          set x1 [x_rot_pix [x_pix $xa] $angle 0 [expr 1.5*$mova] ]
          set y1 [y_rot_pix [y_pix $ya] $angle 0 [expr 1.5*$mova] ]
          set x2 [x_rot_pix [x_pix $xb] $angle 0 [expr 1.5*$movb] ]
          set y2 [y_rot_pix [y_pix $yb] $angle 0 [expr 1.5*$movb] ]
          $c create text $x1 $y1 -text [abs $val1] -fill red -anchor sw
          $c create text $x2 $y2 -text [abs $val2] -fill red -anchor sw

          set x1 [x_pix $xa]
          set y1 [y_pix $ya]
          set x2 [x_rot_pix [x_pix $xa] $angle [expr abs($mova)/2] [expr 3*$mova/4] ]
          set y2 [y_rot_pix [y_pix $ya] $angle [expr abs($mova)/2] [expr 3*$mova/4] ]
            $c create line $x1 $y1 $x2  $y2 -width 1 -fill red

          set x1 [x_pix $xb]
          set y1 [y_pix $yb]
          set x2 [x_rot_pix [x_pix $xb] $angle [expr -abs($mova)/2] [expr 3*$movb/4] ]
          set y2 [y_rot_pix [y_pix $yb] $angle [expr -abs($mova)/2] [expr 3*$movb/4] ]
            $c create line $x1 $y1 $x2  $y2 -width 1 -fill red

          set x1 [x_pix $xa]
          set y1 [y_pix $ya]
          set x2 [x_rot_pix [x_pix $xa] $angle 0 [expr $mova] ]
          set y2 [y_rot_pix [y_pix $ya] $angle 0 [expr $mova] ]
            $c create line $x1 $y1 $x2  $y2 -width 1 -fill red
          set x1 [x_pix $xb]
          set y1 [y_pix $yb]
          set x2 [x_rot_pix [x_pix $xb] $angle 0 [expr $movb] ]
          set y2 [y_rot_pix [y_pix $yb] $angle 0 [expr $movb] ]
            $c create line $x1 $y1 $x2  $y2 -width 1 -fill red
          set x1 [x_rot_pix [x_pix $xa] $angle 0 [expr $mova] ]
          set y1 [y_rot_pix [y_pix $ya] $angle 0 [expr $mova] ]
          set x2 [x_rot_pix [x_pix $xb] $angle 0 [expr $movb] ]
          set y2 [y_rot_pix [y_pix $yb] $angle 0 [expr $movb] ]
            $c create line $x1 $y1 $x2  $y2 -width 1 -fill red
        }
      }
    }
    2 { ;# temperature
    }
  }
}

proc plot_elem_result { elem type max } {
}

proc plot_results {} {
}


# Plots all data (except results):
proc plot_stuff {} {
  global nlen n_id n_x n_y
  global elen e_id e_na e_nb e_h e_E
  global dlen d_id d_n d_dir d_val 
  global flen f_id f_n f_dir f_val
  global def_lc

  # clean canvas:
  .c delete all

  # elements:
  for {set i 0} {$i < $elen} {incr i} {
    set na  [lsearch $n_id [lindex $e_na $i]]
    set nb  [lsearch $n_id [lindex $e_nb $i]]
    plot_element .c [lindex $e_id $i] [x_pix [lindex $n_x $na]] [y_pix [lindex $n_y $na]] [x_pix [lindex $n_x $nb]] [y_pix [lindex $n_y $nb]] 
  }

 # nodes:
  for {set i 0} {$i < $nlen} {incr i} {
    #puts  "Node:"; puts [lindex $n_id $i]; puts [x_pix [lindex $n_x $i]] ; puts [y_pix [lindex $n_y $i]]
    plot_node .c [lindex $n_id $i] [x_pix [lindex $n_x $i]] [y_pix [lindex $n_y $i]] 
  }

  # displacements:
  for {set i 0} {$i < $dlen} {incr i} {
      set na  [lsearch $n_id [lindex $d_n $i]]
      plot_disp .c [lindex $d_id $i] [x_pix [lindex $n_x $na]] [y_pix [lindex $n_y $na]] [lindex $d_dir $i] [lindex $d_val $i] orange
      set na  [lsearch $n_id [lindex $d_n $i]]
      plot_disp .c [lindex $d_id $i] [x_pix [lindex $n_x $na]] [y_pix [lindex $n_y $na]] [lindex $d_dir $i] [lindex $d_val $i] cyan
  }

  # forces:
  for {set i 0} {$i < $flen} {incr i} {
      set na  [lsearch $n_id [lindex $f_n $i]]
      plot_force .c [lindex $f_id $i] [x_pix [lindex $n_x $na]] [y_pix [lindex $n_y $na]] [lindex $f_dir $i] [lindex $f_val $i]
  }

  #.c create text 10 25 -text [format "LC %i" $def_lc] -fill darkgray -anchor sw
}

# Save output to EPS file:
proc save_canvas_eps {canvas file} {
  if [ catch {$canvas postscript -file $file } result] { puts $result ; return -1 }
}

# Mouse tracking - testing function
proc mouse_action {c x y} {
  global workmode pickmode enode_id
  global def_id def_E def_A def_I def_h def_ro def_alpha def_ha def_hb
	global nlen n_id n_x n_y
  global elen e_id e_na e_nb e_ha e_hb e_A e_I e_h e_E e_ro e_alpha
  global dlen d_id d_n d_dir d_val d_lc
  global flen f_id f_n f_dir f_val f_lc
  global def_is1 def_is2 def_is3
  global def_dir1 def_dir2 def_dir3
  global def_val1 def_val2 def_val3
  global def_lc 
	global m_x1 m_y1 m_button m_line

  set gx [what_grid_point_x $x]
  set gy [what_grid_point_y $y] 

  if {$pickmode != 2} {
		set enode_id 0
		set m_button 0
	}

  # nodes:
  if {$pickmode == 1} {
    if {$workmode == 1} {
      if {[if_node_exists  $gx $gy] <= 0} { 
        add_node $gx $gy
      } else {
        tk_messageBox -type ok -icon error -title "Error" -message "There already is a node!"
        return -1
      }
    }
    if {$workmode == 2} {
      node_dialog [what_node_id  [x_real $x] [y_real $y]]
    }
    if {$workmode == 3} {
      delete_node [what_node_id  [x_real $x] [y_real $y]]
    }
  }
  
  # elements:
  if {$pickmode == 2} { 
    if {$workmode != 1} {
			set enode_id 0
			set m_button 0
		}

    if {$workmode == 1} {
      # create:
      if {$enode_id > 0} {
        # second node
				set m_button 0

        if {[set node_b [what_node_id  [x_real $x] [y_real $y]]] <= 0} {
            tk_messageBox -type ok -icon error -title "Error" -message "Invalid nodes!"
            set enode_id -1
            set node_b -1
          return -1
        } else {
          if {[check_elem_exists $enode_id $node_b] > 0} {
            tk_messageBox -type ok -icon error -title "Error" -message "Element already exists!"
            set enode_id -1
            set node_b -1
            return -1
          }
          if {$enode_id == $node_b} {
            tk_messageBox -type ok -icon error -title "Error" -message "Nodes must be different!"
            set enode_id -1
          }
          set def_na $enode_id
          set def_nb $node_b
          set enode_id -1
          set node_b -1
          add_elements $def_na $def_nb $def_ha $def_hb $def_A $def_I $def_h $def_E $def_ro $def_alpha
        }
      } else {
        # first node
        set enode_id [what_node_id  [x_real $x] [y_real $y]] 
				if {$enode_id > 0} {
					set m_x1 [x_pix [lindex $n_x [lsearch $n_id $enode_id]]]
					set m_y1 [y_pix [lindex $n_y [lsearch $n_id $enode_id]]]
					set m_button 1
				}
      }
    }

    if {$workmode == 2} {
      set pos [what_elem_id [x_real $x] [y_real $y]]
      puts $pos
      if {$pos < 0} { return -1 }

      set pos [lsearch $e_id $pos]
      set def_id [lindex $e_id $pos]
      set def_na [lindex $e_na $pos]
      set def_nb [lindex $e_nb $pos]
      set def_ha [lindex $e_ha $pos]
      set def_hb [lindex $e_hb $pos]
      set def_E [lindex $e_E $pos]
      set def_A [lindex $e_A $pos]
      set def_I [lindex $e_I $pos]
      set def_h [lindex $e_h $pos]
      set def_ro [lindex $e_ro $pos]
      set def_alpha [lindex $e_alpha $pos]

      elem_dialog $def_id 
    }
    if {$workmode == 3} {
      delete_elem [what_elem_id [x_real $x] [y_real $y]]
    }
  }

  # displacements/supports
  if {$pickmode == 3} { 

    if {$workmode == 1} {
      set dnode [what_node_id  [x_real $x] [y_real $y]]
      if {$def_is1 > 0} { add_disp $dnode $def_dir1 $def_val1 $def_lc}
      if {$def_is2 > 0} { add_disp $dnode $def_dir2 $def_val2 $def_lc}
      if {$def_is3 > 0} { add_disp $dnode $def_dir3 $def_val3 $def_lc}
    }

    if {$workmode == 2} {
      set dnode [what_node_id  [x_real $x] [y_real $y]]
      if {$dnode > 0} {
        disp_dialog $dnode $def_lc
      }
    }
    
    if {$workmode == 3} {
      set dnode [what_node_id  [x_real $x] [y_real $y]]
      for {set i 0} {$i < $dlen} {incr i} {
        if {$dnode == [lindex $d_n $i]} {
          delete_disp [lindex $d_id $i]
          incr i -1
        }
      }
    }
  }

  # forces
  if {$pickmode == 4} { 
    
    if {$workmode == 1} {
      set dnode [what_node_id  [x_real $x] [y_real $y]]
      add_force $dnode 0 $def_val1 $def_lc
      add_force $dnode 1 $def_val2 $def_lc
      add_force $dnode 2 $def_val3 $def_lc
    }
    
    if {$workmode == 2} {
      set dnode [what_node_id  [x_real $x] [y_real $y]]
      if {$dnode > 0} {
        force_dialog $dnode $def_lc
      }
    }

    if {$workmode == 3} {
      set fnode [what_node_id  [x_real $x] [y_real $y]]
      for {set i 0} {$i < $flen} {incr i} {
        if {$fnode == [lindex $f_n $i]} {
          delete_force [lindex $f_id $i]
          incr i -1
        }
      }
    }
  }

  # element loads
  if {$pickmode == 5} { 
    if {$workmode == 1} {
      set def_le [what_elem_id  [x_real $x] [y_real $y]]
      check_set_eload $def_le $def_lc
    }
  
		if {$workmode == 2} {
      set def_le [what_elem_id [x_real $x] [y_real $y]]
      if {$def_le > 0} {
        select_load_dialog $def_le $def_lc
      }
    }
  }

  plot_stuff
  plot_grid
  make_frame

  #plot_point c [x_pix $gx] [y_pix $gy]
}

proc refresh_canvas {} {
  global c_w c_h
  global max_x max_y

  update idletasks
  set c_w  [winfo width .c]
  set c_h  [winfo height .c]

  if {$max_x > 0} {
    if {$max_y > 0} {
      compute_scale $max_x $max_y ;
    } else {
      compute_scale 6.0 [expr (6.0*$c_h)/$c_w]
    }
  } else {
    compute_scale 6.0 [expr (6.0*$c_h)/$c_w]
  }

  plot_stuff
  plot_grid
  make_frame
}

proc motion_e {c x y} {
  global m_x1 m_y1 m_button m_line
 	global workmode pickmode enode_id

	if {$workmode == 1} {
		if {$pickmode == 2} {
   		if ($m_button) {
       	$c delete $m_line
       	set m_line [$c create line $m_x1 $m_y1 $x $y -fill red]
   		}
		}
	}
}

bind .c <ButtonRelease-1> "mouse_action .c %x %y"
bind .c <Configure> "refresh_canvas"
bind .c <Motion>          { motion_e .c %x %y }

# Menu reactions  ################################################$
proc helpAbout {} {
  tk_messageBox -type ok -title "Ahout" -message "Structural analysis tool.\n(C) Jiri Brozovsky, 2005"
}

proc fileNew {} {
  global data_file

  set data_file ""
  wm title . [format "AxiShell: %s" "(new file)"]
  
  unsetData ; # data are unset here!
  make_empty_data
}

proc fileOpen {} {
  global data_file

  set types {     
    {{Data Files}       {.df}         }
    {{Text Files}       {.txt}        }
    {{All Files}        *             }
  }
  
  set filename [tk_getOpenFile -filetypes $types -defaultextension ".df"]
  
  if {$filename != ""} {     
    if {[file exists $filename] != 1} {
      tk_messageBox -type ok -title "Opening failed!" -icon  "error" -message "Please select existing file."
      return -1
    }
    set data_file $filename

    unsetData ; # data are unset here!
    make_empty_data

    if {[read_data $data_file] != 0} {
      tk_messageBox -type ok -title "Opening failed!" -icon  "error" -message "Reading of file failed.\nPlease try to use different name."
			return -1
    } else { 
      plot_stuff
    }
  } else {
      tk_messageBox -type ok -title "Opening failed!" -icon  "error" -message "No file was opened!"
			return -1
  }

  wm title . [format "AxiShell: %s" $filename]
  fit_replt_canvas
}

proc fileSave {} {
  global data_file

  if {$data_file == ""} {
      tk_messageBox -type ok -title "Save failed!" -icon  "error" -message "No name of file selected!"
			return -1
  }
  if {[write_data $data_file] != 0} {
      tk_messageBox -type ok -title "Save failed!" -icon  "error" -message "Can not write file!"
			return -1
    } 
}

proc fileSaveAs {} {
  global data_file

  set types {     
    {{Data Files}       {.df}         }
    {{Text Files}       {.txt}        }
    {{All Files}        *             }
  }
  
  set filename [tk_getSaveFile -initialfile $data_file -filetypes $types -defaultextension ".df"]
  
  if {$filename != ""} {
    set data_file $filename
    wm title . [format "AxiShell: %s" $filename]

    if {[write_data $data_file] != 0} {
      tk_messageBox -type ok -title "Save failed!" -icon  "error" -message "Can not write file. \nPlease try to select different name or location of file."
			return -1
    } else { 
      plot_stuff
    }
  } else {
      tk_messageBox -type ok -title "Save failed!" -icon  "error" -message "No file name was given!"
			return -1
  }
}

proc new_element {} {
  global statusbar workmode pickmode

  if {[elem_dialog 0] <= 0} {
    set statusbar "Pick two nodes to make an element.."
    set workmode 1 
    set pickmode 2
  } else {
    set statusbar "Failed.."
  }
}


proc new_support {} {
  global statusbar workmode pickmode
  global def_lc

  if {[disp_dialog 0 $def_lc] <= 0} {
    set statusbar "Pick nodes to make a support.."
    set workmode 1 
    set pickmode 3
  } else {
    set statusbar "Failed.."
  }
}

proc new_force {} {
  global statusbar workmode pickmode
  global def_lc

  if {[force_dialog 0 $def_lc] <= 0} {
    set statusbar "Pick nodes to make a force.."
    set workmode 1 
    set pickmode 4
  } else {
    set statusbar "Failed.."
  }
}

proc new_elem_force_x {} {
  global statusbar workmode pickmode
  global def_lc

  if {[load_dialog 0 0 0 $def_lc] <= 0} {
    set statusbar "Pick element to make a force.."
    set workmode 1 
    set pickmode 5
  } else {
    set statusbar "Failed.."
  }
}

proc new_elem_force_y {} {
  global statusbar workmode pickmode
  global def_lc

  if {[load_dialog 0 0 1 $def_lc] <= 0} {
    set statusbar "Pick element to make a force.."
    set workmode 1 
    set pickmode 5
  } else {
    set statusbar "Failed.."
  }
}


proc new_elem_moment_z {} {
  global statusbar workmode pickmode
  global def_lc

  if {[load_dialog 0 0 2 $def_lc] <= 0} {
    set statusbar "Pick element to make a moment.."
    set workmode 1 
    set pickmode 5
  } else {
    set statusbar "Failed.."
  }
}


# Menu system  ####################################################
menu .mbar -tearoff 0
    . configure -menu .mbar

  .mbar add cascade -label "File" -menu .mbar.file
    menu .mbar.file -title "File" -tearoff 1 
    .mbar.file add command -label "New"  -command "init_dialog"
    .mbar.file add separator
    .mbar.file add command -label "Open..."  -command "fileOpen"
    .mbar.file add command -label "Save"  -command "fileSave"
    .mbar.file add command -label "Save As..." -command "fileSaveAs"
    .mbar.file add separator
    .mbar.file add command -label "Quit" -command exit 

 .mbar add cascade -label "Create" -menu .mbar.create
    menu .mbar.create -title "Create" -tearoff 1
    .mbar.create add command -label "Node"  -command {set statusbar "Pick grid point to make node.." ; set workmode 1 ; set pickmode 1}
    .mbar.create add command -label "Element"  -command { set statusbar "-" ; set workmode 1 ; set pickmode 2 ; new_element}
    .mbar.create add separator
    .mbar.create add command -label "Support"  -command { set statusbar "-" ; set workmode 1 ; set pickmode 3 ; new_support}
    .mbar.create add separator
    .mbar.create add command -label "Nodal force"  -command { set statusbar "-" ; set workmode 1 ; set pickmode 4 ; new_force}

 .mbar add cascade -label "Edit" -menu .mbar.edit
    menu .mbar.edit -title "Edit" -tearoff 1
    .mbar.edit add command -label "Node"  -command {set statusbar "Pick node to edit.." ; set workmode 2 ; set pickmode 1}
    .mbar.edit add command -label "Element"  -command {set statusbar "Pick element to edit.." ; set workmode 2 ; set pickmode 2}
    .mbar.edit add command -label "Support"  -command {set statusbar "Pick node to edit support.." ; set workmode 2 ; set pickmode 3}
    .mbar.edit add command -label "Nodal force"  -command {set statusbar "Pick node to edit force.." ; set workmode 2 ; set pickmode 4}

 .mbar add cascade -label "Delete" -menu .mbar.delete
    menu .mbar.delete -title "Delete" -tearoff 1
    .mbar.delete add command -label "Node"  -command {set statusbar "Pick node to delete.." ; set workmode 3 ; set pickmode 1}
    .mbar.delete add command -label "Element"  -command {set statusbar "Pick element to delete.." ; set workmode 3 ; set pickmode 2}
    .mbar.delete add command -label "Support"  -command {set statusbar "Pick node to delete support(s).." ; set workmode 3 ; set pickmode 3}
    .mbar.delete add command -label "Nodal force"  -command {set statusbar "Pick node to delete forces(s).." ; set workmode 3 ; set pickmode 4}

 .mbar add cascade -label "View" -menu .mbar.view
    menu .mbar.view -title "View" -tearoff 1
    .mbar.view add command -label "Fit view"  -command "fit_replt_canvas"
    .mbar.view add separator
    .mbar.view add command -label "Zoom"  -command "zoom_canvas 0.1"
    .mbar.view add command -label "Unzoom"  -command "zoom_canvas -0.1"
    .mbar.view add separator
    .mbar.view add command -label "Set load case"  -command "select_load_case_dialog"
    #.mbar.view add separator
    #.mbar.view add command -label "Save view"  -command "save_canvas_eps .c 1.ps"

if {0} {
 .mbar add cascade -label "Show" -menu .mbar.show
    menu .mbar.show -title "Show" -tearoff 1
    .mbar.show add command -label "New"  -command "NewDoc"

 .mbar add cascade -label "Results" -menu .mbar.results
    menu .mbar.results -title "Results" -tearoff 1
    .mbar.results add command -label "Solve"  -command "NewDoc"
}

 .mbar add cascade -label "Help" -menu .mbar.help
    menu .mbar.help -title "Help" -tearoff 0
    .mbar.help add command -label "About"  -command "helpAbout"

# Statusbar ####################################################
label .statusline -relief sunken -anchor w -textvariable statusbar  
pack .statusline -anchor sw -fill both


# Some inits: ##################################################
make_empty_data ;# Important: makes data
set data_file ""
wm title . "AxiShell"
make_frame

# Command line arguments #######################################
if {$argc > 0} {
  if {[file exists [lindex $argv 0]] == 1} {
    set data_file [lindex $argv 0]
    if {[read_data $data_file] != 0} {
      make_empty_data
    } else {
      wm title . [format "AxiShell: %s" $data_file]
      fit_canvas
      plot_stuff
      plot_grid
      make_frame
    }
  }
}

# Some mess: ###################################################
#tk_bisque
#tk_setPalette background darkslategray foreground wheat activeBackground wheat activeForeground black  selectBackground wheat selectForeground darkslategray
#tk_setPalette background lightsteelblue foreground black

#package require Img
#set im [image create photo -format window -data .c]
#$im write mycanvas.gif -format GIF
