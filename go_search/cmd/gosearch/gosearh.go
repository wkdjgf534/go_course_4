// Main package
package main

import (
	"flag"
	"fmt"
	"go_course_4/go_search/pkg/crawler/spider"
)

const (
	crawlerDepth = 1
	flagMsg      = "Add request for searching in recieved collection"
	errMsg       = "Someting went wrong, check internet connection"
)

func main() {
	urls := [2]string{"https://golang.org", "https://go.dev"}

	s := spider.New()

	for _, url := range urls {
		links, err := s.Scan(url, crawlerDepth)
		if err != nil {
			fmt.Println(errMsg)
		}
		fmt.Println(links)
	}

	sFlag := flag.String("s", "", "Search")

	flag.Parse()

	if len(*sFlag) == 0 {
		fmt.Println(flagMsg)
	}

	fmt.Println(urls, *sFlag)
}
