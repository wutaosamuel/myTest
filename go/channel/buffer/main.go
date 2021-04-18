package main

import (
	"fmt"
)

func main() {
	f1()
	fmt.Println()
	f2()
}

// f1 function 1
func f1() {
	fmt.Println("Basic buffer channel: ")
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		close(ch)
	}()
	fmt.Println("buffer channel int",
		<-ch,
		<-ch,
		<-ch,
		<-ch,
		<-ch)
}

// f2 function 2
func f2() {
	fmt.Println("Try to block buffer channel: not work")
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		close(ch)
	}()
	fmt.Println("buffer channel int block",
		<-ch,
		<-ch)
}
