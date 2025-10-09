package main

import (
	"fmt"
	"project/download-json/bins"
	"project/download-json/file"
	"project/download-json/storage"
)

func main() {

	// Подключаем модули через интерфейсы для DI
	var fileReader = &file.FileService{}
	var dataStorage = storage.NewStorage([]bins.Bin{})

	//Создаем bin

	newBin := bins.NewBin("test-1", false, "Test Bin")

	//Используем bins модуль
	binList := &bins.BinList{}
	binList.Add(newBin)

	// Создаем storage с данными
	storageWithData := storage.NewStorage([]bins.Bin{*newBin})

	// Используем storage через интерфейс

	err := storageWithData.Save("data.json")

	if err != nil {
		fmt.Printf("Ошибка сохранения: %v\n", err)
		return
	}

	// Используем file через интерфейс

	data, err := fileReader.Read("data.json")

	if err != nil {
		fmt.Printf("Ошибка чтения: %v\n", err)
		return
	}

	// Загружаем данные через интерфейс

	err = dataStorage.Load("data.json")

	if err != nil {
		fmt.Printf("Ошибка загрузки: %v\n", err)
		return
	}

	fmt.Printf("Успешно! Все модули подключены через интерфейсы.\n")
	fmt.Printf("Прочитано %d байт\n", len(data))
	fmt.Printf("Создан bin: %s\n", newBin)

}
