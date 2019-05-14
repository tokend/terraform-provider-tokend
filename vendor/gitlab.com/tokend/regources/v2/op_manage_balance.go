package regources

import "gitlab.com/tokend/go/xdr"

//ManageBalanceOp - stores details of create account operation
type ManageBalanceOp struct {
	Key
	Attributes    ManageBalanceOpAttrs     `json:"attributes"`
	Relationships ManageBalanceOpRelations `json:"relationships"`
}

//ManageBalanceOpAttrs - details of ManageBalanceOp
type ManageBalanceOpAttrs struct {
	Action         xdr.ManageBalanceAction `json:"action"`
	BalanceAddress string                  `json:"balance_address"`
}

//ManageBalanceOpAttrs - relationships of ManageBalanceOp
type ManageBalanceOpRelations struct {
	DestinationAccount *Relation `json:"destination_account"`
	Asset              *Relation `json:"asset"`
}
