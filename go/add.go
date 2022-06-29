package main

// #include "../add.c"
import "C"

func Add(a, b int) int {
	sum, _ := C.add(C.int(a), C.int(b))
	return int(sum)
}
