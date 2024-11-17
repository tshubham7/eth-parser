package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tshubham7/eth-parser/internal/mocks"
	"github.com/tshubham7/eth-parser/internal/pkg/constants"
	"github.com/tshubham7/eth-parser/internal/pkg/model"
)

func TestUnit_Subcribe(t *testing.T) {
	mockStore := mocks.NewStore(t)
	mockHttpClient := mocks.NewHttpClient(t)
	pUsecase := NewParserUsecase(mockStore, mockHttpClient)
	ctx := context.Background()

	t.Run("should subscribe when store returns nil", func(t *testing.T) {
		addr := "sometestaddress"
		mockStore.On("AddAddress", addr).Return(nil).Once()

		err := pUsecase.Subscribe(ctx, addr)
		assert.Nil(t, err)
	})
}

func TestUnit_GetCurrentBlock(t *testing.T) {
	mockStore := mocks.NewStore(t)
	mockHttpClient := mocks.NewHttpClient(t)
	pUsecase := NewParserUsecase(mockStore, mockHttpClient)
	ctx := context.Background()

	payload := model.EthClientRequestBody{
		Id:      1,
		Method:  constants.MethodBlockNumber,
		JsonRpc: "2.0",
	}
	respData := "{\"jsonrpc\":\"2.0\",\"result\":\"0x14398ca\",\"id\":83}"
	expectedBlockNumber := 21207242
	t.Run("should return block number when client returns the data", func(t *testing.T) {
		mockHttpClient.On("ExecutePostRequest", ctx, "", payload).Return(respData, nil).Once()

		block, err := pUsecase.GetCurrentBlockNumber(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expectedBlockNumber, block.BlockNum)
	})

	t.Run("should return error when client returns error", func(t *testing.T) {
		mockHttpClient.On("ExecutePostRequest", ctx, "", payload).Return("", errors.New("some error")).Once()

		_, err := pUsecase.GetCurrentBlockNumber(ctx)
		assert.Error(t, err)
	})
}

func TestUnit_GetTransactions(t *testing.T) {
	mockStore := mocks.NewStore(t)
	mockHttpClient := mocks.NewHttpClient(t)
	pUsecase := NewParserUsecase(mockStore, mockHttpClient)
	ctx := context.Background()
	transactions := []model.Transaction{
		{
			To:       "0xa7fb5ca286fc3fd67525629048a4de3ba24cba2e",
			From:     "0xa7fb5ca286fc3fd67525629048a4de3ba24cba2e",
			Value:    "0x0",
			Hash:     "0xe5b57ca9f9b1b65452e0c3c37dc46964ea9577c80c58fe965138ad60bfbdaedc",
			BlockNum: 23332,
		},
	}
	t.Run("should return transactions when stores returns the data", func(t *testing.T) {
		addr := "sometestaddress"
		mockStore.On("IsAddressSubscribed", addr).Return(true, nil).Once()
		mockStore.On("GetTransactions", addr).Return(transactions, nil).Once()

		actualTransactions, err := pUsecase.GetTransactions(ctx, addr)
		assert.Nil(t, err)
		assert.Equal(t, transactions, actualTransactions)
	})

	t.Run("should return error if address is not subscribed", func(t *testing.T) {
		addr := "unsubscribed-address"
		mockStore.On("IsAddressSubscribed", addr).Return(false, nil).Once()

		_, err := pUsecase.GetTransactions(ctx, addr)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrorCodeUnSubscribedAddress, err.Error())
	})
}

func TestUnit_GetAllTransactions(t *testing.T) {
	mockStore := mocks.NewStore(t)
	mockHttpClient := mocks.NewHttpClient(t)
	pUsecase := NewParserUsecase(mockStore, mockHttpClient)
	ctx := context.Background()
	transactions := []model.Transaction{
		{
			To:       "0xa7fb5ca286fc3fd67525629048a4de3ba24cba2e",
			From:     "0xa7fb5ca286fc3fd67525629048a4de3ba24cba2e",
			Value:    "0x0",
			Hash:     "0xe5b57ca9f9b1b65452e0c3c37dc46964ea9577c80c58fe965138ad60bfbdaedc",
			BlockNum: 23332,
		},
	}
	t.Run("should return transactions when stores returns the data", func(t *testing.T) {
		mockStore.On("GetAllTransactions").Return(transactions, nil).Once()

		actualTransactions, err := pUsecase.GetAllTransactions(ctx)
		assert.Nil(t, err)
		assert.Equal(t, transactions, actualTransactions)
	})
}
