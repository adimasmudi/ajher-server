package utils

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"io"
	mathRand "math/rand"
	"time"

	"github.com/google/uuid"
)

func GeneratedUUID() (string, error) {
	// Generate a new UUID
	uuidObj := uuid.New()

	// Convert UUID to byte slice
	uuidBytes, err := uuidObj.MarshalBinary()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(uuidBytes), nil
}
func GenerateRandomString(length int) string {
	// Characters to choose from
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator
	mathRand.Seed(time.Now().UnixNano())

	// Build the random string
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[mathRand.Intn(len(charset))]
	}

	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(cryptoRand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
