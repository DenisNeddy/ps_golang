package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password`
	Url       string    `json: "url`
	CreatedAt time.Time `json: createdAt`
	UpdatedAt time.Time `json: updatedAt`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUWXYZ0123456789-*!")

func (acc *Account) Output() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]

	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	// Валедировать можно входящие данные до их возвращения
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Login:     login,
		Password:  password,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil

}
