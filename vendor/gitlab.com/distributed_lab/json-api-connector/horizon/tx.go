package horizon

type TransactionSubmit struct {
	Tx            string `json:"tx"`
	WaitForIngest bool   `json:"wait_for_ingest"`
}
