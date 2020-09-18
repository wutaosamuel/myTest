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
}