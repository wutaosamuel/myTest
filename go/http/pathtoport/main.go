package main

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

func sayHelloWorld(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Hello world from 1")
}

func sayHello(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Hello from 2")
}

// dial to server and reverse proxy
func Redirect(res http.ResponseWriter, req *http.Request) {
	target := req.Host
	if req.URL.Path == "/helloworld" {
		target += "/"
	}
}

func main() {
	r1 := mux.NewRouter()
	r1.HandleFunc("/", sayHelloWorld)
	go http.ListenAndServe(":9981", r1)

	r2 := mux.NewRouter()
	r2.HandleFunc("/", sayHello)
	go http.ListenAndServe(":9982", r2)


}