package regources

type FeeDataV2 struct {
	FixedFee                  Amount `json:"fixed_fee"`
	ActualPaymentFee          Amount `json:"actual_payment_fee"`
	ActualPaymentFeeAssetCode string `json:"actual_payment_fee_asset_code"`
}
