package regources

import (
	"gitlab.com/tokend/go/xdr"
)

//AssetResponse - response for asset handler
type AssetResponse struct {
	Data     Asset    `json:"data"`
	Included Included `json:"included"`
}

type AssetsResponse struct {
	Links    *Links   `json:"links"`
	Data     []Asset  `json:"data"`
	Included Included `json:"included"`
}

// Asset - Resource object representing AssetEntry
type Asset struct {
	Key
	Attributes    AssetAttrs     `json:"attributes"`
	Relationships AssetRelations `json:"relationships"`
}

type AssetAttrs struct {
	PreIssuanceAssetSigner string                 `json:"pre_issuance_asset_signer" `
	Details                map[string]interface{} `json:"details"`
	MaxIssuanceAmount      Amount                 `json:"max_issuance_amount"`
	AvailableForIssuance   Amount                 `json:"available_for_issuance"`
	Issued                 Amount                 `json:"issued"`
	PendingIssuance        Amount                 `json:"pending_issuance"`
	Policies               xdr.AssetPolicy        `json:"policies"`
	TrailingDigits         uint32                 `json:"trailing_digits"`
	Type                   uint64                 `json:"type"`
}

type AssetRelations struct {
	Owner *Relation `json:"owner"`
}
