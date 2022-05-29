package main

import (
	"fmt"

	"./protocol"
)

type Root struct {
	Name string
	ID   int
}

func NewRoot() *Root {
	return &Root{
		Name: "New User",
		ID:   -1,
	}
}

func (r *Root) Print() {
	fmt.Printf("This a root: name: %s, id: %d\n", r.Name, r.ID)
}

func (r *Root) GetProtocol() protocol.Protocol {
	return &Root{
		Name: r.Name,
		ID:   r.ID,
	}
}
