package main

import (
	"fmt"
)

func main() {
	s := make(map[int][]int, 0)
	s[0] = []int{1, 2, 3}
	s[1] = []int{4, 5, 6}
	s[2] = []int{7, 8, 9}
	fmt.Println("origin:")
	fmt.Println(s)

	fmt.Println()
	fmt.Println("change value")
	s[0] = append(s[0], 0)
	fmt.Println("add s[0] + 0")
	fmt.Println(s)
}
