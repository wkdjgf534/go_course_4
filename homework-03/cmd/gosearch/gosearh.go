// Main package
package main

import (
	"flag"
	"fmt"
	"go-course-4/homework-03/pkg/crawler"
	"go-course-4/homework-03/pkg/crawler/spider"
	"go-course-4/homework-03/pkg/index"
	"strings"
)

const (
	depth = 1
)

func main() {
	urls := []string{"https://golang.org", "https://www.practical-go-lessons.com/"}

	sFlag := flag.String("s", "", "Use parameter -s and add a preferable key word (-s go)")
	flag.Parse()

	if len(*sFlag) == 0 {
		flag.PrintDefaults()
		return
	}

	var docs []crawler.Document
	s := spider.New()
	index := index.New()
	counter := 0

	for _, u := range urls {
		links, err := s.Scan(u, depth)
		if err != nil {
			fmt.Println("Someting went wrong")
		}

		for _, l := range links {
			l.ID = counter
			docs = append(docs, l)
			counter++
		}
	}

	index.Add(&docs)

	idx := index.Ids(strings.ToLower(*sFlag))
	if len(idx) == 0 {
		fmt.Println("Have not found any documents according to your key word")
		return
	}

	fmt.Println("test")
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
}
