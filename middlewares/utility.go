package middlewares

import (
	"crypto/sha256"
	"fmt"
)

func EncryptPassword(input string) string {
	sum := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", sum)
}
