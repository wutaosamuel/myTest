package main

import (
	"fmt"
)

func main() {
	for i := 0; i <3; i++ {
		defer fmt.Println("Defer")
		if i == 2 {
			break
		}
	}
}