package responses

type EstimateTxResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
	Result  struct {
		Commission string `json:"commission"`
	} `json:"result"`
}

type EstimateCoinBuyResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
	Result  struct {
		Commission string `json:"commission"`
		WillPay    string `json:"will_pay"`
	} `json:"result"`
}

type EstimateCoinSellAllResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
	Result  struct {
		WillGet string `json:"will_get"`
	} `json:"result"`
}

type EstimateCoinSellResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
	Result  struct {
		Commission string `json:"commission"`
		WillGet    string `json:"will_get"`
	} `json:"result"`
}

type GasResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Error   *ErrorData `json:"error"`
	Result  string     `json:"result"`
}
