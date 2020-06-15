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

	"github.com/nikolasmelui/go-xml2json-mapper/cache"
	"github.com/nikolasmelui/go-xml2json-mapper/cconfig"
	"github.com/nikolasmelui/go-xml2json-mapper/entity"
	"github.com/nikolasmelui/go-xml2json-mapper/helper"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	req, err := http.NewRequest("GET", entity.ProductsURL, nil)
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
	var productsResponse entity.ProductsResponse
	err = xml.Unmarshal(body, &productsResponse)
	if err != nil {
		log.Printf("error: %v", err)
	}

	var productsCache cache.ProductCache = cache.NewRedisCache(cconfig.Config.RedisHost, cconfig.Config.RedisPassword, cconfig.Config.RedisDB, 100000)

	products := productsResponse.Products
	for i, product := range products {
		fmt.Printf("%d ----------\n", i)

		productCache := productsCache.Get(product.ID)
		// fmt.Println("Old - \n", productCache, "\n", "New - \n", &product)
		// fmt.Println("Old - \n", productCache.Hash, "\n", "New - \n", hash)

		hash := helper.InstanceHash(&product)
		productWithHash := &cache.ProductWithHash{
			Data: product,
			Hash: hash,
		}

		if productCache == nil {
			fmt.Printf("Product %s with index %d not found in cache, inserting...\n+++\n", product.ID, i)
			productsCache.Set(product.ID, productWithHash)
		} else {
			// fmt.Printf("Old - %s, New - %s", productCache.Hash, hash)
			if productCache.Hash == hash {
				fmt.Printf("Product %s with index %d has not changed (same hash)\n", product.ID, i)
				fmt.Printf("%+v\n", productCache)
			} else {

				fmt.Printf("Product %s with index %d has been changed (new hash), upding...\n+++\n", product.ID, i)
				productsCache.Set(product.ID, productWithHash)
			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}
