package client

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func GenerateRandomEmail() string {
	return "test-" + GenerateRandomString(10) + "@example.com"
}
