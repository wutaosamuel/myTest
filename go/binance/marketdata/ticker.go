package marketdata

import (
	"../utils"
)

// TickerURL => https://api.binance.com/api/v3/ticker/price
// TickerURL => https://api.binance.com/api/v3/ticker/price?symbol=BTCBUSD
const TickerURL = "/api/v3/ticker/price"

// Ticker contain symbol and price
type Ticker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// NewTicker
func NewTicker() *Ticker {
	return &Ticker{
		Symbol: "",
		Price:  "",
	}
}

// TickerURL
func (t *Ticker) URL(symbol string) string {
	return utils.API_URL + TickerURL + "?symbol=" + symbol
}
