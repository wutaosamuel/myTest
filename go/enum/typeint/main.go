package main

import (
	"fmt"
)

type TestType int

const (
	Test1 TestType = iota
	Test2
	Test3
)

var TestTypeToString = map[TestType]string{
	Test1: "test1",
	Test2: "test2",
	Test3: "test3",
}

var ToTestType = []TestType{
	Test1,
	Test2,
	Test3,
}

func main() {
	fmt.Println("Test1", int(Test1))
	fmt.Println("Test2", ToTestType[Test2])
}
