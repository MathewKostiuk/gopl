package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	s384 := flag.Bool("SHA384", false, "Print SHA384 hash")
	s512 := flag.Bool("SHA512", false, "Print SHA512 hash")
	flag.Parse()

	for _, arg := range flag.Args() {
		if *s384 {
			sha384 := sha512.Sum384([]byte(arg))
			fmt.Printf("SHA384=%x\n", sha384)
		}

		if *s512 {
			sha512 := sha512.Sum512([]byte(arg))
			fmt.Printf("SHA512=%x\n", sha512)
		}
		s := sha256.Sum256([]byte(arg))
		fmt.Printf("SHA256=%x\n", s)

	}
}
