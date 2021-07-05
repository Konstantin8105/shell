/*
   File name: eshell.h
   Date:      2010/07/21 17:19
   Author:    Jiri Brozovsky

   Copyright (C) 2010 VSB-TU of Ostrava

   This program is free software; you can redistribute it and/or
   modify it under the terms of the GNU General Public License as
   published by the Free Software Foundation; either version 2 of the
   License, or (at your option) any later version.

   This program is distributed in the hope that it will be useful, but
   WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   General Public License for more details.

   You should have received a copy of the GNU General Public License
   in a file called COPYING along with this program; if not, write to
   the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
   02139, USA.

   Axisymetric shell solver.

   See for details:
   viz Schneider, Vykutil: Stavba chemickych zarizeni II.a
       Mikropocitacove aplikace MKP ve statice rotacnich
       skorepin, ES VUT Brno, Brno, Czechoslovakia, 1986

*/


#ifndef __ESHELL_H__
#define __ESHELL_H__

#include "fem.h"
#include "fem_mem.h"
#include "fem_math.h"

#include <string.h>

#ifdef USE_WIN32
#include <windows.h>
#define EXPORT __declspec(dllexport)
#endif

#endif

/* end of eshell.h */
