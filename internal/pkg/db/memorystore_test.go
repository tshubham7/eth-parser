package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tshubham7/eth-parser/internal/pkg/model"
)

func TestIntegration_AddAddress(t *testing.T) {
	store := newMemoryStore()

	address := "testaddress"

	t.Run("should store the address", func(t *testing.T) {
		err := store.AddAddress(address)
		assert.Nil(t, err)
		assert.Equal(t, true, store.addresses[address])
	})
}

func TestIntegration_IsAddressSubscribed(t *testing.T) {
	store := newMemoryStore()
	address := "testaddress"

	t.Run("should returned true", func(t *testing.T) {
		err := store.AddAddress(address)
		assert.Nil(t, err)

		subscribed, err := store.IsAddressSubscribed(address)
		assert.Nil(t, err)
		assert.Equal(t, true, subscribed)
	})
}

func TestIntegration_AddTransaction(t *testing.T) {
	store := newMemoryStore()

	transaction := model.Transaction{
		To:       "0xa7fb5ca286fc3fd67525629048a4de3ba24cba2e",
		From:     "0xa7fb5ca286fc3fd67525629048a4de3ba24cba2e",
		Value:    "0x0",
		Hash:     "0xe5b57ca9f9b1b65452e0c3c37dc46964ea9577c80c58fe965138ad60bfbdaedc",
		BlockNum: 23332,
	}
	store.addresses[transaction.To] = true
	store.addresses[transaction.From] = true
	t.Run("should add a transaction", func(t *testing.T) {
		err := store.AddTransaction(transaction)
		assert.Nil(t, err)

		_, found := store.transactions[transaction.To]
		assert.Equal(t, true, found)

		_, found = store.transactions[transaction.From]
		assert.Equal(t, true, found)
	})
}

func TestIntegration_GetTransaction(t *testing.T) {
	store := newMemoryStore()

	transaction := model.Transaction{
		To:       "0xa7fb5ca286fc3fd67525629048a4de3ba24cba2e",
		From:     "0xa7fb5ca286fc3fd67525629048a4de3ba24cba2e",
		Value:    "0x0",
		Hash:     "0xe5b57ca9f9b1b65452e0c3c37dc46964ea9577c80c58fe965138ad60bfbdaedc",
		BlockNum: 23332,
	}
	store.transactions[transaction.To] = []model.Transaction{transaction}
	t.Run("should returned the data", func(t *testing.T) {
		actualTransactions, err := store.GetTransactions(transaction.To)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(actualTransactions))
	})
}
