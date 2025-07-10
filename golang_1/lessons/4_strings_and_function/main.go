package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {

	fmt.Println("__ Калькулятор индекса массы тела __")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела:  %.0f", imt)
	fmt.Printf(result)

}

func calculateIMT(userKg, userHeight float64) float64 {

	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост  в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Printf("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
