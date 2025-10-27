package main

import (
	"flag"
	"fmt"
	"project/download-json/api"
	"project/download-json/bins"
	"project/download-json/file"
	"project/download-json/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Ошибка")
	}
	action := flag.String("action", "", "Действие: create, get, update")
	name := flag.String("name", "", "Название bin")
	id := flag.String("id", "", "Ваш id")
	fileBin := flag.String("file", "", "Ваш id")
	private := flag.Bool("private", false, "Приватный bin")
	flag.Parse()
	// Подключаем модули через интерфейсы для DI
	var fileReader = &file.FileService{}
	var dataStorage = storage.NewStorage([]bins.Bin{})

	//Создаем bin

	// newBin := bins.NewBin("test-1", false, "Test Bin")

	//Используем bins модуль
	// binList := &bins.BinList{}
	// binList.Add(newBin)

	// Создаем storage с данными
	// storageWithData := storage.NewStorage([]bins.Bin{*newBin})

	// Используем storage через интерфейс

	// err = storageWithData.Save("data.json")

	// if err != nil {
	// 	fmt.Printf("Ошибка сохранения: %v\n", err)
	// 	return
	// }

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
	// api.GetBin("68fa4c66ae596e708f26ad6c")
	api.DeleteBin("68fb629fae596e708f28ce24")

	switch *action {
	case "create":
		api.CreateBin2(dataStorage, *fileBin, *private, *name)

	case "get":
		str, _ := api.GetBin(*id)
		fmt.Println(str.Storage.Bins)
	case "update":
		api.ChangeBin(*fileBin, *id)
	case "delete":
		api.DeleteBin(*id)

	}

	// fmt.Printf("Успешно! Все модули подключены через интерфейсы.\n")
	fmt.Printf("Прочитано %d байт\n", len(data))
	// fmt.Printf("Создан bin: %s\n", newBin)

}
