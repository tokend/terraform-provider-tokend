package xdrbuild

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type Payment struct {
	SourceAccountID *xdr.AccountId
	SourceBalanceID xdr.BalanceId
	Destination     xdr.PaymentOpDestination
	FeeData         xdr.PaymentFeeData
	Amount          uint64
	Subject         string
	Reference       string
}

func (op *Payment) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		SourceAccount: op.SourceAccountID,
		Body: xdr.OperationBody{
			Type: xdr.OperationTypePayment,
			PaymentOp: &xdr.PaymentOp{
				SourceBalanceId: op.SourceBalanceID,
				Destination:     op.Destination,
				Amount:          xdr.Uint64(op.Amount),
				Subject:         xdr.Longstring(op.Subject),
				Reference:       xdr.Longstring(op.Reference),
				FeeData:         op.FeeData,
			},
		},
	}, nil
}

type CreatePaymentForBalanceOpts struct {
	SourceAccountID      *string
	SourceBalanceID      string
	DestinationBalanceID string
	Amount               uint64
	Subject              string
	Reference            string
	Fee                  Fee
}

type Fee struct {
	SourceFixed        uint64
	SourcePercent      uint64
	DestinationFixed   uint64
	DestinationPercent uint64
	SourcePaysForDest  bool
}

func CreatePaymentForBalance(opts CreatePaymentForBalanceOpts) (*Payment, error) {
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

	var sa *xdr.AccountId = nil
	if opts.SourceAccountID != nil {
		sa = &xdr.AccountId{}
		if err := sa.SetAddress(*opts.SourceAccountID); err != nil {
			return nil, errors.Wrap(err, "failed to set op source account id")
		}
	}

	return &Payment{
		SourceAccountID: sa,
		SourceBalanceID: sb,
		Destination: xdr.PaymentOpDestination{
			Type:      xdr.PaymentDestinationTypeBalance,
			BalanceId: &db,
		},
		Amount:    opts.Amount,
		Subject:   opts.Subject,
		Reference: opts.Reference,
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
	}, nil
}

type CreatePaymentForAccountOpts struct {
	SourceAccountID      *string
	SourceBalanceID      string
	DestinationAccountID string
	Amount               uint64
	Subject              string
	Reference            string
	Fee                  Fee
}

func CreatePaymentForAccount(opts CreatePaymentForAccountOpts) (*Payment, error) {
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

	var sa *xdr.AccountId = nil
	if opts.SourceAccountID != nil {
		sa = &xdr.AccountId{}
		if err := sa.SetAddress(*opts.SourceAccountID); err != nil {
			return nil, errors.Wrap(err, "failed to set op source account id")
		}
	}

	return &Payment{
		SourceAccountID: sa,
		SourceBalanceID: sb,
		Destination: xdr.PaymentOpDestination{
			Type:      xdr.PaymentDestinationTypeAccount,
			AccountId: &da,
		},
		Amount:    opts.Amount,
		Subject:   opts.Subject,
		Reference: opts.Reference,
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
	}, nil
}
