package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void Pt() {
	printf("This a Pt func\n");
}

inline void SetString(char* dest, const char* str) {
  size_t len = strlen(str);

  if (strlen(dest) != len) {
    free(dest);
    dest = (char*)malloc(len * sizeof(char));
  }

  strcpy(dest, str);
};
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func malloc() {
	fmt.Println("char* with manually malloc")
	str := C.malloc(C.sizeof_char)
	defer C.free(str)
	val := C.CString("Hello World!")


	C.SetString((*C.char)(str), val)
	fmt.Printf("str: %s, size: %d\n", C.GoString((*C.char)(str)), int(C.strlen((*C.char)(str))))
}

func up() {
	fmt.Println("char* with unsafe Pointer")
	str := C.CString("")
	defer C.free(unsafe.Pointer(str))
	val := C.CString("Hello world!")

	C.SetString(str, val)
	fmt.Printf("str: %s, size: %d\n", C.GoString(str), C.strlen(str))
}

func pt() {
	C.Pt()
}

func main() {
	malloc()
	up()
	fmt.Println()
	C.Pt()
}
