package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"project/download-json/bins"
	"project/download-json/storage"

	"github.com/google/uuid"
)

type Store struct {
	Record storage.Storage
}

// POST - Создание bin

func CreateBin(store *storage.Storage, private bool, name string) error {
	// Проверяем входящий параметр name на заполненность
	if name == "" {

		return errors.New("ошибка: название bin не может быть пустым")
	}

	// Получение ключа
	key, err := GetKey()

	if err != nil {
		return err
	}

	// Создаем новый объект Bin используя конструктор структуры Bin

	newBin := bins.NewBin(uuid.New().String(), private, name)

	// Кладем новый объект в список объектов стурктуры Storage
	store.Bins = append(store.Bins, *newBin)

	binsData := [2]string{newBin.Id, name}
	store.List = append(store.List, binsData)

	// Сохраняем данные структуры Storage в файле data.json , тем самым обновляея локальный файл data.json после оьновления Storage
	// Метод Save возвращает ошибку , если что то пошло не так
	err = store.Save("data.json")
	if err != nil {

		return err
	}

	// Далее переводим обновленный объект Storage в Json формат
	postBody, _ := json.Marshal(store)

	// Создаем запрос (метод, адрес, побготовленные данные). Возвращает запрос или ошибку
	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(postBody))

	if err != nil {
		return err
	}

	// Добавляем в запрос заголовок, который отвечает за передачу данных в json формате
	req.Header.Set("Content-Type", "application/json")
	// Добавляем кастомный заголовок , который требует сервис jsonbin.io , чтобы передавать запросы
	req.Header.Set("X-Master-Key", key)

	// Выводим заголовки и их статусы
	fmt.Println("Request Headers:")
	for key, values := range req.Header {
		fmt.Printf(" %s: %v\n", key, values)
	}
	// Подготовка к отправке запроса. Если позьзуемся функцией NewRequest, то нужно инициализировать отправку через &http.Client{}
	client := &http.Client{}
	// Отправляем нужные данные . Возвращает ответ и ошибку
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}

	// инициализируем закрытие тела запроса перед завершением работы функции CreateBin
	defer req.Body.Close()
	// Читаем тело ответа, чтобы его ввыести в уведомление. Не обязательно для отправки методом POST, обязательно для GET
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	// Вывод уведомлений после отправки запроса
	fmt.Printf("✅ Bin создан: ID=%s, Name=%s\n", newBin.Id, newBin.Name)
	fmt.Println(req.Header.Get("X-Master-Key"))
	fmt.Println(string(body))

	// Возврат nil , если все удачно
	return nil
}

func CreateBin2(store *storage.Storage, file string, private bool, name string) error {
	// Проверяем входящий параметр name на заполненность
	if name == "" {

		return errors.New("ошибка: название bin не может быть пустым")
	}

	if file == "" {

		return errors.New("ошибка: название file не может быть пустым")
	}

	// Получение ключа
	key, err := GetKey()

	if err != nil {
		return err
	}

	// Создаем новый объект Bin используя конструктор структуры Bin

	newBin := bins.NewBin(uuid.New().String(), private, name)

	// Кладем новый объект в список объектов стурктуры Storage
	store.Bins = append(store.Bins, *newBin)

	binsData := [2]string{newBin.Id, name}
	store.List = append(store.List, binsData)

	// Сохраняем данные структуры Storage в файле data.json , тем самым обновляея локальный файл data.json после оьновления Storage
	// Метод Save возвращает ошибку , если что то пошло не так
	err = store.Save(file)
	if err != nil {

		return err
	}

	// Далее переводим обновленный объект Storage в Json формат
	postBody, _ := json.Marshal(store.Bins)

	// Создаем запрос (метод, адрес, побготовленные данные). Возвращает запрос или ошибку
	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(postBody))

	if err != nil {
		return err
	}

	// Добавляем в запрос заголовок, который отвечает за передачу данных в json формате
	req.Header.Set("Content-Type", "application/json")
	// Добавляем кастомный заголовок , который требует сервис jsonbin.io , чтобы передавать запросы
	req.Header.Set("X-Master-Key", key)

	// Выводим заголовки и их статусы
	fmt.Println("Request Headers:")
	for key, values := range req.Header {
		fmt.Printf(" %s: %v\n", key, values)
	}
	// Подготовка к отправке запроса. Если позьзуемся функцией NewRequest, то нужно инициализировать отправку через &http.Client{}
	client := &http.Client{}
	// Отправляем нужные данные . Возвращает ответ и ошибку
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}

	// инициализируем закрытие тела запроса перед завершением работы функции CreateBin
	defer req.Body.Close()
	// Читаем тело ответа, чтобы его ввыести в уведомление. Не обязательно для отправки методом POST, обязательно для GET
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	// Вывод уведомлений после отправки запроса
	fmt.Printf("✅ Bin создан: ID=%s, Name=%s\n", newBin.Id, newBin.Name)
	fmt.Println(req.Header.Get("X-Master-Key"))
	fmt.Println(string(body))

	// Возврат nil , если все удачно
	return nil
}

func GetKey() (string, error) {
	// Получаем ключ доступа из файла .env , который нужен для отправки запроса на сервис jsonbin.io
	key := os.Getenv("KEY")

	// Проверяем что ключ не пустой

	if key == "" {
		return "", errors.New("КЛЮЧ ПУСТОЙ")
	}

	return key, nil
}

func GetBin(id string) (*storage.Storage, error) {
	if id == "" {
		return nil, errors.New("ошибка: название id не может быть пустым")
	}

	// Получение ключа
	key, err := GetKey()

	if err != nil {
		return nil, err
	}

	// id 68f7aa52d0ea881f40b1239b

	req, err := http.NewRequest("GET", "https://api.jsonbin.io/v3/b/"+id, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var store *storage.Storage
	json.Unmarshal(body, &store)
	err = store.Save("data.json")
	if err != nil {
		return nil, err
	}

	return store, nil
}

func DeleteBin(id string) error {
	if id == "" {
		return errors.New("ошибка: название id не может быть пустым")
	}

	// Получение ключа
	key, err := GetKey()

	if err != nil {
		return err
	}

	// id 68f7aa52d0ea881f40b1239b

	req, err := http.NewRequest("DELETE", "https://api.jsonbin.io/v3/b/"+id, nil)

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	var newStorage storage.Storage
	var localStorage *storage.Storage
	var deletedStorage *storage.Storage
	// var newBins bins.BinList
	// var newList [][2]string

	fileContent, err := os.ReadFile("data.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileContent, &localStorage)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &deletedStorage)
	if err != nil {
		return err
	}

	deletedIds := make([]string, 0, len(deletedStorage.Bins))
	for _, v := range deletedStorage.Bins {
		deletedIds = append(deletedIds, v.Id)
	}

	for _, id := range deletedIds {
		newStorage.Bins = bins.BinList{{Id: "fdfdfd"}}
		newStorage.List = removeElem(localStorage.List, id)
	}

	err = newStorage.Save("data.json")
	if err != nil {
		return err
	}

	fmt.Println("Удалено")
	return nil
}

func removeBins(slice bins.BinList, target string) bins.BinList {
	for i, v := range slice {
		if v.Id == target {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func removeElem(slice [][2]string, target string) [][2]string {
	for i, v := range slice {
		if v[0] == target {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func ChangeBin(file string, id string) (*storage.Storage, error) {
	if id == "" {
		return nil, errors.New("ошибка: название id не может быть пустым")
	}

	// Получение ключа
	key, err := GetKey()

	if err != nil {
		return nil, err
	}

	// id 68f7aa52d0ea881f40b1239b\
	fileContent, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", "https://api.jsonbin.io/v3/b/"+id, bytes.NewBuffer(fileContent))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var store *storage.Storage
	json.Unmarshal(body, &store)
	err = store.Save("data.json")
	if err != nil {
		return nil, err
	}

	return store, nil
}
