package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hashing(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	return hashedPassword
}
