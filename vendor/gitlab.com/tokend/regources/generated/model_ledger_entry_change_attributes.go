/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type LedgerEntryChangeAttributes struct {
	// Type of change that was applied to entry * 0: \"created\" * 1: \"updated\" * 2: \"removed\" * 3: \"state\"
	ChangeType xdr.LedgerEntryChangeType `json:"change_type"`
	// Type of entry that was modified * 1:  \"any\" * 2:  \"account\" * 3:  \"signer\" * 4:  \"fee\" * 5:  \"balance\" * 6:  \"payment_request\" * 7:  \"asset\" * 8:  \"reference_entry\" * 9:  \"statistics\" * 10: \"trust\" * 11: \"account_limits\" * 12: \"asset_pair\" * 13: \"offer_entry\" * 15: \"reviewable_request\" * 16: \"external_system_account_id\" * 17: \"sale\" * 18: \"account_kyc\" * 19: \"external_system_account_id_pool_entry\" * 20: \"key_value\" * 22: \"limits_v2\" * 23: \"statistics_v2\" * 24: \"pending_statistics\" * 25: \"contract\" * 26: \"account_role\" * 27: \"account_rule\" * 28: \"atomic_swap_bid\" * 29: \"transaction\" * 30: \"signer_rule\" * 31: \"signer_role\" * 32: \"stamp\" * 33: \"license\" * 34: \"poll\" * 35: \"vote\"
	EntryType xdr.LedgerEntryType `json:"entry_type"`
	Payload   string              `json:"payload"`
}
