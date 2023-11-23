package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() //close connection before exit

	//Buffer for reading
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		n, err := conn.Read(buffer) // Read() blocks until it reads some data from the network and n is the number of bytes read
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		// Print the number of bytes read
		fmt.Printf("Received %d bytes\n", n)

		// Print received data
		fmt.Printf("Received message: %s", buffer[:n]) // :n is a slice operator that returns a slice of the first n bytes of the buffer

		// Send a response back to the client
		response := "Message received successfully\n"
		conn.Write([]byte(response))
	}
}
