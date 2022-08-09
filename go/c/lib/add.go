package lib

/*
#include "./add.h"
*/
import "C"

func Add(a, b int) int {
	sum, _ := C.add(C.int(a), C.int(b))
	return int(sum)
}
