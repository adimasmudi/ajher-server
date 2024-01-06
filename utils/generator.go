package utils

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"io"
	mathRand "math/rand"
	"regexp"
	"strings"
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

func CalculatePoint(sentence string) int {
	// Tokenize the sentence into words
	words := regexp.MustCompile(`\b\w+\b`).FindAllString(strings.ToLower(sentence), -1)

	// Calculate length
	length := len(words)

	// Calculate uniqueness
	uniqueWords := make(map[string]struct{})
	for _, word := range words {
		uniqueWords[word] = struct{}{}
	}
	uniqueness := float64(len(uniqueWords)) / float64(length)

	// Assign points based on length and uniqueness
	var lengthPoints, uniquenessPoints int
	switch {
	case length < 5:
		lengthPoints = 1
	case length < 10:
		lengthPoints = 2
	default:
		lengthPoints = 3
	}

	switch {
	case uniqueness < 0.5:
		uniquenessPoints = 1
	case uniqueness < 0.8:
		uniquenessPoints = 2
	default:
		uniquenessPoints = 3
	}

	// Total points
	totalPoints := lengthPoints + uniquenessPoints

	return totalPoints
}
