package helper

import (
	"crypto/sha256"
	"fmt"
)

// InstanceHash ...
func InstanceHash(instance interface{}) string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v", instance)))
	hashSum := hash.Sum(nil)
	return fmt.Sprintf("%x", hashSum)
}
