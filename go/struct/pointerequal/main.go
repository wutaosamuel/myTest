package main

import (
	"fmt"
)

// Obj a object
type Obj struct {
	Name string
	ID   int
}

func main() {
	o1 := Obj{"o1", 1}
	o2 := Obj{"o1", 1}
	fmt.Println("Test Two Obj value")
	if o1 == o2 {
		fmt.Println("o1 == o2")
	}
	if o1 != o2 {
		fmt.Println("o1 != o2")
	}

	p1 := &Obj{"p1", 1}
	p2 := &Obj{"p1", 1}
	fmt.Println("Test two Obj pointer")
	if p1 == p2 {
		fmt.Println("p1 == p2")
	}
	if p1 != p2 {
		fmt.Println("p1 != p2")
	}

	pv1 := &Obj{"pv1", 1}
	pv2 := &Obj{"pv1", 1}
	fmt.Println("Test two Obj pointer value")
	if *pv1 == *pv2 {
		fmt.Println("*pv1 == *pv2")
	}
	if *pv1 != *pv2 {
		fmt.Println("*pv1 != *pv2")
	}
}
