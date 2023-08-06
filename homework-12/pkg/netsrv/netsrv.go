package netsrv

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"go-course-4/homework-12/pkg/index"
)

// Listen - запуск сетевой службы
func Listen(listener net.Listener, index *index.Index) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error connection: %v\n", err)
			continue
		}
		go handler(conn, index)
	}
}

func handler(conn net.Conn, index *index.Index) error {
	defer conn.Close()
	defer fmt.Println("Connection Closed")

	conn.SetDeadline(time.Now().Add(time.Second * 60))
	r := bufio.NewReader(conn)

	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return err
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
