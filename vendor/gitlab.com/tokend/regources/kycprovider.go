package regources

type KYCProviderResponseAttributes struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type KYCProviderResponseData struct {
	Attributes KYCProviderResponseAttributes `json:"attributes"`
}

type KYCProviderResponse struct {
	Data KYCProviderResponseData `json:"data"`
}
