package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 6, 9, 4, 2}
	fmt.Println(s)
	fmt.Println("Len of s: ", len(s))
	for k, v := range s {
		fmt.Println(k, v)
	}

	empty := []int{}
	isEmpty := false
	fmt.Println()
	fmt.Println("Empty range")
	for k, v := range empty {
		isEmpty = true
		fmt.Println(k, v)
	}
	fmt.Println(isEmpty)
}