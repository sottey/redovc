package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateUniqueID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// Handle error
	}
	return fmt.Sprintf("%x", b)
}
