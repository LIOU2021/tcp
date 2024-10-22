package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	address := "127.0.0.1:8000"
	fmt.Printf("listen: %s\n", address)
	socket, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Accept an incoming connection.
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("new connection: %s\n", conn.RemoteAddr().String())
		// Handle the connection in a separate goroutine.
		go func(conn net.Conn) {
			defer conn.Close()
			// Create a buffer for incoming data.
			buf := make([]byte, 4096)

			// Read data from the connection.
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("got a message")
			// Echo the data bac to the connection.

			response := fmt.Sprintf("From server: %s", string(buf[:n]))
			_, err = conn.Write([]byte(response))
			if err != nil {
				log.Fatal(err)
			}
		}(conn)
	}
}
