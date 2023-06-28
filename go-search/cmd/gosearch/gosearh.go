// Main package
package main

import (
	"flag"
	"fmt"
	"go-course-4/go-search/pkg/crawler"
	"go-course-4/go-search/pkg/crawler/spider"
	"strings"
)

const (
	depth   = 1
	flagMsg = "Add request for searching in recieved collection"
	errMsg  = "Someting went wrong, check internet connection"
)

func main() {
	urls := []string{"https://golang.org", "https://go.dev"}

	sFlag := flag.String("s", "", "go language")
	flag.Parse()

	if len(*sFlag) == 0 {
		fmt.Println(flagMsg)
	}

	var docs []crawler.Document
	s := spider.New()

	for _, u := range urls {
		links, err := s.Scan(u, depth)
		if err != nil {
			fmt.Println(errMsg)
		}
		for _, l := range links {
			docs = append(docs, l)
		}
	}

	for _, d := range docs {
		if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*sFlag)) {
			fmt.Printf("`%s` found: %s\n", *sFlag, d.URL)
		}
	}
}
