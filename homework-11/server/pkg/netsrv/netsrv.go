package netsrv

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	proto = "tcp4"
	addr  = "0.0.0.0:8000"
)

// Start - запуск сетевой службы
func Start() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Minute * 2))

	response := bufio.NewReader(conn)
	fmt.Println(response)

}
