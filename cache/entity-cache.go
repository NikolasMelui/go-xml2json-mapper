package cache

import (
	"crypto/sha256"
	"fmt"

	"github.com/nikolasmelui/go-xml2json-mapper/entity"
)

// Cachable ...
type Cachable interface {
	CreateHash()
}

// ProductCache ...
type ProductCache struct {
	Data entity.Product
	Hash string
}

// CreateHash ...
func (productCache *ProductCache) CreateHash() {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v", productCache)))
	productCache.Hash = fmt.Sprintf("%x", hash.Sum(nil))
}
