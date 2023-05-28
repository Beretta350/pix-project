package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixkey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountId string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pk *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pk)
	if err != nil {
		return err
	}

	if pk.Kind != "email" && pk.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pk.Kind != "active" && pk.Kind != "inactive" {
		return errors.New("invalid status")
	}

	return nil
}

func NewPixKey(kind, key, accountId, status string, account *Account) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return &pixKey, err
	}

	return &pixKey, nil
}
