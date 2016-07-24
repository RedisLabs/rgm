package main

/*
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

*/
import "C"

import (
	"C"
	"errors"
	"strings"
)

// RedisModule is a go wrapper on a redis context
type RedisModule struct {
	ctx *C.RedisModuleCtx
}

// ErrModule represents a generic return from functions that return REDISMODULE_ERROR
var ErrModule = errors.New("Redis Error")

// ReplyWithSimpleString sends a protocol reply with an error or simple string (status message)
func (r *RedisModule) ReplyWithSimpleString(s string) error {
	if rc := C.rm_replyWithSimpleString(r.ctx, C.CString(s)); rc != C.REDISMODULE_OK {
		return ErrModule
	}
	return nil
}

// ReplyWithString sends a protocol reply with a string
func (r *RedisModule) ReplyWithString(s string) error {
	if rc := C.rm_replyWithString(r.ctx, C.CString(s), C.size_t(len(s))); rc != C.REDISMODULE_OK {
		return ErrModule
	}
	return nil
}

func (r *RedisModule) ReplyWithLongLong(l int64) error {
	if rc := C.rm_replyWithLongLong(r.ctx, C.longlong(l)); rc != C.REDISMODULE_OK {
		return ErrModule
	}
	return nil
}

func (r *RedisModule) ReplyWithDouble(d float64) error {
	if rc := C.rm_replyWithDouble(r.ctx, C.double(d)); rc != C.REDISMODULE_OK {
		return ErrModule
	}
	return nil
}

func registerCmd(ctx *C.RedisModuleCtx, cmd, flags string, handler RedisHandler) error {

	if C.rm_CreateCmd(ctx, C.CString(cmd), C.CString(flags), 1, 1, 1) == C.REDISMODULE_ERR {
		return errors.New("Could not register command")
	}

	handlers[strings.ToLower(cmd)] = handler
	return nil
}

//export MODULE_NAME
var ModuleName = "FOO"

/*

 */

