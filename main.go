package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net"
	"os"
)

func main() {
	sendMessage("0.0.0.0:8080", "Hello world")
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

func sendMessage(recipientAddress, message string) {
	conn, _ := net.Dial("tcp", recipientAddress)

	conn.Write([]byte(message))
}
