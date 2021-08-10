package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Obj object
type Obj struct {
	Name string
	Id   int
}

// Objects
type Objects struct {
	Name     string
	Elements []*Obj
}

// main
func main() {
	o1 := &Obj{"o1", 1}
	o2 := &Obj{"o2", 2}
	o3 := &Obj{"o3", 3}

	objects := &Objects{"objects!", []*Obj{o1, o2, o3}}
	fmt.Println("before gob: ")
	fmt.Println(objects)
	for _, v := range objects.Elements {
		fmt.Println(v)
	}
	fmt.Println("encoding ...")
	// b, err := GobEncode(&objects) --> both okay
	b, err := GobEncode(objects)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("decoding ...")
	result := &Objects{}
	// if err := GobDecode(&result, b); err != nil { // --> both okay
	if err := GobDecode(result, b); err != nil {
		fmt.Println(err)
	}
	fmt.Println("after gob: ")
	fmt.Println(result)
	for _, v := range result.Elements {
		fmt.Println(v)
	}
}

// GobEncode
func GobEncode(object interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(object); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return buffer.Bytes(), nil
}

// GobDecode
func GobDecode(object interface{}, b []byte) error {
	buffer := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buffer)
	if err := decoder.Decode(object); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
