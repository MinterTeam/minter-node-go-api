package responses

type ValidatorsResponse struct {
	Jsonrpc    string      `json:"jsonrpc"`
	ID         string      `json:"id"`
	Error      *ErrorData  `json:"error"`
	Validators []Validator `json:"result"`
}

type Validator struct {
	PubKey      string  `json:"pub_key"`
	Signed      *bool   `json:"signed"`
	VotingPower *string `json:"voting_power"`
}
