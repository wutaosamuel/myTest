package main

import (
	"fmt"
)

func main() {
	a00 := make([]string, 0, 0)
	a01 := make([]string, 0, 10)
	a11 := make([]string, 10)
	a12 := make([]string, 5, 10)

	fmt.Println("a00: len=", len(a00))
	fmt.Println("a01: len=", len(a01))
	fmt.Println("a11: len=", len(a11))
	fmt.Println("a12: len=", len(a12))
}