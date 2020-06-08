package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/nikolasmelui/go-xml2json-mapper/cconfig"
	"github.com/nikolasmelui/go-xml2json-mapper/helper"
	"github.com/nikolasmelui/go-xml2json-mapper/products"
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

// NewClient ...
func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func main() {
	req, err := http.NewRequest("GET", products.ProductsURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Accept", "application/xml")
	req.SetBasicAuth(cconfig.Config.BasicAuthLogin, cconfig.Config.BasicAuthPassword)

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
	var data products.Products
	err = xml.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error: %v", err)
	}

	for _, product := range data.Products {
		hash := helper.AddHash(product)
		fmt.Printf("%s\n", hash)
		// product.BeautyPrint()
	}
}
