package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

// HashPassword hashes a plaintext password with bcrypt.
func HashPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcryptCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt hash failed: %w", err)
	}
	return string(hash), nil
}

// VerifyPassword compares a plaintext password against a bcrypt hash.
func VerifyPassword(plain, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}

// HashTokenSHA256 computes a deterministic hex SHA-256 digest for JWT revocation lookup.
func HashTokenSHA256(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

// EncryptPII encrypts sensitive fields (NIK, Phone) using AES-256-GCM with a 32-byte hex key.
func EncryptPII(plaintext, hexKey string) (string, error) {
	if plaintext == "" {
		return "", nil
	}
	key, err := hex.DecodeString(hexKey)
	if err != nil || len(key) != 32 {
		return "", errors.New("invalid AES-256 key (must be 64 hex chars)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptPII decrypts an AES-256-GCM base64 ciphertext back to plaintext.
func DecryptPII(encoded, hexKey string) (string, error) {
	if encoded == "" {
		return "", nil
	}
	key, err := hex.DecodeString(hexKey)
	if err != nil || len(key) != 32 {
		return "", errors.New("invalid AES-256 key")
	}

	raw, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(raw) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := raw[:nonceSize], raw[nonceSize:]
	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}
