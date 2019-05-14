package regources

import "gitlab.com/tokend/go/xdr"

//ManageAssetOp - details of corresponding op
type ManageAssetOp struct {
	Key
	Attributes    ManageAssetOpAttrs     `json:"attributes"`
	Relationships ManageAssetOpRelations `json:"relationships"`
}

//ManageAssetOpAttrs - details of corresponding op
type ManageAssetOpAttrs struct {
	AssetCode         string                `json:"asset_code,omitempty"`
	Type              uint64                `json:"type,omitempty"`
	Action            xdr.ManageAssetAction `json:"action"`
	Policies          *xdr.AssetPolicy      `json:"policies,omitempty"`
	CreatorDetails    Details               `json:"creator_details,omitempty"`
	PreissuedSigner   string                `json:"preissuance_signer,omitempty"`
	MaxIssuanceAmount Amount                `json:"max_issuance_amount,omitempty"`
}

//ManageAssetRelations - details of corresponding op
type ManageAssetOpRelations struct {
	Request *Relation `json:"request"`
}
