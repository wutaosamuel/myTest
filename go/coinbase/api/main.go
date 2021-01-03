package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

// SpotPrice is product_id info.
type SpotPrice struct {
	SpotData	Data	`json:"data"`
}

// Data is data
type Data struct {
	Base	string `json:"base"`
	Currency string `json:"currency"`
	Price string `json:"amount"`
}

// Error is error
type Error struct {
	Message string `json:"message"`
}

func main() {
	// var data []byte
	url := "https://api.coinbase.com/v2/prices/BTC-USD/spot"
	body := bytes.NewReader(make([]byte, 0))
	client := &http.Client{}
	// stats := &Stats{}
	ticker := &SpotPrice{}

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
