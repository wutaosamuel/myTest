package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("End of main")
	fmt.Println("Start of main")
	panic("Panic stop")
}