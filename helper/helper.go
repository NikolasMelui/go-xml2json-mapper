package helper

import (
	"crypto/sha256"
	"fmt"
)

// AddHash ...
func AddHash(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
