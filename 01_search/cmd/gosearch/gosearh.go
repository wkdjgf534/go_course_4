// Main package
package main

import (
	"flag"
	"fmt"
	"search/pkg/crawler"
	"search/pkg/crawler/spider
)

func main() {
	urls := [2]string{"https://golang.org", "https://go.dev"}

	sFlag := flag.String("s", "", "Search")

	flag.Parse()

	if len(*sFlag) == 0 {
		fmt.Println("Add request for searching")
	}

	fmt.Println(urls, *sFlag)
}
