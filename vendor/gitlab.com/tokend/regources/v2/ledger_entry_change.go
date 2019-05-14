package regources

import "gitlab.com/tokend/go/xdr"

// LedgerEntryChange - resource that represents the ledger entry change
type LedgerEntryChange struct {
	Key
	Attributes LedgerEntryChangeAttrs `json:"attributes"`
}

// LedgerEntryChangeAttrs - attributes of the ledger entry change resource
type LedgerEntryChangeAttrs struct {
	Payload    string                    `json:"payload"`
	ChangeType xdr.LedgerEntryChangeType `json:"change_type"`
	EntryType  xdr.LedgerEntryType       `json:"entry_type"`
}
