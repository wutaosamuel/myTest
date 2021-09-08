package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	filename := filepath.Dir("./")
	fmt.Println(filename)
	abs, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(abs)
}