package main

import (
	"fmt"

	"../../object"
)

func main() {
	o1 := object.Obj{Name: "o1", Id: 0}
	fmt.Println("o1 ->", o1)
	(&o1).SetId(100)
	fmt.Println("use address (&o1) set id to 100 -> ")
	fmt.Println(o1)
}