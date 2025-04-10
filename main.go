package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	fmt.Println(generateKeys())
}

func generateKeys() (*rsa.PrivateKey, error) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	privateFile, _ := os.Create("private.pem")

	pem.Encode(privateFile, &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	return privateKey, nil
}
