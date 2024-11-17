package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/tshubham7/eth-parser/internal/pkg/client"
	"github.com/tshubham7/eth-parser/internal/pkg/constants"
	"github.com/tshubham7/eth-parser/internal/pkg/db"
	"github.com/tshubham7/eth-parser/internal/pkg/model"
	"github.com/tshubham7/eth-parser/internal/pkg/utils"
)

type parser struct {
	store     db.Store
	apiClient client.HttpClient
}

func NewParserUsecase(store db.Store, apiClient client.HttpClient) ParserUsecase {
	return &parser{store: store, apiClient: apiClient}
}

func (p *parser) GetCurrentBlockNumber(ctx context.Context) (model.Block, error) {
	blockNum, err := p.fetchCurrentBlockNumber(ctx)
	if err != nil {
		return model.Block{}, err
	}
	return model.Block{BlockNum: blockNum}, nil
}

func (p *parser) Subscribe(ctx context.Context, address string) error {
	return p.store.AddAddress(address)
}

func (p *parser) GetTransactions(ctx context.Context, address string) ([]model.Transaction, error) {
	subscribed, err := p.store.IsAddressSubscribed(address)
	if err != nil {
		return nil, err
	}
	if !subscribed {
		return nil, errors.New(constants.ErrorCodeUnSubscribedAddress)
	}
	return p.store.GetTransactions(address)
}

func (p *parser) GetAllTransactions(ctx context.Context) ([]model.Transaction, error) {
	return p.store.GetAllTransactions()
}

func (p *parser) Process(ctx context.Context, closeCh <-chan struct{}) {
	log := utils.GetCurrentLogger(ctx)
	for {
		select {
		case <-closeCh:
			log.Infof("Stopping processing parsing...")
			return
		default:
			block, err := p.fetchCurrentBlockNumber(ctx)
			if err != nil {
				log.Errorf("Error fetching current block: %v", err)
				continue
			}
			if err := p.parseBlock(ctx, block); err != nil {
				log.Errorf("Error parsing block: %v", err)
			}
			time.Sleep(time.Second * 1)
		}
	}
}

func (p *parser) parseBlock(ctx context.Context, blockNumber int) error {
	transactions, err := p.fetchBlock(ctx, blockNumber)
	if err != nil {
		return err
	}

	for _, tx := range transactions {
		fromSubscribed, _ := p.store.IsAddressSubscribed(tx.From)
		toSubscribed, _ := p.store.IsAddressSubscribed(tx.To)
		if fromSubscribed || toSubscribed {
			p.store.AddTransaction(tx)
		}
	}

	return nil
}
