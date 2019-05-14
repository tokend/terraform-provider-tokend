package regources

type StampOp struct {
	Key
	Attributes StampOpAttributes `json:"attributes"`
}

type StampOpAttributes struct {
	LedgerHash  string `json:"ledger_hash"`
	LicenseHash string `json:"license_hash"`
}
