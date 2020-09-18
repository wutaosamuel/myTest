package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/usr/bin/wget","-p", "google.com")
	cmd.Dir = "/home/pi/project/test/go/command/env/"
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env,
		"http_proxy=http://172.16.8.220:10089",
		"https_proxy=http://172.16.8.220:10089")

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}