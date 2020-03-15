package main

import (
	"os"
	"time"
	"fmt"
)

func main() {
	f, err := os.Stat("./dir")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(f.ModTime())
	fmt.Println(time.Now())
}