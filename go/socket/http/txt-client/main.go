package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	socketPath := "c:/d/tmp/http.socket"
	client := http.Client{
		Transport: &http.Transport{
			Dial: func(_, _ string) (net.Conn, error) {
				return net.Dial("unix", socketPath)
			},
		},
	}

	//res, err := client.Get("unix://http.socket")
	res, err := client.Get("http://http.socket/hello")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		print("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
		os.Exit(1)
	}
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	fmt.Println(body)
}
