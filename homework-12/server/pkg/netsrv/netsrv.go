package netsrv

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"go-course-4/homework-11/server/pkg/index"
)

const (
	proto = "tcp4"
	addr  = "0.0.0.0:8000"
)

// Listen - запуск сетевой службы
func Listen(index *index.Index) {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error connection: %v\n", err)
			continue
		}
		go handler(conn, index)
	}
}

func handler(conn net.Conn, index *index.Index) {
	defer conn.Close()
	defer fmt.Println("Connection Closed")

	conn.SetDeadline(time.Now().Add(time.Second * 60))
	r := bufio.NewReader(conn)

	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		search := index.Search(string(msg))

		if len(msg) == 0 || len(search) == 0 {
			fmt.Fprintf(conn, "Nothing found, repeat again\n")
		}

		for _, d := range search {
			fmt.Fprintf(conn, "Found article %s - %s\n", d.Title, d.URL)
		}

		conn.SetDeadline(time.Now().Add(time.Second * 30))
	}

}
