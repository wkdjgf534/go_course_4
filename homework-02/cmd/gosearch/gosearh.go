// Main package
package main

import (
	"flag"
	"fmt"
	"go-course-4/homework-02/pkg/crawler"
	"go-course-4/homework-02/pkg/crawler/spider"
	"strings"
)

const (
	depth   = 2
	flagMsg = "Use parameter -s and add a preferable string for searching"
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
	//i := index.New()
	count := 0

	for _, u := range urls {
		links, err := s.Scan(u, depth)
		if err != nil {
			fmt.Println(errMsg)
		}

		for _, l := range links {
			l.ID = count
			docs = append(docs, l)
			count++
		}
	}

	for _, d := range docs {
		if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*sFlag)) {
			fmt.Printf("%s found: %v\n", *sFlag, d)
		}
	}
}
