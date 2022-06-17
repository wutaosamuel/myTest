package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func runPWD() string {
	cmd := exec.Command("pwd")
	output, err := cmd.Output()
	if err != nil {
		return "pwd error"
	}
	return string(output)
}

func runLocal() string {
	filename := filepath.Dir("./")
	abs, err := filepath.Abs(filename)
	if err != nil {
		return "local error"
	}
	return abs
}

func main() {
	fmt.Println("They are the same, outputting path with executing location")
	fmt.Println("Execute pwd")
	pwd := runPWD()
	fmt.Println(pwd)

	fmt.Println("local path with ./")
	local := runLocal()
	fmt.Println(local)
}
