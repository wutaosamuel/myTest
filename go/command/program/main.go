package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("git", "--version")
	cmd2 := exec.Command("gi", "-v")
	if err := cmd1.Run(); err != nil {
		panic(err)
	}
	fmt.Println("cmd1 work")
	if err := cmd2.Run(); err != nil {
		panic(err)
	}
	fmt.Println("cmd2 work")
}
