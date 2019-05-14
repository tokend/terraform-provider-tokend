package xdrbuild

import (
	"encoding/json"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateIssuanceRequest struct {
	Reference string
	Receiver  string
	Asset     string
	Amount    uint64
	Details   json.Marshaler
	AllTasks  *uint32
}

func (op CreateIssuanceRequest) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.Reference, validation.Required, validation.Length(1, 64)),
		validation.Field(&op.Asset, validation.Required, validation.Length(1, 16)),
		validation.Field(&op.Amount, validation.Required),
		validation.Field(&op.Receiver, validation.Required),
		validation.Field(&op.Details, validation.Required),
	)
}

func (op CreateIssuanceRequest) XDR() (*xdr.Operation, error) {
	var receiver xdr.BalanceId
	if err := receiver.SetString(op.Receiver); err != nil {
		return nil, errors.Wrap(err, "failed to set receiver")
	}

	var allTasksXDR xdr.Uint32
	var allTasksXDRPointer *xdr.Uint32

	if op.AllTasks != nil {
		allTasksXDR = xdr.Uint32(*op.AllTasks)
		allTasksXDRPointer = &allTasksXDR
	} else {
		allTasksXDRPointer = nil
	}

	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}


	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateIssuanceRequest,
			CreateIssuanceRequestOp: &xdr.CreateIssuanceRequestOp{
				Reference: xdr.String64(op.Reference),
				Request: xdr.IssuanceRequest{
					Asset:           xdr.AssetCode(op.Asset),
					Amount:          xdr.Uint64(op.Amount),
					Receiver:        receiver,
					CreatorDetails: xdr.Longstring(details),
				},
				AllTasks: allTasksXDRPointer,
			},
		},
	}, nil
}
