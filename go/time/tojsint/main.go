package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	t := time.Now()
	fmt.Println(t.UTC().UnixNano())
	// to millisecond
	mt := t.UTC().UnixNano() / int64(time.Millisecond)
	fmt.Println(mt)
	// to string
	st := strconv.FormatInt(mt, 10)
	fmt.Println(st)
}