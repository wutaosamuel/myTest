package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	socketFile := "c:/d/tmp/go.socket"

	// listen
	l, err := net.Listen("unix", socketFile)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer l.Close()

	// remove the socket file
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(socketFile)
		os.Exit(1)
	}()

	// handler
	for {
		conn, err := l.Accept()
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		go echoServer(conn)
	}
}

func echoServer(c net.Conn) {
	defer c.Close()

	for {
		buf := make([]byte, 1024)

		bufLen, err := c.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// receive message
		data := buf[:bufLen]
		fmt.Println("receive incoming data: ")
		fmt.Println(string(data))
		fmt.Println()

		// send message
		_, err = c.Write([]byte("receive ok"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
