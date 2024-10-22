package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type ConnPool struct {
	Dial     func() (net.Conn, error) // 連線方式
	MaxIdle  int                      // 最大閒置連線數量
	MinIdle  int                      // 最小閒置連線數量
	conns    []*connWithTime
	mu       sync.Mutex
	idleTime time.Duration // 閒置時間
}

func (p *ConnPool) CreatePool() {
	for i := 0; i < p.MinIdle; i++ {
		conn, err := p.Dial()
		if err != nil {
			log.Fatal(err)
			return
		}

		con := &connWithTime{
			Conn: conn,
			t:    time.Now(),
		}

		p.conns = append(p.conns, con)
	}
}

func (p *ConnPool) Get() (net.Conn, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.conns) < 1 {
		fmt.Println("create one connect from Get method")
		conn, err := p.Dial()
		if err != nil {
			return nil, err
		}

		return conn, nil
	}

	conn := p.conns[0]
	p.conns = p.conns[1:]
	if p.idleTime > 0 && time.Since(conn.t) > p.idleTime {
		conn.Close()
		return p.Get()
	}

	return conn.Conn, nil
}

func (p *ConnPool) put(conn net.Conn) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.MaxIdle > 0 && len(p.conns) >= p.MaxIdle {
		fmt.Println("delete idle connection")
		conn.Close()
		return nil
	}
	p.conns = append(p.conns, &connWithTime{conn, time.Now()})
	return nil
}

type connWithTime struct {
	net.Conn
	t time.Time
}

func main() {
	address := "127.0.0.1:8000"
	fmt.Printf("connecting to: %s\n", address)

	pool := &ConnPool{
		Dial: func() (net.Conn, error) {
			return net.Dial("tcp", address)
		},
		MaxIdle:  10,
		MinIdle:  2,
		idleTime: 10 * time.Second,
	}

	fmt.Println("create pool")
	pool.CreatePool()
	fmt.Println("init pool: ", len(pool.conns))

	go func() {
		for {
			fmt.Printf("connections: %d\n", len(pool.conns))
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("test concurrency")
	go func() {
		for i := 0; i < 15; i++ {
			go func(n int) {
				for {
					log.Printf("index: %d, try get conn\n", n)
					conn, err := pool.Get()
					if err != nil {
						log.Printf("index: %d, pool get fail: %v\n", n, err)
					}
					log.Printf("index: %d, get conn. current conns: %d\n", n, len(pool.conns))

					time.Sleep(2 * time.Second)

					// conn.Write([]byte("ping"))
					conn.Write([]byte(fmt.Sprintf("echo %d", n)))

					time.Sleep(2 * time.Second)
					log.Printf("index: %d, put conn\n", n)
					pool.put(conn)

					time.Sleep(3 * time.Second)
				}
			}(i)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}

}
