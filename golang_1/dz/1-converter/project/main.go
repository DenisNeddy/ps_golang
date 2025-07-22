package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var currentCurrency string
	var quantity float64
	var targetCurrency string

	for {
		fmt.Print("1. Ввод исходной валюты(USD,RUB,EUR): ")
		result, err := getCurrencyInput()

		if err != nil {
			fmt.Println("Не корректный тип Валюты")
			fmt.Println("Попробуйте ввести заново!")
			continue
		}
		currentCurrency = result
		break
	}

	for {
		fmt.Print("2. Ввод числа: ")

		result, err := getQuantity()
		if err != nil {
			fmt.Println("--------------------")
			fmt.Println("Введите правильный тип значения количества средств")
			fmt.Println("Попробуйте еще раз!")
			fmt.Println("--------------------")
			continue
		}
		quantity = result
		break
	}

	for {
		switch {
		case currentCurrency == "RUB":
			fmt.Print("3. Ввод целевой валюты(USD,EUR): ")
		case currentCurrency == "EUR":
			fmt.Print("3. Ввод целевой валюты(USD,RUB): ")
		default:
			fmt.Print("3. Ввод целевой валюты(RUB,EUR): ")
		}

		result, err := getCurrencyInput()

		if result == currentCurrency {
			fmt.Println("--------------------")
			fmt.Println("Желаемая валюта не должна совподать с текущей")
			fmt.Println("Попробуйте еще раз!")
			fmt.Println("--------------------")
			continue
		}

		if err != nil {
			fmt.Println("--------------------")
			fmt.Println("Не корректный тип Валюты")
			fmt.Println("Попробуйте еще раз!")
			fmt.Println("--------------------")
			continue
		}

		targetCurrency = result
		break
	}

	fmt.Printf("Результат конвертации: %.2f", calcCurrency(currentCurrency, quantity, targetCurrency))
	fmt.Print(" ", targetCurrency)

}

func getCurrencyInput() (string, error) {
	var currentCurrency string

	fmt.Scan(&currentCurrency)

	if currentCurrency == "EUR" || currentCurrency == "RUB" || currentCurrency == "USD" {
		return currentCurrency, nil
	}
	return "", errors.New("INCORECT_PARAM")
}

func getQuantity() (float64, error) {

	var quantity string

	fmt.Scan(&quantity)
	money, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		return 0.0, errors.New("INCORECT_QUANTITY")
	}
	return money, nil
}

func calcCurrency(currentCurrency string, quantity float64, targetCurrency string) float64 {

	rates := map[string]float64{
		"USD": 78.0,
		"RUB": 1.0,
		"EUR": 91.0,
	}

	rub := quantity * rates[currentCurrency]
	return rub / rates[targetCurrency]

}
