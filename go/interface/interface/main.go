package main

import (
	"../broker"
	"../strategy"
)

func main() {
	keep := &strategy.KeepStrategy{}
	binance := &broker.Binance{}
	binance.Set(keep)
	binance.Begin()
}
