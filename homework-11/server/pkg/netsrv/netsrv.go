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

// Start - запуск сетевой службы
func Start(ind *index.Index) {
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
		handler(conn, ind)
	}
}

func handler(conn net.Conn, ind *index.Index) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Minute * 2))

	response := bufio.NewReader(conn)
	fmt.Println(response)

	res, _, err := response.ReadLine()
	if err != nil {
		return
	}

	fmt.Printf("%T", ind)
	result := ind.Search("go")

	fmt.Println(result)
	//for _, d := range result {
	//	fmt.Println(d)
	//}

	if len(res) == 0 {
		_, err := fmt.Fprintf(conn, "Nothing found\n")
		if err != nil {
			return
		}
	}
}
