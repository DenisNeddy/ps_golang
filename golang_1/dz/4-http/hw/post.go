// **Описание**: Создайте функцию, которая принимает данные
// для отправки и URL, выполняет POST запрос с JSON телом
// и возвращает декодированный ответ.
// **Входные данные**: data (map[string]interface{}) - данные
// для отправки в JSON, url (строка) - URL для POST запроса
// **Выходные данные**: map[string]interface{} - декодированный
// JSON ответ, error - ошибка выполнения запроса
// **Ограничения**: data не должен быть nil, url должен быть
// валидным HTTP/HTTPS адресом
// **Примеры**:
// Input: data = map[string]interface{}{"username": "john",
// "email": "john@example.com"}, url = "https://httpbin.org/post"
// Output: map[string]interface{}{"json":
// map[string]interface{}{"username": "john",
// "email": "john@example.com"}, "url":
// "https://httpbin.org/post"}, nil
//
// Input: data = map[string]interface{}{"message": "Hello World",
// "priority": 1}, url = "https://api.example.com/messages"
// Output: map[string]interface{}{"id": 12345,
// "status": "created"}, nil

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func sendPostRequest(data map[string]interface{}, urlStr string) (map[string]interface{}, error) {

	if data == nil {
		return nil, errors.New("data is nil")
	}

	if urlStr == "" {
		return nil, errors.New("url is empty")
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, errors.New("invalid URL scheme")
	}

	postBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New("received non-2xx response")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var requestMap map[string]interface{}

	err = json.Unmarshal(body, &requestMap)
	if err != nil {
		return nil, err
	}
	return requestMap, nil
}