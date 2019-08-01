package xdrbuild

import (
	"encoding/hex"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type ReviewDetails struct {
	TasksToAdd      uint32
	TasksToRemove   uint32
	ExternalDetails string
}

type ReviewRequest struct {
	ID            uint64
	Hash          *string
	Action        xdr.ReviewRequestOpAction
	Type          xdr.ReviewableRequestType
	Reason        string
	ReviewDetails ReviewDetails
}

func (op ReviewRequest) XDR() (*xdr.Operation, error) {
	reviewDetails := xdr.ReviewDetails{
		TasksToAdd:      xdr.Uint32(op.ReviewDetails.TasksToAdd),
		TasksToRemove:   xdr.Uint32(op.ReviewDetails.TasksToRemove),
		ExternalDetails: op.ReviewDetails.ExternalDetails,
	}

	reviewRequest := xdr.ReviewRequestOp{
		RequestType:   op.Type,
		RequestId:     xdr.Uint64(op.ID),
		Action:        op.Action,
		Reason:        xdr.Longstring(op.Reason),
		ReviewDetails: reviewDetails,
	}

	xdrOp := &xdr.Operation{
		Body: xdr.OperationBody{
			Type:            xdr.OperationTypeReviewRequest,
			ReviewRequestOp: &reviewRequest,
		},
	}

	if op.Hash != nil {
		hashBB, err := hex.DecodeString(*op.Hash)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode hash")
		}

		copy(xdrOp.Body.ReviewRequestOp.RequestHash[:], hashBB[:32])
	}
	return xdrOp, nil
}
