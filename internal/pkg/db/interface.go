package db

import "github.com/tshubham7/eth-parser/internal/pkg/model"

type Store interface {
	AddAddress(address string) error
	IsAddressSubscribed(address string) (bool, error)

	AddTransaction(tx model.Transaction) error
	GetTransactions(address string) ([]model.Transaction, error)
	GetAllTransactions() ([]model.Transaction, error)
}
