package regources

//PayoutOpAttrs - details of corresponding op
type PayoutOp struct {
	Key
	Attributes    PayoutOpAttrs     `json:"attributes"`
	Relationships PayoutOpRelations `json:"relationships"`
}

//PayoutOpAttrs - details of corresponding op
type PayoutOpAttrs struct {
	MaxPayoutAmount      Amount `json:"max_payout_amount"`
	MinAssetHolderAmount Amount `json:"min_asset_holder_amount"`
	MinPayoutAmount      Amount `json:"min_payout_amount"`
	ExpectedFee          Fee    `json:"expected_fee"`
	ActualFee            Fee    `json:"actual_fee"`
	ActualPayoutAmount   Amount `json:"actual_payout_amount"`
}

type PayoutOpRelations struct {
	SourceAccount *Relation `json:"source_account"`
	SourceBalance *Relation `json:"source_balance"`
	Asset         *Relation `json:"asset"`
}
