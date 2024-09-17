package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	c, err := net.Dial("unix", "C:/d/tmp/go.socket")
	if err != nil {
		println(err.Error())
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
		fmt.Println(err.Error())
		return
	}

	// receive message
	buf := make([]byte, 1024)
	bufLen, err := c.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data := buf[:bufLen]
	fmt.Println("receive response:")
	fmt.Println(string(data))
	fmt.Println()
}
