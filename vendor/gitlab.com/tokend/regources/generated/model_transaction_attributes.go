/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type TransactionAttributes struct {
	// close time of ledger where transaction belongs to
	CreatedAt time.Time `json:"created_at"`
	// base-64 encoded XDR representation of transaction itself
	EnvelopeXdr string `json:"envelope_xdr"`
	// hash of transaction
	Hash string `json:"hash"`
	// sequence of ledger where transaction belongs to
	LedgerSequence int32 `json:"ledger_sequence"`
	// base-64 encoded XDR representation of core response's meta information
	ResultMetaXdr string `json:"result_meta_xdr"`
	// base-64 encoded XDR representation of core response
	ResultXdr string `json:"result_xdr"`
}
