package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWitchTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]

	}
	acc.password = string(res)
}

// 1. Если логина нет, то ошибка
// 2. Если нет пароля, то генерим
// func newAccount(login, password, urlString string) (*account, error) {
// 	// Валедировать можно входящие данные до их возвращения
// 	if login == "" {
// 		return nil, errors.New("INVALID_LOGIN")
// 	}
// 	_, err := url.ParseRequestURI(urlString)

// 	if err != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}

// 	newAcc := &account{
// 		url:      urlString,
// 		login:    login,
// 		password: password,
// 	}

// 	if password == "" {
// 		newAcc.generatePassword(12)
// 	}

// 	return newAcc, nil

// }

func newAccountWithTimeStamp(login, password, urlString string) (*accountWitchTimeStamp, error) {
	// Валедировать можно входящие данные до их возвращения
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWitchTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUWXYZ0123456789-*!")

func main() {

	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := newAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Print("Не верный формат URL или Логин")
		return
	}
	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
