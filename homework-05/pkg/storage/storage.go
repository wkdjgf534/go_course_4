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

// LoadFrom - загружает результат поиска из файла.
func (s *Service) LoadFrom(r io.Reader) ([]crawler.Document, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var docs []crawler.Document
	err = json.Unmarshal(data, &docs)
	if err != nil {
		return nil, err
	}
	return docs, nil
}
