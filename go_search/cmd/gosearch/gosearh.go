// Main package
package main

import (
	"flag"
	"fmt"
	"go_course_4/go_search/pkg/crawler"
	"go_course_4/go_search/pkg/crawler/spider"
)

const (
	depth   = 1
	flagMsg = "Add request for searching in recieved collection"
	errMsg  = "Someting went wrong, check internet connection"
)

func main() {
	urls := []string{"https://golang.org", "https://go.dev"}

	sFlag := flag.String("s", "", "Search")
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
		fmt.Println(links)
	}
	fmt.Println(docs)
}
