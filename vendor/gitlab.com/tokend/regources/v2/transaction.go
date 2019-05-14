package regources

import "time"

// TransactionsResponse - response for get transactions request
type TransactionsResponse struct {
	Links    *Links                  `json:"links"`
	Data     []Transaction           `json:"data"`
	Included Included                `json:"included"`
	Meta     TransactionResponseMeta `json:"meta"`
}

// TransactionsResponseMeta - meta information for transactions response
type TransactionResponseMeta struct {
	LatestLedgerSequence  int32     `json:"latest_ledger_sequence"`
	LatestLedgerCloseTime time.Time `json:"latest_ledger_close_time"`
}

// Transaction - resource that represents transaction
type Transaction struct {
	Key
	Attributes    TransactionAttrs     `json:"attributes"`
	Relationships TransactionRelations `json:"relationships"`
}

// TransactionAttrs - attributes of the transaction resource
type TransactionAttrs struct {
	Hash           string    `json:"hash"`
	LedgerSequence int32     `json:"ledger_sequence"`
	CreatedAt      time.Time `json:"created_at"`
	EnvelopeXdr    string    `json:"envelope_xdr"`
	ResultXdr      string    `json:"result_xdr"`
	ResultMetaXdr  string    `json:"result_meta_xdr"`
}

// TransactionRelations - relationships of the transaction resource
type TransactionRelations struct {
	Source             *Relation           `json:"source"`
	Operations         *RelationCollection `json:"operations"`
	LedgerEntryChanges *RelationCollection `json:"ledger_entry_changes"`
}
