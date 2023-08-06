// Main package
package main

import (
	"fmt"

	"go-course-4/homework-11/pkg/crawler"
	"go-course-4/homework-11/pkg/crawler/spider"
	"go-course-4/homework-11/pkg/index"
	"go-course-4/homework-11/pkg/netsrv"
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
	netsrv.Listen(ind)
}
