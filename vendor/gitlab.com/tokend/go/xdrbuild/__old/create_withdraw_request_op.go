package __old

import (
	"encoding/json"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type WithdrawRequestDetails interface {
	WithdrawRequestDetails() string
}

type ETHWithdrawRequestDetails struct {
	Address string `json:"address"`
}

func (d *ETHWithdrawRequestDetails) WithdrawRequestDetails() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal details"))
	}
	return string(bytes)
}

func (d ETHWithdrawRequestDetails) Validate() error {
	return validation.ValidateStruct(&d,
		// TODO check valid ETH address
		validation.Field(&d.Address, validation.Required),
	)
}

// TODO tests
type CreateWithdrawRequestOp struct {
	Balance  string
	Asset    string
	Amount   uint64
	AllTasks *uint32
	Details  WithdrawRequestDetails
}

func (op CreateWithdrawRequestOp) Validate() error {
	return validation.ValidateStruct(&op,
		// TODO check valid balance ID
		validation.Field(&op.Balance, validation.Required),
		validation.Field(&op.Asset, validation.Required, validation.Length(1, 16)),
		validation.Field(&op.Amount, validation.Required),
		validation.Field(&op.Details, validation.Required),
	)
}

func (op CreateWithdrawRequestOp) XDR() (*xdr.Operation, error) {
	var balanceID xdr.BalanceId
	if err := balanceID.SetString(op.Balance); err != nil {
		return nil, errors.Wrap(err, "failed to set receiver")
	}

	var allTasksXDR *xdr.Uint32
	if op.AllTasks != nil {
		allTasksXDR = new(xdr.Uint32)
		*allTasksXDR = xdr.Uint32(*op.AllTasks)
	}

	operation := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateWithdrawalRequest,
			CreateWithdrawalRequestOp: &xdr.CreateWithdrawalRequestOp{
				Request: xdr.WithdrawalRequest{
					Balance:        balanceID,
					Amount:         xdr.Uint64(op.Amount),
					CreatorDetails: xdr.Longstring(op.Details.WithdrawRequestDetails()),
				},
				AllTasks: allTasksXDR,
			},
		},
	}

	if op.AllTasks != nil {
		var allTasksXDR xdr.Uint32
		var allTasksXDRPointer *xdr.Uint32
		allTasksXDR = xdr.Uint32(*op.AllTasks)
		allTasksXDRPointer = &allTasksXDR
		operation.Body.CreateWithdrawalRequestOp.AllTasks = allTasksXDRPointer
		operation.Body.CreateWithdrawalRequestOp.Ext = xdr.CreateWithdrawalRequestOpExt{
			V: xdr.LedgerVersionEmptyVersion,
		}
	}

	return operation, nil
}
