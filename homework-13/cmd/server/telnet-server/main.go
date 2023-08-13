// Main package
package main

import (
	"fmt"
	"net"

	"go-course-4/homework-13/pkg/crawler"
	"go-course-4/homework-13/pkg/crawler/spider"
	"go-course-4/homework-13/pkg/index"
	"go-course-4/homework-13/pkg/netsrv"
)

const (
	proto = "tcp4"
	addr1 = ":8000"
	depth = 1
)

var urls = []string{"https://golang.org", "https://www.practical-go-lessons.com/"}

func main() {
	s := spider.New()
	ind := index.New()

	var docs []crawler.Document
	for _, u := range urls {
		links, err := s.Scan(u, depth)
		if err != nil {
			fmt.Printf("We got an error: %s\n", err)
			continue
		}
		docs = append(docs, links...)
	}
	ind.AddDocuments(docs)

	listener, err := net.Listen(proto, addr1)
	if err != nil {
		fmt.Printf("Something went wrong with server on %s: %s\n", addr1, err)
		return
	}
	defer listener.Close()

	err = netsrv.Start(listener, ind)
	if err != nil {
		fmt.Printf("Error from the server1: %s", err)
		return
	}
}
