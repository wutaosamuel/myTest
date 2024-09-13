package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	c, err := net.Dial("unix", "/c/d/tmp/go.socket")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	defer c.Close()

	for {
		sendWithResponse(c)
		time.Sleep(5 * time.Second)
	}
}

func sendWithResponse(c net.Conn) {
	// send message
	_, err := c.Write([]byte("Hello from client!"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// receive message
	buf := make([]byte, 1024)
	bufLen, err := c.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := buf[0:bufLen]
	fmt.Println("receive response:")
	fmt.Println(data)
	fmt.Println()
}
