package xdrbuild

import (
	"encoding/hex"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type ReviewRequest struct {
	ID              uint64
	Hash            string
	Action          xdr.ReviewRequestOpAction
	Reason          string
	TasksToAdd      uint64
	TasksToRemove   uint64
	ExternalDetails string
}

func (op *ReviewRequest) XDR() (*xdr.Operation, error) {
	reviewRequestOp := xdr.ReviewRequestOp{
		RequestId:       xdr.Uint64(op.ID),
		Action:          op.Action,
		Reason:          xdr.Longstring(op.Reason),
		TasksToAdd:      xdr.Uint64(op.TasksToAdd),
		TasksToRemove:   xdr.Uint64(op.TasksToRemove),
		ExternalDetails: xdr.Longstring(op.ExternalDetails),
	}

	hash, err := hex.DecodeString(op.Hash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode hash hex string")
	}
	copy(reviewRequestOp.RequestHash[:], hash)

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:            xdr.OperationTypeReviewRequest,
			ReviewRequestOp: &reviewRequestOp,
		},
	}, nil

}
