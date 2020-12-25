package main

import (
	"testing"
)

// TestIsExist test IsExist
func TestIsExist(t *testing.T) {
	fileName := "utils.go"
	isExist, err := IsExist(fileName)
	if err != nil {
		t.Fatal(err)
	}
	if isExist {
		t.Log(fileName, " exist")
	}
	if !isExist {
		t.Log(fileName, " not exist")
	}
}