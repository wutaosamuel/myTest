package marketdata

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"
)

// TestTickers
func TestTickers(t *testing.T) {
	// proxy
	httpProxy := "http://172.16.8.168:10089"
	proxy, err := url.Parse(httpProxy)
	if err != nil {
		t.Fatal(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}

	tickers := NewTickers()
	url := tickers.URL()
	body := bytes.NewReader(make([]byte, 0))
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 20,
	}

	t.Log("Tickers URL -> " + url)

	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		res.Body.Close()
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		binanceError := NewError()
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(binanceError); err != nil {
			t.Fatal(err)
		}
		t.Fatal(binanceError)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(tickers); err != nil {
		t.Fatal(err)
	}

	//t.Log(tickers)
	t.Log("Full tickers is too many, check length ->", len(*tickers))
}
