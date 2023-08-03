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
		log.Print(err)
		return
	}
	defer conn.Close()

	go readFromSrv(conn)

	r := bufio.NewReader(os.Stdin)
	fmt.Println("Input your search data")
	for {
		text, _ := r.ReadString('\n')
		req := []byte(text)

		_, err := conn.Write(req)
		if err != nil {
			return
		}

		if text == "quit\n" || text == "exit\n" {
			fmt.Println("exiting...")
			return
		}
	}
}

func readFromSrv(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		fmt.Println(string(msg))
	}
}
