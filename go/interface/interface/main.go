package main

import (
	"fmt"

	"../broker"
	"../strategy"
)

func main() {
	keep := strategy.NewKeepStrategy()
	binance := broker.NewBinance()
	binance.Set(keep)
	binance.Begin()

	fmt.Println()
	fmt.Println("test gob save and restore")
	if err := binance.SaveGob("save.gob"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("save done")
	restore := &broker.Binance{Strategy: &strategy.KeepStrategy{}}
	if err := restore.ReadGob("save.gob"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Restore")
	restore.Begin()
}
