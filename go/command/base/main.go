package main

import (
	"os"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	gitPath := "/home/pi/project/test"
	cmd := "git remote show all"
	notwin(gitPath, cmd)
}

// notwin not run on windows
// //+build !windows
func notwin(path, command string) {
	// replace ${var} or $var in shell
	commandToRun := os.ExpandEnv(command)
	// separate command to args in terms of spaces
	args := strings.Fields(strings.TrimSpace(commandToRun))
	fmt.Println(args)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = path
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

// windows is run on windows
// //+build windows
func windows(command string) {
	// replace ${var} or $var in shell
	commandToRun := os.ExpandEnv(command)
	// separate command to args in terms of spaces
	args := strings.Fields(strings.TrimSpace(commandToRun))
	cmd := exec.Command(args[0], args[1:]...)
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}