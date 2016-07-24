/* Created by "go tool cgo" - DO NOT EDIT. */

/* package _/home/dvirsky/code/go-redis-modules-sdk/go */

/* Start of preamble from import "C" comments.  */


#line 3 "/home/dvirsky/code/go-redis-modules-sdk/go/dispatch.go"

#include "../redismodule.h"
static char *rm_string(RedisModuleString **s, int offset) {
	return (char*)RedisModule_StringPtrLen(s[offset], NULL);
}


#line 3 "/home/dvirsky/code/go-redis-modules-sdk/go/redis.go"

#include <stdlib.h>
#include "../redismodule.h"



extern int GoDispatch(RedisModuleCtx* p0, RedisModuleString** p1, int p2);

static int rm_CreateCmd(RedisModuleCtx *ctx, char *cmd, char *flags, int i, int j, int k) {
	return RedisModule_CreateCommand(ctx, cmd, GoDispatch, flags, i,j,k);
}

static int rm_replyWithSimpleString(RedisModuleCtx *ctx, char *str) {
	return RedisModule_ReplyWithSimpleString(ctx, str);
}

static int rm_replyWithError(RedisModuleCtx *ctx, char *str) {
	return RedisModule_ReplyWithError(ctx, str);
}

static int rm_replyWithArray(RedisModuleCtx *ctx, long len) {
	return RedisModule_ReplyWithArray(ctx, len);
}

static void rm_setArrayLength(RedisModuleCtx *ctx, long len) {
	 RedisModule_ReplySetArrayLength(ctx, len);
}

static int rm_replyWithString(RedisModuleCtx *ctx, const char *buf, size_t len) {
	return RedisModule_ReplyWithStringBuffer(ctx, buf, len);
}

static int rm_replyWithNull(RedisModuleCtx *ctx) {
	return RedisModule_ReplyWithNull(ctx);
}

static int rm_replyWithDouble(RedisModuleCtx *ctx, double d) {
	return RedisModule_ReplyWithDouble(ctx, d);
}

static int rm_replyWithLongLong(RedisModuleCtx *ctx, long long l) {
	return RedisModule_ReplyWithLongLong(ctx, l);
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


extern int GoDispatch(RedisModuleCtx* p0, RedisModuleString** p1, int p2);

extern int GoOnLoad(RedisModuleCtx* p0);

#ifdef __cplusplus
}
#endif
