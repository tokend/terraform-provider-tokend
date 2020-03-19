package xdrbuild

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateRedemptionRequest struct {
	SourceBalanceID      string
	DestinationAccountID string
	Amount               uint64
	Reference            string
	Details              json.Marshaler
	AllTasks             *uint32
}

func (op CreateRedemptionRequest) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.SourceBalanceID, validation.Required),
		validation.Field(&op.DestinationAccountID, validation.Required),
		validation.Field(&op.Amount, validation.Required),
		validation.Field(&op.Reference, validation.Required, validation.Length(1, 64)),
		validation.Field(&op.AllTasks, validation.NilOrNotEmpty),
	)
}

func (op CreateRedemptionRequest) XDR() (*xdr.Operation, error) {
	var source xdr.BalanceId
	if err := source.SetString(op.SourceBalanceID); err != nil {
		return nil, errors.Wrap(err, "failed to set source balance")
	}

	var destination xdr.AccountId
	if err := destination.SetAddress(op.DestinationAccountID); err != nil {
		return nil, errors.Wrap(err, "failed to set destination account")
	}

	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateRedemptionRequest,
			CreateRedemptionRequestOp: &xdr.CreateRedemptionRequestOp{
				Reference: xdr.String64(op.Reference),
				AllTasks:  (*xdr.Uint32)(op.AllTasks),
				RedemptionRequest: xdr.RedemptionRequest{
					SourceBalanceId: source,
					Destination:     destination,
					Amount:          xdr.Uint64(op.Amount),
					CreatorDetails:  xdr.Longstring(details),
				},
			},
		},
	}, nil
}
