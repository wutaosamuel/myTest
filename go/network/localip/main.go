package main

import (
	"fmt"
	"net"
)

func main() {
	iFaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range iFaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Println("address: ", addr)
			fmt.Println("ip: ", ip)
		}
	}
}
