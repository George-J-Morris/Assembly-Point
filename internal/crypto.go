package internal

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func GenerateKeyPair(outPath string) error {
	// Server: generate an RSA keypair.
	sk, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Printf("failed to generate RSA key: %v", err)
	}

	publicKey, _ := x509.MarshalPKIXPublicKey(&sk.PublicKey)
	privateKey := x509.MarshalPKCS1PrivateKey(sk)

	pubBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKey,
	}

	privBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKey,
	}

	privOut, err := os.Create(outPath + "/" + "private.pem")
	if err != nil {
		log.Fatalf("Failed to open private.pem for writing: %v", err)
	}

	pubOut, err := os.Create(outPath + "/" + "public.pem")
	if err != nil {
		log.Fatalf("Failed to open public.pem for writing: %v", err)
	}

	pem.Encode(pubOut, pubBlock)
	pem.Encode(privOut, privBlock)

	return nil
}

func LoadRsaPublicKey(dir string) (any, error) {
	dat, err := os.ReadFile(dir)
	if err != nil {
		log.Printf("Failed to load public key: %v", err)
		return nil, err
	}

	decoded, _ := pem.Decode(dat)

	publicKey, err := x509.ParsePKIXPublicKey(decoded.Bytes)
	if err != nil {
		log.Printf("Failed to parse public key: %v", err)
		return nil, err
	}

	return publicKey, nil
}

func LoadRsaPrivateKey(dir string) (any, error) {
	dat, err := os.ReadFile(dir)
	if err != nil {
		log.Printf("Failed to load private key: %v", err)
		return nil, err
	}

	decoded, _ := pem.Decode(dat)

	publicKey, err := x509.ParsePKCS1PrivateKey(decoded.Bytes)
	if err != nil {
		log.Printf("Failed to parse private key: %v", err)
		return nil, err
	}

	return publicKey, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
