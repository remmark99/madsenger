package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	command := os.Args[1]

	switch command {
	case "send":
		recepient := os.Args[2]
		message := strings.Join(os.Args[3:]," ")
		sendMessage(recepient, message)
	}
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

	fmt.Println("Message sent!")
}
