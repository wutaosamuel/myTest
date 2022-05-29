package main

import (
	"fmt"

	"./protocol"
)

type User struct {
	Name string
	Age  int
}

func NewUser() *User {
	return &User{
		Name: "New User",
		Age:  -1,
	}
}

func (u *User) Print() {
	fmt.Printf("This a User: name: %s, age: %d\n", u.Name, u.Age)
}

func (u *User) GetProtocol() protocol.Protocol {
	return &User{
		Name: u.Name,
		Age:  u.Age,
	}
}
