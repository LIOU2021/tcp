package main

import (
	"fmt"
	"net"
	"strings"
)

func connHandler(c net.Conn) {
	fmt.Printf("Connect  %s >> %s\n", c.RemoteAddr(), c.LocalAddr())
	defer fmt.Printf("Close connection from %v\n", c.RemoteAddr())

	if c == nil {
		return
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		inputs := strings.Split(inStr, " ")
		switch inputs[0] {
		case "ping":
			fmt.Printf("ping from : %v\n", c.RemoteAddr())
			c.Write([]byte("pong\n"))
		case "echo":
			echoStr := strings.Join(inputs[1:], " ") + "\n"
			fmt.Printf("remoteAddr: %v, echo from : %s", c.RemoteAddr(), echoStr)
			c.Write([]byte(echoStr))
		case "quit":
			c.Close()
			break
		default:
			c.Write([]byte("welcome to my tcp server connect\n"))
		}
	}
}

func main() {
	address := "127.0.0.1:8000"
	fmt.Printf("listen: %s\n", address)

	server, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Fail to start server, %s\n", err)
	}
	fmt.Println("Server Started ...")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect, %s\n", err)
			break
		}
		go connHandler(conn)
	}
}
