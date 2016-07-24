package main

/*
#include "common.h"

static char *rm_string(RedisModuleString **s, int offset) {
	return (char*)RedisModule_StringPtrLen(s[offset], NULL);
}
*/
import "C"

type RedisHandler func(*RedisModule, []string) error

var handlers = map[string]RedisHandler{}

// convertArgs converts a redis argument list to a go string slice
func convertArgs(argv **C.RedisModuleString, argc int) []string {
	args := make([]string, 0, argc)
	for i := 0; i < argc; i++ {

		arg := C.rm_string(argv, C.int(i))
		args = append(args, C.GoString(arg))
	}
	return args
}

//export goDispatch
func goDispatch(ctx *C.RedisModuleCtx, argv **C.RedisModuleString, argc C.int) C.int {

	args := convertArgs(argv, int(argc))

	r := &RedisModule{ctx}
	h := handlers[args[0]]

	h(r, args)
	//fmt.Println(args)

	return C.REDISMODULE_OK // C.CString(fmt.Sprintf("got %d args, command was %s", argc, args[0]))
}
