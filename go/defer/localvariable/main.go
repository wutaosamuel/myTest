package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("work, it's not order queue")
	t := map[string]string{"0": "this is 0", "1": "it's 1"}
	ch := make(chan error, len(t))
	for _, v := range t {
		go func(v string) {
			time.Sleep(3 * time.Second)
			ch <- errors.New(v)
		}(v)
	}
	go func(v string) {
		time.Sleep(5 * time.Second)
		ch <- errors.New(v)
	}("error not in list")
	for i := 0; i < len(t)+1; i++ {
		err := IsChannelError(ch)
		fmt.Println(i, " -> ", err)
	}
}

func IsChannelError(err chan error) error {
	if e, ok := <-err; ok {
		return e
	} else {
		return errors.New("unexpected error channel closed")
	}
}
