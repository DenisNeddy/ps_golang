package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operationStatus string
	var numbers = []int{}
	for {

		for {
			fmt.Print("Введите нужный тип операции(SUM,AVG,MED): ")
			result, err := getOperationStatus()

			if err != nil {
				fmt.Println("Не корректный тип операции")
				fmt.Println("Попробуйте ввести заново!")
				continue
			}
			operationStatus = result
			break
		}
		fmt.Println("Ваш статус операции ", operationStatus)

		for {
			fmt.Print("Введите нужный числа: ")
			result, err := getNumbers()

			if err != nil {
				fmt.Println("Не корректный тип данных")
				fmt.Println("Попробуйте ввести заново!")
				continue
			}
			numbers = result
			break
		}

		fmt.Println("Введенные вами числа", numbers)

		switch operationStatus {
		case "SUM":
			fmt.Println("Результат суммирования: ", getSum(numbers))
		case "AVG":
			fmt.Println("Результат вычисления среднего значения: ", getAverage(numbers))
		case "MED":
			fmt.Println("Результат вычисления медианы: ", getMedian(numbers))
		}

		var answer string
		fmt.Print("Хотите продолжить? (y/n): ")
		fmt.Scan(&answer)
		if strings.ToLower(answer) != "y" {
			break
		}
	}

}

func getOperationStatus() (string, error) {
	operations := [3]string{"SUM", "AVG", "MED"}
	var op string
	fmt.Scan(&op)
	for _, value := range operations {
		if op == value {
			return value, nil
		}

	}
	return "", errors.New("INCORECT_PARAM")

}

func getNumbers() ([]int, error) {
	var numbersString string
	fmt.Scan(&numbersString)
	list := strings.Split(numbersString, ",")
	numbers := []int{}

	for _, value := range list {
		num, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println("Ошибка при преобразовании строки:", err)
			return nil, errors.New("INCORECT INT")
		}
		numbers = append(numbers, num)

	}

	return numbers, nil

}

func getSum(numbers []int) int {
	var total int
	for _, value := range numbers {
		total += value
	}

	return total
}

func getAverage(numbers []int) float64 {
	sortNumbers := make([]int, len(numbers))
	copy(sortNumbers, numbers)
	numbersLength := len(sortNumbers)

	total := getSum(numbers)

	return float64(total) / float64(numbersLength)
}

func getMedian(numbers []int) float64 {
	sortNumbers := make([]int, len(numbers))
	copy(sortNumbers, numbers)
	numbersLength := len(sortNumbers)
	var median float64

	sort.Ints(sortNumbers)

	if numbersLength%2 != 0 {
		median = float64(sortNumbers[numbersLength/2])
	} else {

		median = (float64(sortNumbers[numbersLength/2-1]) + float64(sortNumbers[numbersLength/2])) / 2.0
	}

	return float64(median)
}
