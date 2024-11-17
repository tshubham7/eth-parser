package db

import "github.com/tshubham7/eth-parser/internal/pkg/model"

type MemoryStore struct {
	addresses    map[string]bool
	transactions map[string][]model.Transaction
}

func newMemoryStore() *MemoryStore {
	return &MemoryStore{
		addresses:    make(map[string]bool),
		transactions: make(map[string][]model.Transaction),
	}
}

func (s *MemoryStore) AddAddress(address string) error {
	if _, exists := s.addresses[address]; exists {
		return nil
	}
	s.addresses[address] = true
	return nil
}

func (s *MemoryStore) IsAddressSubscribed(address string) (bool, error) {
	return s.addresses[address], nil
}

func (s *MemoryStore) AddTransaction(tx model.Transaction) error {
	if s.addresses[tx.From] {
		s.transactions[tx.From] = append(s.transactions[tx.From], tx)
	}
	if s.addresses[tx.To] {
		s.transactions[tx.To] = append(s.transactions[tx.To], tx)
	}
	return nil
}

func (s *MemoryStore) GetTransactions(address string) ([]model.Transaction, error) {
	return s.transactions[address], nil
}

// TODO: remove duplicate records
func (s *MemoryStore) GetAllTransactions() ([]model.Transaction, error) {
	var transactions []model.Transaction
	for _, v := range s.transactions {
		transactions = append(transactions, v...)
	}
	return transactions, nil
}
