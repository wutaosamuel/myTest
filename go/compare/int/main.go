package main

import "fmt"

func main() {
	fmt.Printf("0 > 0 -> %t\n", 0 > 0)
	fmt.Printf("0 > 1 -> %t\n", 0 > 1)
	fmt.Printf("58 < 59 -> %t\n", 58 < 59)
	fmt.Printf("59 < 59 -> %t\n", 59 < 59)
	fmt.Printf("60 < 59 -> %t\n", 60 < 59)
}