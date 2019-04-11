package responses

import "time"

type BlockResponse struct {
	Response
	Result struct {
		Hash         string        `json:"hash"`
		Size         string        `json:"size"`
		Height       string        `json:"height"`
		TxCount      string        `json:"num_txs"`
		TotalTx      string        `json:"total_txs"`
		Proposer     string        `json:"proposer"`
		BlockReward  string        `json:"block_reward"`
		Time         time.Time     `json:"time"`
		Transactions []Transaction `json:"transactions"`
		Validators   []Validator   `json:"validators"`
	} `json:"result"`
}
