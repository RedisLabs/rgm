#set environment variable RM_INCLUDE_DIR to the location of redismodule.h
ifndef RM_INCLUDE_DIR
	RM_INCLUDE_DIR=../
endif

ifndef RMUTIL_LIBDIR
	RMUTIL_LIBDIR=../rmutil
endif

# find the OS
uname_S := $(shell sh -c 'uname -s 2>/dev/null || echo not')

# Compile flags for linux / osx
ifeq ($(uname_S),Linux)
	SHOBJ_CFLAGS ?=  -fno-common -g -ggdb
	SHOBJ_LDFLAGS ?= -shared -Bsymbolic
else
	SHOBJ_CFLAGS ?= -dynamic -fno-common -g -ggdb
	SHOBJ_LDFLAGS ?= -bundle -undefined dynamic_lookup
endif
CFLAGS = -I$(RM_INCLUDE_DIR) -Wall -g -fPIC -lc -lm -Og -std=gnu99  
CC=gcc

.PHONY: module.go
all: module.so 

module.so: module.o module.go
	$(LD) -o $@ module.o module.a $(SHOBJ_LDFLAGS) $(LIBS) -lc 

module.go:
	CGO_CFLAGS="-fPIC" go build -o module.so -gcflags=-shared -asmflags -shared -installsuffix _shared -buildmode=c-shared redis.go module.go onload.go

clean:
	rm -rf *.xo *.so *.o *.a
