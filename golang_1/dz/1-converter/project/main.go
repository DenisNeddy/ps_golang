package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	for {
		fmt.Print("1. Ввод исходной валюты(USD,RUB,EUR): ")
		currentCurrency, err := getCurrencyInput()
		if err != nil {

			fmt.Println("Не корректный тип Валюты")
			fmt.Println("Попробуйте ввести заново!")
			continue
		}

		quantity, err := getQuantity()

		if err != nil {
			fmt.Println("--------------------")
			fmt.Println("Введите правильный тип значения количества средств")
			fmt.Println("Попробуйте еще раз!")
			fmt.Println("--------------------")

			isRepeat := checkRepeatCalc()
			if isRepeat {
				break
			}
			continue
		}

		switch {
		case currentCurrency == "RUB":
			fmt.Print("3. Ввод целевой валюты(USD,EUR): ")
		case currentCurrency == "EUR":
			fmt.Print("3. Ввод целевой валюты(USD,RUB): ")
		default:
			fmt.Print("3. Ввод целевой валюты(RUB,EUR): ")
		}

		targetCurrency, err := getCurrencyInput()

		if targetCurrency == currentCurrency {
			fmt.Println("--------------------")
			fmt.Println("Желаемая валюта не должна совподать с текущей")
			fmt.Println("Попробуйте еще раз!")
			fmt.Println("--------------------")

			isRepeat := checkRepeatCalc()
			if isRepeat {
				break
			}
			continue
		}

		if err != nil {
			fmt.Println("--------------------")
			fmt.Println("Не корректный тип Валюты")
			fmt.Println("Попробуйте еще раз!")
			fmt.Println("--------------------")

			isRepeat := checkRepeatCalc()
			if isRepeat {
				break
			}
			continue
		}

		// result := fmt.Sprint(calcCurrency(currentCurrency, quantity, targetCurrency))
		fmt.Printf("Результат конвертации: %.2f", calcCurrency(currentCurrency, quantity, targetCurrency))
		fmt.Print(" ", targetCurrency)
		break
	}
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
	fmt.Print("2. Ввод числа: ")
	fmt.Scan(&quantity)
	money, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		return 0.0, errors.New("INCORECT_QUANTITY")
	}
	return money, nil
}

func calcCurrency(currentCurrency string, quantity float64, targetCurrency string) float64 {
	const USDInEUR float64 = 0.94
	const USDInRUB float64 = 78
	const EURInRUB float64 = USDInRUB / USDInEUR
	var result float64

	if currentCurrency == "RUB" {
		if targetCurrency == "EUR" {
			result = quantity / EURInRUB
		}
		if targetCurrency == "USD" {
			result = quantity / USDInRUB
		}
	}

	if currentCurrency == "USD" {
		if targetCurrency == "RUB" {
			result = USDInRUB * quantity
		}
		if targetCurrency == "EUR" {
			result = USDInEUR * quantity
		}
	}

	if currentCurrency == "EUR" {
		if targetCurrency == "RUB" {
			result = EURInRUB * quantity
		}
		if targetCurrency == "USD" {
			result = quantity / USDInEUR
		}
	}

	return result
}

func checkRepeatCalc() bool {
	var userChoise string
	fmt.Print("Вы хотите выйти? (y/n): ")
	fmt.Scan(&userChoise)

	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
