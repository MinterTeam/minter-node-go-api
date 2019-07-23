package responses

import "encoding/json"

type TransactionResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Error   *ErrorData  `json:"error"`
	Result  Transaction `json:"result"`
}

type TransactionsResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Error   *ErrorData    `json:"error"`
	Result  []Transaction `json:"result"`
}

type Transaction struct {
	Hash        string             `json:"hash"`
	Height      string             `json:"height"`
	From        string             `json:"from"`
	Type        uint8              `json:"type"`
	Nonce       string             `json:"nonce"`
	GasPrice    uint32             `json:"gas_price"`
	GasCoin     string             `json:"gas_coin"`
	GasUsed     string             `json:"gas_used"`
	Gas         string             `json:"gas"`
	Payload     string             `json:"payload"`
	ServiceData string             `json:"service_data"`
	RawTx       string             `json:"raw_tx"`
	Log         *string            `json:"log"`
	Data        json.RawMessage    `json:"data"`
	IData       interface{}        `json:"-"`
	Tags        *map[string]string `json:"tags"`
}

type SendTransactionResponse struct {
	Jsonrpc string                 `json:"jsonrpc"`
	ID      string                 `json:"id"`
	Error   *ErrorData             `json:"error"`
	Result  *SendTransactionResult `json:"result"`
}

type SendTransactionResult struct {
	Code int32  `json:"code"`
	Data string `json:"data"`
	Log  string `json:"log"`
	Hash string `json:"hash"`
}
