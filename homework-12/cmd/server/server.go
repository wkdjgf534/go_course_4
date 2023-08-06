// Main package
package main

import (
	"fmt"
	"net"

	"go-course-4/homework-12/pkg/crawler"
	"go-course-4/homework-12/pkg/crawler/spider"
	"go-course-4/homework-12/pkg/index"
	"go-course-4/homework-12/pkg/netsrv"
)

const (
	proto = "tcp4"
	addr  = "0.0.0.0:8000"
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

	listener, err := net.Listen(proto, addr)
	if err != nil {
		fmt.Printf("Something went wrong with server on %s: %s\n", addr, err)
		return
	}
	defer listener.Close()

	err = netsrv.Listen(listener, ind)
	if err != nil {
		fmt.Printf("Yet another error from the server: %s", err)
		return
	}

}
