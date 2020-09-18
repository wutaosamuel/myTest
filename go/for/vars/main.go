package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i, j := 0, 0; i <= 10 && j <=5; i, j = i+1, j+1 {
		fmt.Println(i, j)
		sum = i+j
		fmt.Println("sum", sum)
	}
	for i, j := 0, 0; i <= 10 && j <=5; i++ {
		fmt.Println(i, j)
		if i % 2 == 0 {
			j++
		}
	}
}