package index

import (
	"go-course-4/homework-02/pkg/crawler"
	"strings"
)

// Service - служба инвертированного индексирования.
type Service struct {
	words map[string][]int
}

// New - конструктор службы инвертированного индексирования.
func New() *Service {
	i := Service{}
	i.words = map[string][]int{}
	return &i
}

// Add - добавление слова и id документа в map
func (i *Service) Add(str string, docs *[]crawler.Document) {
	strs := strings.Fields(str) // Строку преобразуем в массив
	for _, s := range strs {
		for _, d := range *docs {
			if strings.Contains(strings.ToLower(d.Title), strings.ToLower(s)) {
				i.words[s] = append(i.words[s], d.ID)
			}
		}
	}
}

// Search - бинарный поиск по индексам документов
func (i *Service) Search() {
}
