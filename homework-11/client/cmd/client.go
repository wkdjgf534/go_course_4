package main

import (
	"fmt"
	"io"
	"log"
	"net"
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

	msg, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ответ от сервера:", string(msg))
}
