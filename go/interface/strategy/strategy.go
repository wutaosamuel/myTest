package strategy

import (
	"fmt"
)

type Strategy interface {
	Begin()
}

// KeepStrategy
type KeepStrategy struct {
	Message string
}

// NewKeepStrategy
func NewKeepStrategy() *KeepStrategy {
	return &KeepStrategy{
		Message: "KeepStrategy Begin",
	}
}

// Begin
func (ks *KeepStrategy) Begin() {
	fmt.Println(ks.Message)
}
