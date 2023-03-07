package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateSignature(key string, dataSignature []byte) string {
	mac := hmac.New(sha256.New, []byte(key))
	// Cannot return error
	if _, err := mac.Write(dataSignature); err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
