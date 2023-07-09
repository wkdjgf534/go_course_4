package storage

import (
	"encoding/json"
	"fmt"
	"go-course-4/homework-05/pkg/crawler"
	"os"
)

// Service - служба хранения результата службы сканирования.
type Service struct{}

// New - конструктор службы хранения.
func New() *Service {
	s := Service{}
	return &s
}

// Save - сохраняет результат поиска в файл.
func (s *Service) Save(docs *[]crawler.Document, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer f.Close()

	j, err := json.Marshal(docs)
	if err != nil {
		return err
	}
	fmt.Println(j)

	//d := make([]crawler.Document, 0)
	//json.Unmarshal(j, d)
	//fmt.Println(d)
	return nil
}

// Load - загружает результат поиска из файла.
//func (s *Service) Load(name string) ([]crawler.Document, error) {
//	return s
//}
