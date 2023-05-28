package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (a *Account) isValid() error {
	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}

	return nil
}

func NewAccount(owner string, number string, bank *Bank) (*Account, error) {
	account := Account{
		OwnerName: owner,
		Number:    number,
		Bank:      bank,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return &account, err
	}

	return &account, nil
}
