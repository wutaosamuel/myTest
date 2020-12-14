package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This is the test!\nThis is second line"
	fmt.Println(s)
	s = strings.ReplaceAll(s, "\n", " ")
	ss := strings.Split(s, " ")
	for k, v := range ss {
		fmt.Println(k, v)
	}

	n := strings.Split(s, "@")
	fmt.Println()
	fmt.Println("No match: ")
	for k, v := range n {
		fmt.Println(k, v)
	}
}