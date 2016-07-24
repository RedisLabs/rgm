package module

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
import (
	"fmt"
	"reflect"
)

// Redis is a go wrapper on a redis context
type Redis struct {
	ctx *C.RedisModuleCtx
}

// ErrModule represents a generic return from functions that return Redis_ERROR
var ErrModule = errors.New("Redis Error")

// ReplyWithSimpleString sends a protocol reply with an error or simple string (status message)
func (r *Redis) ReplyWithSimpleString(s string) error {
	if rc := C.rm_replyWithSimpleString(r.ctx, C.CString(s)); rc != C.REDISMODULE_OK {
		return ErrModule
	}
	return nil
}

// ReplyWithString sends a protocol reply with a string
func (r *Redis) ReplyWithString(s string) error {
	if rc := C.rm_replyWithString(r.ctx, C.CString(s), C.size_t(len(s))); rc != C.REDISMODULE_OK {
		return ErrModule
	}
	return nil
}

func (r *Redis) ReplyWithLongLong(l int64) error {
	if rc := C.rm_replyWithLongLong(r.ctx, C.longlong(l)); rc != C.REDISMODULE_OK {
		return ErrModule
	}
	return nil
}

func (r *Redis) ReplyWithDouble(d float64) error {
	if rc := C.rm_replyWithDouble(r.ctx, C.double(d)); rc != C.REDISMODULE_OK {
		return ErrModule
	}
	return nil
}

func (r *Redis) ReplyWithArray(len int) error {
	if rc := C.rm_replyWithArray(r.ctx, C.long(len)); rc != C.REDISMODULE_OK {
		return ErrModule
	}
	return nil
}

// Reply is a generic reply function that tries to guess the reply type expected based on the value type
func (r *Redis) Reply(v interface{}) error {
	switch val := v.(type) {
	case string:
		return r.ReplyWithString(val)
	case int:
		return r.ReplyWithLongLong(int64(val))
	case int32:
		return r.ReplyWithLongLong(int64(val))
	case int64:
		return r.ReplyWithLongLong(int64(val))
	case int16:
		return r.ReplyWithLongLong(int64(val))
	case int8:
		return r.ReplyWithLongLong(int64(val))
	case float32:
		return r.ReplyWithDouble(float64(val))
	case float64:
		return r.ReplyWithDouble(val)
	case []byte:
		return r.ReplyWithString(string(val))
	case []string:
		if err := r.ReplyWithArray(len(val)); err != nil {
			return err
		}
		for _, s := range val {
			if err := r.ReplyWithString(s); err != nil {
				return err
			}
		}
	case []interface{}:
		if err := r.ReplyWithArray(len(val)); err != nil {
			return err
		}
		for _, x := range val {
			if err := r.Reply(x); err != nil {
				return err
			}
		}

	case error:
		r.ReplyWithSimpleString(val.Error())
	default:
		return fmt.Errorf("Could not encode value of type %s", reflect.TypeOf(val))
	}

	return nil
}
