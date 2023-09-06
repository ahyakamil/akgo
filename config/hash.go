package config

import (
	"akgo/env"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/argon2"
)

func IsRSAPrivateKey(keyString string) bool {
	block, _ := pem.Decode([]byte(keyString))
	if block == nil {
		return false // Not a valid PEM-encoded key
	}

	parsedKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return err == nil && parsedKey != nil
}

func HashPassword(password string) (string, error) {
	timeCost := 1
	memoryCost := 10 // 10KB
	threads := 4
	keyLen := 32 // 256 bits

	salt := env.PasswordPrivateKey

	hash := argon2.IDKey([]byte(password), []byte(salt), uint32(timeCost), uint32(memoryCost), uint8(threads), uint32(keyLen))
	saltHashStr := fmt.Sprintf("%s", fmt.Sprintf("%x", hash))
	return saltHashStr, nil
}

func VerifyPassword(password, storedHash string) (bool, error) {
	salt := env.PasswordPrivateKey
	timeCost := 1
	memoryCost := 10 // 10KB
	threads := 4
	keyLen := 32 // 256 bits

	hash := argon2.IDKey([]byte(password), []byte(salt), uint32(timeCost), uint32(memoryCost), uint8(threads), uint32(keyLen))
	saltHashStr := fmt.Sprintf("%s", fmt.Sprintf("%x", hash))
	return saltHashStr == storedHash, nil
}
