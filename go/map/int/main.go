package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2, 6, 5, 3, 7, 6, 9, 2}
	sort.Ints(s)
	fmt.Println(s)
	var m map[int]bool
	m = make(map[int]bool, 0)
	for _, v := range s {
		m[v] = true
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}