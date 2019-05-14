package operations

type ReviewRequest struct {
	ID            string `json:"id"`
	PT            string `json:"paging_token"`
	TransactionID string `json:"transaction_id"`
	Action        string `json:"action"`
	RequestID     uint64 `json:"request_id"`
	RequestType   string `json:"request_type"`
	IsFulfilled   bool   `json:"is_fulfilled"`
}

func (op ReviewRequest) PagingToken() string {
	return op.PT
}
