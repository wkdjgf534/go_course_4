package index

import (
	"go-course-4/homework-02/pkg/crawler"
	"strings"
)

// Service - служба инвертированного индексирования.
type Service struct {
	words map[string][]int
}

// strToSlice - Преобразование строки в slice строк
func strToSlice(str string) []string { return strings.Fields(strings.ToLower(str)) }

// New - конструктор службы инвертированного индексирования.
func New() *Service {
	i := Service{}
	i.words = map[string][]int{}
	return &i
}

// Add - добавление слов из заголовка и id документа в map
func (i *Service) Add(docs *[]crawler.Document) {
	for _, d := range *docs {
		arrStr := strToSlice(d.Title)
		for _, s := range arrStr {
			i.words[s] = append(i.words[s], d.ID)
		}
	}
}

// Ids - получение ID документов по ключевому слову
func (i *Service) Ids(str string) []int {
	var ids = []int{}
	arrStr := strToSlice(str)

	for _, s := range arrStr {
		ids = append(ids, i.words[s]...)
	}

	return ids
}
