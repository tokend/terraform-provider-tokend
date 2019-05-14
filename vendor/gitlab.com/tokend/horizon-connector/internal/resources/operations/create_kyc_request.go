package operations

type CreateKYCRequest struct {
	ID                 string                 `json:"id"`
	RequestID          uint64                 `json:"request_id"`
	AccountToUpdateKYC string                 `json:"account_to_update_kyc"`
	AccountTypeToSet   int32                  `json:"account_type_to_set"`
	KYCData            map[string]interface{} `json:"kyc_data"`
	PT                 string                 `json:"paging_token"`
	TransactionID      string                 `json:"transaction_id"`
}

func (op CreateKYCRequest) PagingToken() string {
	return op.PT
}
