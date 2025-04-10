package main

import (
	"log"
	"madsenger/internal/handlers"
	"net"
)

func main() {
	go startListening(":8080", handlers.HandleClientConnection)
	go startListening(":8081", handlers.HandleGTUIConnection)
}

func startListening(port string, handler func(net.Conn)) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v\n", port, err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v\n", err)
		}

		go handler(conn)
	}
}
