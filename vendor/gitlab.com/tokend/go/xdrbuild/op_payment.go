package xdrbuild

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type Payment struct {
	SecurityType    uint32
	SourceBalanceID xdr.BalanceId
	Destination     xdr.MovementDestination
	FeeData         xdr.PaymentFeeData
	Amount          uint64
	Subject         string
	Reference       string
	Source          string
}

func (op *Payment) XDR() (*xdr.Operation, error) {
	var source *xdr.AccountId
	if op.Source != "" {
		source = new(xdr.AccountId)
		err := source.SetAddress(op.Source)
		if err != nil {
			return nil, errors.Wrap(err, "failed to set source account id")
		}
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypePayment,
			PaymentOp: &xdr.PaymentOp{
				SourceBalanceId: op.SourceBalanceID,
				SecurityType:    xdr.Uint32(op.SecurityType),
				Destination:     op.Destination,
				Amount:          xdr.Uint64(op.Amount),
				FeeData:         op.FeeData,
				Subject:         xdr.Longstring(op.Subject),
				Reference:       xdr.Longstring(op.Reference),
				Ext:             xdr.PaymentOpExt{},
			},
		},
		SourceAccount: source,
	}, nil
}

type CreatePaymentForBalanceOpts struct {
	SecurityType         uint32
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

	return &Payment{
		SecurityType:    opts.SecurityType,
		SourceBalanceID: sb,
		Destination: xdr.MovementDestination{
			Type:      xdr.DestinationTypeBalance,
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
	SecurityType         uint32
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

	return &Payment{
		SecurityType:    opts.SecurityType,
		SourceBalanceID: sb,
		Destination: xdr.MovementDestination{
			Type:      xdr.DestinationTypeAccount,
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
