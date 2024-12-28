
package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var secretKey = []byte("your-secret-key-32chars")

// Encrypt encrypts a string using AES
func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	stream := cipher.NewGCMWithNonceSize(block, 12)
	ciphertext := stream.Seal(nil, nonce, []byte(text), nil)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts a string using AES
func Decrypt(encrypted string) (string, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	data, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	nonce := data[:12]
	stream := cipher.NewGCMWithNonceSize(block, 12)
	plaintext, err := stream.Open(nil, nonce, data[12:], nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
