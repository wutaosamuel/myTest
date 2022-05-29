package main

import (
	"fmt"
	"./protocol"
)

func main() {
	user := NewUser();
	user.Name = "User1"
	user.Age = 22
	protocol.ProtocolPrint(user)

	root := NewRoot();
	root.Name = "RootAdmin"
	root.ID = 0
	protocol.ProtocolPrint(root)

	fmt.Println()
	fmt.Println("private func print")
	print(user)
	print(root)
}

func print(ptl protocol.Protocol) {
	ptl.Print()
}