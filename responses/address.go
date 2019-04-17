package responses

type AddressResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
	Result  Balance    `json:"result"`
}
