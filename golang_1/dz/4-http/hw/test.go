// package main

// import (
// 	"fmt"
// 	"net/url"
// )

// func buildURLWithParams(baseURL string, params map[string]string) string {
// 	// **Описание**: Создайте функцию, которая принимает базовый URL
// 	// и map с параметрами, а затем формирует полный URL с query
// 	// параметрами.
// 	// **Входные данные**: baseURL (строка) - базовый URL,
// 	// params (map[string]string) - параметры для добавления в query
// 	// **Выходные данные**: строка - полный URL с закодированными
// 	// query параметрами
// 	// **Ограничения**: baseURL не пустой, params может быть пустым
// 	// **Примеры**:
// 	// Input: baseURL = "https://api.example.com/search",
// 	// params = map[string]string{"q": "golang", "limit": "10"}
// 	// Output: "https://api.example.com/search?limit=10&q=golang"
// 	//
// 	// Input: baseURL = "https://api.weather.com/current",
// 	// params = map[string]string{"city": "New York", "units": "metric"}
// 	// Output: "https://api.weather.com/current?city=New+York&units=metric"
//     myUrl, err := url.Parse(baseURL)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return ""
// 	}

//   param := url.Values{}
//     for k, v := range params {
//       param.Add(k, fmt.Sprint(v))
//     }
//       myUrl.RawQuery = param.Encode()
// 	return myUrl.String()
// }

// func main() {
// 	params := map[string]string{
// 		"q":     "golang",
// 		"limit": "10",
// 	}
// 	result := buildURLWithParams("https://api.example.com/search", params)
// 	fmt.Println(result)
// }

// package main

// import (
// 	"io"
// 	"strings"
// )

// func countChunks(text string) int {
// 	// **Описание**: Создайте функцию, которая читает данные из
// 	// строки порциями по 8 байт и возвращает количество
// 	// прочитанных порций.
// 	// **Входные данные**: text (строка) - исходный текст для
// 	// чтения
// 	// **Выходные данные**: int - количество полных порций по
// 	// 8 байт, которые удалось прочитать
// 	// **Ограничения**: text может быть пустой строкой,
// 	// длина text <= 1000 символов
// 	// **Примеры**:
// 	// Input: text = "Hello, World! This is a test string."
// 	// Output: 4
// 	//
// 	// Input: text = "Go"
// 	// Output: 0

// 	if len(text) == 0 {
// 		return 0
// 	}

// 	block := make([]byte, 8)
// 	r := strings.NewReader(text)

// 	var counter int
// 	for {
// 		n, err := r.Read(block)
// 		if err == io.EOF {
// 			break
// 		}
// 		if n == 8 {

// 			counter++
// 		}

// 	}
// 	return counter
// }

// package main

// import (
// 	"encoding/json"
// )

// func serializeToJSON(data map[string]interface{}) (string, error) {
// 	// **Описание**: Создайте функцию, которая принимает map с
// 	// данными и возвращает JSON строку.
// 	// **Входные данные**: data (map[string]interface{}) - данные
// 	// для сериализации
// 	// **Выходные данные**: string - JSON строка, error - ошибка
// 	// сериализации
// 	// **Ограничения**: map может содержать строки, числа, булевы
// 	// значения и nil
// 	// **Примеры**:
// 	// Input: data = map[string]interface{}{"name": "John",
// 	// "age": 30, "active": true}
// 	// Output: "{\"active\":true,\"age\":30,\"name\":\"John\"}",
// 	// nil
// 	//
// 	// Input: data = map[string]interface{}{"city": "Moscow",
// 	// "population": 12500000}
// 	// Output: "{\"city\":\"Moscow\",\"population\":12500000}",
// 	// nil

//     jsonData, err := json.Marshal(data)

//     if err != nil {
//       return "", err
//     }

// 	return string(jsonData), nil
// }