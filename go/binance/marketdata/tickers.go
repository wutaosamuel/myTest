package marketdata

import (
	"../utils"
)

// TickerURL => https://api.binance.com/api/v3/ticker/price
// TickerURL => https://api.binance.com/api/v3/ticker/price?symbol=BTCBUSD
const TickersURL = "/api/v3/ticker/price"

// Tickers contain ticker
//type Tickers struct {
	//Contents []Ticker `json:""`
//}
type Tickers []struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// NewTickers
func NewTickers() *Tickers {
	return &Tickers{}
}

// TickersURL
func (ts *Tickers) URL() string {
	return utils.API_URL + TickersURL
}
