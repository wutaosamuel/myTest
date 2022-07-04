package main

import (
	"fmt"
)

func main() {
	fmt.Println("start to test goto keyword")
	fmt.Println("start test a")
	s := pt("a")
	fmt.Println("result a is", s)
}

func pt(s string) string {
	if s == "a" {
		goto aTag
	}

	if s == "b" {
		goto bTag
	}

	if s == "c" {
		goto cTag
	}

	return "not a or b or c"

aTag:
	a()
bTag:
	b()
cTag:
	c()

	return "run a or b or c"
}

func a() {
	fmt.Println("this is a func print")
}

func b() {
	fmt.Println("this is b func print")
}

func c() {
	fmt.Println("this is c func print")
}
