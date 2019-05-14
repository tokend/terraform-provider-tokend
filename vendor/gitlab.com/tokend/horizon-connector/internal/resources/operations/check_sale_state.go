package operations

type CheckSaleState struct {
	SaleID        uint64 `json:"sale_id"`
	Effect        string `json:"effect"`
	TransactionID string `json:"transaction_id"`
	PT            string `json:"paging_token"`
}

func (op CheckSaleState) PagingToken() string {
	return op.PT
}
