// Main package
package main

import (
	"flag"
	"fmt"
	"go-course-4/go-search/pkg/crawler"
	"go-course-4/go-search/pkg/crawler/spider"
	"log"
	"strings"
)

const depth = 1

func main() {
	urls := []string{"https://golang.org", "https://go.dev"}

	// third parameter set a default value
	sFlag := flag.String("s", "", "go")
	flag.Parse()

	if len(*sFlag) == 0 {
		flag.PrintDefaults()
	}

	var docs []crawler.Document
	s := spider.New()

	for _, u := range urls {
		links, err := s.Scan(u, depth)
		if err != nil {
			log.Fatal(err)
		}
		for _, l := range links {
			docs = append(docs, l)
		}
	}

	for _, d := range docs {
		if strings.Contains(d.Title, strings.ToLower(*sFlag)) {
			fmt.Printf("`%s` found: %s\n", *sFlag, d.URL)
		}
	}
}
