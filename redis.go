package main

/*
#include <stdlib.h>
#include "common.h"



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
