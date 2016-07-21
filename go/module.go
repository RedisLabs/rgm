package main

/*
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

*/
import "C"
import (
	"errors"
	"fmt"
	"strings"
)

// RedisCtx is a go wrapper on a redis context
type RedisCtx struct {
	ctx *C.RedisModuleCtx
}

func (r *RedisCtx) ReplyWithSimpleString(s string) error {
	C.rm_replySimpleString(r.ctx, C.CString(s))
	return nil

}

type RedisHandler func(*RedisCtx, []string) error

var handlers = map[string]RedisHandler{}

func HandleFoo(ctx *RedisCtx, args []string) error {
	ctx.ReplyWithSimpleString("It worked!")
	return nil
}

//export GoDispatch
func GoDispatch(ctx *C.RedisModuleCtx, argv **C.RedisModuleString, argc C.int) C.int {

	fmt.Println("Dispatch!")

	args := make([]string, 0, argc)
	for i := 0; i < int(argc); i++ {

		arg := C.rm_string(argv, C.int(i))
		args = append(args, C.GoString(arg))
	}

	r := &RedisCtx{ctx}
	h := handlers[args[0]]

	h(r, args)
	//fmt.Println(args)

	return C.REDISMODULE_OK // C.CString(fmt.Sprintf("got %d args, command was %s", argc, args[0]))
}

func registerCmd(ctx *C.RedisModuleCtx, cmd, flags string, handler RedisHandler) error {

	if C.rm_createCmd(ctx, C.CString(cmd), C.CString(flags), 1, 1, 1) == C.REDISMODULE_ERR {
		return errors.New("Could not register command")
	}

	handlers[strings.ToLower(cmd)] = handler
	return nil
}

//export GoOnLoad
func GoOnLoad(ctx *C.RedisModuleCtx) C.int {

	if err := registerCmd(ctx, "go.foo", "readonly", HandleFoo); err != nil {
		return C.REDISMODULE_ERR
	}

	return C.REDISMODULE_OK

}

func main() {}
