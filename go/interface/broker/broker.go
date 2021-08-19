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
}

// Set
func (b *Binance) Set(strategy st.Strategy) {
	b.Strategy = strategy
}

// Begin
func (b *Binance) Begin() {
	b.Strategy.Begin()
	fmt.Println("binance begin")
}
