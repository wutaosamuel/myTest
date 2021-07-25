package main

import (
	"fmt"

	"../../object"
)

func main() {
	m := make(map[int]object.Obj, 3)
	m[0] = object.Obj{Name: "o0", Id: 1}
	m[1] = object.Obj{Name: "o1", Id: 11}
	m[2] = object.Obj{Name: "o2", Id: 111}

	fmt.Println("init map: ")
	for k, v := range(m) {
		fmt.Println(k, v)
	}

	for k, v := range(m) {
		(&v).SetId(k)
	}
	fmt.Println("(&v) reset value -> not work")
	for k, v := range(m) {
		fmt.Println(k, v)
	}

	for k, v := range(m) {
		v.Id = k
		m[k] = v
	}
	fmt.Println("m[k] = v -> work")
	for k, v := range(m) {
		fmt.Println(k, v)
	}
}