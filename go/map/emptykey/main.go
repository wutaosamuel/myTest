package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string, 3)

	// FIXME: works
	m[""] = "empty"
	m["1"] = ""
	m["2"] = "full"
	for k, v := range m {
		fmt.Println(k, v)
	}
}