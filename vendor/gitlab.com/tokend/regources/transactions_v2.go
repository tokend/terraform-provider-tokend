package regources

import "time"

// TransactionV2 represents a single, successful transaction with ledger changes
type TransactionV2 struct {
	ID              string                `json:"id"`
	PT              string                `json:"paging_token"`
	Hash            string                `json:"hash"`
	LedgerCloseTime time.Time             `json:"created_at"`
	LedgerSequence  int32                 `json:"ledger_seq"`
	EnvelopeXDR     string                `json:"envelope_xdr"`
	ResultXDR       string                `json:"result_xdr"`
	Changes         []LedgerEntryChangeV2 `json:"changes"`
}

// PagingToken implementation for hal.Pageable
func (t TransactionV2) PagingToken() string {
	return t.PT
}

// LedgerEntryChangeV2 represents what happened with entry with xdr ledger entry or ledger key
type LedgerEntryChangeV2 struct {
	Effect    int32  `json:"effect"`
	EntryType int32  `json:"entry_type"`
	Payload   string `json:"payload"`
}
