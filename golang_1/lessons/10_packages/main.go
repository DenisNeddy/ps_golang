package main

import (
	"demo/acc_app/account"
	"demo/acc_app/files"
	"fmt"
)

func main() {

	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Print("Не верный формат URL или Логин")
		return
	}

	myAccount.OutputPassword()
	files.WriteFile()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
