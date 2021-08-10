package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// main
func main() {
	err := errors.New("test error to string")
	fmt.Println(err.Error())
	fmt.Println("test error format by string -> ")
	err = ToError()
	fmt.Println("the error", err)
	fmt.Println("to format", ToFormat(err))
}

// ToFormat
func ToFormat(err error) []string {
	s := err.Error()
	result := make([]string, 0)
	elements := strings.Split(s, "&")
	for _, element := range elements {
		pair := strings.Split(element, "=")
		fmt.Println(element, "len:", len(pair))
		if pair[0] == "Code" {
			result = append(result, "Code")
			_, e := strconv.Atoi(pair[1])
			if e != nil {
				fmt.Println(e)
			}
			result = append(result, pair[1])
		}
		if pair[0] == "Message" {
			result = append(result, "Message")
			result = append(result, pair[1])
		}
	}
	return result
}

// ToError
func ToError() error {
	code := "Code"
	//codeValue := 1
	message := "Message"
	messageValue := "message"

	eString := code + "="
	eString += "&" + message + "=" + messageValue
	return errors.New(eString)
}
