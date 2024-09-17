package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	socketFile := "c:/d/tmp/http.socket"

	// local socket listen
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

	http.HandleFunc("/hello", helloHandler)
	http.Serve(l, nil)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("receive request")
	fmt.Fprintf(w, "Hello world!\n")
}
