package usecase

import (
	"context"

	"github.com/tshubham7/eth-parser/internal/pkg/model"
)

type ParserUsecase interface {
	GetCurrentBlockNumber(ctx context.Context) (model.Block, error)
	Subscribe(ctx context.Context, address string) error
	GetTransactions(ctx context.Context, address string) ([]model.Transaction, error)

	Process(ctx context.Context, closeCh <-chan struct{})

	// admin
	GetAllTransactions(ctx context.Context) ([]model.Transaction, error)
}
