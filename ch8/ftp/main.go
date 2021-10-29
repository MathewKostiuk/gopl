package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
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
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	for input.Scan() {
		cmd := strings.Split(input.Text(), " ")
		switch cmd[0] {
		case "close":
			c.Close()
			return
		case "ls":
			command := exec.Command("ls")
			command.Stdout = c
			command.Run()
			continue
		case "cd":
			if err := os.Chdir(cmd[1]); err != nil {
				log.Printf("command finished with error: %v\n", err)
			}
			continue
		case "get":
			f, err := os.Open(cmd[1])
			if err != nil {
				log.Printf("get finished with error: %v\n", err)
			}
			mustCopy(c, f)

			continue
		default:
			resp := "available commands are: ls, cd, get, and close\n"
			r := strings.NewReader(resp)
			mustCopy(c, r)
			continue
		}
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("error during copy: %v\n", err)
	}
}
