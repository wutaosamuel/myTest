package strategy

import (
	"fmt"
)

type Strategy interface {
	Begin()
}

// KeepStrategy
type KeepStrategy struct {
}

// Begin
func (ks *KeepStrategy) Begin() {
	fmt.Println("KeepStrategy Begin")
}
