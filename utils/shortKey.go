package utils

import (
	"math/rand"
)

func GenerateShortKey() string {

	const (
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		keyLen  = 5
	)

	shortKey := make([]byte, keyLen)

	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortKey)
}
