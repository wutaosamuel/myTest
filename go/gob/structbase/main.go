package main

import (
	"fmt"
	"os"
	"path"
	"encoding/gob"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println("Root dir: ", dir)
	gob := path.Join(dir, "test.gob")
	fmt.Println("Gob file: ", gob)

	o1 := &Obj{"o1"}
	o2 := &Obj{"o2"}
	o3 := &Obj{"o3"}

	o := []*Obj{o1, o2, o3}

	object := &Object{}
	object.List = o
	fmt.Println("input: ")
	for k, v := range object.List {
		fmt.Println(k, v)
	}
	fmt.Println()
	fmt.Println("Save")
	saveGob(gob, object)
	fmt.Println("Read: ")
	r := readGob(gob) 
	for k, v := range r.List {
		fmt.Println(k, v)
	}
}

// Obj object
type Obj struct {
	Name string
}

// Object object
type Object struct {
	List []*Obj
}

// saveGob save as gob file
func saveGob(file string, data *Object) {
	f, _ := os.OpenFile(file, os.O_RDWR | os.O_CREATE, 0777)
	encode := gob.NewEncoder(f)
	if err := encode.Encode(*data); err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	return
}

// readGob read gob file
func readGob(file string) *Object {
	var result Object

	f, _ := os.Open(file)
	decode := gob.NewDecoder(f)
	if err := decode.Decode(&result); err != nil {
		fmt.Println(err)
	}
	os.Remove(file)

	return &result
}
