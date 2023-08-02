package index

import (
	"fmt"
	"strings"

	"go-course-4/homework-11/server/pkg/crawler"
)

// Index - служба инвертированного индексирования.
type Index struct {
	Docs  []crawler.Document
	Words map[string][]int
}

// strToSlice - Преобразование строки в slice строк
func strToSlice(str string) []string { return strings.Fields(strings.ToLower(str)) }

// New - конструктор службы инвертированного индексирования.
func New() *Index {
	return &Index{
		Words: map[string][]int{},
		Docs:  []crawler.Document{},
	}
}

// AddDocuments - добавление документов
func (i *Index) AddDocuments(links []crawler.Document) {
	i.Docs = append(i.Docs, links...)
	i.addIndex()
}

// addIndex - добавление слов из заголовка и id документа в map
func (i *Index) addIndex() error {
	for _, d := range i.Docs {
		arrStr := strToSlice(d.Title)
		for _, s := range arrStr {
			i.Words[s] = append(i.Words[s], d.ID)
		}
	}
	return nil
}

// Search - поиск результата
func (i *Index) Search(str string) {
	idx := i.Words[str]
	for _, ind := range idx {
		min, max := i.Docs[0].ID, i.Docs[len(i.Docs)-1].ID
		for min <= max {
			mid := (min + max) / 2
			if i.Docs[mid].ID == ind {
				fmt.Println("found document id:", ind, "URL:", i.Docs[mid].URL)
				break
			} else if i.Docs[mid].ID < ind {
				min = mid + 1
			} else {
				max = mid - 1
			}
		}
	}

}
