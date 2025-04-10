package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := listener.Accept()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	fmt.Printf("%s says: %s\n", getSender(conn), string(buf[:n]))
}

func getSender(conn net.Conn) (string) {
	return conn.RemoteAddr().(*net.IPAddr).IP.String()
}
