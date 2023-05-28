package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionConfirmed string = "confirmed"
	TransactionError     string = "error"
)

type TransactionRepositoryInterface interface {
	RegisterTransaction(transaction *Transaction) error
	SaveCompletedTransaction(transaction *Transaction) error
	FindTransaction(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"-"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}

	if t.Amount <= 0 {
		return errors.New("invalid amount transaction")
	}

	if t.Status != TransactionConfirmed && t.Status != TransactionCompleted && t.Status != TransactionPending && t.Status != TransactionError {
		return errors.New("invalid status transaction")
	}

	if t.PixKeyTo.AccountId == t.AccountFrom.ID {
		return errors.New("transaction to the same account")
	}

	return nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.Description = description
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func NewTransaction(account *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom:       account,
		Amount:            amount,
		PixKeyTo:          pixKeyTo,
		Status:            TransactionPending,
		Description:       description,
		CancelDescription: "",
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return &transaction, err
	}

	return &transaction, nil
}
