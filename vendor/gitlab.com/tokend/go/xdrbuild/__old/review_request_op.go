package __old

import (
	"encoding/hex"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type ReviewRequestDetails interface {
	ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails
}

// TODO research why does it exist
type ReviewRequestOpDetails struct {
	Type       xdr.ReviewableRequestType
	Withdrawal *ReviewRequestOpWithdrawalDetails
	Issuance   *ReviewRequestOpIssuanceDetails
}

type ReviewRequestOpWithdrawalDetails struct {
	ExternalDetails string
}

type ReviewRequestOpIssuanceDetails struct{}
type IssuanceDetails struct{}
type AtomicSwapDetails struct{}
type ChangeRoleDetails struct{}

func (d IssuanceDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeCreateIssuance,
	}
}

func (d AtomicSwapDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeCreateAtomicSwap,
	}
}

type ReviewDetails struct {
	TasksToAdd      uint32
	TasksToRemove   uint32
	ExternalDetails string
}

type ReviewRequestOp struct {
	ID uint64
	// Hash optional, not a pointer for backwards compatibility
	Hash          string
	Action        xdr.ReviewRequestOpAction
	Details       ReviewRequestDetails
	Reason        string
	ReviewDetails *xdr.ReviewDetails
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

func (d ChangeRoleDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeChangeRole,
	}
}

func (op ReviewRequestOp) XDR() (*xdr.Operation, error) {
	if op.ReviewDetails == nil {
		op.ReviewDetails = &xdr.ReviewDetails{}
	}
	xdrop := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeReviewRequest,
			ReviewRequestOp: &xdr.ReviewRequestOp{
				RequestId: xdr.Uint64(op.ID),
				Action:    op.Action,
				Reason:    xdr.Longstring(op.Reason),
				ReviewDetails: xdr.ReviewDetails{
					TasksToAdd:      xdr.Uint32(op.ReviewDetails.TasksToAdd),
					TasksToRemove:   xdr.Uint32(op.ReviewDetails.TasksToRemove),
					ExternalDetails: op.ReviewDetails.ExternalDetails,
				},
			},
		},
	}

	if op.Hash != "" {
		hash, err := hex.DecodeString(op.Hash)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode hash")
		}
		copy(xdrop.Body.ReviewRequestOp.RequestHash[:], hash[:32])
	}

	if op.Details != nil {
		xdrop.Body.ReviewRequestOp.RequestDetails = op.Details.ReviewRequestDetails()
	}

	return xdrop, nil
}
