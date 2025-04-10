package handlers

import (
	"fmt"
	"net"
)

func HandleClientConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	fmt.Printf("%s says: %s\n", getSender(conn), string(buf[:n]))
}

func getSender(conn net.Conn) (string) {
	return conn.RemoteAddr().(*net.IPAddr).IP.String()
}
