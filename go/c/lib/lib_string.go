package lib

/*
#include <stdio.h>
#include <stdlib.h>
#include "./lib_string.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func SetString() {
	fmt.Println("set string")
	str := C.malloc(C.sizeof_char)
	defer C.free(str)
	val := C.CString("Hello world!")
	defer C.free(unsafe.Pointer(val))
	e := C.utils_SetString((*C.char)(str), val)
	fmt.Printf("e: %d, str: %s\n", int(e), C.GoString((*C.char)(str)))

	s := C.CString("")
	defer C.free(unsafe.Pointer(s))
	C.utils_SetString(s, val)
	fmt.Printf("s: %s\n", C.GoString(s))
}
