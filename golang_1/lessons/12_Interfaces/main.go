package main

import (
	"demo/acc_app/account"
	"demo/acc_app/files"
	"demo/acc_app/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// 1. Создать аккаунт
	// 2. Найти аккаунт
	// 3. Удалить аккаунт
	// 4. Выход
	color.Cyan("__Менеджер паролей__")
	vault := account.NewVault(files.NewJsonDb("data.json"))
	// vault := account.NewVault(cloud.NewCloudDb("ad.ru"))
Menu:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}

	}
	// createAccount()

}

func findAccount(vault *account.VaultWithDb) {
	// url по которому искать
	// Поисе
	// Вывод
	url := promptData([]string{"Введите URL для поиска"})
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено!")
	}
	for _, account := range accounts {
		account.Output()
	}

}

func deleteAccount(vault *account.VaultWithDb) {

	// Получить url
	url := promptData([]string{"Введите URL для поиска"})
	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
	// Вызвать метод vault
	// Удалено или нет
}

// Функция принимала слайс любого типа
// Выводит строкой каждый элемент, а последний через Printf и к последнему добавляется ":"
func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	var res string
	fmt.Scanln(&res)
	return res
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		output.PrintError("Не верный формат URL или Логин")
		return
	}

	vault.AddAccount(*myAccount)

}
