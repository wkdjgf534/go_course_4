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

// read - Чтение из файла
func read(name string, st *storage.Service) ([]crawler.Document, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	docs, err := st.Load(f)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

// write - Запись в файл
func write(docs *[]crawler.Document, name string, st *storage.Service) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	err = st.Save(docs, f)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	depth := 1
	fName := "./backup.txt"
	urls := []string{"https://golang.org", "https://www.practical-go-lessons.com/"}
	sFlag := flag.String("s", "", "Use parameter -s and add a preferable key word (-s go)")
	flag.Parse()

	if len(*sFlag) == 0 {
		flag.PrintDefaults()
		return
	}

	var docs []crawler.Document
	s := spider.New()
	i := index.New()
	st := storage.New()

	if _, err := os.Stat(fName); err == nil {
		docs, err := read(fName, st)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("test", docs)
	}

	if len(docs) == 0 {
		for _, u := range urls {
			links, err := s.Scan(u, depth)
			if err != nil {
				fmt.Printf("We got an error: %s\n", err)
				continue
			}
			docs = append(docs, links...)
			write(&docs, fName, st)
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
