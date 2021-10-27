package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	const format = "%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Timezone", "Time")
	fmt.Fprintf(tw, format, "--------", "--------")

	first := strings.Split(os.Args[1], "=")
	for _, arg := range os.Args[2:] {
		port := strings.Split(arg, "=")
		conn, err := net.Dial("tcp", port[1])
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(tw, conn)
	}

	conn, err := net.Dial("tcp", first[1])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(tw, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
