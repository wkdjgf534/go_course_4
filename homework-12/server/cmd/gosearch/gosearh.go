// Main package
package main

import (
	"fmt"

	"go-course-4/homework-12/server/pkg/crawler"
	"go-course-4/homework-12/server/pkg/crawler/spider"
	"go-course-4/homework-12/server/pkg/index"
	"go-course-4/homework-12/server/pkg/netsrv"
	"go-course-4/homework-12/server/pkg/webapp"
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
	go webapp.Listen(ind)
	netsrv.Listen(ind)
}
