package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{3, 2, 4, 8, 9, 6, 1, 1}
	sort.Ints(s)
	fmt.Println(s)
}