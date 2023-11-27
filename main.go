package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	plaintext := "HELLO"
	key := generateRandomKey(len(plaintext))
	ciphertext := encryptOTP(plaintext, key)
	decryptedText := decryptOTP(ciphertext, key)

	fmt.Println("Plaintext:", plaintext)
	fmt.Println("Key:", key)
	fmt.Println("Ciphertext:", ciphertext)
	fmt.Println("Decrypted Text:", decryptedText)
}

func generateRandomKey(length int) []byte {
	key := make([]byte, length)
	for i := 0; i < length; i++ {
		key[i] = byte(rand.Intn(256)) // Generating random byte values
	}
	return key
}

func encryptOTP(plaintext string, key []byte) []byte {
	if len(plaintext) != len(key) {
		panic("Plaintext and key must have the same length")
	}

	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key[i] // XOR operation
	}
	return ciphertext
}

func decryptOTP(ciphertext []byte, key []byte) string {
	if len(ciphertext) != len(key) {
		panic("Ciphertext and key must have the same length")
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i++ {
		plaintext[i] = ciphertext[i] ^ key[i] // XOR operation
	}
	return string(plaintext)
}
