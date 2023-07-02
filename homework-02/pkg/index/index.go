package index

import (
	"fmt"
	"go-course-4/homework-02/pkg/crawler"
	"strings"
)

// Service - служба инвертированного индексирования.
type Service struct {
	words map[string][]int
}

// Преобразование строки в slice строк
func strToSlice(str string) []string { return strings.Fields(strings.ToLower(str)) }

// New - конструктор службы инвертированного индексирования.
func New() *Service {
	i := Service{}
	i.words = map[string][]int{}
	return &i
}

// Add - добавление слов и ids документов в map
func (i *Service) Add(docs *[]crawler.Document) {
	for _, d := range *docs {
		arrStr := strToSlice(d.Title)
		for _, s := range arrStr {
			i.words[s] = append(i.words[s], d.ID)
		}
	}
}

// Search - поиск в map по ключевым словам и получение ids документов
func (i *Service) Search(str string) {
	var ids = []int{}
	arrStr := strToSlice(str)
	for _, s := range arrStr {
		fmt.Printf("%T", i.words[s])
		ids = append(ids, i.words[s])
	}
	fmt.Println(ids)
}
