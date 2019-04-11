package responses

type ValidatorsResponse struct {
	Response
	Validators []Validator `json:"result"`
}

type Validator struct {
	PubKey      string  `json:"pub_key"`
	Signed      *bool   `json:"signed"`
	VotingPower *string `json:"voting_power"`
}
