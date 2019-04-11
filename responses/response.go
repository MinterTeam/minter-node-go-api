package responses

type Response struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
}

type ErrorData struct {
	Code     int32     `json:"code"`
	Message  string    `json:"message"`
	Data     *string   `json:"data"`
	TxResult *txResult `json:"tx_result"`
}

type txResult struct {
	Code int32  `json:"code"`
	Log  string `json:"log"`
}
