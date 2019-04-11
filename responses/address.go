package responses

type AddressResponse struct {
	Response
	Result Balance `json:"result"`
}
