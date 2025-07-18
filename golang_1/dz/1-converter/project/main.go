package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	for {
		fmt.Print("Какая у вас валюта? Введите USD, или RUB, или EUR: ")
		currentCurrency, err := getCurrentCurrencyInput()
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
			fmt.Print("Какую Валюту вы хотите получить?Введите USD, или EUR: ")
		case currentCurrency == "EUR":
			fmt.Print("Какую Валюту вы хотите получить?Введите USD, или RUB: ")
		default:
			fmt.Print("Какую Валюту вы хотите получить?Введите RUB, или EUR: ")
		}

		targetCurrency, err := getCurrentCurrencyInput()

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

		fmt.Print(calcCurrency(currentCurrency, quantity, targetCurrency))
		break
	}
}

func getCurrentCurrencyInput() (string, error) {
	var currentCurrency string

	fmt.Scan(&currentCurrency)

	if currentCurrency == "EUR" || currentCurrency == "RUB" || currentCurrency == "USD" {
		return currentCurrency, nil
	}
	return "", errors.New("INCORECT_PARAM")
}

func getQuantity() (float64, error) {

	var quantity string
	fmt.Print("Введите сумму: ")
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

func getUserInput() (string, float64, string) {
	var currentCurrency string
	var quantity float64
	var targetCurrency string

	fmt.Print("Какая у вас валюта?Введите: ")
	fmt.Scan(&currentCurrency)
	fmt.Print("Введите сумму: ")
	fmt.Scan(&quantity)
	fmt.Print("Какую валюты вы хотите?Введите: ")
	fmt.Scan(&targetCurrency)
	return currentCurrency, quantity, targetCurrency
}
