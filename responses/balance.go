package responses

type Balance struct {
	Balance          map[string]string `json:"balance"`
	TransactionCount string            `json:"transaction_count"`
	Address          string            `json:"address"`
}

type BalancesResponse struct {
	Response
	Result []Balance `json:"result"`
}
