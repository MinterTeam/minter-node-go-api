package responses

type CoinInfoResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
	Result  struct {
		Name           string `json:"name"`
		Symbol         string `json:"symbol"`
		Volume         string `json:"volume"`
		Crr            string `json:"crr"`
		ReserveBalance string `json:"reserve_balance"`
	} `json:"result"`
}
