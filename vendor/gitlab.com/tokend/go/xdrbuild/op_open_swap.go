package xdrbuild

import (
	"encoding/json"
	"time"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type OpenSwap struct {
	SourceBalanceID xdr.BalanceId
	Destination     xdr.OpenSwapOpDestination
	FeeData         xdr.PaymentFeeData
	Details         json.Marshaler
	Amount          uint64
	SecretHash      [32]byte
	LockTime        time.Time
}

func (op *OpenSwap) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeOpenSwap,
			OpenSwapOp: &xdr.OpenSwapOp{
				SourceBalance: op.SourceBalanceID,
				Amount:        xdr.Uint64(op.Amount),
				Destination:   op.Destination,
				FeeData:       op.FeeData,
				SecretHash:    op.SecretHash,
				LockTime:      xdr.Int64(op.LockTime.UTC().Unix()),
				Details:       xdr.Longstring(details),
				Ext:           xdr.EmptyExt{},
			},
		},
	}, nil
}

type CreateOpenSwapForBalanceOpts struct {
	SourceBalanceID      string
	DestinationBalanceID string
	Amount               uint64
	Asset                string
	Fee                  Fee
	SecretHash           [32]byte
	LockTime             time.Time
	Details              json.Marshaler
}

func CreateOpenSwapForBalance(opts CreateOpenSwapForBalanceOpts) (*OpenSwap, error) {
	var db xdr.BalanceId
	err := db.SetString(opts.DestinationBalanceID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set source balance id")
	}

	var sb xdr.BalanceId
	err = sb.SetString(opts.SourceBalanceID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set source balance id")
	}

	return &OpenSwap{
		SourceBalanceID: sb,
		Destination: xdr.OpenSwapOpDestination{
			Type:      xdr.PaymentDestinationTypeBalance,
			BalanceId: &db,
		},
		FeeData: xdr.PaymentFeeData{
			SourceFee: xdr.Fee{
				Fixed:   xdr.Uint64(opts.Fee.SourceFixed),
				Percent: xdr.Uint64(opts.Fee.SourcePercent),
			},
			DestinationFee: xdr.Fee{
				Fixed:   xdr.Uint64(opts.Fee.DestinationFixed),
				Percent: xdr.Uint64(opts.Fee.DestinationPercent),
			},
			SourcePaysForDest: opts.Fee.SourcePaysForDest,
		},
		Amount:     opts.Amount,
		SecretHash: opts.SecretHash,
		LockTime:   opts.LockTime,
		Details:    opts.Details,
	}, nil
}

type CreateOpenSwapForAccountOpts struct {
	SourceBalanceID      string
	DestinationAccountID string
	Amount               uint64
	Asset                string
	Fee                  Fee
	SecretHash           [32]byte
	LockTime             time.Time
	Details              json.Marshaler
}

func CreateOpenSwapForAccount(opts CreateOpenSwapForAccountOpts) (*OpenSwap, error) {
	var sb xdr.BalanceId
	err := sb.SetString(opts.SourceBalanceID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set source balance id")
	}

	var da xdr.AccountId
	err = da.SetAddress(opts.DestinationAccountID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get destination account id")
	}

	return &OpenSwap{
		SourceBalanceID: sb,
		Destination: xdr.OpenSwapOpDestination{
			Type:      xdr.PaymentDestinationTypeAccount,
			AccountId: &da,
		},
		FeeData: xdr.PaymentFeeData{
			SourceFee: xdr.Fee{
				Fixed:   xdr.Uint64(opts.Fee.SourceFixed),
				Percent: xdr.Uint64(opts.Fee.SourcePercent),
			},
			DestinationFee: xdr.Fee{
				Fixed:   xdr.Uint64(opts.Fee.DestinationFixed),
				Percent: xdr.Uint64(opts.Fee.DestinationPercent),
			},
			SourcePaysForDest: opts.Fee.SourcePaysForDest,
		},
		Amount:     opts.Amount,
		SecretHash: opts.SecretHash,
		LockTime:   opts.LockTime,
		Details:    opts.Details,
	}, nil
}
