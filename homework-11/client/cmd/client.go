package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	defaultProtocol = "tcp4"
	defaultSocket   = "localhost:8000"
)

func main() {
	conn, err := net.Dial(defaultProtocol, defaultSocket)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ответ от сервера:", string(msg))
}
