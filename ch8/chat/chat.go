package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	name string
	ch   chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming messages to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}
		case nc := <-entering:
			clients[nc] = true
			for cli := range clients {
				if cli.name != nc.name {
					nc.ch <- cli.name + " is also here"
				}
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)       // outgoing client messages
	active := make(chan struct{}) // listen for not-idle signal
	go clientWriter(conn, ch)
	go handleIdleClient(conn, active)

	name := conn.RemoteAddr().String()
	client := client{name, ch}
	ch <- "You are " + name
	messages <- name + " has arrived"
	entering <- client

	input := bufio.NewScanner(conn)
	for input.Scan() {
		active <- struct{}{} // user is active
		messages <- name + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client
	messages <- name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func handleIdleClient(conn net.Conn, ch chan struct{}) {
	select {
	case <-ch:
		handleIdleClient(conn, ch)
		break
	case <-time.After(10 * time.Second):
		conn.Close()
	}
}
