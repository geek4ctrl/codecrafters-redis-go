package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
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
	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Failed to handle the connection", err.Error())
			os.Exit(1)
		}

		fmt.Println("Show me the buffer 1: ", string(buffer));
		fmt.Println("Show me the buffer 2: ", string(buffer)[2]);
		fmt.Println("Show me the buffer 3: ", string(buffer[2]));

		arrayOfElement := string.Split(string(buffer), "\r\n");
		firstArrayOfElement, restOfArrayOfElement := arrayOfElement[0][1:], arrayOfElement[1:];

		fmt.Println("Show me the buffer 4: ", arrayOfElement);
		fmt.Println("Show me the buffer 5: ", firstArrayOfElement);
		fmt.Println("Show me the buffer 6: ", restOfArrayOfElement);

		fmt.Println(strings.Split(string(buffer), "\r\n"));

		switch string(buffer) {
		case "ping":
			conn.Write([]byte("+PONG\r\n"))
		}

		conn.Write([]byte("+PONG\r\n"))
	}
}

