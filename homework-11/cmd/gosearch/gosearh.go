// Main package
package main

import (
	"fmt"

	"go-course-4/homework-05/pkg/crawler"
	"go-course-4/homework-05/pkg/crawler/spider"
	"go-course-4/homework-05/pkg/index"
)

const depth = 2

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
}
