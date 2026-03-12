package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func hashClientIP(ip string) string {
	if strings.TrimSpace(ip) == "" {
		return ""
	}
	sum := sha256.Sum256([]byte(strings.TrimSpace(ip)))
	return hex.EncodeToString(sum[:])
}
