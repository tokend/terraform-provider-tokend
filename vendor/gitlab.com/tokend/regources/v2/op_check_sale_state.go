package regources

import "gitlab.com/tokend/go/xdr"

//CheckSaleStateAttrs - details of corresponding op
type CheckSaleStateOp struct {
	Key
	Attributes    CheckSaleStateOpAttrs     `json:"attributes"`
	Relationships CheckSaleStateOpRelations `json:"relationships"`
}

//CheckSaleStateAttrs - details of corresponding op
type CheckSaleStateOpAttrs struct {
	Effect xdr.CheckSaleStateEffect `json:"effect"`
}

//CheckSaleStateAttrs - relationships of the operation
type CheckSaleStateOpRelations struct {
	Sale *Relation `json:"sale"`
}
