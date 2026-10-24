package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8080", 5*time.Second)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server!")
	fmt.Println("My endpoint:", conn.LocalAddr())
	fmt.Println("Server endpoint:", conn.RemoteAddr())

	// Write a message to the server.
	fmt.Fprintf(conn, "GET /index.html\n")

	//read the response from the server
	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err != nil {
		log.Fatal("Read failed:", err)
	}

	// Print the response from the server.
	fmt.Println("Received from server:", string(bs[:n]))
}
