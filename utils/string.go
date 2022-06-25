package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func IsStringUUID(str string) bool {
	_, err := uuid.Parse(str)
	if err != nil {
		return false
	}
	return true
}

// generate random character
var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randSeq(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// end of generate random character

func GenerateRandomChars(length int) string {
	rand.Seed(time.Now().UnixNano())
	return randSeq(length)
}
