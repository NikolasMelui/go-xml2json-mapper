package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Client ...
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// BaseURL ...
const BaseURL = "https://yandex.ru"

// NewClient ...
func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func main() {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/", BaseURL), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "text/html")
	req.Header.Set("Accept", "text/html")

	client := &http.Client{
		Timeout: time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			panic(errors.New(errRes.Message))
		}
		panic(fmt.Errorf("Unknown error, status code: %d", res.StatusCode))
	}

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
