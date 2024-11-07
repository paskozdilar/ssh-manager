package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type Key struct {
	pvt    ssh.Signer
	pub    ssh.PublicKey
	pvtPEM []byte
	pubPEM []byte
}

func NewKey() (*Key, error) {
	// Generate keys
	key, err := rsa.GenerateKey(rand.Reader, 3072)
	if err != nil {
		return nil, fmt.Errorf("generate RSA private key: %w", err)
	}

	pvt, err := ssh.NewSignerFromKey(key)
	if err != nil {
		return nil, fmt.Errorf("SSH signer from key: %w", err)
	}

	pub, err := ssh.NewPublicKey(&key.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("SSH pubkey from RSA pubkey: %w", err)
	}

	// Serialize to PEM
	pvtPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	pubPEM := ssh.MarshalAuthorizedKey(pub)

	return &Key{
		pvt:    pvt,
		pub:    pub,
		pvtPEM: pvtPEM,
		pubPEM: pubPEM,
	}, nil
}
