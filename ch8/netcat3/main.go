package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type Closer interface {
	CloseWrite() error
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Println(err)
		}
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	shutdownWrite(conn)

	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func shutdownWrite(conn net.Conn) {
	v, ok := conn.(Closer)
	if ok {
		v.CloseWrite()
	}
}
