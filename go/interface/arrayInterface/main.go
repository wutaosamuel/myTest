package main

import (
	"./protocol"
	"fmt"
)

type Manager struct {
	Arr []protocol.Protocol
}

func NewManager() *Manager {
	return &Manager{
		Arr: make([]protocol.Protocol, 0),
	}
}

func (m *Manager) Push(prt protocol.Protocol) {
	m.Arr = append(m.Arr, prt)
}

func (m *Manager) Print() {
	for _, v := range m.Arr {
		v.Print()
	}
}

func main() {
	manager := NewManager()

	user := NewUser()
	user.Name = "User1"
	user.Age = 22

	root := NewRoot()
	root.Name = "Admin"
	root.ID = 0

	fmt.Println("direct push back to array --> work")
	manager.Push(user)
	manager.Push(root)
	manager.Print()

	fmt.Println()
	fmt.Println("function push back to array --> work")
	manager.Push(user.GetProtocol())
	manager.Push(root.GetProtocol())
	manager.Print()
}
