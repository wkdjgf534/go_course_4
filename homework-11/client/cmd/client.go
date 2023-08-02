package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	proto = "tcp4"
	addr  = "localhost:8000"
)

func main() {
	conn, err := net.Dial(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	server := bufio.NewReader(conn)

	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatal(err)
		}

		b := append(scanner.Bytes(), '\r', '\n')
		_, err = conn.Write(b)
		if err != nil {
			log.Fatal(err)
		}
		for {
			b, _, err = server.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			if len(b) == 0 {
				break
			}
			fmt.Println(string(b))
		}
	}
}
