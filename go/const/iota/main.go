package main

import (
	"fmt"
)

// ConstType const type
type ConstType uint

const (
	one ConstType = iota
	two
	three
	four
	five
)

func main() {
	var one ConstType = one
	var two ConstType = two
	var three ConstType = three
	var four ConstType = four

	fmt.Println(one)
	a := one & two
	fmt.Print("0 & 1: ")
	fmt.Println(a)
	b := one | two
	fmt.Print("0 | 1: ")
	fmt.Println(b)
	
	if (three == 2) {
		fmt.Println("const type of uint can match uint number")
	}
	if (four != 3) {
		fmt.Println("const type of uint can not match uint number")
	}
}