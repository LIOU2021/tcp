package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "quit" {
			fmt.Println("close connect !")
			conn.Close()
			break
		}
		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			fmt.Printf("Fail to send data, %s\n", err)
			break
		}
		buf := make([]byte, 4096)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read from server: %s\n", err)
			break
		}
		fmt.Println(string(buf[0:cnt]))
	}
}
