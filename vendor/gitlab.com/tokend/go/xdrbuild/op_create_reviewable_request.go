package xdrbuild

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateReviewableRequest struct {
	SecurityType uint32
	Operations   []Operation
}

func (op *CreateReviewableRequest) XDR() (*xdr.Operation, error) {
	operations := make([]xdr.ReviewableRequestOperation, 0, len(op.Operations))
	for _, operation := range op.Operations {
		xdrOp, err := operation.XDR()
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert operation to xdr struct")
		}
		if xdrOp == nil {
			return nil, errors.New("unexpected nil result from XDR method")
		}

		revOp, err := xdrOp.Body.ToReviewableRequestOp()
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert operation body to reviewable request op")
		}

		operations = append(operations, revOp)
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateReviewableRequest,
			CreateReviewableRequestOp: &xdr.CreateReviewableRequestOp{
				SecurityType: xdr.Uint32(op.SecurityType),
				Operations:   operations,
				Ext:          xdr.EmptyExt{},
			},
		},
	}, nil
}
