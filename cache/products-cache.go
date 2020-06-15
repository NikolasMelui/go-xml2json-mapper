package cache

import "github.com/nikolasmelui/go-xml2json-mapper/entity"

// EcommerceData ...
// type EcommerceData interface {
// }

// ProductWithHash ...
type ProductWithHash struct {
	Data entity.Product
	Hash string
}

// ProductCache ...
type ProductCache interface {
	Set(key string, value *ProductWithHash)
	Get(key string) *ProductWithHash
}
