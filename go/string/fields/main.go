package main

import (
	"fmt"
	"strings"
)

func main() {
	command := "echo $GOPATH"
	fmt.Print("TrimSpace: ")
	fmt.Println(strings.TrimSpace(command))
	fmt.Print("Fields: ")
	fmt.Println(strings.Fields(strings.TrimSpace(command)))
}