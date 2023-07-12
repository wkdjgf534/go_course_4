// Main package
package main

import (
	"flag"
	"fmt"
	"go-course-4/homework-05/pkg/crawler"
	"go-course-4/homework-05/pkg/crawler/spider"
	"go-course-4/homework-05/pkg/index"
	"go-course-4/homework-05/pkg/storage"
	"os"
	"strings"
)

const (
	depth = 1
	fName = "./backup.txt"
)

var urls = []string{"https://golang.org", "https://www.practical-go-lessons.com/"}

// read - Чтение из файла
func read(name string) ([]crawler.Document, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	docs, err := storage.LoadFrom(f)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

// write - Запись в файл
func write(name string, docs *[]crawler.Document) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	err = storage.Save(docs, f)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	sFlag := flag.String("s", "", "Use parameter -s and add a preferable key word (-s go)")
	flag.Parse()

	if len(*sFlag) == 0 {
		flag.PrintDefaults()
		return
	}

	var docs []crawler.Document
	s := spider.New()
	i := index.New()

	docs, err := read(fName)
	if err != nil {
		fmt.Printf("Error: %s\nWe have to download new data again\n", err)
	}

	if len(docs) == 0 || err != nil {
		for _, u := range urls {
			links, err := s.Scan(u, depth)
			if err != nil {
				fmt.Printf("We got an error: %s\n", err)
				continue
			}
			docs = append(docs, links...)
		}
		err := write(fName, &docs)
		if err != nil {
			fmt.Printf("Error: %s\nWe can not write new data to the file\n", err)
		}
	}

	i.Add(&docs)

	idx := i.Ids(strings.ToLower(*sFlag))

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
