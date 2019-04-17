package responses

type Balance struct {
	Balance          map[string]string `json:"balance"`
	TransactionCount string            `json:"transaction_count"`
	Address          string            `json:"address"`
}

type BalancesResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
	Result  []Balance  `json:"result"`
}
