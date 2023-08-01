package index

import (
	"strings"

	"go-course-4/homework-05/pkg/crawler"
)

// Index - служба инвертированного индексирования.
type Index struct {
	words map[string][]int
}

// strToSlice - Преобразование строки в slice строк
func strToSlice(str string) []string { return strings.Fields(strings.ToLower(str)) }

// New - конструктор службы инвертированного индексирования.
func New() *Index {
	return &Index{
		words: map[string][]int{},
	}
}

// Add - добавление слов из заголовка и id документа в map
func (i *Index) Add(docs *[]crawler.Document) error {
	for _, d := range *docs {
		arrStr := strToSlice(d.Title)
		for _, s := range arrStr {
			i.words[s] = append(i.words[s], d.ID)
		}
	}
	return nil
}

// Ids - получение ID документов по ключевому слову
func (i *Index) Ids(str string) []int { return i.words[str] }
