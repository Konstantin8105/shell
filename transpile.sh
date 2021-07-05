#!/bin/bash

c4go transpile -o base.go   \
	-clang-flag="-lm"       \
	./c-src/shell/fem_mem.c \
	./c-src/shell/eshell.h  \
	./c-src/shell/eshell.c  \
	./c-src/shell/fem.h     \
	./c-src/shell/fem_math.c\
	./c-src/shell/fem_eqs.c
