/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "gitlab.com/tokend/go/xdr"

type LedgerEntryChangeAttributes struct {
	// Type of change that was applied to entry * 0: \"created\" * 1: \"updated\" * 2: \"removed\" * 3: \"state\"
	ChangeType xdr.LedgerEntryChangeType `json:"change_type"`
	// Type of entry that was modified * 1:  \"any\" * 2:  \"account\" * 3:  \"signer\" * 4: \"reviewable_request\" * 5: \"account_kyc\" * 6: \"key_value\" * 7: \"account_role\" * 8: \"account_rule\" * 9: \"transaction\" * 10: \"signer_rule\" * 11: \"signer_role\"
	EntryType xdr.LedgerEntryType `json:"entry_type"`
	Payload   string              `json:"payload"`
}
