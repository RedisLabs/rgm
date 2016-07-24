/* Created by "go tool cgo" - DO NOT EDIT. */

/* package github.com/RedisLabs/rgm/module */

/* Start of preamble from import "C" comments.  */


#line 3 "/home/dvirsky/go/src/github.com/RedisLabs/rgm/module/dispatch.go"

#include "common.h"

static char *rm_string(RedisModuleString **s, int offset) {
	return (char*)RedisModule_StringPtrLen(s[offset], NULL);
}


#line 3 "/home/dvirsky/go/src/github.com/RedisLabs/rgm/module/module.go"

// The pure C part of the module's initialization callback
#include "common.h"

static int rm_CreateCmd(RedisModuleCtx *ctx, char *cmd, char *flags, int i, int j, int k) {
	return RedisModule_CreateCommand(ctx, cmd, goDispatch, flags, i,j,k);
}



/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

typedef struct { const char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


extern int goDispatch(RedisModuleCtx* p0, RedisModuleString** p1, int p2);

extern char* getModuleName();

extern int goOnLoad(RedisModuleCtx* p0);

#ifdef __cplusplus
}
#endif
