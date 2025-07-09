package main

import "fmt"

func main() {
	const USDInEUR float64 = 0.94
	const USDInRUB float64 = 78
	const EURInRUB float64 = USDInRUB / USDInEUR

	fmt.Print(EURInRUB)
}
