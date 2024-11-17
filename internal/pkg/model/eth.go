package model

type Transaction struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Value    string `json:"value"`
	Hash     string `json:"hash"`
	BlockNum int    `json:"blockNumber"`
}

type Block struct {
	BlockNum int `json:"blockNumber"`
}
