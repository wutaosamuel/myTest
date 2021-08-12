package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Request is http request for website
/*
 * - method -> allow 9 methods, e.g. "GET", "POST" and "HEAD"
 * - url 		-> target url for fetching
 * - params -> add more params(map[key]value) into req.Header
 * - result -> return result by json
 */
func Request(method, url string, params map[string]string,
	weberrors, result interface{}) (*http.Response, error) {
	body := bytes.NewBuffer(make([]byte, 0))
	client := &http.Client{}

	if err := checkMethod(method); err != nil {
		fmt.Println(err)
		return nil, Errs("Request Error: ", err)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return nil, Errs("New Request Error: ", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if params == nil || len(params) != 0 {
		for key, value := range params {
			req.Header.Add(key, value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		res.Body.Close()
		return res, Errs("Client Call Request Error: ", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 && weberrors != nil {
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(weberrors); err != nil {
			fmt.Println(err)
			return res, Errs("WebError Decode Error: ", err)
		}
		return res, nil
	}

	decoder := json.NewDecoder(res.Body)
	if result != nil {
		if err = decoder.Decode(result); err != nil {
			fmt.Println(err)
			return res, err
		}
	}

	return res, nil
}

// checkMethod check http method
// TODO: fit to http0.9, http1.0, http1.1 ...
func checkMethod(method string) error {
	if method == "GET" ||
		method == "HEAD" ||
		method == "POST" ||
		method == "PUT" ||
		method == "DELETE" ||
		method == "CONNECT" ||
		method == "OPTIONS" ||
		method == "TRACE" ||
		method == "PATCH" {
		return nil
	}
	return errors.New("http method not allow")
}
