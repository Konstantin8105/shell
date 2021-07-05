/*
   File name: eshell.c
   Date:      2010/07/21 17:18
   Author:    Jiri Brozovsky

   Copyright (C) 2010 VSB-TU of Ostrava

   This program is free software; you can redistribute it and/or
   modify it under the terms of the GNU General Public License as
   published by the Free Software Foundation; either version 2 of the
   License.

   This program is distributed in the hope that it will be useful, but
   WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   General Public License for more details.

   You should have received a copy of the GNU General Public License
   in a file called COPYING along with this program; if not, write to
   the Free Software Foundation, Inc., 675 Mass Ave, Cambridge, MA
   02139, USA.


   Axisymetric shell solver. Use: eshell <input >output

   See for details:
   viz Schneider, Vykutil: Stavba chemickych zarizeni II.a
       Mikropocitacove aplikace MKP ve statice rotacnich
       skorepin, ES VUT Brno, Brno, Czechoslovakia, 1986
*/

#include "eshell.h"

#ifdef _USE_WITH_MONTE_
/* DATA STRUCTURES */

long monte_i_len = 0 ; /* input variables */
long monte_o_len = 0 ; /* output variables */

#endif

/* INPUT DATA: */

long n_m = 0 ; /* number of materials */
long n_n = 0 ; /* number of nodes */
long n_e = 0 ; /* number of elements */
long n_d = 0 ; /* number of displacements/supports */
long n_f = 0 ; /* number of loads */

long n_r_inp = 0 ; /* number of random input data */
long n_r_opt = 0 ; /* number of optim input data */

/* materials */
double *m_E1   = NULL ; /* E1 (bulk modullus)   */
double *m_E2   = NULL ; /* E2 (bulk modullus)   */
double *m_G    = NULL ; /* G (shear modullus)   */
double *m_nu1  = NULL ; /* nu1 (poisson ratio)  */
double *m_nu2  = NULL ; /* nu2 (poisson ratio)  */
double *m_q    = NULL ; /* volume gravity force */
double *m_vp   = NULL ; /* volume unit  price   */
double *m_t    = NULL ; /* width (if >=0 then ovewrites e_t[] data) */

/* nodes */
double *n_x = NULL ; /* x coordinates */
double *n_y = NULL ; /* y coordinates */

/*elements */
long   *e_n1  = NULL ; /* first nodes <0, n_n-1> */
long   *e_n2  = NULL ; /* second nodes  <0, n_n-1> */
long   *e_mat = NULL ; /* material numbers  <0, n_m-1> */
double *e_t   = NULL ; /* element widths (constatnt on element) */

/* displacements */
long   *d_n   = NULL ; /* nodes <0, n_n-1> */
long   *d_dir = NULL ; /* orientation w=0, u=1, pho=2, Ez=3, Ex=4, Erot=5 */
double *d_val = NULL ; /* size of displacement or stiffness */

/* forces in nodes */
long   *f_n   = NULL ; /* nodes <0, n_n-1> */
long   *f_dir = NULL ; /* orientation Fw=0, Fu=1, Mpho=2 */
double *f_val = NULL ; /* size of the force */

/* water load: */
double  w_top = 0.0 ; /* water level */
double  w_bot = 0.0 ; /* bottom of water */
double  w_val = 0.0 ; /* volume weight in N/m^3 - negative: <-, positive: -> */
long    w_min = -1  ; /* minimal element number for water load */
long    w_max = -1  ; /* maximal element number for water load */

/* random input data */
static long   *rand_type = NULL ; /* type of data (see README.RANDOM) */
static long   *rand_pos  = NULL ; /* index of data  */
static long   *rand_indx = NULL ; /* data index - if applicable */

/* optim input data */
long   *opt_type = NULL ; /* type of data (see README.RANDOM) */
long   *opt_pos  = NULL ; /* index of data  */
long   *opt_indx = NULL ; /* data index - if applicable */
double *opt_data = NULL ; /* data for replacing */

/* failure condition data */
long    fail_type = 0 ;    /* type of failure condition        */
long    n_fail    = 0 ;    /* number of failure condition data */
double *fail_data = NULL ; /* failure condition data           */

/* SOLUTION DATA */
tMatrix K ;
tVector u ;
tVector F ;

tMatrix Ke ;  /* 6x6 */
tMatrix D ;   /* 5x5 */
tMatrix B ;   /* 5x6 */
tMatrix Bt ;  /* 6x5 */
tMatrix BtD ; /* 6x5 */
tMatrix DB ;  /* 5x6 */
tVector Fe ;  /* 5 */
tVector ue ;  /* 6 */

/* result helpers data */
long  n_en = 0 ;
long *en_num = NULL;
long *en_frm = NULL;
long *en_pos = NULL ;

/* program constants */
int solution_only = 1 ;
int random_only   = 1 ;
int price_only    = 1 ;
int write_only    = 0 ;

/* frees input data */
void free_input_data(void)
{
  if (m_E1  != NULL){femDblFree(m_E1);}
  if (m_E2  != NULL){femDblFree(m_E2);}
  if (m_G   != NULL){femDblFree(m_G);}
  if (m_nu1 != NULL){femDblFree(m_nu1);}
  if (m_nu2 != NULL){femDblFree(m_nu2);}
  if (m_q   != NULL){femDblFree(m_q);}
  if (m_vp  != NULL){femDblFree(m_vp);}
  if (m_t  != NULL){femDblFree(m_t);}

  if (n_x   != NULL){femDblFree(n_x);}
  if (n_y   != NULL){femDblFree(n_y);}

  if (e_n1  != NULL){femIntFree(e_n1);}
  if (e_n2  != NULL){femIntFree(e_n2);}
  if (e_mat != NULL){femIntFree(e_mat);}
  if (e_t   != NULL){femDblFree(e_t);}

  if (d_n   != NULL){femIntFree(d_n);}
  if (d_dir != NULL){femIntFree(d_dir);}
  if (d_val != NULL){femDblFree(d_val);}

  if (n_f > 0) 
  {
    if (f_n   != NULL){femIntFree(f_n);}
    if (f_dir != NULL){femIntFree(f_dir);}
    if (f_val != NULL){femDblFree(f_val);}
  }

  if (n_r_inp > 0)
  {
    if (rand_type != NULL){femIntFree(rand_type);}
    if (rand_pos  != NULL){femIntFree(rand_pos);}
    if (rand_indx != NULL){femIntFree(rand_indx);}
  }

  if (n_r_opt > 0)
  {
    if (opt_type != NULL){femIntFree(opt_type);}
    if (opt_pos  != NULL){femIntFree(opt_pos);}
    if (opt_indx != NULL){femIntFree(opt_indx);}
    if (opt_data != NULL){femDblFree(opt_data);}
  }

  if (n_en > 0)
  {
    if (en_num != NULL){femIntFree(en_num);}
    if (en_frm != NULL){femIntFree(en_frm);}
    if (en_pos != NULL){femIntFree(en_pos);}
  }

  if (n_fail > 0)
  {
    if (fail_data != NULL){femDblFree(fail_data);}
  }

  n_m       = 0 ;
  n_n       = 0 ;
  n_e       = 0 ;
  n_d       = 0 ;
  n_f       = 0 ;
  n_r_inp   = 0 ;
  n_r_opt   = 0 ;

  n_en      = 0 ;

  fail_type = 0 ;
  n_fail    = 0 ;
}

/* first node must be always under the second - it exchanges them */
void check_elem_data(void)
{
  long i ;
  long tmp = 0 ;
  
  for (i=0; i<n_e; i++)
  {
    if (n_y[e_n1[i]] > n_y[e_n2[i]])
    {
      tmp = e_n1[i] ;
      e_n1[i] = e_n2[i] ;
      e_n2[i] = tmp ;
    }
  }
}

/* will prepare element nodes filed for optimised result output */
int get_enode_fields(void)
{
  long i, j ;

  if (en_num == NULL) { return(AF_ERR_VAL); }
  if (en_frm == NULL) { return(AF_ERR_VAL); }

  for (i=0; i<n_e; i++)
  {
    en_num[e_n1[i]]++ ;
    en_num[e_n2[i]]++ ;
  }

  n_en = 0 ;
  for (i=0; i<n_n; i++) 
  { 
    en_frm[i] = n_en ;

    n_en += en_num[i] ; 
  }

  if ((en_pos = femIntAlloc(n_en)) == NULL) { goto memFree; }

  for (i=0; i<n_en; i++) { en_pos[i] = -1 ; }

  for (i=0; i<n_e; i++)
  {
    for (j=0; j<en_num[e_n1[i]]; j++)
    {
      if (en_pos[en_frm[e_n1[i]]+j] == -1)
      {
        en_pos[en_frm[e_n1[i]]+j] = i ;
        break ;
      }
    }

    for (j=0; j<en_num[e_n2[i]]; j++)
    {
      if (en_pos[en_frm[e_n2[i]]+j] == -1)
      {
        en_pos[en_frm[e_n2[i]]+j] = i ;
        break ;
      }
    }
  }

  return(AF_OK);
memFree:
  return(AF_ERR_MEM);
}

/* reads data from stream
 * @param fw stream for reading
 */
int read_input_data(FILE *fw)
{
  long i ;

  if (fscanf(fw, "%li", &n_m) <= 0) {goto memFree;}
  if (n_m < 1) { fprintf(msgout,"Invalid number of materials!\n"); goto memFree ;}

  if (fscanf(fw, "%li", &n_n) <= 0) {goto memFree;}
  if (n_n < 2) { fprintf(msgout,"Invalid number of nodes!\n"); goto memFree ;}

  if (fscanf(fw, "%li", &n_e) <= 0) {goto memFree;}
  if (n_e < 2) { fprintf(msgout,"Invalid number of elements!\n"); goto memFree ;}

  if (fscanf(fw, "%li", &n_d) <= 0) {goto memFree;}
  if (n_d < 3) { fprintf(msgout,"Invalid number of supports!\n"); goto memFree ;}

  if (fscanf(fw, "%li", &n_f) < 0) {goto memFree;}
  if (n_f < 0) { fprintf(msgout,"Invalid number of forces!\n"); goto memFree ;}

#ifdef DEVEL_VERBOSE
  fprintf(msgout, 
      "\nMaterials: %li; Nodes: %li; Elements: %li; Supports: %li; Forces: %li\n",
      n_m, n_n, n_e, n_d, n_f);
#endif

  /* data allocations */
  if ((m_E1  = femDblAlloc(n_m)) == NULL) {goto memFree;}
  if ((m_E2  = femDblAlloc(n_m)) == NULL) {goto memFree;}
  if ((m_G   = femDblAlloc(n_m)) == NULL) {goto memFree;}
  if ((m_nu1 = femDblAlloc(n_m)) == NULL) {goto memFree;}
  if ((m_nu2 = femDblAlloc(n_m)) == NULL) {goto memFree;}
  if ((m_q   = femDblAlloc(n_m)) == NULL) {goto memFree;}
  if ((m_vp  = femDblAlloc(n_m)) == NULL) {goto memFree;}
  if ((m_t  = femDblAlloc(n_m)) == NULL) {goto memFree;}

  if ((n_x   = femDblAlloc(n_n)) == NULL) {goto memFree;}
  if ((n_y   = femDblAlloc(n_n)) == NULL) {goto memFree;}

  if ((e_n1  = femIntAlloc(n_e)) == NULL) {goto memFree;}
  if ((e_n2  = femIntAlloc(n_e)) == NULL) {goto memFree;}
  if ((e_mat = femIntAlloc(n_e)) == NULL) {goto memFree;}
  if ((e_t   = femDblAlloc(n_e)) == NULL) {goto memFree;}

  if ((d_n   = femIntAlloc(n_d)) == NULL) {goto memFree;}
  if ((d_dir = femIntAlloc(n_d)) == NULL) {goto memFree;}
  if ((d_val = femDblAlloc(n_d)) == NULL) {goto memFree;}

  if (n_f > 0)
  {
    if ((f_n   = femIntAlloc(n_f)) == NULL) {goto memFree;}
    if ((f_dir = femIntAlloc(n_f)) == NULL) {goto memFree;}
    if ((f_val = femDblAlloc(n_f)) == NULL) {goto memFree;}
  }

  if ((en_num = femIntAlloc(n_n)) == NULL) {goto memFree;}
  if ((en_frm = femIntAlloc(n_n)) == NULL) {goto memFree;}

  /* reading of data: */

  for (i=0; i<n_m; i++) /* materials */
  {
    if (fscanf(fw, " %lf %lf %lf %lf %lf %lf %lf %lf", 
        &m_E1[i], &m_E2[i], &m_G[i], &m_nu1[i], &m_nu2[i], &m_q[i], &m_vp[i], &m_t[i]) <= 0)
       { goto memFree ; }
    if ((m_E1[i] == m_E2[i])||(m_E2[i] <= 0.0)) /* isotropic*/
    {
      m_E2[i]  = m_E1[i] ;
      m_nu2[i] = m_nu1[i]; 
      if (m_G[i] <= 0.0) {m_G[i] = m_E1[i] / (2.0*(1.0 + m_nu1[i])); }
    }
    else
    {
      if ((m_E1[i] <= 0.0)|| (m_E2[i] <= 0.0)|| (m_G[i] <= 0.0)|| (m_nu1[i] <= 0.0)|| (m_nu2[i] <= 0.0))
      {
        fprintf(msgout,"Invalid or incomplete data for material %li\n", i);
        goto memFree;
      }
    }
  }

  for (i=0; i<n_n; i++) /* nodes */
  {
    if (fscanf(fw, "%lf %lf", &n_x[i], &n_y[i]) <= 0) { goto memFree ; }
  }

  for (i=0; i<n_e; i++) /* elements */
  {
    if (fscanf(fw, "%li %li %li %lf",
          &e_n1[i], &e_n2[i], &e_mat[i], &e_t[i]) <= 0) 
       { goto memFree ; }
    if ((e_n1[i] <0)||(e_n1[i] >= n_n)){fprintf(msgout,"Invalid n1 in element %li\n",i); goto memFree;}
    if ((e_n2[i] <0)||(e_n2[i] >= n_n)){fprintf(msgout,"Invalid n2 in element %li\n",i); goto memFree;}
    if ((e_n1[i] == e_n2[i])){fprintf(msgout,"Invalid nodes in element %li\n",i); goto memFree;}
    if ((e_mat[i] <0)||(e_mat[i] >= n_m)){fprintf(msgout,"Invalid material in element %li\n",i); goto memFree;}
    if (e_t[i] <= 0.0){fprintf(msgout,"Invalid width in element %li\n",i); goto memFree;}
  }

  for (i=0; i<n_d; i++) /* displacements */
  {
    if (fscanf(fw, "%li %li %lf", &d_n[i], &d_dir[i], &d_val[i]) <= 0)
       { goto memFree ; }
    if ((d_n[i] <0)||(d_n[i] >= n_n)){fprintf(msgout,"Invalid node in support %li\n",i); goto memFree;}
    if ((d_dir[i] <0)||(d_dir[i] >= 6)){fprintf(msgout,"Invalid direction in support %li\n",i); goto memFree;}
    if ((d_dir[i] >2)&&(d_val[i] < 0.0)){fprintf(msgout,"Invalid stiffness in support %li\n",i); goto memFree;}
  }

  for (i=0; i<n_f; i++) /* forces */
  {
    if (fscanf(fw, "%li %li %lf", &f_n[i], &f_dir[i], &f_val[i]) <= 0)
       { goto memFree ; }
    if ((f_n[i] <0)||(f_n[i] >= n_n)){fprintf(msgout,"Invalid node for force %li\n",i); goto memFree;}
    if ((f_dir[i] <0)||(f_dir[i] >= 3)){fprintf(msgout,"Invalid direction for force %li\n",i); goto memFree;}
  }

  /* water pressure data */
  if (fscanf(fw, "%lf %lf %lf %li %li", &w_top, &w_bot, &w_val, &w_min, &w_max) <= 0)
       { goto memFree ; }

#ifdef DEVEL_VERBOSE /* control output of input data */
  fprintf(msgout,
      "\nMATERIALS (%li)\n      E1          E2         G          nu1   nu2    q    price\n",n_m);
  for (i=0; i<n_m; i++)
  {
    fprintf(msgout,"%e %e %e %1.3f %1.3f %1.3f %4.1f\n",
        m_E1[i],m_E2[i],m_G[i],m_nu1[i],m_nu2[i],m_q[i],m_vp[i]);
  }

  fprintf(msgout, "\nNODES (%li)\n      X         Y\n",n_n);
  for (i=0; i<n_n; i++) { fprintf(msgout,"%e %e\n", n_x[i],n_y[i]); }

  fprintf(msgout,
      "\nELEMENTS (%li)\n    n1     n2 material width\n",n_e);
  for (i=0; i<n_e; i++)
  {
    fprintf(msgout,"%6li %6li %6li   %e\n", e_n1[i],e_n2[i],e_mat[i],e_t[i]);
  }

  fprintf(msgout, "\nSUPPORTS (%li)\n     n dir  size\n",n_d);
  for (i=0; i<n_d; i++)
    { fprintf(msgout,"%6li %2li  %e\n", d_n[i],d_dir[i],d_val[i]); }

  fprintf(msgout, "\nFORCES (%li)\n     n dir  size\n",n_f);
  for (i=0; i<n_f; i++)
    { fprintf(msgout,"%6li %2li  %e\n", f_n[i],f_dir[i],f_val[i]); }


  fprintf(msgout, "\nWATER PRESSURE\n     top       bottom         size     min  max\n");
  fprintf(msgout, "%e %e %e %4li %4li\n\n",w_top, w_bot, w_val, w_min, w_max) ;
#endif

  check_elem_data(); /* check of element nodes */

  if (get_enode_fields() != AF_OK) {goto memFree;}


  /* failure condition data: */
  if (fscanf(fw, "%li", &fail_type) <= 0) 
  {
    /* that's great, no failure is needed */
    fail_type = 0 ;
    n_fail    = 0 ;
  }
  else
  {
     if (fail_type > 0)
    {
      if (fscanf(fw, "%li", &n_fail) <= 0) 
      {
        fail_type = 0 ;
      }
      else
      {
        if((fail_data = femDblAlloc(n_fail)) == NULL)
        {
          fprintf(msgout, "Cannot allocate memory for failure data!\n");
          goto memFree;
        }
        for (i=0; i< n_fail; i++)
        {
          if (fscanf(fw,"%lf", &fail_data[i]) <= 0)
          {
            fprintf(msgout, "Invalid failure data!\n");
            goto memFree;
          }
        }
      }
    }
  }

  /* random variables: */
  if (fscanf(fw, "%li", &n_r_inp) <= 0) 
  {
#if _USE_WITH_MONTE_
    goto memFree;
#else
    n_r_inp = 0 ;
    /* fprintf(msgout, "No random data found.\n"); */
    return(AF_OK) ;
#endif
  }

  if (n_r_inp < 1) 
#if _USE_WITH_MONTE_
  { fprintf(msgout,"Invalid number of random inputs!\n"); goto memFree ;}
#else
  {return(AF_OK);}
#endif

  if ((rand_type = femIntAlloc(n_r_inp)) == NULL) {goto memFree;}
  if ((rand_pos  = femIntAlloc(n_r_inp)) == NULL) {goto memFree;}
  if ((rand_indx = femIntAlloc(n_r_inp)) == NULL) {goto memFree;}

  for (i=0; i<n_r_inp; i++)
  {
    if (fscanf(fw, "%li %li %li", &rand_type[i], &rand_pos[i], &rand_indx[i]) <= 0) { goto memFree ; }
  }

  /* optimized variables: ------------------------------------- */
  if (fscanf(fw, "%li", &n_r_opt) <= 0) 
  {
    n_r_opt = 0 ;
#ifndef _USE_WITH_MONTE_
    /*fprintf(msgout, "No optim. data found.\n"); */
#endif
    return(AF_OK) ;
  }

  if (n_r_opt < 1) 
#ifndef _USE_WITH_MONTE_
  { fprintf(msgout,"Invalid number of optim. inputs!\n"); return(AF_OK) ;}
#else
  {return(AF_OK);}
#endif

  if ((opt_type = femIntAlloc(n_r_opt)) == NULL) {goto memFree;}
  if ((opt_pos  = femIntAlloc(n_r_opt)) == NULL) {goto memFree;}
  if ((opt_indx = femIntAlloc(n_r_opt)) == NULL) {goto memFree;}
  if ((opt_data = femDblAlloc(n_r_opt)) == NULL) {goto memFree;}

  for (i=0; i<n_r_opt; i++)
  {
    if (fscanf(fw, "%li %li %li", &opt_type[i], &opt_pos[i], &opt_indx[i]) <= 0) { goto memFree ; }
  }

  for (i=0; i<n_r_opt; i++)
  {
    if (fscanf(fw, "%lf", &opt_data[i]) <= 0)
    {
      femDblFree(opt_data) ;
      femIntFree(opt_indx) ;
      femIntFree(opt_pos) ;
      femIntFree(opt_type) ;
      n_r_opt = 0 ;
    }
  }

  return(AF_OK) ;

memFree:
  free_input_data();
  fprintf(msgout, "Error when reading input!\n");
  return(AF_ERR_IO);
}

/* Writes input data to stream ------------------ */
int write_input_data(FILE *fw)
{
  long i ;

  /* sizes */
  fprintf(fw, "%li %li %li %li %li\n", n_m, n_n, n_e, n_d, n_f );

  for (i=0; i<n_m; i++) /* materials */
  {
    fprintf(fw," %e %e %e %e %e %e %e %e\n",
      m_E1[i], m_E2[i], m_G[i], m_nu1[i], m_nu2[i], m_q[i], m_vp[i], m_t[i] );
  }

  for (i=0; i<n_n; i++) /* nodes */
  { fprintf(fw, "%e %e\n", n_x[i], n_y[i]); }

  for (i=0; i<n_e; i++) /* elements */
  { fprintf(fw, "%li %li %li %e\n", e_n1[i], e_n2[i], e_mat[i], e_t[i]); }

  for (i=0; i<n_d; i++) /* displacements */
  { fprintf(fw, "%li %li %e\n", d_n[i], d_dir[i], d_val[i]); }

  for (i=0; i<n_f; i++) /* supports */
  { fprintf(fw, "%li %li %e\n", f_n[i], f_dir[i], f_val[i]); }

  /* water pressure data */
  fprintf(fw, "%e %e %e %li %li\n", w_top, w_bot, w_val, w_min, w_max);

  /* failure condition data: */
  fprintf(fw, "%li\n", fail_type) ;
  if (fail_type > 0)
  {
    fprintf(fw, "%li\n", n_fail) ;
    for (i=0; i< n_fail; i++) { fprintf(fw,"%e ", fail_data[i]); }
    fprintf(fw,"\n");
  }

  return(AF_OK) ;
}

/* Frees data used by solver */
void free_solver_data(void)
{
  femMatFree(&Ke) ;
  femMatFree(&D) ;
  femMatFree(&B) ;
  femMatFree(&Bt) ;
  femMatFree(&BtD) ;
  femMatFree(&DB) ;

  femVecFree(&ue);
  femVecFree(&Fe);

  femMatFree(&K); 
  femVecFree(&u);
  femVecFree(&F);
}

/* Allocates data for f.e. solver (K,u,F)*/
int alloc_solver_data(void)
{
  long i, j ;
  long *n_field = NULL ;
  long *alloc_field = NULL ;

  femMatNull(&K);
  femVecNull(&u);
  femVecNull(&F);

  femMatNull(&Ke) ;
  femMatNull(&D) ;
  femMatNull(&B) ;
  femMatNull(&Bt) ;
  femMatNull(&BtD) ;
  femMatNull(&DB) ;
  femVecNull(&Fe);
  femVecNull(&ue);

  if (femFullMatInit( &Ke, 6, 6) != AF_OK) {goto memFree;}
  if (femFullMatInit( &D, 5, 5) != AF_OK) {goto memFree;}
  if (femFullMatInit( &B, 5, 6) != AF_OK) {goto memFree;}
  if (femFullMatInit( &Bt, 6, 5) != AF_OK) {goto memFree;}
  if (femFullMatInit( &BtD, 6, 5) != AF_OK) {goto memFree;}
  if (femFullMatInit( &DB, 5, 6) != AF_OK) {goto memFree;}
  if (femVecFullInit(&Fe, 5) != AF_OK) { goto memFree ; }
  if (femVecFullInit(&ue, 6) != AF_OK) { goto memFree ; }

  /* Compute allocation vector */
  if ((n_field     = femIntAlloc(n_n)) == NULL) {goto memFree;}
  if ((alloc_field = femIntAlloc(n_n*3)) == NULL) {goto memFree;}

  for (i=0; i<n_n; i++)
  {
    for (j=0; j<n_e; j++)
    {
      if (e_n1[j] == i) {n_field[i]++;}
      if (e_n2[j] == i) {n_field[i]++;}
    }
  }

  for (i=0; i<n_n; i++)
  {
    for (j=0; j<3; j++)
    {
      alloc_field[3*i+j] = 3 * 6 * n_field[i] ; /* is "6" enough?*/
    }
  }

  /* alloc K, u, F */
  if (femSparMatInitDesc(&K, n_n*3, n_n*3, alloc_field) != AF_OK) { goto memFree ; }
  if (femVecFullInit(&F, n_n*3) != AF_OK) { goto memFree ; }
  if (femVecFullInit(&u, n_n*3) != AF_OK) { goto memFree ; }

  femIntFree(alloc_field);
  femIntFree(n_field);
  return(AF_OK);
memFree:
  if (alloc_field != NULL) femIntFree(alloc_field);
  if (n_field != NULL) femIntFree(n_field);
  free_solver_data();
  fprintf(msgout,"Out of memory!");
  return(AF_ERR_MEM);
}

/** computes material stiffness matrix of elemen
 * @param i element nomber <0..n_e-1>
 * @param t eleemnt width
 * @param D pointer to allocated (!) D matrix
 */
void get_D_matrix(long i, double t, tMatrix *D)
{
  double E1, E2, nu1, nu2, G, mult ;

  E1  = m_E1[e_mat[i]] ;
  E2  = m_E2[e_mat[i]] ;
  G   = m_G[e_mat[i]] ;
  nu1 = m_nu1[e_mat[i]] ;
  nu2 = m_nu2[e_mat[i]] ;

  mult = t / (1.0 - nu1*nu2) ;

  femMatPut( D, 1,1, E1 * mult );
  femMatPut( D, 1,2, nu2 * mult );
  femMatPut( D, 2,1, nu2 * mult );
  femMatPut( D, 2,2, E2 * mult );

  femMatPut( D, 3,3, (E1*t*t)/(12.0) * mult );
  femMatPut( D, 4,4, (E2*t*t)/(12.0) * mult );

  femMatPut( D, 3,4, nu2 * (E1*t*t)/(12.0) * mult );
  femMatPut( D, 4,3, nu2 * (E1*t*t)/(12.0) * mult );

  femMatPut( D, 5,5, (5.0/6.0) * G/t );


}

/** computes B matrix
 * @param i element number
 * @param B pointer to allocated (!) B matrix
 * @param Lc element length (result)
 * @param Rc average distance from axis or revolution
 */
void get_B_matrix(long i, tMatrix *B, double *Lc, double *Rc)
{
  double L,C,S,R, dx, dy;

  dx = n_x[e_n2[i]]-n_x[e_n1[i]] ;
  dy = n_y[e_n2[i]]-n_y[e_n1[i]] ;
  L = sqrt( pow(dx,2) + pow(dy,2));
  R = 0.5 * (  n_x[e_n1[i]] + n_x[e_n2[i]] ) ;

  S = -1.0 * dx / L ; 
  C = -1.0 * dy / L ;

  /* B matrix: */
  femMatPut(B, 1,1, -1.0*C/L ) ;
  femMatPut(B, 1,2, -1.0*S/L ) ;
  femMatPut(B, 1,4,  1.0*C/L ) ;
  femMatPut(B, 1,5,  1.0*S/L ) ;

  femMatPut(B, 2,2,  1.0/(2.0*R) ) ;
  femMatPut(B, 2,5,  1.0/(2.0*R) ) ;

  femMatPut(B, 3,3, -1.0/L ) ;
  femMatPut(B, 3,6,  1.0/L ) ;

  femMatPut(B, 4,3,  S/(2.0*R) ) ;
  femMatPut(B, 4,6,  S/(2.0*R) ) ;

  femMatPut(B, 5,1, -1.0*S/L ) ;
  femMatPut(B, 5,2,  1.0*C/L ) ;
  femMatPut(B, 5,3,  1.0/2.0 ) ;
  femMatPut(B, 5,4,  1.0*S/L ) ;
  femMatPut(B, 5,5, -1.0*C/L ) ;
  femMatPut(B, 5,6,  1.0/2.0 ) ;

  *Lc = L ;
  *Rc = R ;
}

/** creates stiffness matrix */
int get_matrix(void)
{
  double t ;
  double L, R, F2, q;
  long i, j, k, posj, posk ;

  femMatSetZero(&K);
  femVecSetZero(&u);
  femVecSetZero(&F);

  for (i=0; i<n_e; i++)
  {
    /* if material width is specified then use element width: */
    if ((t = m_t[e_mat[i]]) <= 0.0) { t = e_t[i] ; } 
    t = e_t[i] ;

    femMatSetZero(&Ke);
    femMatSetZero(&B);
    femMatSetZero(&Bt);
    femMatSetZero(&BtD);
    femMatSetZero(&D);

    /* material stiffness matrix D: */
    get_D_matrix(i, t, &D);

    /* B matrix */
    get_B_matrix(i, &B, &L, &R);

    /* transpose of B */
    femMatTran(&B, &Bt);

    /* matrix multiplications (Bt*D*B): */
    femMatMatMult(&Bt, &D, &BtD); /* => BtD*/
    femMatMatMult(&BtD, &B, &Ke); /* => Ke  without L*R */

    /* element stifness matrix Ke: */
    femValMatMultSelf(R*L, &Ke) ;

    /* localisation to "K": */
    for (j=1; j<=6; j++)
    {
      if (j<4) { posj = (e_n1[i]*3) + j ; }
      else { posj = (e_n2[i]*3) + j - 3 ; }

      for (k=1; k<=6; k++)
      {
        if (k<4) { posk = (e_n1[i]*3) + k ; }
        else { posk = (e_n2[i]*3) + k - 3 ; }

        femMatAdd(&K, posj, posk, femMatGet(&Ke, j, k) ) ; 
      }
    }

    /* gravitation */
    if (fabs( (q=m_q[e_mat[i]]) ) > FEM_ZERO)
    {
      F2 = (-0.5) * q * t * L ;
      femVecAdd(&F, 3*e_n1[i], F2) ;
      femVecAdd(&F, 3*e_n2[i], F2) ;
    }
  }

  return(AF_OK);
}

/** generates water pressure load */
int generate_water_load_x(void)
{
  /* it goes through elements and decides if they are under the
   * water level (or over the bottom) then it computes horizontal
   * pressure on the element nodes
   */
  long   i ;
  double y1, y2, dx, L, val1, val2 ;
  long   from, to ;
  long   down = 1 ;
  long   use_1 = 1 ; /* don't ignore this node */
  long   use_2 = 1 ; /* don't ignore this node */
  long   pos1, pos2 ;
  double y_max, y_min ; /* real limits of water position */
  double a,b ; /* hydrostatic pressures on element - top, bot */

  if ( fabs(w_val) > (100*FEM_ZERO) )
  {
    /* limits for element testing (probably unused): */
    if ((w_max-w_min) == 0)
    {
      from = 0 ; 
      to = n_e ;
    }
    else
    {
      if ( (w_min < 0) || (w_min >= n_e) )
          { from = 0 ; }
      else { from = w_min ; }
      if ( (w_max < 0) || (w_max > n_e) )
          { to = n_e ; }
      else { to = w_max ; }
    }

    /* setting of unreachable limits for water */
    y_min = n_y[e_n1[from]] ; y_max = y_min ;
    for (i=from; i<to; i++)
    {
      if (y_min > n_y[e_n1[i]]) { y_min = n_y[e_n1[i]] ; }
      if (y_min > n_y[e_n2[i]]) { y_min = n_y[e_n2[i]] ; }
      
      if (y_max < n_y[e_n1[i]]) { y_max = n_y[e_n1[i]] ; }
      if (y_max < n_y[e_n2[i]]) { y_max = n_y[e_n2[i]] ; }
    }

    /* adjusting limits: */
    if (w_top < y_max) {y_max = w_top ;}
    if (w_bot > y_min) {y_min = w_bot ;}

    for (i=from; i<to; i++)
    {
      y1 = n_y[e_n1[i]] ;
      y2 = n_y[e_n2[i]] ;

      /* geometric features: */
      if ((y1 > y_max) || (y1 < y_min)) {use_1 = 0 ;}
      if ((y2 > y_max) || (y2 < y_min)) {use_2 = 0 ;}
      if ((use_1 == 0) && (use_2 == 0)) {continue;}
      
      if (y1 > y2) 
      { 
        down = 2 ; 
        val1 = y1 ;
        y1   = y2 ;
        y2   = val1 ;
      }
      if (y1 < y_min) {y1 = y_min ;}
      if (y2 > y_max) {y2 = y_max ;}

      dx = fabs(n_x[e_n2[i]] - n_x[e_n1[i]]) ;
      L = sqrt ( dx*dx + pow(y2-y1, 2) ) ;

      if (pow(y2-y1, 2) < FEM_ZERO) {continue;} /* nothing to do */

      /* TODO: compute limit values */
      b = (y_max-y1) * w_val ;
      a = (y_max-y2) * w_val ;

#ifdef DEVEL_VERBOSE
      fprintf(msgout,"Y: %e %e, a=%e b=%e\n", y1,y2,a,b);
#endif

      /* set values in nodes: */
      if (use_1 == 0)
      {
        val2 = ( a + 0.5 * (b-a) ) * L  ;
        val1 = 0.0 ;
      }
      else
      {
        if (use_2 == 0)
        {
          val1 =  ( a + 0.5 * (b-a) ) * L  ;
          val2 = 0.0 ;
        }
        else
        {
          val1 = 0.5*a*L + 0.25*(b-a)*L + 0.125*(b-a)*L ;
          val2 = 0.5*a*L + 0.125*(b-a)*L ;
        }
      }

      /* positions of loads */
      if (down == 1)
      {
        /* val1 (lower) is at n1 */
        pos1 = e_n1[i]*3 + 1 ;
        pos2 = e_n2[i]*3 + 1 ;
      }
      else
      {
        /* val1 is at n2 */
        pos1 = e_n2[i]*3 + 1 ;
        pos2 = e_n1[i]*3 + 1 ;
      }

      /* adding of loads: */
      femVecAdd(&F, pos1, val1 ) ;
      femVecAdd(&F, pos2, val2 ) ;

#ifdef DEVEL_VERBOSE
      fprintf(msgout, "ADDED: e[%li] f%li(%li)<- %e, f%li(%li)<- %e, L=%e dx=%e\n", i, pos1,e_n1[i], val1, pos2,e_n2[i], val2, L, dx);
#endif
    }
  }

  return(AF_OK) ;
}

/* applies supports in nodes */
int get_loads_and_supports(void)
{
  long i, j, pos ;

  for (i=0; i<n_f; i++)
      { femVecAdd(&F, f_n[i]*3 + f_dir[i]+1, f_val[i] ) ; }

  for (i=0; i<n_d; i++)
  {
    if (d_dir[i] > 2)
    {
      /* stifnesses */
      pos = d_n[i]*3 + d_dir[i]-2 ;
      femMatAdd(&K, pos, pos, d_val[i]);
    }
    else
    {
      /* displacements */
      pos = d_n[i]*3 + d_dir[i]+1 ;

      if (fabs(d_val[i]) <= FEM_ZERO)
      {
        femMatSetZeroCol(&K, pos) ;
        femMatSetZeroRow(&K, pos) ;
        femVecPut(&u, pos, 0.0 );
        femVecPut(&F, pos, 0.0 ); /* yes, it deletes force in support */
        femMatPut(&K, pos, pos, 1.0) ;
      }
      else
      {
        for (j=1; j<=n_n*3; j++)
        {
          femVecAdd(&F, j, -1.0*femMatGet(&K, j, pos)*d_val[i]) ;
        }
        femMatSetZeroCol(&K, pos) ;
        femMatSetZeroRow(&K, pos) ;
        femVecPut(&u, pos, d_val[i] );
        femMatPut(&K, pos, pos, femVecGet(&F,pos) / d_val[i]) ;
      }
    }
  }

  return(AF_OK);
}


/** computes internal force is nodes
 * @param el element number <0..n_e-1>
 * @param N1 meridian force
 * @param N2 perpendicular force
 * @param M1 meridian moment
 * @param M2 perpendicular force
 * @param Q tangent force
 * @return status
 */
void get_int_forces(long el, 
    double *N1,
    double *N2,
    double *M1,
    double *M2,
    double *Q
    )
{
  double t, L, R ;
  long j, posj ;

  femMatSetZero(&D);
  femMatSetZero(&B);
  femMatSetZero(&DB);

  femVecSetZero(&ue) ;
  femVecSetZero(&Fe) ;

  /* get local stiffness vector */
  for (j=1; j<=6; j++)
  {
    if (j<4) { posj = (e_n1[el]*3) + j ; }
    else { posj = (e_n2[el]*3) + j - 3 ; }
    femVecPut(&ue, j, femVecGet(&u, posj) ) ; 
  }

  /* get B and D */
  t   = e_t[el] ;
  get_D_matrix(el, t, &D);
  get_B_matrix(el, &B, &L, &R);
  femMatMatMult(&D, &B, &DB);

  /* get vector */
  femMatVecMult(&DB, &ue, &Fe);

 *N1 =  femVecGet(&Fe, 1) ;
 *N2 =  femVecGet(&Fe, 2) ;
 *M1 =  femVecGet(&Fe, 3) ;
 *M2 =  femVecGet(&Fe, 4) ;
 *Q  =  femVecGet(&Fe, 5) ;
}

int print_result(FILE *fw)
{
#ifdef RUN_VERBOSE
  long i, j;
  long count = 0;
  double N1, N2, Q, M1, M2 ;
  double sN1, sN2, sQ, sM1, sM2 ;

  N1=0.0; N2=0.0; M1=0.0; M2=0.0; Q=0.0;
  sN1=0.0; sN2=0.0; sM1=0.0; sM2=0.0; sQ=0.0;

  fprintf(fw, "#  X     Y        w            u           angle            N1          N2           M1          M2          Q\n");
  for (i=0; i<n_n; i++)
  {
    sN1=0.0; sN2=0.0; sM1=0.0; sM2=0.0; sQ=0.0;
#if 1
    count = 0 ;
    for (j=0; j<n_e; j++)
    {
      if ((e_n1[j] == i)||(e_n2[j] == i))
      {
        get_int_forces(j, &N1, &N2, &M1, &M2, &Q ); /* internal forces in centroid */
        sN1 += N1 ; sN2 += N2 ; sM1 += M1 ; sM2 += M2 ; sQ += Q ;
        count++ ;
      }
    }
#else
    for (j=0; j<en_num[i]; j++)
    {
      count = en_num[i] ;
      get_int_forces(en_pos[en_frm[i]]+j, &N1, &N2, &M1, &M2, &Q );
      sN1 += N1 ; sN2 += N2 ; sM1 += M1 ; sM2 += M2 ; sQ += Q ;
    }
#endif
    if (count > 0)
    {
      sN1 /= ((double)count) ;
      sN2 /= ((double)count) ;
      sM1 /= ((double)count) ;
      sM2 /= ((double)count) ;
      sQ  /= ((double)count) ;
    }

    fprintf(fw, "%2.3f %2.3f %e %e %e %e %e %e %e %e \n",
        n_x[i], n_y[i], 
        femVecGet(&u, 3*i+1),
        femVecGet(&u, 3*i+2),
        femVecGet(&u, 3*i+3),
        sN1, sN2, sM1, sM2, Q
        );
  }
#endif
  return(AF_OK);
}

/* generates output variable list for Monte input file */
void generate_rand_out_file(FILE *fw)
{
  long i ;

  fprintf(fw,"%li\n", n_n*8 + 1 );

  fprintf(fw,"FAIL 3 2\n");

  for (i=0; i<n_n; i++)
  {
    fprintf(fw,"UY%li 2\n", i);
    fprintf(fw,"UX%li 2\n", i);
    fprintf(fw,"RT%li 2\n", i);
    fprintf(fw,"NX%li 2\n", i);
    fprintf(fw,"NY%li 2\n", i);
    fprintf(fw,"MX%li 2\n", i);
    fprintf(fw,"MY%li 2\n", i);
    fprintf(fw,"QQ%li 2\n", i);
  }
  fprintf(fw,"0\n"); /* no correlations at all */
}

/* generates textual symbol for displacement */
char *generate_d_type(int type)
{
  switch (type)
  {
    case 0: return("UY"); break ;
    case 1: return("UX"); break ;
    case 2: return("RT"); break ;
    case 3: return("EY"); break ;
    case 4: return("EX"); break ;
    case 5: return("ER"); break ;
  }
  return("XX");
}

/* generates textual symbol for force */
char *generate_f_type(int type)
{
  switch (type)
  {
    case 0: return("FY"); break ;
    case 1: return("FX"); break ;
    case 2: return("MT"); break ;
  }
  return("XX");
}

/* generates textual symbol for water load */
char *generate_w_type(int type)
{
  switch (type)
  {
    case 0: return("TOP"); break ;
    case 1: return("BOT"); break ;
    case 2: return("SIZE"); break ;
  }
  return("XX");
}

/* generates textual symbol for failure criteria */
char *generate_fc_type(int type)
{
  switch (fail_type)
  {
    case 1: /* concrete cracking limit */
      switch (type)
      {
        case 0: return("COMPR"); break ;
        case 1: return("TENS"); break ;
        default: return("UNKNOWN"); break ;
      }
      break;

    default:
      return("XX");
      break ;
  }
  return("XX");
}


/** Writes input data for Monte 
 * @param fw file stream to write data 
 * @return status
 */
void generate_rand_input_file(FILE *fw)
{
  long i ;

  fprintf(fw, "%li\n", n_r_inp); 

  for (i=0; i<n_r_inp; i++)
  {
    switch (rand_type[i])
    {
      case 0: /* material */
        switch (rand_indx[i])
        {
          case 0: 
            fprintf(fw,"MAT%li_E1 %e 1 normal-1-02.dis\n",
               rand_pos[i], m_E1[rand_pos[i]]);
               break ;
          case 1:
            fprintf(fw,"MAT%li_E2 %e 1 normal-1-02.dis\n",
               rand_pos[i], m_E2[rand_pos[i]]);
               break ;
          case 2: 
            fprintf(fw,"MAT%li_G %e 1 normal-1-02.dis\n",
               rand_pos[i], m_G[rand_pos[i]]);
               break ;
          case 3:
            fprintf(fw,"MAT%li_NU1 %e 1 normal-1-02.dis\n",
               rand_pos[i], m_nu1[rand_pos[i]]);
               break ;
          case 4:
            fprintf(fw,"MAT%li_NU2 %e 1 normal-1-02.dis\n",
               rand_pos[i], m_nu2[rand_pos[i]]);
               break ;
          case 5:
            fprintf(fw,"MAT%li_VF %e 1 normal-1-02.dis\n",
               rand_pos[i], m_vp[rand_pos[i]]);
               break ;
          case 6:
            fprintf(fw,"MAT%li_T %e 1 normal-1-02.dis\n",
               rand_pos[i], m_t[rand_pos[i]]);
               break ;
        }
        break;
      case 1: /* node */
        switch (rand_indx[i])
        {
          case 0:
            fprintf(fw,"N%li_X %e 1 normal-1-02.dis\n",
               rand_pos[i], n_x[rand_pos[i]]);
               break ;
          case 1: 
            fprintf(fw,"N%li_Y %e 1 normal-1-02.dis\n",
               rand_pos[i], n_y[rand_pos[i]]);
               break ;
        }
        break;
      case 2: /* element width */
            fprintf(fw,"E%li_WIDTH %e 1 normal-1-02.dis\n",
               rand_pos[i], e_t[rand_pos[i]]);
          break;
      case 3: /* displacement */
            fprintf(fw,"D%li_%s_SIZE %e 1 normal-1-02.dis\n",
               rand_pos[i], generate_d_type(rand_indx[i]), d_val[rand_pos[i]]);
          break;
      case 4: /* force */
            fprintf(fw,"F%li_%s_SIZE %e 1 normal-1-02.dis\n",
               rand_pos[i],  generate_f_type(rand_indx[i]), f_val[rand_pos[i]]);
          break;
      case 5: /* node */
        switch (rand_indx[i])
        {
          case 0:
            fprintf(fw,"W_%s %e 1 normal-1-02.dis\n",
               generate_w_type(rand_indx[i]),w_top);
               break ;
          case 1: 
            fprintf(fw,"W_%s %e 1 normal-1-02.dis\n",
               generate_w_type(rand_indx[i]),w_bot);
               break ;
          case 2: 
            fprintf(fw,"W_%s %e 1 normal-1-02.dis\n",
               generate_w_type(rand_indx[i]),w_val);
               break ;
        }
        break;
      case 6: /* failure critical */
            fprintf(fw,"FC_%s_%li %e 1 normal-1-02.dis\n",
              generate_fc_type(rand_indx[i]),
              rand_indx[i],
              fail_data[rand_indx[i]]);
        break;

      default:
          fprintf(msgout,"Unused input random variable %li!\n", i);
          break;
    }
  }
    
}

/* ** FAILURE CRITERIA DEFINITIONS ** */

/** 
 *  provides failure testing*/
long fail_test_concrete(void)
{
  double N1, N2, Q, M1, M2, h ;
  double I1, J2, J3, alpha, beta, lambda, k, cos3f, c1, c2, fc ;
  double s1, s2, sm, tmp ;
  long i ;

  k = fail_data[1] / fail_data[0] ;

  for (i=0; i<n_e; i++)
  {
    get_int_forces(i, &N1, &N2, &M1, &M2, &Q ); /* internal forces in centroid */
    h = e_t[i] ;
    s1 = (6.0*M1) / h  + (N1) / h ;
    s2 = (6.0*M2) / h  + (N2) / h ;
		if (s1 < s2) {tmp=s1; s1=s2; s2=tmp;}

    I1 = s1+s2 ;
    sm = I1 / 3.0 ;
    J3 = (s1-sm)*(s2-sm) ;
    J2 = (1.0/6.0) * ( pow(s1-s2,2) + s1*s1 + s2*s2 ) ;
    alpha = 1.0 / (9.0 * pow(k, 1.4) ) ;
    beta = 1.0 / (3.7 * pow(k, 1.1) ) ;
    cos3f = ( 3.0*pow(3.0,0.5)/2.0 ) * ( J3 / pow(J2, 1.5)) ;

    c1 = 1.0 / (0.7 * pow(k, 1.1)) ;
    c2 = 1.0 - 6.8* pow(k-0.07, 2) ;

    if (cos3f < 0.0)
    {
      lambda = c1 * cos(FEM_PI/3.0 - (1.0/3.0)*acos(0.0-c2*cos3f) ) ;
    }
    else
    {
      lambda = c1 * cos( (1.0/3.0)*acos(0.0 -c2*cos3f) ) ;
    }

    fc = alpha*(J2/pow(fail_data[0],2)) 
       + lambda * (sqrt(J2)/fail_data[0])
       + beta * (I1/fail_data[0]) ;

#ifdef DEVEL_VERBOSE
    fprintf(stderr,"[%li] fc = %e, ft = %e\n", i, fail_data[0], fail_data[1]);
    fprintf(stderr,"[%li] s1: %e, s2: %e \n sm: %e I1: %e, J2: %e, J3: %e\n",
        i, s1, s2, sm, I1, J2, J3 );

    fprintf(stderr,"[%li] alpha: %e, beta: %e, lambda: %e cos3f: %e\n c1: %e, c2: %e => fc: %e\n",
        i, alpha, beta, lambda, cos3f, c1, c2, fc );
#endif

    if (fc > 1.0)
    {
      /* failed */
#ifdef DEVEL_VERBOSE
      fprintf(stderr,"Element %li FAILED\n", i);
#endif
      return(1);
    }
  }
  return(0);
}

/** runs failure test
 * @return 1 for failure, 0 for tother cases
 */
long fail_test(void)
{
  switch (fail_type)
  {
    case 1: /* concrete: no-crack allowed */
      return(fail_test_concrete());
      break ;
    case 0:  /* no criteria -> no fail */
    default:
      return(0);
      break ;
  }
  return(0);
}

/** Computes price of the structure based on material volume */
double compute_price(void)
{
  static double price ;
  double volume, dx, dpx, dy ;
  long i ;

  price = 0.0 ;

  for (i=0; i<n_e; i++)
  {
    dx = fabs( n_x[e_n2[i]] - n_x[e_n1[i]] ); /* R-r */
    dpx = n_x[e_n2[i]] + n_x[e_n1[i]] ; /* R+r */
    dy = fabs( n_y[e_n2[i]] - n_y[e_n1[i]] );

    if (dx <= FEM_ZERO) /* cillinder */
    {
      volume = dy * ( 2.0 * FEM_PI * n_x[e_n2[i]] ) ; /* 2*pi*r */
    }
    else
    {
      if (dy <= FEM_ZERO) /* circle in plane */
      {
        volume = FEM_PI*fabs(pow(n_x[e_n2[i]],2)-pow(n_x[e_n1[i]],2));
      }
      else /* arbitrary shape */
      {
        volume = FEM_PI*dpx*sqrt(dy*dy + dx*dx );
      }
    }
    price += e_t[i]* volume * m_vp[e_mat[i]] ;
  }
  return(price);
}

/* replace f.e. input  data with their optimized counterparts */
int optim_replace_data(double *ifld)
{
  long i ;

  if ((ifld == NULL) || (n_r_opt < 1)) {return(AF_OK);}

  for (i=0; i<n_r_opt; i++)
  {
    switch (opt_type[i])
    {
      case 0: /* material */
        switch (opt_indx[i])
        {
          case 0: m_E1[opt_pos[i]] = ifld[i] ; break ;
          case 1: m_E2[opt_pos[i]] = ifld[i] ; break ;
          case 2: m_G[opt_pos[i]] = ifld[i] ; break ;
          case 3: m_nu1[opt_pos[i]] = ifld[i] ; break ;
          case 4: m_nu2[opt_pos[i]] = ifld[i] ; break ;
          case 5: m_q[opt_pos[i]] = ifld[i] ; break ;
          case 6: m_t[opt_pos[i]] = ifld[i] ; break ;
        }
        break;
      case 1: /* node */
        switch (opt_indx[i])
        {
          case 0: n_x[opt_pos[i]] = ifld[i] ; break ;
          case 1: n_y[opt_pos[i]] = ifld[i] ; break ;
        }
        break;
      case 2: /* element width */
          e_t[opt_pos[i]] = ifld[i] ; 
          break;
      case 3: /* displacement */
          d_val[opt_pos[i]] = ifld[i] ; 
          break;
      case 4: /* force */
          f_val[opt_pos[i]] = ifld[i] ; 
          break;
      case 5: /* material */
        switch (opt_indx[i])
        {
          case 0: w_top = ifld[i] ; break ;
          case 1: w_bot = ifld[i] ; break ;
          case 2: w_val = ifld[i] ; break ;
        }
        break;
      case 6: /* failure condition */
          if (opt_indx[i] < n_fail)
          {
            fail_data[opt_indx[i]] = ifld[i] ;
          }
          break ;

      default:
          fprintf(msgout,"Unused input optim variable %li!\n", i);
          break;
    }
  }
  return(AF_OK);
}

#ifdef _USE_WITH_MONTE_
/* ======================================================== */
/* FUNCTIONS for Monte interaction  */

/* replace f.e. input  data with their random counterparts */
int monte_replace_data(double *ifld)
{
  long i ;

  for (i=0; i<n_r_inp; i++)
  {
    switch (rand_type[i])
    {
      case 0: /* material */
        switch (rand_indx[i])
        {
          case 0: m_E1[rand_pos[i]] = ifld[i] ; break ;
          case 1: m_E2[rand_pos[i]] = ifld[i] ; break ;
          case 2: m_G[rand_pos[i]] = ifld[i] ; break ;
          case 3: m_nu1[rand_pos[i]] = ifld[i] ; break ;
          case 4: m_nu2[rand_pos[i]] = ifld[i] ; break ;
          case 5: m_q[rand_pos[i]] = ifld[i] ; break ;
          case 6: m_t[rand_pos[i]] = ifld[i] ; break ;
        }
        break;
      case 1: /* node */
        switch (rand_indx[i])
        {
          case 0: n_x[rand_pos[i]] = ifld[i] ; break ;
          case 1: n_y[rand_pos[i]] = ifld[i] ; break ;
        }
        break;
      case 2: /* element width */
          e_t[rand_pos[i]] = ifld[i] ; 
          break;
      case 3: /* displacement */
          d_val[rand_pos[i]] = ifld[i] ; 
          break;
      case 4: /* force */
          f_val[rand_pos[i]] = ifld[i] ; 
          break;
      case 5: /* material */
        switch (rand_indx[i])
        {
          case 0: w_top = ifld[i] ; break ;
          case 1: w_bot = ifld[i] ; break ;
          case 2: w_val = ifld[i] ; break ;
        }
        break;
      case 6: /* failure condition */
          if (rand_indx[i] < n_fail)
          {
            fail_data[rand_indx[i]] = ifld[i] ;
          }
          break ;

      default:
          fprintf(msgout,"Unused input random variable %li (type %li)!\n", i,rand_type[i]);
          break;
    }
  }
  return(AF_OK);
}

/* gets output data from f.e. solution */
int monte_collect_results(double *ofld)
{
  long i, j ;
  long count = 0;
  double N1, N2, Q, M1, M2 ;
  double sN1, sN2, sQ, sM1, sM2 ;

  N1=0.0; N2=0.0; M1=0.0; M2=0.0; Q=0.0;
  sN1=0.0; sN2=0.0; sM1=0.0; sM2=0.0; sQ=0.0;

  for (i=0; i<n_n; i++)
  {
    sN1=0.0; sN2=0.0; sM1=0.0; sM2=0.0; sQ=0.0;
#if 0
    count = 0 ;
    for (j=0; j<n_e; j++)
    {
      if ((e_n1[j] == i)||(e_n2[j] == i))
      {
        get_int_forces(j, &N1, &N2, &M1, &M2, &Q ); /* internal forces in centroid */
        sN1 += N1 ; sN2 += N2 ; sM1 += M1 ; sM2 += M2 ; sQ += Q ;
        count++ ;
      }
    }
#else
    for (j=0; j<en_num[i]; j++)
    {
      count = en_num[i] ;
      get_int_forces(en_pos[en_frm[i]]+j, &N1, &N2, &M1, &M2, &Q );
      sN1 += N1 ; sN2 += N2 ; sM1 += M1 ; sM2 += M2 ; sQ += Q ;
    }
#endif
    if (count > 0)
    {
      sN1 /= ((double)count) ;
      sN2 /= ((double)count) ;
      sM1 /= ((double)count) ;
      sM2 /= ((double)count) ;
      sQ  /= ((double)count) ;
    }

    ofld[8*i+1] = femVecGet(&u, 3*i+1);
    ofld[8*i+2] = femVecGet(&u, 3*i+2);
    ofld[8*i+3] = femVecGet(&u, 3*i+3);
    ofld[8*i+4] = sN1;
    ofld[8*i+5] = sN2;
    ofld[8*i+6] = sM1;
    ofld[8*i+7] = sM2;
    ofld[8*i+8] = Q;
  }

  ofld[0] = fail_test() ;

  return(AF_OK);
}

/* simulation data initial NULL-ing */
void monte_io_null(void)
{
  monte_i_len = 0 ; /* input variables */
  monte_o_len = 0 ; /* output variables */
}

/* returns number of variables */
#ifdef USE_WIN32
EXPORT void monte_nums_of_vars(char *param, long *ilen, long *olen, long *ffunc)
#else
void monte_nums_of_vars(char *param, long *ilen, long *olen, long *ffunc)
#endif
{
  *ilen = monte_i_len ; /* required number of input variables */
  *olen = monte_o_len ; /* returned number of output variables */
  *ffunc = 0 ; /* currently not available */
  return ;
}

/* allocation of simulation data */
int monte_io_alloc(long ilen, long olen)
{
  double n ;

  monte_io_null();

  n = ilen + olen ;

  if (n <= 0) {return(AF_OK);}
  
  monte_i_len = ilen ;
  monte_o_len = olen ;

  return(AF_OK);
}

/* interface type definition for Monte (2 is for the advanced type) */
#ifdef USE_WIN32
EXPORT long monte_dlib_interface_type(void) { return(2) ; }
#else
long monte_dlib_interface_type(void) { return(2) ; }
#endif

/* allocation of structural and f.e. data  */
#ifdef USE_WIN32
EXPORT int monte_init_lib_stuff(char *param)
#else
int monte_init_lib_stuff(char *param)
#endif
{
  FILE *fr = NULL ;

#ifdef RUN_VERBOSE
  msgout = stderr ; /* for output from "fprintf(msgout,...)" */
#endif

  if (param == NULL) { return(-1); }
  if (strlen(param) < 1) { return(-1); }
  
  if ((fr = fopen(param,"r")) == NULL) { goto memFree; }
  if (read_input_data(fr) != AF_OK) { goto memFree ; }
  fclose(fr);
  if (monte_io_alloc(n_r_inp, n_n*8+1) != AF_OK)  { goto memFree ; }
  if (alloc_solver_data() != AF_OK) { goto memFree ; }
  
  if (opt_data != NULL) {optim_replace_data(opt_data);}

  return(0);
memFree:
#ifdef RUN_VERBOSE
  fprintf(msgout, "Invalid or non-existant data file!\n") ;
#endif
  return(-1);
}

/* cleaning of structural and f.e. data
 * @param param input data file name
 */
#ifdef USE_WIN32
EXPORT int monte_clean_lib_stuff(char *param)
#else
int monte_clean_lib_stuff(char *param)
#endif
{
#if 0 /* it makes segfaults in monte => disabled */
  free_solver_data();
  free_input_data();
  monte_io_null();
#endif
  return(0);
}

/* f.e. solution solution
 * @param para, input data file name (unused here)
 * @param ifld random data input
 * @param ofld random data output
 */
#ifdef USE_WIN32
EXPORT int monte_solution(char *param, double *ifld, double *ofld)
#else
int monte_solution(char *param, double *ifld, double *ofld)
#endif
{
  int rv = 0;

  rv = monte_replace_data(ifld) ;

  if (rv == AF_OK) rv = get_matrix();
  if (rv == AF_OK) rv = generate_water_load_x();
  if (rv == AF_OK) rv = get_loads_and_supports();
  if (rv == AF_OK) rv = femEqsCGwJ(&K, &F, &u, 1e-9, 6*3*n_n) ;

  if (rv == AF_OK) rv = monte_collect_results(ofld) ;

  return(rv);
}

/* ======================================================== */
#else

/** Prints simple help to stdout
 * @param argc the same as "argc" from main
 * @param argv the same as "argv" from main
 */
void print_help(int argc, char *argv[])
{
  printf("\neSHELL 1.0: axisymetric shells solver\n");
  printf("(C) 2010 VSB-TU of Ostrava \n");
  printf("(C) 2003-2010 Jiri Brozovsky (uFEM libraries)\n");
  printf("\nThis is free software licensed under GNU GPL 2.0\n");
  printf("\nUsage: %s [parameters] <input >output\n", argv[0]);
  printf("\nParameters:\n");
  printf("   -s        ... force solution only output\n");
  printf("   -g        ... generate random data only \n");
  printf("   -p        ... compute price function only\n");
  printf("   -w        ... write input data and finish\n");
  printf("   -h        ... print this help\n");
}

/** Understands command line parameters */
int cmd_param(int argc, char *argv[])
{
  int i;

  for (i=1; i<argc; i++)
  {
    if ((strcmp(argv[i],"-h") == 0)||(strcmp(argv[i],"--help") == 0))
       { print_help(argc, argv) ; exit(AF_OK); }

    if ((strcmp(argv[i],"-s") == 0)||(strcmp(argv[i],"--solution") == 0))
       { solution_only = 1 ; price_only = 0 ; random_only = 0 ; }
    if ((strcmp(argv[i],"-g") == 0)||(strcmp(argv[i],"-r") == 0)||(strcmp(argv[i],"--random") == 0))
       { solution_only = 0 ; price_only = 0 ; random_only = 1 ; }
    if ((strcmp(argv[i],"-p") == 0)||(strcmp(argv[i],"--price") == 0))
       { solution_only = 0 ; price_only = 1 ; random_only = 0 ; }
    if ((strcmp(argv[i],"-w") == 0)||(strcmp(argv[i],"--price") == 0))
       { write_only = 1 ; }
  }
  return(AF_OK);
}

/** main() routine for standalone program only. */
int main(int argc, char *argv[])
{
  int stat = 0 ;
  msgout = stderr ;

  cmd_param(argc, argv);

  stat += read_input_data(stdin);
  stat += alloc_solver_data();

  stat += optim_replace_data(opt_data);

  if (write_only == 1) { write_input_data(stdout); return(0); }

  if (solution_only == 1) 
  {
    stat += get_matrix();
    stat += generate_water_load_x();
    stat += get_loads_and_supports();
    stat = femEqsCGwJ(&K, &F, &u, 1e-9, 6*3*n_n);
  }

  if ((n_r_inp > 0)&&(random_only == 1))
  {
    if (solution_only != 0) { print_result(stderr);}
    generate_rand_input_file(stdout);
    generate_rand_out_file(stdout);
  }
  else
  {
    if (solution_only == 1) { print_result(stdout); }
  }

  if (solution_only == 1) 
  {
    if (fail_test() != 0) { fprintf(stderr,"# Structure FAILED\n"); }
  }

  if (price_only == 1)
  {
    if (solution_only == 1)
         { fprintf(msgout,"# Price is %lf\n", compute_price()); }
    else { fprintf(stdout,"%e\n", compute_price()); }
  }

#if 0
  free_input_data();
  free_solver_data();
#endif

  return(AF_OK);
}
#endif

/* end of eshell.c */
