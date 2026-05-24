package password

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultCost = bcrypt.DefaultCost
	MinCost     = bcrypt.MinCost
	MaxCost     = bcrypt.MaxCost
)

type BcryptConfig struct {
	Cost int
}

// HashBcrypt hashes the given password using bcrypt with the default cost.
// It returns the hashed password as a string or an error if hashing fails.
func HashBcrypt(password string) (string, error) {
	return HashBcryptWithConfig(password, BcryptConfig{Cost: DefaultCost})
}

// HashBcryptWithConfig hashes the given password using bcrypt with the specified configuration.
// It returns the hashed password as a string or an error if hashing fails.
func HashBcryptWithConfig(pwd string, config BcryptConfig) (string, error) {
	if pwd == "" {
		return "", errors.New("password manager: password cannot be empty")
	}

	cost := config.Cost
	if cost == 0 {
		cost = DefaultCost
	} else if cost < DefaultCost {
		cost = DefaultCost
	}

	if cost < MinCost || cost > MaxCost {
		return "", fmt.Errorf("password manager: invalid bcrypt cost %d (must be between %d and %d)", cost, MinCost, MaxCost)
	}

	// Pre-hash and encode to fit within bcrypt's 72-byte limit safely
	sum := sha256.Sum256([]byte(pwd))
	preHashed := base64.StdEncoding.EncodeToString(sum[:])

	hashed, err := bcrypt.GenerateFromPassword([]byte(preHashed), cost)
	if err != nil {
		return "", fmt.Errorf("password manager: failed to hash password: %w", err)
	}

	return string(hashed), nil
}

// VerifyBcrypt compares the given password with the provided bcrypt hash.
// It returns nil if the password matches the hash, or an error if it does not match or if verification fails.
func VerifyBcrypt(password, hash string) error {
	sum := sha256.Sum256([]byte(password))
	preHashed := base64.StdEncoding.EncodeToString(sum[:])

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(preHashed))
	if err != nil {
		return fmt.Errorf("password manager: failed to verify password: %w", err)
	}

	return nil
}
