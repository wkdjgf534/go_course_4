package storage

import (
	"encoding/json"
	"go-course-4/homework-05/pkg/crawler"
	"io"
)

// Service - служба хранения результата службы сканирования.
type Service struct{}

// New - конструктор службы хранения.
func New() *Service {
	s := Service{}
	return &s
}

// Save - сохраняет результат поиска в файл.
func (s *Service) Save(docs *[]crawler.Document, w io.Writer) error {
	j, err := json.Marshal(*docs)
	if err != nil {
		return err
	}

	_, err = w.Write(j)
	if err != nil {
		return err
	}

	return nil
}

/*
func (s *Service) Save(docs *[]crawler.Document, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer f.Close()
	fmt.Println(*docs)

	j, err := json.Marshal(*docs)
	if err != nil {
		return err
	}
	fmt.Println(j)

	var d []crawler.Document
	json.Unmarshal(j, d)
	fmt.Println(d)
	return nil
}

// Load - загружает результат поиска из файла.
//func (s *Service) Load(name string) ([]crawler.Document, error) {
//	return s
//}
*/
