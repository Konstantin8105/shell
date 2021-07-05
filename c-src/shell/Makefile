# Makefile for the "eshell": solver of axisymetric thin eshells
#
# (C) 2010 Jiri Brozovsky
#
#
all: eshell libshell.so

CC=gcc -fPIC
LD=ld -shared

#CC=/home/jirka/src/qi/openwrt-xburst/staging_dir/toolchain-mipsel_gcc-4.3.3+cs_uClibc-0.9.30.1/usr/bin/mipsel-openwrt-linux-uclibc-gcc -fPIC
#LD=/home/jirka/src/qi/openwrt-xburst/staging_dir/toolchain-mipsel_gcc-4.3.3+cs_uClibc-0.9.30.1/usr/bin/mipsel-openwrt-linux-uclibc-ld -shared

# Win32:
#CC=/usr/bin/i586-mingw32msvc-gcc -DUSE_WIN32
#LD=/usr/bin/i586-mingw32msvc-gcc -shared -Wl,--out-implib,shell.a -DUSE_WIN32


CFLAGS=-Wall -O0 -DRUN_VERBOSE -g -DDEVEL_VERBOSE
#CFLAGS=-Wall -O0 -DRUN_VERBOSE -g
CFLAGS=-Wall -O2 -DRUN_VERBOSE

LIBS=-lm -L/usr/local/lib #-lefence
OBJECTS=fem_eqs.o fem_math.o fem_mem.o eshell.o

eshell: $(OBJECTS)
	$(CC) $(CFLAGS) $(OBJECTS) $(LIBS) -o $(@)


eshell.o: eshell.h eshell.c fem.h
	$(CC) -c $(CFLAGS) eshell.c

fem_eqs.o: fem_math.h fem_eqs.c fem.h
	$(CC) -c $(CFLAGS) fem_eqs.c

fem_math.o: fem_math.h fem_math.c fem.h
	$(CC) -c $(CFLAGS) fem_math.c

fem_mem.o: fem_mem.h fem_mem.c fem.h
	$(CC) -c $(CFLAGS) fem_mem.c

eshellib.o:  eshell.h eshell.c fem.h
	$(CC) -c $(CFLAGS) -D_USE_WITH_MONTE_ eshell.c -o $(@)

libshell.so: eshellib.o fem_eqs.o fem_math.o fem_mem.o
		$(LD) eshellib.o fem_eqs.o fem_math.o fem_mem.o -o $(@)

shell.dll: eshellib.o fem_eqs.o fem_math.o fem_mem.o
		$(LD) eshellib.o fem_eqs.o fem_math.o fem_mem.o -o $(@)



clean:
	-rm -f *.o eshell libshell.so eshell.exe core *.tk *.dis so_locations *.fem
