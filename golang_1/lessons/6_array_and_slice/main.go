package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recover ", r)
		}
	}()

	fmt.Println("__ Калькулятор индекса массы тела __")
	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			panic("Не заданы параметры расчёта")
		}
		outputResult(IMT)

		isRepeateCalc := checkRepeatCalculation()

		if !isRepeateCalc {
			break
		}

	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела:  %.0f", imt)
	fmt.Println(result)

	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userKg, userHeight float64) (float64, error) {
	if userHeight <= 0 || userKg <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}

	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчёт (y/n): ")
	fmt.Scan(&userChoise)

	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false

}
