/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type CreateAccountSpecificRuleData struct {
	AccountId *string `json:"account_id,omitempty"`
	Forbids   bool    `json:"forbids"`
	// Ledger entry key
	LedgerKey xdr.LedgerKey `json:"ledger_key"`
}
