package cryptography

import (
	"strings"
	"testing"
)

func TestSum224(t *testing.T) {
	msg := "ABCD"
	res := CalculateChecksum([]byte(msg))

	if !strings.HasSuffix(res, "60b6") {
		t.Error("Expected input to generate another checksum...")
	}
}
