package main

import (
	"regexp"
	"fmt"
)

const (
	Str = "{{ city }}, { state } {{ zip }}"
	Pattern = `{{ ([^{}]*) }}`
)

var KV = map[string]string{
	"city":   "Shang Hai",
	"state":  "Shang Hai",
	"zip":    "30000",
}

func main() {
	fmt.Println("String:", Str)
	fmt.Println()

	findAllString(Str, Pattern)
}

func findAllString(str, pattern string) {
	//var result []byte = []byte{}

	fmt.Println("Pattern:", pattern)
	reg := regexp.MustCompile(pattern)
	reg.ReplaceAllStringFunc(str, func(s string) string {
		fmt.Println(s)
		key := reg.FindStringSubmatch(s)[1]
		fmt.Println(key)
		return ""
	})
}
