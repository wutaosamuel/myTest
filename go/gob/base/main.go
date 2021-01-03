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
	m := make(map[string]string, 3)
	m["1"] = "test1"
	m["gob"] = "gob"
	m["map"] = "is map"

	saveGob(gob, m)

	for k, v := range readGob(gob) {
		fmt.Println(k, v)
	}
}

// saveGob save as gob file
func saveGob(file string, data map[string]string) {
	f, _ := os.OpenFile(file, os.O_RDWR | os.O_CREATE, 0777)
	encode := gob.NewEncoder(f)
	if err := encode.Encode(data); err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	return
}

// readGob read gob file
func readGob(file string) map[string]string{
	var result map[string]string

	f, _ := os.Open(file)
	decode := gob.NewDecoder(f)
	if err := decode.Decode(&result); err != nil {
		fmt.Println(err)
	}
	os.Remove(file)

	return result
}
