// **Описание**: Создайте функцию, которая принимает данные
// для геолокации и интегрируется с внешним REST API для
// получения координат по названию города.
//
// **Входные данные**: cityName (строка) - название города
// для поиска координат
//
// **Выходные данные**: latitude (float64) - широта,
// longitude (float64) - долгота, error - ошибка выполнения
// запроса
//
// **Ограничения**: cityName не должен быть пустым, должен
// содержать только буквы и пробелы
//
// **Примеры**:
// Input: cityName = "London"
// Output: latitude = 51.5074, longitude = -0.1278, error = nil
//
// Input: cityName = "Tokyo"
// Output: latitude = 35.6762, longitude = 139.6503, error = nil
import (
	"errors"
	"regexp"
	"encoding/json"

	"io"
	"net/http"
	"strconv"
)

type Data struct {
	Lon string `json:"lon"`
	Lat string `json:"lat"`
}

func getCoordinates(cityName string) (float64, float64, error) {
	if cityName == "" {

		return 0.0, 0.0, errors.New("Param is empty")
	}

	validate := `^[a-zA-Z ]+$`
	matched, err := regexp.MatchString(validate, cityName)
	if err != nil {
		return 0.0, 0.0, err
	}

	if !matched {
		return 0.0, 0.0, errors.New("incorrect symbols in cityName")
	}

	resp, err := http.Get("https://nominatim.openstreetmap.org/search?q=" + cityName + "&format=json&limit=1")

	if err != nil {

		return 0.0, 0.0, err
	}

	if resp.StatusCode != 200 {
		return 0.0, 0.0, errors.New("NOT200")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return 0.0, 0.0, err
	}

	var geo []Data

	err = json.Unmarshal(body, &geo)
	if err != nil {
		return 0.0, 0.0, err
	}

	if len(geo) > 0 {
		lat, err := strconv.ParseFloat(geo[0].Lat, 64)
		if err != nil {
			return 0.0, 0.0, err
		}
		lon, err := strconv.ParseFloat(geo[0].Lon, 64)
		if err != nil {
			return 0.0, 0.0, err
		}
		return lat, lon, nil

	}

	return 0.0, 0.0, errors.New("response is empty")

}