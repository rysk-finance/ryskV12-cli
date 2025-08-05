package main

type Transfer struct {
	Asset     string `json:"asset"`
	ChainID   int    `json:"chainId"`
	Amount    string `json:"amount"`
	IsDeposit bool   `json:"isDeposit"`
	Nonce     string `json:"nonce"`
	Signature string `json:"signature"`
}

type Quote struct {
	AssetAddress string `json:"assetAddress"`
	ChainID      int    `json:"chainId"`
	Expiry       int64  `json:"expiry"`
	IsPut        bool   `json:"isPut"`
	IsTakerBuy   bool   `json:"isTakerBuy"`
	Maker        string `json:"maker"`
	Nonce        string `json:"nonce"`
	Price        string `json:"price"`
	Quantity     string `json:"quantity"`
	Strike       string `json:"strike"`
	Signature    string `json:"signature"`
	ValidUntil   int64  `json:"validUntil"`
	USD          string `json:"usd"`
}

type QuoteNotification struct {
	RequestID string `json:"rfqId"`
	Asset     string `json:"assetAddress"`
	ChainID   int    `json:"chainId"`
	NewBest   string `json:"newBest"`
	Yours     string `json:"yours"`
}
