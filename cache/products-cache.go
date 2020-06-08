package cache

import "github.com/nikolasmelui/go-xml2json-mapper/entity"

// ProductCache ...
type ProductCache struct {
	Data entity.Product
	Hash string
}

// ProductsCache ...
type ProductsCache interface {
	Set(key string, value *ProductCache)
	Get(key string) *ProductCache
}
