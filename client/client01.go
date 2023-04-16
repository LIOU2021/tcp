package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("client start")
	defer fmt.Println("client end")
	for i := 0; i < 3; i++ {
		conn, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			log.Fatal(err)
		}

		msg := "I'm a Dev"
		if _, err := conn.Write([]byte(msg)); err != nil {
			log.Fatal(err)
		}

		// buf := make([]byte, len(msg))
		buf := make([]byte, 1000)

		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(string(buf[:n]))

		_ = n
		fmt.Println(string(buf))
	}
}
