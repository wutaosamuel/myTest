package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Object {
	char* Name;
	int ID;
};

inline struct Object* NewObject() {
	struct Object* obj = (struct Object*) malloc(sizeof(struct Object));
	return obj;
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

type Object struct {
	Obj *C.struct_Object
}

func NewObject() *Object {
	return &Object{
		Obj: C.NewObject(),
	}
}

func (o *Object) Free() {
	C.free(unsafe.Pointer(o.Obj))
}

func main() {
	obj := NewObject()
	defer obj.Free()
	obj.Obj.ID = C.int(2)
	fmt.Printf("id: %d\n", int(obj.Obj.ID))
}
