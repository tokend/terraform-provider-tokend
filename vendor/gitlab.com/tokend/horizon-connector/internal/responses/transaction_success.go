package responses

type TransactionSuccess struct {
	Hash     string `json:"hash"`
	Ledger   int32  `json:"ledger"`
	Envelope string `json:"envelope_xdr"`
	Result   string `json:"result_xdr"`
	Meta     string `json:"result_meta_xdr"`
}
