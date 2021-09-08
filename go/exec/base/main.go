package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/pwd")
	//err := cmd.Run()
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("the result is %s\n", out)
}
