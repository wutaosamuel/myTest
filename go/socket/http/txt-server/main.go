package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	// local socket listen
	l, err := net.Listen("unix", "/c/d/tmp/http.socket")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	defer l.Close()

	http.HandleFunc("/hello", helloHandler)
	http.Serve(l, nil)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}
