package model

import "github.com/tshubham7/eth-parser/internal/pkg/constants"

type EthClientRequestBody struct {
	Id      int                    `json:"id"`
	JsonRpc string                 `json:"jsonrpc"`
	Method  constants.EthApiMethod `json:"method"`
	Params  []any                  `json:"params"`
}
