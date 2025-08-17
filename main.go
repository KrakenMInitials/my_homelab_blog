package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // closes conncection when returned from handleConnection()

	// Setup a reader to process w/ buffer incoming connections
	reader := bufio.NewReader(conn)

	// Reader will be tcp oriented and therefore have segments etc
	// Because a single message may contain multiple commands, we manually parse with delims
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(conn, "Error reading command: %v\n", err)
		return
	}

	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	if len(parts) != 2 {
		fmt.Fprintf(conn, "Invalid command format. Expected format: COMMAND:RESOURCE\n")
		return
	}

	command := parts[0]
	resource := parts[1]
	log.Printf("Recieved command: %s %s \n", command, resource)

	switch command {
	case "GET":
		handleGet(conn, resource)
	default:
		fmt.Fprintf(conn, "Unknown command: %s\n", command)
	}
}

func handleGet(conn net.Conn, resource string){
	fmt.Fprintf(conn, "GET command recieved for: %s\n", resource)
}