package main

import (
	"fmt"

	"./utils"
)

func main() {
	utils.GlobalSet()
	fmt.Println(utils.Auth)
}