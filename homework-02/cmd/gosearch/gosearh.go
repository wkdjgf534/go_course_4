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
	flagMsg = "Use parameter -s and add a preferable key words (-s \"go test rest\")"
	errMsg  = "Someting went wrong"
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
	min, max := docs[0].ID, docs[len(docs)-1].ID

	for _, i := range idx {
		fmt.Println(i)
		for min <= max {
			fmt.Println(min, max)
			mid := (min + max) / 2
			if docs[mid].ID == i {
				fmt.Println(docs[mid].URL)
			}
			if docs[mid].ID < i {
				min = mid + 1
			} else {
				max = mid - 1
			}
		}
	}
}

/*
func Binary(data []int, item int) int {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == item {
			return mid
		}
		if data[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

*/
