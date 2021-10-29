package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	go handleResponse(os.Stdout, conn)
	handleConn(conn, os.Stdin)
}

func handleConn(dst io.Writer, src io.Reader) {
	x, err := io.Copy(dst, src)
	fmt.Println(x)
	if err != nil {
		log.Fatal(err)
	}
}

func handleResponse(dst io.Writer, c net.Conn) {
	_, err := io.Copy(dst, c)
	if err != nil {
		log.Fatal(err)
	}
	c.Close()
}
