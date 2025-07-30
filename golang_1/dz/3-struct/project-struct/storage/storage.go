package storage

import (
	"encoding/json"
	"errors"
	"os"
	"project/download-json/bins"
)

type Storage struct {
	Bins bins.BinList `json: "bin"`
}

func NewStorage(bins []bins.Bin) *Storage {
	return &Storage{
		Bins: bins,
	}
}

func (storage *Storage) Save(path string) error {
	file, err := json.Marshal(storage)
	if err != nil {
		return errors.New("не удалось преобразовать данные")
	}

	err = os.WriteFile(path, file, 0644)
	if err != nil {
		return errors.New("не удалось записать данные")
	}
	return nil

}

func (storage *Storage) Load(path string) error {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return errors.New("не смог прочитать файл")
	}

	err = json.Unmarshal(fileContent, &storage)

	if err != nil {
		return errors.New("не удалось разобрать файл")
	}

	return nil

}
