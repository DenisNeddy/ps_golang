// // массивы это структура занимающая фиксированное место в памяти, а слайс это ссылка на этот массив. или его часть.
// // По этой причиние слайс легковестный, при изменение слайса, меняется данные зависимого массива с которого был сделан слайс
// // массивы - это некоторая последовательность однотипных значений ФИКСИРОВАННОЙ ДЛИННЫ

package main

import "fmt"

// В цикле спрашиваем ввод транзакций: -10 , 10, 40.5
// Добавлять каждую в массив транзаций
// Вывести массив
// Вывести сумму баланса в консоль

func main() {

	// // transaction := [6]int{1, 2, 3, 4, 5, 6}
	// // Слайсы
	// // patrial := transaction[1:4]

	// transaction := [6]int{1, 2, 3, 4, 5, 6}
	// transationNew := transaction[:]
	// fmt.Print(transationNew)

	// // объявление слайса
	// transactions1 := []int{1, 2, 3}            // у него нет фиксированной длинны в отличае от мпссива
	// transactions1 = append(transactions1, 100) // напрямую добавление элемента

	// newTransaction := append(transactions1, 100) // Добавление элемента слайсу

	tr := make([]string, 0, 2)
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "1")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "2")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "3")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "4")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "5")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "6")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "4")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "5")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "6")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "6")
	fmt.Println(tr)

	transactions := []float64{}

	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
	}

	fmt.Printf("Ваш баланс: %.2f", calculateBalance(transactions))

}

func scanTransaction() float64 {
	var transaction float64

	fmt.Print("Введите тразакцию (n для выхода): ")
	fmt.Scan(&transaction)

	return transaction
}

func calculateBalance(transactions []float64) float64 {
	var total float64
	for _, value := range transactions {
		total += value
	}
	return total
}
