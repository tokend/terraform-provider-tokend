package responses

type TransactionBadRequest struct {
	Type   string                      `json:"type"`
	Extras TransactionBadRequestExtras `json:"extras"`
}

type TransactionResultCodes struct {
	Transaction string   `json:"transaction"`
	Operations  []string `json:"operations"`
	Messages    []string `json:"messages"`
}

type TransactionBadRequestExtras struct {
	EnvelopeXDR string                 `json:"envelope_xdr"`
	ResultXDR   string                 `json:"result_xdr"`
	ResultCodes TransactionResultCodes `json:"result_codes"`
}
