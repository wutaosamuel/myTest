package main

import (
	"regexp"
	"fmt"
)

const (
	Str = "{{city}}, { state } {{ zip }}"
	Pattern1 = `{([^{}]*)}`
	Pattern2 = `{{([^{}]*)}}`
	Pattern3 = `{{ ([^{}]*) }}`
)

func main() {
	fmt.Println("String:", Str)
	fmt.Println()

	findAllString(Pattern1)
	fmt.Println()
	findAllString(Pattern2)
	fmt.Println()
	findAllString(Pattern3)
	fmt.Println()
}

func findAllString(pattern string) {
	fmt.Println("Pattern:", pattern)
	reg := regexp.MustCompile(pattern)
	matches := reg.FindAllStringSubmatch(Str, -1)
	for _, v := range matches {
		fmt.Println(v[1])
	}
}
