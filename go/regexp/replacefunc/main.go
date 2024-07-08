package main

import (
	"regexp"
	"fmt"
)

func main() {
	// phrase flags
	// read input
	// replace

	reg := regexp.MustCompile(`{{ ([^{}]*) }}`)
	str := reg.ReplaceAllStringFunc("{{city}}, {{ state }} {{ zip }}", replace)
	fmt.Println()
	fmt.Println(str)
}

func replace(str string) string {
	fmt.Println(str)
	return ""
}