package signer

import (
	"crypto/hmac"
	"crypto/sha256"
)

var Hash = sha256.New

const HashSize = sha256.Size

func Sign(message, key []uint8) []uint8 {
	h := hmac.New(Hash, key)
	h.Write(message)
	return h.Sum(nil)
}

func Verify(message, key, expectedSignature []uint8) bool {
	actualSignature := Sign(message, key)
	return hmac.Equal(actualSignature, expectedSignature)
}
