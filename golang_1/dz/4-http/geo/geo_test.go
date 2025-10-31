package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка,  expected результат, данные для функции
	city := "Moscow"

	expected := geo.GeoData{
		City: "Moscow",
	}

	// Act - выаоняем функцию

	got, err := geo.GetMyLocation(city)

	if err != nil {
		t.Error(err)
	}

	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}

	// Assert - поверка результата с expected
}

func TestGetmyLocationNoCity(t *testing.T) {
	city := "L"

	_, err := geo.GetMyLocation(city)

	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrorNoCity, err)
	}
}
