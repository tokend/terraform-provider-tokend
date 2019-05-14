package xdrbuild

import (
	. "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/tokend/go/xdr"
)

type LicenseOp struct {
	AdminCount      uint64
	DueDate         uint64
	LedgerHash      xdr.Hash
	PrevLicenseHash xdr.Hash
	Signatures      []xdr.DecoratedSignature
}

func (op LicenseOp) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeLicense,
			LicenseOp: &xdr.LicenseOp{
				LedgerHash: op.LedgerHash,
				PrevLicenseHash: op.PrevLicenseHash,
				DueDate: xdr.Uint64(op.DueDate),
				AdminCount: xdr.Uint64(op.AdminCount),
				Signatures: op.Signatures,
			},
		},
	}, nil
}

func (op LicenseOp) Validate() error {
	return ValidateStruct(&op,
		Field(&op.AdminCount, Required, Min(2)),
		Field(&op.LedgerHash, Required),
		Field(&op.DueDate, Required),
		Field(&op.PrevLicenseHash, Required),
		Field(&op.Signatures, Required, Length(2, 2)),
		)
}




