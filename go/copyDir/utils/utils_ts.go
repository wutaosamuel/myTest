package utils

import (
	"os"
	"io/ioutil"
	"fmt"
)

// CopyDirTreeTs test copy dir tree
func CopyDirTreeTs() {
	fmt.Println("start to copy dir tree")
	fmt.Println("create test dir")
	if err := os.MkdirAll("./testA/testB", 0777); err != nil {
		fmt.Println("create test fail")
		return
	}
	defer os.RemoveAll("./testA")
	testA := []byte("under TestA dir")
	if err := ioutil.WriteFile("./testA/testA.txt", testA, 0777); err != nil {
		fmt.Println("create testA.txt fail")
		return
	}
	testB := []byte("under testA/testB dir")
	if err := ioutil.WriteFile("./testA/testB/testB.txt", testB, 0777); err != nil{
		fmt.Println("create testB.txt fail")
		return
	}

	fmt.Println("copy file now")
	if err := CopyDirTree("./testCopy", "./testA"); err != nil {
		fmt.Println("copy dir fail")
		return
	}
	fmt.Println("Done, remember delete ./testCopy")
}