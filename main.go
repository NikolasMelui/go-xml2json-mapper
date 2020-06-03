package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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

// BasicAuth ...
type BasicAuth struct {
	Login    string
	Password string
}

// BaseURL ...
const BaseURL = "https://your.xml.server.com"

var basicAuth = &BasicAuth{
	Login:    "login",
	Password: "password",
}

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
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/html")
	req.Header.Set("Accept", "text/html")
	req.SetBasicAuth(basicAuth.Login, basicAuth.Password)

	client := &http.Client{
		Timeout: time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			log.Fatal(errors.New(errRes.Message))
		}
		log.Fatal(fmt.Errorf("Unknown error, status code: %d", res.StatusCode))
	}

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(body)
}
