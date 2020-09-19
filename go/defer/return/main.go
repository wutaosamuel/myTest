package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start to test no return")
	noReturn()
	fmt.Println()
	fmt.Println("Start to test have return")
	haveReturn()
}

func noReturn() {
	fmt.Println("noReturn: Start")
	defer fmt.Println("noReturn: End - defer")
}

func haveReturn() {
	fmt.Println("haveReturn: Start")
	defer fmt.Println("haveReturn: End - defer")
	return
}