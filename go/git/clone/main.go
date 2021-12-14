package main

import (
	"fmt"
)

func main() {
	path := "/dev/shm"
	gitURL := "https://github.com/vortexgpgpu/vortex.git"
	proxy := "http://172.16.8.168:8123"

	fmt.Println("start to test proxy git clone")
	ProxyClone(path, gitURL, proxy)
	fmt.Println("done")
}