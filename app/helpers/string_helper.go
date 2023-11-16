package helpers

import (
	"math/rand"
	"strings"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateRandomFileName(filename string) string {
	// Get type file
	splitFileName := strings.Split(filename, ".")
	lenghtOfSplit := len(splitFileName)

	typeFile := splitFileName[lenghtOfSplit - 1]

	// Random
	random := generateRandomString(24)

	name := random + "." + typeFile
	return name
}