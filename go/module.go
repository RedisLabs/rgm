package main

/*
#include <stdlib.h>
*/
import "C"

//export AGoFunction
func AGoFunction() *C.char {
	return C.CString("foo motherfucker")
}

func main() {}
