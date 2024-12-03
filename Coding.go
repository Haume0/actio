package actio

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

// Sha256 generates a sha256 hash from a string
func Sha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// GenerateCode generates a random 16 character code
func GenerateCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	code := string(b)
	encodedCode := url.QueryEscape(code)
	return encodedCode
}

// MapToXML converts a map to an XML string
func MapToXML(data map[string]string, h string) []byte {
	xmlString := fmt.Sprintf("<%v>", h)
	for k, v := range data {
		xmlString += fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	xmlString += fmt.Sprintf("</%v>", h)
	return []byte(xmlString)
}
