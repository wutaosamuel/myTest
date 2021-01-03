package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string, 5)
	m["first"] = "1"
	m["second"] = "2"
	m["third"] = "3"
	m["fourth"] = "4"
	m["fifth"] = "5"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
