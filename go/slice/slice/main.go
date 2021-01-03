package main

import "fmt"

func main() {
	s := "0123456789"
	fmt.Println("s[3:] -> ", s[3:])
	fmt.Println("s[:3] -> ", s[:3])
	fmt.Println("s[len(s)-3:] -> ", s[len(s)-3:])
}
