package storage

import (
	"go-course-4/homework-05/pkg/crawler"
	"os"
)

// Save - сохраняет результат поиска в файл
func Save(source []crawler.Document, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer f.Close()
	return nil
}

// Load - загружает результат поиска из файла
func Load(name string) ([]crawler.Document, error) {
	//file

}
