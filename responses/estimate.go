package responses

type EstimateTxResponse struct {
	Response
	Result struct {
		Commission string `json:"commission"`
	} `json:"result"`
}

type EstimateCoinBuyResponse struct {
	Response
	Result struct {
		Commission string `json:"commission"`
		WillPay    string `json:"will_pay"`
	} `json:"result"`
}

type EstimateCoinSellResponse struct {
	Response
	Result struct {
		WillGet string `json:"will_get"`
	} `json:"result"`
}

type EstimateCoinSellAllResponse struct {
	Response
	Result struct {
		Commission string `json:"commission"`
		WillGet    string `json:"will_get"`
	} `json:"result"`
}

type GasResponse struct {
	Response
	Result string `json:"result"`
}
