package tools

import "testing"

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey("test", 32)
	if err != nil {
		t.Error(err)
	}

	if len(key) != 32 {
		t.Error("Key length is not 32")
	}

}

func TestEncryptDecrypt(t *testing.T) {
	testData := "This is a test message."
	keyLength := 16 // Example key length, adjust according to your requirements
	key, err := GenerateKey(testData, keyLength)
	if err != nil {
		t.Errorf("Failed to generate key: %v", err)
	}

	ciphertext, err := Encrypt(testData, key)
	if err != nil {
		t.Errorf("Encryption failed: %v", err)
	}

	decryptedData, err := Decrypt(ciphertext, key)
	if err != nil {
		t.Errorf("Decryption failed: %v", err)
	}

	if decryptedData != testData {
		t.Errorf("Decrypted data does not match original data. Expected: %s, Got: %s", testData, decryptedData)
	}
}
