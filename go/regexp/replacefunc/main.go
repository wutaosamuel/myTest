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
	reg.ReplaceAllStringFunc("{{city}}, {{ state }} {{ zip }}", replace)
}

func replace(str string) string {
	fmt.Println(str)
	return ""
}