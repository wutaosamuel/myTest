package main

import (
	"fmt"
)

// Object struct
type Object struct {
	Name string
}

func main() {
	m := make(map[string]Object, 0)
	test := m["test"]
	test.Name = "test"
	m["test"] = test
	fmt.Println(m)

	// FIXME: not work
	// p := make(map[string]*Object, 0)
	// pointer := p["pointer"]
	// pointer.Name = "pointer"
	// fmt.Println(p)
}