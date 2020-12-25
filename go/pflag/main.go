package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	var (
		addFlag = pflag.BoolP("add", "a", false, "add func")
	  intFlag = pflag.IntSliceP("ints", "i", make([]int, 0), "ints")
		stringsFlag []string
	)

	pflag.StringSliceVarP(&stringsFlag, "strings", "s", make([]string, 0), "strings test")
	
	pflag.Parse()

	fmt.Print("add: ")
	fmt.Println(*addFlag)
	fmt.Print("strings slice: ")
	fmt.Println(stringsFlag)
	fmt.Print("ints: ")
	fmt.Println(*intFlag)
}
