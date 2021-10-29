package main

import (
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
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	shutdownWrite(conn)
	<-done
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
