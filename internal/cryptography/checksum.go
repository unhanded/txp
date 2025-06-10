package cryptography

import (
	"crypto/sha256"
	"fmt"
)

func CalculateChecksum(data []byte) string {
	sum := sha256.Sum224(data)

	return fmt.Sprintf("%x", sum)
}
