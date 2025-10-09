package main

import (
	"demo/acc_app/account"
	"demo/acc_app/files"
	"demo/acc_app/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по Url",
	"3. Найти аккаунт по Login",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

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
		variant := promptData(
			menuVariants...,
		)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }

	}
	// createAccount()

}

func findAccountByUrl(vault *account.VaultWithDb) {
	// url по которому искать
	// Поисе
	// Вывод
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)

}

func findAccountByLogin(vault *account.VaultWithDb) {
	// url по которому искать
	// Поисе
	// Вывод
	login := promptData("Введите Login для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)

}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов не найдено!")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {

	// Получить url
	url := promptData("Введите URL для поиска")
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
func promptData(prompt ...string) string {
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
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		output.PrintError("Не верный формат URL или Логин")
		return
	}

	vault.AddAccount(*myAccount)

}
