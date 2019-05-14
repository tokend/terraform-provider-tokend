package regources

import "gitlab.com/tokend/go/xdr"

//SetFeeOp - stores details of create account operation
type SetFeeOp struct {
	Key
	Attributes SetFeeOpAttrs `json:"attributes"`
}

//SetFeeOpAttrs - details of SetFeeOp
type SetFeeOpAttrs struct {
	AssetCode      string      `json:"asset_code"`
	FixedFee       Amount      `json:"fixed_fee"`
	PercentFee     Amount      `json:"percent_fee"`
	FeeType        xdr.FeeType `json:"fee_type"`
	AccountAddress *string     `json:"account_address,omitempty"`
	AccountRole    *xdr.Uint64 `json:"account_role,omitempty"`
	Subtype        int64       `json:"subtype"`
	LowerBound     Amount      `json:"lower_bound"`
	UpperBound     Amount      `json:"upper_bound"`
	IsDelete       bool        `json:"is_delete"`
	// FeeAsset deprecated
}
