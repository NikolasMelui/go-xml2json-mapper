package cache

import "github.com/nikolasmelui/go-xml2json-mapper/products"

// ProductCache ...
type ProductCache struct {
	Data products.Product
	Hash string
}

// ProductsCache ...
type ProductsCache interface {
	Set(key string, value *ProductCache)
	Get(key string) *ProductCache
}
