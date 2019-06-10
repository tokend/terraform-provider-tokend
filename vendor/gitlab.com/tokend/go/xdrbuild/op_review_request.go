package xdrbuild

import (
	"encoding/hex"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type ReviewRequestDetailsProvider interface {
	ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails
}

type ReviewRequestOpDetails struct {
	Type       xdr.ReviewableRequestType
	Withdrawal *ReviewRequestOpWithdrawalDetails
	Issuance   *ReviewRequestOpIssuanceDetails
}

type ReviewRequestOpWithdrawalDetails struct {
	ExternalDetails string
}

type ReviewRequestOpIssuanceDetails struct{}

type ReviewDetails struct {
	TasksToAdd      uint32
	TasksToRemove   uint32
	ExternalDetails string
}

type ReviewRequest struct {
	ID            uint64
	Hash          *string
	Action        xdr.ReviewRequestOpAction
	Details       ReviewRequestDetailsProvider
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

	if op.Details != nil {
		xdrOp.Body.ReviewRequestOp.RequestDetails = op.Details.ReviewRequestDetails()
	}

	return xdrOp, nil
}

type WithdrawalDetails struct {
	ExternalDetails string
}

func (d WithdrawalDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeCreateWithdraw,
		Withdrawal: &xdr.WithdrawalDetails{
			ExternalDetails: d.ExternalDetails,
		},
	}
}

type ChangeRoleDetails struct{}

func (d ChangeRoleDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeChangeRole,
	}
}

type AtomicSwapDetails struct{}

func (d AtomicSwapDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeCreateAtomicSwapBid,
	}
}

type IssuanceDetails struct{}

func (d IssuanceDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeCreateIssuance,
	}
}
