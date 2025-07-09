package main // пакет main это точка входа в программу

import (
	"fmt"
	"math"
) // необходимых пакетов

func main() {
	// var userHeight = 1.9
	// var userKg float64 = 95
	// var IMT = userKg / math.Pow(userHeight, 2)
	// fmt.Print(IMT)

	// var userHeight, userKg = 1.9, 95

	const IMTPower float64 = 2

	userHeight := 1.9
	userKg := 95.0
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)

} // вся логика пишется внутри функции main

// ИМТ = вес / рост ^ 2
