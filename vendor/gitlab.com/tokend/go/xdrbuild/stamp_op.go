package xdrbuild

import "gitlab.com/tokend/go/xdr"

type StampOp struct {}

func (op StampOp) Validate() error {
	return nil
}

func (op StampOp) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeStamp,
			StampOp: &xdr.StampOp{},
		},
	}, nil
}



