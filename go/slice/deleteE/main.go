package main

import "fmt"

func main() {
	//s := []int{0,1,2,3,4,5,6,7,8,9}
	//i := 4
	s := []int{0,1,2}
	i := 1
	result := append(s[:i], s[len(s)-1-i:]...)
	fmt.Println("len: ", len(s))
	fmt.Println(result)
}
