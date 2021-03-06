package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup

	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go func(shout string) {
			fmt.Fprintln(c, "\t", strings.ToUpper(shout))
			time.Sleep(1 * time.Second)
			fmt.Fprintln(c, "\t", shout)
			time.Sleep(1 * time.Second)
			fmt.Fprintln(c, "\t", strings.ToLower(shout))
			wg.Done()
			select {
			case <-time.After(10 * time.Second):
				c.Close()
			}
		}(input.Text())
	}
	// closer
	go func() {
		wg.Wait()
		shutdownWrite(c)
	}()
}

func shutdownWrite(conn net.Conn) {
	v, ok := conn.(*net.TCPConn)
	if ok {
		v.CloseWrite()
	}
}
