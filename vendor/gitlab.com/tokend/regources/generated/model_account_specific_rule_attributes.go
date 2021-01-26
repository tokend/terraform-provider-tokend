/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type AccountSpecificRuleAttributes struct {
	// Defines whether rule is applied to specific user, or to all users
	AccountId *string `json:"account_id,omitempty"`
	// defines whether or not it is forbidden to access resource
	Forbids bool `json:"forbids"`
	// Ledger entry key
	LedgerKey xdr.LedgerKey `json:"ledger_key"`
}
