package main

import (
	"fmt"
	gogit "github.com/go-git/go-git"
)

func main() {
	dir := "./tmp"
	_, err := gogit.PlainInit(dir, true)
	if err != nil {
		fmt.Println(err)
		return
	}
}