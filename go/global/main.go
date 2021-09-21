package main

import (
	"fmt"

	"./utils"
)

func main() {
	utils.GlobalSet()
	fmt.Println(utils.Auth)

	testStruct()
}

func testStruct() {
	fmt.Println("get problem @ struct")
	o1 := utils.NewO["object"]
	obj1, ok := o1.(*utils.Object)
	if !ok {
		fmt.Println("not object struct")
	}
	obj1.Name = "name"
	obj1.ID = 1
	fmt.Println("obj1", obj1)

	o2 := utils.NewO["object"]
	obj2, ok := o2.(*utils.Object)
	if !ok {
		fmt.Println("not object struct")
	}
	obj2.Name = "n"
	obj2.ID = 999
	fmt.Println()
	fmt.Println("obj2", obj2)
	fmt.Println("obj1", obj1)
}