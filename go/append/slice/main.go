package main

import "fmt"

// Obj Obj object
type Obj struct {
	Name string
	Content []string
}

// Object object object
type Object struct {
	Content []Obj
}

func main() {
	o := Obj{"name", []string{"a", "b"}}
	fmt.Println("Start -> ", o)
	o.Content = append(o.Content, "c")
	fmt.Println("append -> ", o)
	fmt.Println("work")

	object := Object{
		[]Obj {
			Obj{"1", []string{"a", "b"}},
			Obj{"2", []string{"c", "d"}},
		},
	}
	fmt.Println("Start -> ", object)
	for _, v := range object.Content {
		v.Content = append(v.Content, "z")
	}
	fmt.Println("append -> ", object)
	fmt.Println("not work")
}