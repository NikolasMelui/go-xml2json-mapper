package cache

import "github.com/nikolasmelui/go-xml2json-mapper/entity"

// ProductWithHash ...
type ProductWithHash struct {
	Data entity.Product
	Hash string
}

// ProductCache ...
type ProductCache interface {
	Set(key string, value *entity.Product)
	Get(key string) *ProductWithHash
}
