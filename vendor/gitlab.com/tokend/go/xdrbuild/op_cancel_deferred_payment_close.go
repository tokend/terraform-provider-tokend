package xdrbuild

import "gitlab.com/tokend/go/xdr"

type CancelDeferredPaymentClose struct {
	RequestID uint64
}

func (op CancelDeferredPaymentClose) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCancelCloseDeferredPaymentRequest,
			CancelCloseDeferredPaymentRequestOp: &xdr.CancelCloseDeferredPaymentRequestOp{
				RequestId: xdr.Uint64(op.RequestID),
				Ext:       xdr.CancelCloseDeferredPaymentRequestOpExt{},
			},
		},
	}, nil
}
