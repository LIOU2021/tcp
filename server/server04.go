package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func connHandler(c net.Conn) {
	fmt.Printf("Connect  %s >> %s\n", c.RemoteAddr(), c.LocalAddr())
	defer func() {
		fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
	}()

	readWriteCh := make(chan struct{}, 100)
	go func() {
		for {
			select {
			case <-time.After(5 * time.Second):
				log.Println(c.RemoteAddr(), ":超过5秒没收到心跳")
				close(readWriteCh) // 关闭心跳机制
				c.Close()
				return
			case <-readWriteCh:
				log.Println(c.RemoteAddr(), ":收到讯息")
			}
		}
	}()

	if c == nil {
		return
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := c.Read(buf)
		log.Println("c.Read: ", string(buf[0:cnt]), "cnt:", cnt, "err:", err)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}

		readWriteCh <- struct{}{} // 通知心跳

		inStr := strings.TrimSpace(string(buf[0:cnt]))

		c.Write([]byte(inStr))
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
