#! /usr/bin/env tclsh
#
# Simple wrapper for Monte and shell
#
# (C) Jiri Brozovsky, VSB-TU of Ostrava 2010
#
#  This program is free software; you can redistribute it and/or
#  modify it under the terms of the GNU General Public License as
#  published by the Free Software Foundation; either version 2 of the
#  License.
#
#  This program is distributed in the hope that it will be useful, but
#  WITHOUT ANY WARRANTY; without even the implied warranty of
#  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
#  General Public License for more details.
#
#  You should have received a copy of the GNU General Public License
#  in a file called COPYING along with this program; if not, write to
#  the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
#  02139, USA.
#
#
# Description:
# ============
# This program will read files with optimisation data
# and wil run Monte and eshell on these data
# It should select best set of optimised paramaters
#

# Initial prints:
puts " "
puts "OPTIMIST: software for structural optimalization by Monte Carlo"
puts "  Version 1.0.0 "
puts "  (C) 2010 VSB-TU Ostrava, FAST"
puts "  (C) 2006, 2010 Jiri Brozovsky (simulation engine, libraries)"
puts " This is free software available under GNU GPL 2.0"
puts " "
puts "INIT:"
puts " "

#### TK initialization (if impossible then it runs in batch mode)
if [catch {package require Tk} ] {
  puts " OK, we DON'T have Tk: batch mode."
  set Have_Tk 0
} else {
  #puts " Great, we will use an ice GUI!"
  set Have_Tk 1
}

set Have_Tk 0 ; # No gui for authorised software!
               # They haven't paid us enough...

#### Test platform/OS
set my_platform $tcl_platform(platform)
puts " Computing environment: $my_platform ($tcl_platform(os), $tcl_platform(machine))."

#### PATHS AND EXECUTABLES (edit them to fit your needs...)
set bin_dir "" ;# directory where histograms are stored 
set his_dir "" ;# directory where histograms are stored 

# read environment variables
if [ catch { set myenv $env(OPTIM_SHELL_DIR) } ] {
  set bin_dir ""
} else {
  set bin_dir myenv
}
if [ catch { set myenv $env(OPTIM_SHELL_HIS_DIR) } ] {
  if [ catch { set myenv $env(MONTE) } ] {
    set his_dir myenv
  } else {
    set his_dir ""
  }
} else {
  set his_dir myenv
}


# set system dependent paths:
if { $my_platform == "unix" } {
  # linux, mac os x etc:
  if { $bin_dir != ""} {
    if { [ file exists "$bin_dir/eshell" ] == 0 } { set bin_dir "" }
  }
  if { $bin_dir == ""} {
    set homedir $env(HOME)
    if { [file exists "$homedir/progs/shell/eshell"] } {
      set bin_dir "$homedir/progs/shell"
    } else {
      if { [file exists "$homedir/optshell/eshell"] } {
      set bin_dir "$homedir/optshell"
      } else {
        if { [file exists "/opt/optshell/eshell"] } {
          set bin_dir "$homedir/opt/optshell"
        } else {
          if { [file exists "/usr/bin/eshell"] } {
            set bin_dir "/usr/bin"
          } else {
            set bin_dir "."
          }
        }
      }
    }
  }
  set monte_exe "$bin_dir/monte"
  set shell_price "$bin_dir/eshell"
  set lib_shell "$bin_dir/libshell.so"
  # his_dir:
  if { $his_dir != ""} {
    if { [ file exists "$bin_dir/histograms" ] == 1 } {
      set his_dir "$bin_dir/his_dir"
    } else {
      if { [ file exists "$bin_dir/normal.dis" ] == 1 } {
        set his_dir $bin_dir
      } else {
        set his_dir "."
      }
    }
  }
} else {
  # assumed to be "windows":
  if { $bin_dir != ""} {
    if { [file exists "$bin_dir\\eshell.exe"] == 0 } { set bin_dir "" }
  }
  if { $bin_dir == ""} {
    set homedir $env(HOME)
    if { [ file exists "$homedir\\optshell\\eshell.exe"] } {
      set bin_dir "$homediropt\\optshell"
    } else {
      if { [file exists "C:\\Program Files\\optshell\\eshell.exe"] } {
      set bin_dir "C:\\Program Files\\optshell"
      } else {
        if { [file exists "C:\\optshell\\eshell.exe"] } {
          set bin_dir "C:\\optshell"
        } else {
          if { [file exists "C:\\eshell.exe"] } {
            set bin_dir "C:\\"
          } else {
            set bin_dir "."
          }
        }
      }
    }
  }
  set monte_exe "$bin_dir\\monte.exe" ;# Monte binary
  set shell_price "$bin_dir\\eshell.exe" ;# eshell binary
  set lib_shell "$bin_dir\\libshell.dll"
  # his_dir:
  if { $his_dir != ""} {
    if { [ file exists "$bin_dir\\histograms" ] == 1 } {
      set his_dir "$bin_dir\\his_dir"
    } else {
      if { [ file exist "$bin_dir\\normal.dis" ] == 1 } {
        set his_dir $bin_dir
      } else {
        set his_dir "."
      }
    }
  }
}
puts " Binary files: $bin_dir"
puts " Histograms:   $his_dir"
#### old code commented out:
#set bin_dir "." ;# directory where histograms are stored 
#set his_dir "." ;# directory where histograms are stored 
#set monte_exe "$bin_dir/monte" ;# Monte binary
#set shell_price "$bin_dir/eshell" ;# eshell binary
#set lib_shell "$bin_dir/libshell.so"

#################################################################

#### FILES:

set i_file_name  "default.opt" ;#"default.opt"
set i_monte_name "" ;#"default.mon"
set i_shell_name "" ;#"default.txt"
set i_shell_it_name "" ;#"default.tmp"
set o_stat_name  "" ;#"default.rop"


#### SIMULATION 

set sim_number  10000
set sim_verbose 1000
set sim_test    100
set sim_wall    15
set sim_savesim 0
set sim_prob    0.0001


#### DATA DEFINITIONS

# Optimised parameters set data:
set sets_len 0
set set_num [list]
set set_type0 [list]
set set_type1 [list]
set set_type2 [list]
set set_data [list]

# Optimalisation functions
set func_len 0
set func_type [list]
set func_def [list]

# Data sets 
set data_len 0
set data_data [list]
set data_price [list]
set data_sort [list]

# Results:
set final_list [list]
set prob_fail 0.0

# External data storage
set var_shell_data "" ; # shell input file content

#### DATA MANIPULATING FUNCTIONS:

# reads input data
proc read_input_file {} {
  global Have_Tk
  global i_file_name
  global sets_len set_num set_data 
  global set_type0 set_type1 set_type2
  global func_len func_type func_def 
  global sim_prob sim_number

  # check file name:
  if {$i_file_name == ""} {
    if {$Have_Tk == 1} {
      tk_messageBox -icon error -type ok -message "Need to set name of input file!"
    } else { puts  "Need to set name of input file!" }
    return -1
  }

  # open file:
  if [ catch {set fr [open $i_file_name "r"]} result] {
    if {$Have_Tk == 1} {
      tk_messageBox -icon error -type ok -message "Can not open file!"
    } else { puts "Can not open file!" }
    return -1 
  }

  # clear data fields
  unset sets_len
  unset set_num
  unset set_type0
  unset set_type1
  unset set_type2
  unset set_data

  # make empty data fields
  set sets_len 0
  set set_num [list]
  set set_type0 [list]
  set set_type1 [list]
  set set_type2 [list]
  set set_data [list]
  set line [list]

  # read data sets here
  gets $fr sets_len
  if {$sets_len > 0} {
    for {set i 0} {$i < $sets_len} {incr i} {
      # parameter type and size:
      gets $fr line
      set type0 [lindex $line 0]
      set type1 [lindex $line 1]
      set type2 [lindex $line 2]
      set num   [lindex $line 3]

      lappend set_num $num
      lappend set_type0 $type0
      lappend set_type1 $type1
      lappend set_type2 $type2

      # parameter values:
      if {$num > 0} {
        gets $fr line
        lappend set_data $line
      } else {
        lappend set_data "0"
      }
    }
  }

  # functions: 
  gets $fr func_len
  if {$func_len > 0} {
    for {set i 0} {$i < $func_len} {incr i} {
      gets $fr line
      if {[lindex $line 0] == "FUNCTION"} {
        set type [lindex $line 1]
        lappend func_type $type
        if {$type == "0"} {
          set line2 0
          gets $fr line2
          lappend func_def $line2
        } else {
          lappend func_def "0"
        }
      }
    }
  }

  # simulation parameters:
  gets $fr line
  set sim_number [lindex $line 0]
  set sim_prob [lindex $line 1]
  if { $sim_number < 1 } { set sim_number 1 }
  if {$sim_prob <= 0 } { set sim_prob 1e-9 }

  # close file:
  if [ catch [close $fr] result] { 
    if {$Have_Tk == 1} {
      tk_messageBox -icon error -type ok -message "Can not close file!"
    } else { puts "Can not close file!" }
    return -1 
  }

  return 0;
}

#### prints input data to output
proc print_input_file {} {
  global i_file_name
  global sets_len set_num set_data 
  global set_type0 set_type1 set_type2
  global func_len func_type func_def 
  global sim_number sim_prob

  puts " "
  puts "INPUT DATA:"
  puts " "
  puts " Number of data sets: $sets_len"
  puts " "
  if {$sets_len > 0} {
    for {set i 0} {$i < $sets_len} {incr i} {
      puts " Type $i: [lindex $set_type0  $i ] [lindex $set_type1  $i ] [lindex $set_type2  $i ], num: [lindex $set_num $i ] " 
      puts " Data $i: [lindex $set_data $i ] "
    }
  } else {
    puts "No data to print"
  }

  if {$func_len > 0} {
    puts " "
    for {set i 0} {$i < $func_len} {incr i} {
      puts " Function $i, type: [lindex $func_type $i] "
      if {[lindex $func_type $i] == 0} {
        puts "  Definition $i: [lindex $func_def $i]"
      }
    }
  }

  puts " "
  puts " Simulations: $sim_number"
  puts " Failure prob.: $sim_prob"

  return 0;
}


# write input data
proc write_input_file {} {
  global Have_Tk
  global i_file_name
  global sets_len set_num set_data 
  global set_type0 set_type1 set_type2
  global func_len func_type func_def 
  global sim_number sim_prob

  # check file name:
  if {$i_file_name == ""} {
    if {$Have_Tk == 1} {
      tk_messageBox -icon error -type ok -message "Need to set name of input file!"
    } else { puts  "Need to set name of input file!" }
    return -1
  }

  # open file:
  if [ catch {set fr [open $i_file_name "w"]} result] {
    if {$Have_Tk == 1} {
      tk_messageBox -icon error -type ok -message "Can not open file!"
    } else { puts "Can not open file!" }
    return -1 
  }

  # data sets 
  puts $fr $sets_len
  if {$sets_len > 0} {
    for {set i 0} {$i < $sets_len} {incr i} {
      # parameter type and size:
      puts $fr "[lindex $set_type0 $i] [lindex $set_type1 $i] [lindex $set_type2 $i] [lindex $set_num $i]"
      # parameter values:
      if {$set_num > 0} { puts $fr [lindex $set_data $i] }
    }
  }

  # functions: 
  puts $fr $func_len
  if {$func_len > 0} {
    for {set i 0} {$i < $func_len} {incr i} {
      puts $fr "FUNCTION [lindex $func_type $i]"
      if {$func_type == "0"} {
        puts $fr [lindex $func_def $i]
      } 
    }
  }

  puts $fr "$sim_number $sim_prob"
  
  # close file:
  if [ catch [close $fr] result] { 
    if {$Have_Tk == 1} {
      tk_messageBox -icon error -type ok -message "Can not close file!"
    } else { puts "Can not close file!" }
    return -1 
  }

  return 0;
}

#### Read "shell" input file into variable
proc get_shell_input {} {
  global Have_Tk i_shell_name var_shell_data 

  # check file name:
  if {$i_shell_name == ""} {
    if {$Have_Tk == 1} {
      tk_messageBox -icon error -type ok -message "Incorect name of FEM file!"
    } else { puts  "Incorect name of FEM file!" }
    return -1
  }

  # open file:
  if [ catch {set frr [open $i_shell_name "r"]} result] {
    if {$Have_Tk == 1} {
      tk_messageBox -icon error -type ok -message "Can not open FEM file!"
    } else { puts "Can not open FEM file!" }
    return -1 
  }

  # read the whole file into variable:
  set var_shell_data [read $frr] ; #puts $var_shell_data

  # close file:
  if [ catch [close $frr] result] { 
    if {$Have_Tk == 1} {
      tk_messageBox -icon error -type ok -message "Can not close FEM file!"
    } else { puts "Can not close FEM file!" }
    return -1 
  }
}

#### Set default file names
proc set_file_names {} {
  global i_file_name i_monte_name i_shell_name i_shell_it_name o_stat_name

  if {[string length $i_file_name] > 0 } {
    set i_monte_name ""
    set i_shell_name ""
    set i_shell_it_name ""
    set o_stat_name  ""

    append i_monte_name  [file rootname $i_file_name] ".mon"
    append i_shell_name  [file rootname $i_file_name] ".txt"
    append i_shell_it_name  [file rootname $i_file_name] ".tmp"
    append o_stat_name  [file rootname $i_file_name] ".rop"
  } else {
    return -1
  }
}

#### Optimalization function (making of n-ties of input data)
proc make_data_sets {} {
  global sets_len set_num set_data 
  global set_type0 set_type1 set_type2
  global func_len func_type func_def 
  global data_len data_data

  set line [list]
  unset data_data
  set data_data [list]

  set line_mult [list]

  # generate indexes
  for {set i 0} {$i < $sets_len} {incr i} {
    set multpl 1
    for {set j [expr $i + 1]} {$j < $sets_len} {incr j} {
      set multpl [expr $multpl * [lindex $set_num $j]]
			#puts "$i $j $multpl [lindex $set_num $j]"
    }
    lappend line_mults $multpl
  }


  set data_len_all 1
  for {set i 0} {$i < $sets_len} {incr i} { 
    set data_len_all [expr $data_len_all * [lindex $set_num $i] ]
  }

  # actual creation of data sets:
  for {set i 0} {$i < $sets_len} {incr i} {
    unset line
    set line [list]
    set repp [expr $data_len_all / [lindex $line_mults $i]]
    for {set l 0} {$l < $repp} {incr l} {
      for {set j 0} {$j < [lindex $set_num $i] } {incr j} {
        for {set k 0} {$k < [lindex $line_mults $i]} {incr k} {
          lappend line [lindex $set_data $i $j]
        }
      }
    }
    lappend data_data $line
  }

  set data_len $data_len_all

	# use of "FUNCTION" data-chaging functions:
  for {set i 0} {$i < $data_len_all} {incr i} {
		for {set j 0} {$j < $func_len } {incr j} {
			if { [lindex $func_type $j] == "0"  } {
			  eval [lindex $func_def $j]
		  }
		}
	}

  # list of data sets:
  puts " "
  puts "PRICE DATA:"
  puts " "
  puts " Size of opt. data sets: $data_len"
  return 0 
}

#### Get optim. data into variable
proc get_opt_data_line {ii} {
  global sets_len set_num set_data 
  global set_type0 set_type1 set_type2
  global data_len data_data
  
  set g ""

  # get opt lines 
  append g "$sets_len \n"
  for {set i 0} {$i < $sets_len} {incr i} {
    append g "[lindex $set_type0 $i] "
    append g "[lindex $set_type1 $i] "
    append g "[lindex $set_type2 $i] \n"
  }

  # get ii-th optim. data set
  for {set i 0} {$i < $sets_len} {incr i} {
    append g "[lindex $data_data $i $ii ] "
  }
  append g "\n"

  return $g
}

#### Price function: compute prices of variants and their order
proc get_prices {} {
  global sets_len set_num set_data 
  global set_type0 set_type1 set_type2
  global func_len func_type func_def 
  global data_len data_data data_price data_sort
  global var_shell_data
  global shell_price i_file_name

  unset data_price
  set data_price [list]
  unset data_sort
  set data_sort [list]
  
  #set data_len 1
  for {set i 0} {$i < $data_len} {incr i} {
    set inp "" ; # variable with input data 
    append inp $var_shell_data
    append inp "\n"
    append inp [get_opt_data_line $i]
    #set d [exec tee /tmp/b <<$inp] ; # easiest execution
    #catch [set d [exec tee /tmp/b <<$inp ] ] ; # to mask use of stderr (!?)
    catch [set d [exec $shell_price -p <<$inp ] ] 
    #set d [exec $shell_price -p <<$inp ] 
    
    lappend data_price $d
    lappend data_sort -1
  }

  # get total maximum first:
  set max -1 
  for {set i 0} {$i < $data_len} {incr i} {
    set val [lindex $data_price $i]
    if {$max <= $val} { 
      set max $val 
      set pos $i 
    }
  }
  set min $max
  set count 0
  lset data_sort $pos 0
  set pr_maxprice $max

  # get data positions:
  set pos -1
  set min 0
  for {set i 1} {$i < $data_len} {incr i} {
    set max -1
    for {set j 0} {$j < $data_len} {incr j} {
      set val [lindex $data_price $j]
      if {$max <= $val} { 
        if {[lindex $data_sort $j] < 0} {
          set max $val 
          set pos $j 
        }
      }
    }
    lset data_sort $pos $i
    set min $max
  }
  set pr_minprice [lindex $data_price $pos] 
  puts " Min. price: $pr_minprice"
  puts " Max. price: $pr_maxprice"
}

#### Executes the monte solver 
proc run_monte_solver { i_fem_file} {
  global monte_exe his_dir lib_shell i_monte_name i_shell_it_name
  global sim_number sim_verbose sim_test sim_wall sim_prob
  set res_sim_line [list]
  set res_sim 0

  # run the Monte solver (quietly!):
  catch [ set res [exec $monte_exe -ld $lib_shell -d  $his_dir -lda $i_fem_file -i $i_monte_name -fpf $i_shell_it_name \
  -s $sim_number -wall $sim_wall -fon $sim_test ] ]

  # open file:
  if [ catch {set frs [open $i_shell_it_name "r"]} result] { return -1 }
    gets $frs  res_sim_line
  if [ catch [close $frs] result] { return -1 }

  set res_sim [lindex $res_sim_line 0]

  # data test:
  if {[lindex $res_sim_line 0] < $sim_prob} {
    puts " "
    puts " Probability of failure [lindex $res_sim_line 0] < $sim_prob."
    return 1 ; # best data found: we are done
  } else {
      puts -nonewline " ([lindex $res_sim_line 0] ) =>"
	}

  return 0 
}

#### Runs actual solution for given data sets
proc run_solution {} {
  global sets_len set_num set_data 
  global set_type0 set_type1 set_type2
  global func_len func_type func_def 
  global data_len data_data
  global data_len data_data data_price data_sort
  global var_shell_data
  global sim_number sim_verbose sim_test sim_wall 
  global i_file_name Have_Tk shell_price 

  puts "\n"
  puts "SOLUTION for $data_len alternatives: "
  puts " "
  for {set i 0} {$i < $data_len} {incr i} {
    for {set j 0} {$j < $data_len} {incr j} {
      # find data by price order:
      if { [lindex $data_sort $j] == [expr $data_len - $i - 1 ] } {
         # variable with input data:
        set inp ""
        append inp $var_shell_data
        append inp "\n"
        append inp [get_opt_data_line $i]

        # ## create input file for solver:
        # set file name:
        set i_sim_name ""
        append i_sim_name  [file rootname $i_file_name] "-$i.fem"
        # open file:
        if [ catch {set frs [open $i_sim_name "w"]} result] {
          if {$Have_Tk == 1} {
            tk_messageBox -icon error -type ok -message "Can not open file for simultaion model!"
          } else { puts "Can not open file for simulation model!" }
          return -1 
        }
        puts $frs $inp
        flush $frs
        # close file:
        if [ catch [close $frs] result] { 
          if {$Have_Tk == 1} {
            tk_messageBox -icon error -type ok -message "Can not close file!"
          } else { puts "Can not close file!" }
          return -1 
        }

        # "Monte" runs:
        puts -nonewline " Solving alternative [expr $i + 1] / $data_len:"
        set res 0
         set res [ run_monte_solver $i_sim_name]
        if { $res == 1} {
          #puts " OK." 
          puts " This is the best alternative!"
          puts " "
          puts " File: $i_sim_name"
          return 0 ;# structure not failed => OK, we are done!
        } else {
          puts " failed."
        }

        break
      }
    }
  }
}

#### Reads parameters from command line:
# 0 ... file.opt name
# 1 ... if -c then GUI is not started
proc read_cmd_line {} {
  global argc argv Have_Tk i_file_name

  if { $argc > 0} {
    set file_name [lindex $argv 0]
    puts " Parameters from command line:  $file_name"
    # file must end with ".opt":
    if { [regexp "\.opt" $file_name ] == "1" } {
      set i_file_name $file_name
      puts " Input data file name: $i_file_name"
    } else {
      puts " Invalid extension - using default file name!"
    }

    if { $argc > 1 } {
      if { [lindex $argv 1] == "-c" } {
        # run tui only
        if { $Have_Tk == "1" } { 
          set Have_Tk 0 
        }
      }
    }
  } else {
    puts " No command line parameters."
  }
}

#############################################################
#############################################################
#
# GUI will be here (TODO)
#
# GUI description:
#

### Dialog windows and support functions for GUI:
if { $Have_Tk == "1" } {

## About dialog:
proc helpAbout {} {
	tk_messageBox -type ok -title "Ahout" -message "OPTIMIST:\nStructural Optimization Tool\n(C) 2010 VSB-TU of Ostrava, FAST"
}

# openes existing *.opt file:
proc fileOpen {} {
  global i_file_name

  # opening:
  set name ""
  
  set types { 
    {{Monte input files}       {.opt}        TEXT} 
    {{Monte input files}       {.OPT}        TEXT} 
  }

  set name [tk_getOpenFile -defaultextension  ".mon" -filetypes $types]

  if {$name == ""} { return } 

  set i_file_name $name

  if {[read_input_file] != 0} {
    set i_file_name ""
    clean_all_data
    return -1
  } else {
		set_file_names
	wm title . "OPTIMIST: $i_file_name"
	}
}
	
} ; # END of support functions

#############################################################
#
proc program_run_solution {} {
  read_cmd_line    ;# reads parameters from command line
  set_file_names   ;# set names of files for I/O
  read_input_file  ;# reads "default.opt" file 
  print_input_file ;# prints input data report do stdout

  get_shell_input  ;# reads shell (fem) model
  make_data_sets   ;# creates testing data sets

  get_prices       ;# computes prices of structure
  run_solution     ;# runs Monte Carlo simulation hell
}
#
#############################################################
program_run_solution ; # non-interactive mode
puts " "
puts "OPTIMIST run is finished. Thanks for using!"
puts " "
exit 0
############################################################
