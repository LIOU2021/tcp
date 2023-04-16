package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for i := 0; i < 3; i++ {
		conn, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			log.Fatal(err)
		}

		msg := "I'm a Kungfu Dev"
		if _, err := conn.Write([]byte(msg)); err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, len(msg))
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(buf[:n]))
	}
}
