// Main package
package main

import (
	"flag"
	"fmt"
	"go-course-4/homework-02/pkg/crawler"
	"go-course-4/homework-02/pkg/crawler/spider"
	"go-course-4/homework-02/pkg/index"
)

const (
	depth   = 1
	flagMsg = "Use parameter -s and add a preferable key word (-s go)"
	errMsg  = "Someting went wrong"
	unsMsg  = "Have not found any documents according to your key word"
)

func main() {
	urls := []string{"https://golang.org", "https://www.practical-go-lessons.com/"}

	sFlag := flag.String("s", "", flagMsg)
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
			fmt.Println(errMsg)
		}

		for _, l := range links {
			l.ID = counter
			docs = append(docs, l)
			counter++
		}
	}

	index.Add(&docs)

	idx := index.Ids(*sFlag)
	if len(idx) == 0 {
		fmt.Println(unsMsg)
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
