package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

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
