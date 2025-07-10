package main

import "fmt"

func main() {
	const USDInEUR float64 = 0.94
	const USDInRUB float64 = 78
	const EURInRUB float64 = USDInRUB / USDInEUR

	fmt.Print(EURInRUB)
}

func getUserInput() (float64, string, string) {
	var quantity float64
	var currentCurrency string
	var targetCurrency string

	fmt.Print("Введите сумму: ")
	fmt.Scan(&quantity)
	fmt.Print("Какая у вас валюта?Введите: ")
	fmt.Scan(&currentCurrency)
	fmt.Print("Какую валюты вы хотите?Введите: ")
	fmt.Scan(&targetCurrency)
	return quantity, currentCurrency, targetCurrency
}

func calcCurrency(quantity float64, currentCurrency string, targetCurrency string) {}
