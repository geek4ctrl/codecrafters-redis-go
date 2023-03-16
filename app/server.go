package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")

	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting connection: ", err.Error())
				os.Exit(1)
			}

			go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	var sb strings.Builder
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Failed to handle the connection", err.Error())
		os.Exit(1)
	}

	fmt.Println("Show me the conn buffer: ", sb[buffer[:n]]);

	conn.Write([]byte("+PONG\r\n"))
	conn.Close()
}
