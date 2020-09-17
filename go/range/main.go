package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 6, 9, 4, 2}
	for k, v := range s {
		fmt.Print(k)
		fmt.Print(" ")
		fmt.Println(v)
	}
}