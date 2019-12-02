package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type Issuance struct {
	SecurityType uint32
	Reference    string
	Destination  xdr.MovementDestination
	Asset        string
	Amount       uint64
	Details      json.Marshaler
	Source       string
}

func (op Issuance) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	var source *xdr.AccountId
	if op.Source != "" {
		source = new(xdr.AccountId)
		err = source.SetAddress(op.Source)
		if err != nil {
			return nil, errors.Wrap(err, "failed to set source account id")
		}
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeIssuance,
			IssuanceOp: &xdr.IssuanceOp{
				SecurityType:   0,
				Asset:          xdr.AssetCode(op.Asset),
				Amount:         xdr.Uint64(op.Amount),
				Destination:    op.Destination,
				Reference:      xdr.Longstring(op.Reference),
				CreatorDetails: xdr.Longstring(details),
				Fee:            xdr.Fee{},
				Ext:            xdr.EmptyExt{},
			},
		},
		SourceAccount: source,
	}, nil
}

type CreateIssuanceForBalanceOpts struct {
	SecurityType         uint32
	DestinationBalanceID string
	Amount               uint64
	Reference            string
	Asset                string
	Details              json.Marshaler
}

func CreateIssuanceForBalance(opts CreateIssuanceForBalanceOpts) (*Issuance, error) {
	var db xdr.BalanceId
	err := db.SetString(opts.DestinationBalanceID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set source balance id")
	}

	return &Issuance{
		SecurityType: opts.SecurityType,
		Reference:    opts.Reference,
		Destination: xdr.MovementDestination{
			Type:      xdr.DestinationTypeBalance,
			BalanceId: &db,
		},
		Asset:   opts.Asset,
		Amount:  opts.Amount,
		Details: opts.Details,
	}, nil
}

type CreateIssuanceForAccountOpts struct {
	SecurityType         uint32
	DestinationAccountID string
	Amount               uint64
	Reference            string
	Asset                string
	Details              json.Marshaler
}

func CreateIssuanceForAccount(opts CreateIssuanceForAccountOpts) (*Issuance, error) {
	var da xdr.AccountId
	err := da.SetAddress(opts.DestinationAccountID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get destination account id")
	}

	return &Issuance{
		SecurityType: opts.SecurityType,
		Reference:    opts.Reference,
		Destination: xdr.MovementDestination{
			Type:      xdr.DestinationTypeAccount,
			AccountId: &da,
		},
		Asset:   opts.Asset,
		Amount:  opts.Amount,
		Details: opts.Details,
	}, nil
}
