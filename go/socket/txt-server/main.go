package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("unix", "/c/d/tmp/go.socket")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			println(err)
			os.Exit(1)
		}

		go echoServer(conn)
	}
}

func echoServer(c net.Conn) {
	buf := make([]byte, 1024)

	bufLen, err := c.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// receive message
	data := buf[0:bufLen]
	fmt.Println("receive incoming data: ")
	fmt.Println(string(data))
	fmt.Println()

	// send message
	_, err = c.Write([]byte("receive ok"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
