package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

func UnHashing(password string, hashing string) bool {
	hash := sha256.Sum256([]byte(password))
	hashPWD := hex.EncodeToString(hash[:])

	return hashPWD == hashing
}
