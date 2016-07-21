/* Created by "go tool cgo" - DO NOT EDIT. */

/* package command-line-arguments */

/* Start of preamble from import "C" comments.  */


#line 3 "/home/dvirsky/code/go-redis-modules-sdk/go/module.go"

#include <stdlib.h>
#include "../redismodule.h"

static char *rm_string(RedisModuleString **s, int offset) {
	return (char*)RedisModule_StringPtrLen(s[offset], NULL);
}

extern int GoDispatch(RedisModuleCtx* p0, RedisModuleString** p1, int p2);

static int rm_createCmd(RedisModuleCtx *ctx, char *cmd, char *flags, int i, int j, int k) {
	return RedisModule_CreateCommand(ctx, cmd, GoDispatch, flags, i,j,k);
}

static int rm_replySimpleString(RedisModuleCtx *ctx, char *str) {
	return RedisModule_ReplyWithSimpleString(ctx, str);
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
typedef __complex float GoComplex64;
typedef __complex double GoComplex128;

// static assertion to make sure the file is being used on architecture
// at least with matching size of GoInt.
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

typedef struct { char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


extern int GoDispatch(RedisModuleCtx* p0, RedisModuleString** p1, int p2);

extern int GoOnLoad(RedisModuleCtx* p0);

#ifdef __cplusplus
}
#endif
