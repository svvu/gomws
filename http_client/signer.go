package mwsHttpClient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func SignV2(stringToSign string, secret string) string {
	hmac256 = ComputeHmac256(stringToSign, secret)
	return base64.StdEncoding.EncodeToString(hmac256)
}

func computeHmac256(message string, secret string) []byte {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	return hash.Sum(nil)
}
