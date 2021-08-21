package broker

import (
	"fmt"

	st "../strategy"
)

// Broker
type Broker interface {
	Set(strategy st.Strategy)
	Begin()
}

// Binance
type Binance struct {
	Strategy st.Strategy
	Message string
}

// NewBinance
func NewBinance() *Binance {
	return &Binance{
		Strategy: nil,
		Message: "binance begin",
	}
}

// Set
func (b *Binance) Set(strategy st.Strategy) {
	b.Strategy = strategy
}

// Begin
func (b *Binance) Begin() {
	fmt.Println(b.Message)
	b.Strategy.Begin()
}
