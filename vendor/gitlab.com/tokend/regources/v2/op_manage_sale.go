package regources

import (
	"gitlab.com/tokend/go/xdr"
)

//ManageSaleOp - details of corresponding op
type ManageSaleOp struct {
	Key
	Attributes ManageSaleOpAttrs `json:"attributes"`
}

//ManageSaleOpAttrs - details of corresponding op
type ManageSaleOpAttrs struct {
	SaleID uint64               `json:"sale_id"`
	Action xdr.ManageSaleAction `json:"action"`
}
