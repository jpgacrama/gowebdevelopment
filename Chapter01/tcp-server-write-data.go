package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	listener, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		logErrorAndCloseListener("Error starting TCP server: %v", err, listener)
	}
	defer closeListener(listener)

	log.Printf("Listening on %s:%s", connHost, connPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logErrorAndCloseListener("Error accepting connection: %v", err, listener)
		}
		go handleRequest(conn)
	}
}

func logErrorAndCloseListener(format string, err error, listener net.Listener) {
	log.Printf(format, err)
	closeListener(listener)
}

func closeListener(listener net.Listener) {
	if err := listener.Close(); err != nil {
		log.Printf("Error closing listener: %v", err)
	}
}

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Print("Message Received from the client: ", string(message))
	conn.Close()
}
