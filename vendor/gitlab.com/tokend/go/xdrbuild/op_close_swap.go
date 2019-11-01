package xdrbuild

import "gitlab.com/tokend/go/xdr"

type CloseSwap struct {
	SwapID uint64
	Secret *[32]byte
}

func (op CloseSwap) XDR() (*xdr.Operation, error) {
	var secret *xdr.Hash
	if op.Secret != nil {
		temp := xdr.Hash(*op.Secret)
		secret = &temp
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCloseSwap,
			CloseSwapOp: &xdr.CloseSwapOp{
				SwapId: xdr.Uint64(op.SwapID),
				Secret: secret,
				Ext:    xdr.EmptyExt{},
			},
		},
	}, nil
}
