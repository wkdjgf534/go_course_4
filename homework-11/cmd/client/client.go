package main

import (
	"bufio"
	"fmt"
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
		fmt.Println(err)
		return
	}
	defer conn.Close()

	r := bufio.NewReader(os.Stdin)
	fmt.Println("Input your search data:")

	for {
		text, err := r.ReadString('\n')
		if err != nil {
			return
		}

		req := []byte(text)

		_, err = conn.Write(req)
		if err != nil {
			return
		}

		if text == "quit\n" || text == "exit\n" {
			fmt.Println("exiting...")
			return
		}

		r := bufio.NewReader(conn)
		for {
			msg, _, err := r.ReadLine()
			if err != nil {
				return
			}

			fmt.Println(string(msg))
		}
	}
}
