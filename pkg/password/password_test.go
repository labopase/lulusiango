package password

import (
	"errors"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashAndVerifyBcrypt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		password  string
		cost      int
		expectErr bool
	}{
		{
			name:      "Happy Path - Default Cost",
			password:  "mysecurepassword123",
			cost:      DefaultCost,
			expectErr: false,
		},
		{
			name:      "Happy Path - Custom Cost",
			password:  "anothersecurepwd",
			cost:      bcrypt.MinCost, // Use min cost to keep tests fast
			expectErr: false,
		},
		{
			name:      "Empty Password",
			password:  "",
			cost:      DefaultCost,
			expectErr: true,
		},
		{
			name:      "Very Long Password (72+ Bytes)",
			password:  "averylongpasswordthatgoesonandonoandislongerthanseventytwocharactersjusttotestbcryptlimitmitigation",
			cost:      bcrypt.MinCost,
			expectErr: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			hash, err := HashBcryptWithConfig(tc.password, BcryptConfig{Cost: tc.cost})
			if tc.expectErr {
				if err == nil {
					t.Fatalf("expected error for password %q, got nil", tc.password)
				}
				return
			}

			if err != nil {
				t.Fatalf("failed to hash password: %v", err)
			}

			if hash == "" {
				t.Fatal("expected non-empty hash string")
			}

			// Verify with correct password
			err = VerifyBcrypt(tc.password, hash)
			if err != nil {
				t.Fatalf("verification failed for correct password: %v", err)
			}

			// Verify with incorrect password
			err = VerifyBcrypt(tc.password+"wrong", hash)
			if err == nil {
				t.Fatal("expected verification failure for mismatched password, got nil")
			}
			if !errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				t.Errorf("expected error wrapping bcrypt.ErrMismatchedHashAndPassword, got %v", err)
			}
		})
	}
}
