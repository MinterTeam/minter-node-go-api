package responses

type CoinInfoResponse struct {
	Response
	Result struct {
		Name           string `json:"name"`
		Symbol         string `json:"symbol"`
		Volume         string `json:"volume"`
		Crr            string `json:"crr"`
		ReserveBalance string `json:"reserve_balance"`
	} `json:"result"`
}
