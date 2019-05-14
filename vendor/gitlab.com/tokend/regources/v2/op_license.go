package regources

import "time"

//LicenseOpAttrs - details of corresponding op
type LicenseOp struct {
	Key
	Attributes LicenseOpAttrs `json:"attributes"`
}

//LicenseOpAttrs - details of corresponding op
type LicenseOpAttrs struct {
	DueDate         time.Time `json:"due_date"`
	AdminCount      uint64    `json:"admin_count"`
	LedgerHash      string    `json:"ledger_hash"`
	PrevLicenseHash string    `json:"prev_license_hash"`
	Signatures      []string  `json:"signatures"`
}
