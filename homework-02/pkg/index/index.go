package index

import (
	"fmt"
)

// Service - служба инвертированного индексирования.
type Service struct {
}

// New - конструктор службы инвертированного индексирования.
func New() *Service {
	s := Service{}
	return &s
}

// Add - добавление слова и номер документа в выборку
func (i *Service) Add(s string) {
	fmt.Println(s)
}

// Search - бинарный поиск проиндексированных значений
func (i *Service) Search(s string) {
	fmt.Println(s)
}
