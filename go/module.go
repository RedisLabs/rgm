package main

/*
#include <stdlib.h>
#include "../redismodule.h"

static char *rm_string(RedisModuleString **s, int offset) {
	return (char*)RedisModule_StringPtrLen(s[offset], NULL);
}


*/
import "C"
import "fmt"

//export GoDispatch
func GoDispatch(ctx *C.RedisModuleCtx, argv **C.RedisModuleString, argc C.size_t) *C.char {

	args := make([]string, 0, argc)
	for i := 0; i < int(argc); i++ {

		arg := C.rm_string(argv, C.int(i))
		args = append(args, C.GoString(arg))
	}

	fmt.Println(args)

	return C.CString(fmt.Sprintf("got %d args, command was %s", argc, args[0]))
}

func main() {}
