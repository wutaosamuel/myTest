package main

import (
	"fmt"
)

// Obj obj object
type Obj struct {
	Name	string
	ID	int
}

func main() {
	m := make(map[int]Obj, 3)
	mp := make(map[int]*Obj, 3)
	m[0] = Obj{"o0", 1}
	m[1] = Obj{"o1", 11}
	m[2] = Obj{"o2", 111}

	mp[0] = &Obj{"op0", 1}
	mp[1] = &Obj{"op1", 11}
	mp[2] = &Obj{"op2", 111}
	fmt.Println("m:", m)
	fmt.Println("mp:", mp)

	fmt.Println()
	fmt.Println("m change value")
	for k, v := range m {
		o := &v
		o.ID = k
	}
	fmt.Println("m:", m)

	fmt.Println()
	fmt.Println("mp change value")
	for k, v := range mp {
		v.ID = k
	}
	for k, v := range mp {
		fmt.Println(k, v)
	}
}