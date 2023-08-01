// Main package
package main

import (
	"io"
	"log"
	"net"
)

const (
	defaultProtocol = "tcp4"
	defaultSocket   = "0.0.0.0:8000"
	depth           = 1
)

var urls = []string{"https://golang.org", "https://www.practical-go-lessons.com/"}

func main() {
	listener, err := net.Listen(defaultProtocol, defaultSocket)
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

	/*
		var docs []crawler.Document
		s := spider.New()
		ind := index.New()

		for _, u := range urls {
			links, err := s.Scan(u, depth)
			if err != nil {
				fmt.Printf("We got an error: %s\n", err)
				continue
			}
			docs = append(docs, links...)
		}

		ind.Add(&docs)

		/*
			idx := ind.Ids(strings.ToLower(clentQuery))

			for _, i := range idx {
				min, max := docs[0].ID, docs[len(docs)-1].ID
				for min <= max {
					mid := (min + max) / 2
					if docs[mid].ID == i {
						fmt.Println("found document id:", i, "URL:", docs[mid].URL)
						break
					} else if docs[mid].ID < i {
						min = mid + 1
					} else {
						max = mid - 1
					}
				}
			}
	*/
}

func handler(conn io.ReadWriteCloser) {
	msg := "Test"
	conn.Write([]byte(msg))
	conn.Close()
}
