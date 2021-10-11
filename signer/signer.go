package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

// Signer -
type Signer struct{}

// New -
func New() *Signer {
	return &Signer{}
}

// Do -
func (s *Signer) Do(secret, url, data string) string {
	h := hmac.New(sha256.New, []byte(secret))

	message := fmt.Sprintf("%s?%s", url, data)

	log.Println("message: ", message)

	h.Write([]byte(message))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
