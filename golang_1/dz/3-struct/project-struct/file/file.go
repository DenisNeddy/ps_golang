package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Интерфейс для DI
type FileReader interface {
	Read(path string) ([]byte, error)
}

// Реализация интерфейса
type FileService struct{}

func (fs *FileService) Read(path string) ([]byte, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return nil, errors.New("__не удалось прочесть файл__")
	}
	return file, nil
}

// Оригинальная функция (отсавляем)
func Read(path string) ([]byte, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return nil, errors.New("__не удалось прочесть файл__")
	}
	return file, nil
}

func isJsonFile(path string) bool {
	return strings.EqualFold(filepath.Ext(path), ".json")
}
