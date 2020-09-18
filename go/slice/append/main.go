package main

import (
	"fmt"
)

func main() {
	s := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := make([]int, 0)
	fmt.Println("Get first 2", s[:2])
	fmt.Println("Get remove first 3", s[3:])
	result = append(result, s[:2]...)
	result = append(result, s[3:]...)
	fmt.Println("remove 2", result)
}