package xdrbuild

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateWithdrawalRequest struct {
	SourceBalanceID string
	Amount          uint64
	UniversalAmount uint64
	FeeData         xdr.PaymentFeeData
	Details         json.Marshaler
	AllTasks        *uint32
}

func (op CreateWithdrawalRequest) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.SourceBalanceID, validation.Required),
		validation.Field(&op.Amount, validation.Required),
	)
}

func (op CreateWithdrawalRequest) XDR() (*xdr.Operation, error) {
	var balance xdr.BalanceId
	if err := balance.SetString(op.SourceBalanceID); err != nil {
		return nil, errors.Wrap(err, "failed to set source balance")
	}

	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateWithdrawalRequest,
			CreateWithdrawalRequestOp: &xdr.CreateWithdrawalRequestOp{
				Request: xdr.WithdrawalRequest{
					Balance:         balance,
					Amount:          xdr.Uint64(op.Amount),
					UniversalAmount: xdr.Uint64(op.UniversalAmount),
					Fee:             op.FeeData.SourceFee,
					CreatorDetails:  xdr.Longstring(details),
					Ext:             xdr.WithdrawalRequestExt{},
				},
				AllTasks: (*xdr.Uint32)(op.AllTasks),
				Ext:      xdr.CreateWithdrawalRequestOpExt{},
			},
		},
	}, nil
}
