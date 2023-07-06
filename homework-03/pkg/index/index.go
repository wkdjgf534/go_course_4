package index

import (
	"errors"
	"go-course-4/homework-03/pkg/crawler"
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
func (i *Service) Add(docs *[]crawler.Document) error {
	if len(*docs) == 0 {
		return errors.New("Check you internet connection or URL")
	}

	for _, d := range *docs {
		arrStr := strToSlice(d.Title)
		for _, s := range arrStr {
			i.words[s] = append(i.words[s], d.ID)
		}
	}
	return nil
}

// Ids - получение ID документов по ключевому слову
func (i *Service) Ids(str string) ([]int, error) {
	ids := i.words[str]

	if ids == nil {
		return nil, errors.New("Have not found any documents according to your key word")
	}
	return ids, nil
}
