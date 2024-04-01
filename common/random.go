package common

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

func generateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func GenSalt(length int) []byte {
	if length < 0 {
		length = 50
	}
	return generateRandomSalt(length)
}

type hashPassword struct{}

func NewHashPassword() *hashPassword {
	return &hashPassword{}
}

func (h *hashPassword) Hash(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func (h *hashPassword) ValidatePassword(hashPassword, currentPassword string, salt []byte) bool {
	var currentPasswordHash = h.Hash(currentPassword, salt)
	return hashPassword == currentPasswordHash
}
