package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"../marketdata"
	"../utils"
)

func main() {
	url := utils.API_URL + marketdata.TickerURL
	body := bytes.NewReader(make([]byte, 0))
	client := &http.Client{}
	ticker := marketdata.NewTicker()

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
		binanceError := marketdata.NewError()
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(binanceError); err != nil {
			panic(err)
		}
		fmt.Println(binanceError)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(ticker); err != nil {
		fmt.Println("decode error")
		panic(err)
	}

	fmt.Println(ticker)
}
