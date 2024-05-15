package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
)

// This function generates a key based on the data and key_length for encryption and decryption keys
func GenerateKey(data string, key_length int) (string, error) {
	// Hash the data
	hashedData := Hashing(data)

	// Ensure the hashed data is at least key_length long
	if len(hashedData) < key_length {
		return "", errors.New("hashed data is too short to generate a key of the desired length")
	}

	// Trim the hashed data to the desired key_length
	key := hashedData[:key_length]

	return key, nil
}

// This function takes a string as input and returns the SHA-256 hash of the string as a hexadecimal string
func Hashing(data string) string {
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

// Encrypt function encrypts data using AES encryption with the provided key
func Encrypt(data, key string) ([]byte, error) {
	// Create a new AES cipher block using the key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	// Create a new GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Generate a random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the data using AES-GCM
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)

	return ciphertext, nil
}

// Decrypt function decrypts data using AES encryption with the provided key
func Decrypt(ciphertext []byte, key string) (string, error) {
	// Create a new AES cipher block using the key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Create a new GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Get the nonce size
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext is too short")
	}

	// Split the nonce and ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data using AES-GCM
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
