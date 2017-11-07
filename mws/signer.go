package mws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// SignV2 sign the message with the key by using hmac256.
func SignV2(stringToSign string, secret string) string {
	hmac256 := computeHmac256(stringToSign, secret)
	return base64.StdEncoding.EncodeToString(hmac256)
}

func computeHmac256(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}
