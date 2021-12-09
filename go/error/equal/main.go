package main

import (
	"errors"
	"fmt"

	gogit "github.com/go-git/go-git"
)

func main() {
	s1 := "first error"
	e1 := errors.New(s1)

	fmt.Println("equal --> pass")
	fmt.Println()
	if e1 != errors.New(s1) {
		fmt.Println("errors.New() not equal")
	}
	if e1.Error() == s1 {
		fmt.Println("err.Error() can be equal")
	}

	if _, err := gogit.PlainOpen("./name"); err != nil {
		if err == gogit.ErrRepositoryNotExists {
			fmt.Println("var ErrNotExist can be equal")
		}
	}
}