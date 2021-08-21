package tiger

import (
	"fmt"

	st "../../strategy"
)

// Tiger
type Tiger struct {
	Strategy st.Strategy
	Message  string
}

// NewTiger
func NewTiger() *Tiger {
	return &Tiger{
		Strategy: nil,
		Message:  "tiger begin",
	}
}

// Set
func (t *Tiger) Set(strategy st.Strategy) {
	t.Strategy = strategy
}

// Begin
func (t *Tiger) Begin() {
	fmt.Println(t.Message)
}
