/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type AssetAttributes struct {
	// Asset volume authorized to be issued by an asset owner
	AvailableForIssuance Amount  `json:"available_for_issuance"`
	Details              Details `json:"details"`
	// Asset volume that is currently in circulation
	Issued Amount `json:"issued"`
	// Max volume of an asset that can be in circulation
	MaxIssuanceAmount Amount `json:"max_issuance_amount"`
	// Asset volume to be distributed via [asset sale↪](https://tokend.gitbook.io/knowledge-base/platform-features/crowdfunding) but currently locked by the system
	PendingIssuance Amount          `json:"pending_issuance"`
	Policies        xdr.AssetPolicy `json:"policies"`
	// address of the signer responsible for pre-issuance. [Details↪](https://tokend.gitbook.io/knowledge-base/technical-details/key-entities/asset#pre-issued-asset-signer)
	PreIssuanceAssetSigner string `json:"pre_issuance_asset_signer"`
	// Number of significant digits after the point
	TrailingDigits uint32 `json:"trailing_digits"`
	// Numeric type of asset
	Type uint64 `json:"type"`
}
