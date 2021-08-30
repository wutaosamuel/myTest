package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- 10
	}()
	if i, ok := <- ch; ok {
		fmt.Println(i)
	}
}