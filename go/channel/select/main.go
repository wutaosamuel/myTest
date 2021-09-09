package main

import (
	"fmt"
	"time"
)

func main() {
	MulCh()
}

func MulCh() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go func() {
		p(ch1)
	}()
	go func() {
		time.Sleep(5 * time.Second)
		p(ch2)
	}()
	go func() {
		time.Sleep(10 * time.Second)
		p(ch3)
	}()

	fmt.Println("Start")
	select {
	case c1 := <-ch1:
		if c1 {
			fmt.Println("ch1 working")
		} else {
			fmt.Println("ch1 not working")
		}
	case c2 := <-ch2:
		if c2 {
			fmt.Println("ch2 working")
		} else {
			fmt.Println("ch2 not working")
		}
	case c3 := <-ch3:
		if c3 {
			fmt.Println("ch3 working")
		} else {
			fmt.Println("ch3 not working")
		}
	}
	fmt.Println("Done")
}

func p(ch chan bool) {
	time.Sleep(5 * time.Second)
	ch <- true
}
