package usecase

import (
	"context"
	"encoding/json"
	"os"

	"github.com/tshubham7/eth-parser/internal/pkg/constants"
	"github.com/tshubham7/eth-parser/internal/pkg/model"
	"github.com/tshubham7/eth-parser/internal/pkg/utils"
)

func (p *parser) fetchCurrentBlockNumber(ctx context.Context) (int, error) {
	log := utils.GetCurrentLogger(ctx)
	log.Info("fetching recent block number...")

	url := os.Getenv(constants.EnvEthEndpoint)
	payload := model.EthClientRequestBody{
		Id:      1,
		Method:  constants.MethodBlockNumber,
		JsonRpc: "2.0",
	}
	respBody, err := p.apiClient.ExecutePostRequest(ctx, url, payload)
	if err != nil {
		log.Errorf("failed to execute post api call: %v", err)
		return 0, err
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(respBody), &result); err != nil {
		log.Errorf("failed to unmarshal response data: %v", err)
		return 0, err
	}

	blockHex := result["result"].(string)
	return utils.HexToInt(blockHex)
}

func (p *parser) fetchBlock(ctx context.Context, blockNumber int) ([]model.Transaction, error) {
	log := utils.GetCurrentLogger(ctx)
	log.Info("fetching block details...")

	blockHex := utils.IntToHex(blockNumber)
	url := os.Getenv(constants.EnvEthEndpoint)
	payload := model.EthClientRequestBody{
		Id:      1,
		Method:  constants.MethodGetBlockByNumber,
		JsonRpc: "2.0",
		Params:  []any{blockHex, true},
	}
	respBody, err := p.apiClient.ExecutePostRequest(ctx, url, payload)
	if err != nil {
		log.Errorf("failed to execute post api call: %v", err)
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(respBody), &result); err != nil {
		log.Errorf("failed to unmarshal response data: %v", err)
		return nil, err
	}

	blockData := result["result"].(map[string]interface{})
	txs := blockData["transactions"].([]interface{})
	transactions := []model.Transaction{}
	for _, tx := range txs {
		txMap := tx.(map[string]interface{})
		if txMap["to"] == nil || txMap["from"] == nil {
			continue
		}
		transactions = append(transactions, model.Transaction{
			From:     txMap["from"].(string),
			To:       txMap["to"].(string),
			Value:    txMap["value"].(string),
			Hash:     txMap["hash"].(string),
			BlockNum: blockNumber,
		})
	}
	return transactions, nil
}
