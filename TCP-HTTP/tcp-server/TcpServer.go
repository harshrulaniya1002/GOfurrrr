package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	conn, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fmt.Println("TCP server listening on 127.0.0.1:8080")

	for {
		conn, err := conn.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		// Handle each client concurrently.
		// multiplexing: each client connection is handled in its own goroutine.
		go handleClient(conn)
	}

}

func handleGet(conn net.Conn, resource string) {
	// Simulate fetching the resource.
	// In a real application, you would fetch the resource from a database or file system.
	fmt.Fprintf(conn, "GET Command received for resource: %s\n", resource)

	// Simulate sending a response back to the client.
	// response := fmt.Sprintf("Fetched resource: %s\n", resource)
	// _, err := conn.Write([]byte(response))
	// if err != nil {
	// 	log.Println("Write error:", err)
	// 	return
	// }
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())
	fmt.Println("Server endpoint:", conn.LocalAddr())

	reader := bufio.NewReader(conn)


	// for loop to continuously read messages from the client
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(conn, "Error reading command: %v\n", err)
			return
		}

		// Split the message into command and arguments
		// for eg: if the message is "ECHO Hello", then parts[0] will be "ECHO" and parts[1] will be "Hello"
		parts := strings.SplitN(strings.TrimSpace(message)," ", 2)
		if len(parts) < 2 {
			fmt.Fprintf(conn, "Invalid command format\n")
			return
		}

		command := parts[0]
		resource:= parts[1]
		log.Printf("Recieved Command: %s, %s\n",command, resource)

		switch command {
		case "GET":
			// Handle GET request
			handleGet(conn, resource)
		default:
			fmt.Fprintf(conn, "Unknown command: %s\n", command)
		}
	}
}
