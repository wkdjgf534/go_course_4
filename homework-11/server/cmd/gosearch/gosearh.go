// Main package
package main

import (
	"fmt"

	"go-course-4/homework-11/server/pkg/crawler/spider"
	"go-course-4/homework-11/server/pkg/index"
	"go-course-4/homework-11/server/pkg/netsrv"
)

const (
	proto = "tcp4"
	addr  = "0.0.0.0:8000"
	depth = 2
)

var urls = []string{"https://golang.org", "https://www.practical-go-lessons.com/"}

func main() {
	s := spider.New()
	ind := index.New()

	for _, u := range urls {
		links, err := s.Scan(u, depth)
		if err != nil {
			fmt.Printf("We got an error: %s\n", err)
			continue
		}
		ind.AddDocuments(links)
	}

	netsrv.Listen(ind)
}
