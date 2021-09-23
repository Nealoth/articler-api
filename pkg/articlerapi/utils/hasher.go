package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

type Hasher struct {
	salt string
}

func NewHasher(salt string) *Hasher {
	return &Hasher{
		salt: salt,
	}
}

func (h *Hasher) Hash(str string) string {
	sha := sha512.New()
	sha.Write([]byte(h.salt + str))
	return hex.EncodeToString(sha.Sum(nil))
}
