package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	secret := "secret"
	data := "test"
	fmt.Println("secret:", secret, "data:", data)
	encodeString := encode(secret, data)
	fmt.Println("encoding -> ", encodeString)
}

func encode(secret, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func decode() {

}