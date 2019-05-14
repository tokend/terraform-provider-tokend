/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManageAssetOpAttributes struct {
	// * 0: \"create_asset_creation_request\" * 1: \"create_asset_update_request\" * 2: \"cancel_asset_request\" * 3: \"change_preissued_asset_signer\" * 4: \"update_max_issuance\"
	Action xdr.ManageAssetAction `json:"action"`
	// Asset to manage
	AssetCode         string  `json:"asset_code"`
	CreatorDetails    Details `json:"creator_details"`
	MaxIssuanceAmount Amount  `json:"max_issuance_amount"`
	// Bit mask. * 1:  \"transferable\" * 2:  \"base_asset\" * 4:  \"stats_quote_asset\" * 8:  \"withdrawable\" * 16: \"issuance_manual_review_required\" * 32: \"can_be_base_in_atomic_swap\" * 64: \"can_be_quote_in_atomic_swap\"
	Policies *xdr.AssetPolicy `json:"policies,omitempty"`
	// Address of pre-issuance signer
	PreIssuanceSigner string `json:"pre_issuance_signer"`
}
