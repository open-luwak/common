package common

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateTraceID generates a 128-bit Trace ID
func GenerateTraceID() string {
	randomBytes, _ := generateRandomBytes(16)
	return hex.EncodeToString(randomBytes)
}

// GenerateSpanID generates a 64-bit Span ID
func GenerateSpanID() string {
	randomBytes, _ := generateRandomBytes(8)
	return hex.EncodeToString(randomBytes)
}

// generates a random byte slice of the specified length
func generateRandomBytes(length int) (randomBytes []byte, err error) {
	if length < 1 {
		return nil, fmt.Errorf("length must be at least 1")
	}
	if length > 1024 {
		return nil, fmt.Errorf("length must not exceed 1024")
	}
	randomBytes = make([]byte, length)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("crypto/rand.Read() panic: %v", r)
		}
	}()
	_, err = rand.Read(randomBytes)
	return randomBytes, err
}
