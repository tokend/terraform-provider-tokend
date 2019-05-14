package operations

import "gitlab.com/tokend/regources"

type PaymentV2 struct {
	PT                 string              `json:"paging_token"`
	TransactionID      string              `json:"transaction_id"`
	PaymentID          uint64              `json:"payment_id"`
	From               string              `json:"from,omitempty"`
	To                 string              `json:"to,omitempty"`
	Asset              string              `json:"asset"`
	Subject            string              `json:"subject"`
	Amount             regources.Amount    `json:"amount"`
	SourceFeeData      regources.FeeDataV2 `json:"source_fee_data"`
	DestinationFeeData regources.FeeDataV2 `json:"destination_fee_data"`
	SourcePaysForDest  bool                `json:"source_pays_for_dest"`
}

func (op PaymentV2) PagingToken() string {
	return op.PT
}
