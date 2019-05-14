package regources

// CreateWithdrawalRequest - represents details of the `withdrawal` reviewable request
type CreateWithdrawalRequest struct {
	Key
	Attributes    CreateWithdrawalRequestAttrs     `json:"attributes"`
	Relationships CreateWithdrawalRequestRelations `json:"relationships"`
}

// CreateWithdrawalRequestAttrs - attributes of the `withdrawal` reviewable request
type CreateWithdrawalRequestAttrs struct {
	Fee            Fee     `json:"fee"`
	Amount         Amount  `json:"amount"`
	CreatorDetails Details `json:"creator_details"`
}

// CreateWithdrawalRequestRelations - relationships of the `withdrawal` reviewable request
type CreateWithdrawalRequestRelations struct {
	Balance *Relation `json:"balance"`
}
