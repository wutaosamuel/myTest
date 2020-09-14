package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Stats is struct
type Stats struct {
	Low         string `json:"low"`
	High        string `json:"high"`
	Open        string `json:"open"`
	Volume      string `json:"volume"`
	Last        string `json:"last"`
	Volume30Day string `json:"volume_30day"`
}

// Ticker is product_id info.
type Ticker struct {
	TradeID int       `json:"trade_id,number"`
	Price   string    `json:"price"`
	Size    string    `json:"size"`
	Time    time.Time `json:"time,string"`
	Bid     string    `json:"bid"`
	Ask     string    `json:"ask"`
	Volume  string    `json:"volume"`
}

// Error is error
type Error struct {
	Message string `json:"message"`
}

func main() {
	// var data []byte
	url := "https://api-public.sandbox.pro.coinbase.com/products/BTC-USD/ticker"
	body := bytes.NewReader(make([]byte, 0))
	client := &http.Client{}
	// stats := &Stats{}
	ticker := &Ticker{}

	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		res.Body.Close()
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Cannot open")
		coinbaseError := Error{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&coinbaseError); err != nil {
			panic(err)
		}
		fmt.Println(coinbaseError.Message)
	}

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(ticker); err != nil {
		fmt.Println("decode error")
		panic(err)
	}

	fmt.Println(ticker)
}
