package main

import (
	"fmt"
	"flag"
)

func main() {
	var intFlag int
	flag.IntVar(&intFlag, "i", 0, "i for int number")

	var strFlag = flag.String("str", "", "str for string")

	flag.Parse()
	fmt.Println("entered int number:", intFlag)
	fmt.Println("entered string:", *strFlag)
}