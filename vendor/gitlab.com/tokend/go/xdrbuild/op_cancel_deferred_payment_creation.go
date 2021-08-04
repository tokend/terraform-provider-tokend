package xdrbuild

import "gitlab.com/tokend/go/xdr"

type CancelDeferredPaymentCreation struct {
	RequestID uint64
}

func (op CancelDeferredPaymentCreation) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCancelDeferredPaymentCreationRequest,
			CancelDeferredPaymentCreationRequestOp: &xdr.CancelDeferredPaymentCreationRequestOp{
				RequestId: xdr.Uint64(op.RequestID),
				Ext:       xdr.CancelDeferredPaymentCreationRequestOpExt{},
			},
		},
	}, nil
}
