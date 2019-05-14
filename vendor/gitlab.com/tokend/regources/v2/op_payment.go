package regources

// PaymentOp - stores details of payment operation
type PaymentOp struct {
	Key
	Attributes    PaymentOpAttrs     `json:"attributes"`
	Relationships PaymentOpRelations `json:"relationships"`
}

// PaymentOpAttrs - stores details of payment operation
type PaymentOpAttrs struct {
	Amount                  Amount `json:"amount"`
	SourceFee               Fee    `json:"source_fee"`
	DestinationFee          Fee    `json:"destination_fee"`
	SourcePayForDestination bool   `json:"source_pay_for_destination"`
	Subject                 string `json:"subject"`
	Reference               string `json:"reference"`
}

// PaymentOpRelations - relationships of the operation
type PaymentOpRelations struct {
	AccountFrom *Relation `json:"account_from"`
	AccountTo   *Relation `json:"account_to"`
	BalanceFrom *Relation `json:"balance_from"`
	BalanceTo   *Relation `json:"balance_to"`
	Asset       *Relation `json:"asset"`
}
