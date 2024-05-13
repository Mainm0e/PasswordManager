package data

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashing function
// This function takes a string as input and returns the SHA-256 hash of the string as a hexadecimal string
func hashing(data string) string {
	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Write the data to the hasher
	hasher.Write([]byte(data))

	// Get the hashed bytes
	hashedBytes := hasher.Sum(nil)

	// Convert the hashed bytes to a hexadecimal string
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString
}
