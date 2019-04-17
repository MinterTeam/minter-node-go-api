package responses

type EventsResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
	Result  struct {
		Events []event `json:"events"`
	} `json:"result"`
}

type event struct {
	Type  string     `json:"type"`
	Value eventValue `json:"value"`
}

type eventValue struct {
	Role            string `json:"role"`
	Coin            string `json:"coin"`
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ValidatorPubKey string `json:"validator_pub_key"`
}
