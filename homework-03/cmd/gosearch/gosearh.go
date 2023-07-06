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

func main() {
	depth := 2
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
			fmt.Printf("We got an error: %s\n", err)
			continue
		}

		for _, l := range links {
			l.ID = counter
			docs = append(docs, l)
			counter++
		}
	}

	err := index.Add(&docs)
	if err != nil {
		fmt.Print(err)
		return
	}

	idx, err := index.Ids(strings.ToLower(*sFlag))
	if err != nil {
		fmt.Println(err)

	}

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
