package account

import (
	"demo/acc_app/output"
	"encoding/json"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json: updatedAt`
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDb {

	file, err := db.Read()
	if err != nil {

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		output.PrintError("Не удалось разобрать файл data.json")

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) FindAccountByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)

		}

	}

	return accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue

		}
		isDeleted = true

	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *VaultWithDb) save() {

	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError("Не удалось предобразовать")
	}

	vault.db.Write(data)

}
