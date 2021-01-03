package main

import (
	"fmt"
)

// Obj object
type Obj struct {
	Name map[int][]int
}

// Objects objects
type Objects struct {
	Obj map[int]Obj
}

// NewObjects new objects
func NewObjects() *Objects {
	o := make(map[int][]int, 1)
	o[0] = []int{1,2,3}
	o[1] = []int{1,2,3}

	os := make(map[int]Obj, 1)
	os[0] = Obj{o}
	os[1] = Obj{o}

	return &Objects{os}
}

func main() {
	s := make(map[int][]int, 0)
	s[0] = []int{1, 2, 3}
	s[1] = []int{4, 5, 6}
	s[2] = []int{7, 8, 9}
	fmt.Println("origin:")
	fmt.Println(s)

	fmt.Println()
	fmt.Println("change value")
	s[0] = append(s[0], 0)
	fmt.Println("add s[0] + 0")
	fmt.Println(s)

	fmt.Println()
	fmt.Println("No pointer: ")
	os := NewObjects()
	fmt.Println(os)
	fmt.Println("add 4")
	//os.Obj[0].Name[0] = append(os.Obj[0].Name[0], 4) // work
	for _, v := range os.Obj {		// not work
		for _, n := range v.Name {
			n = append(n, 4)
		}
	}
	fmt.Println(os)
}
