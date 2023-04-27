package main

import (
	"strings"
	"testing"
)

func TestGenerateGPGKey(t *testing.T) {
	name := "John Doe"
	email := "johndoe@example.com"
	passphrase := "my-secret-passphrase"

	// Generate key pair
	keyPair, err := generateGPGKey(name, email, passphrase)
	if err != nil {
		t.Errorf("Failed to generate key pair: %v", err)
	}

	// Verify private key
	if !strings.HasPrefix(keyPair.Private, "-----BEGIN PGP PRIVATE KEY BLOCK-----") {
		t.Errorf("Private key should start with '-----BEGIN PGP PRIVATE KEY BLOCK-----'")
	}
	if !strings.HasSuffix(keyPair.Private, "-----END PGP PRIVATE KEY BLOCK-----") {
		t.Errorf("Private key should end with '-----END PGP PRIVATE KEY BLOCK-----'")
	}

	// Verify public key
	if !strings.HasPrefix(keyPair.Public, "-----BEGIN PGP PUBLIC KEY BLOCK-----") {
		t.Errorf("Public key should start with '-----BEGIN PGP PUBLIC KEY BLOCK-----'")
	}
	if !strings.HasSuffix(keyPair.Public, "-----END PGP PUBLIC KEY BLOCK-----") {
		t.Errorf("Public key should end with '-----END PGP PUBLIC KEY BLOCK-----'")
	}
}
