// Main package
package main

import (
	"fmt"
	"io"

	"go-course-4/homework-11/server/pkg/crawler"
	"go-course-4/homework-11/server/pkg/crawler/spider"
	"go-course-4/homework-11/server/pkg/index"
	"go-course-4/homework-11/server/pkg/netsrv"
)

const (
	proto = "tcp4"
	addr  = "0.0.0.0:8000"
	depth = 1
)

var urls = []string{"https://golang.org", "https://www.practical-go-lessons.com/"}

func main() {
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

	netsrv.Start()
}

func handler(conn io.ReadWriteCloser) {
	msg := "Test"
	conn.Write([]byte(msg))
	conn.Close()
}
