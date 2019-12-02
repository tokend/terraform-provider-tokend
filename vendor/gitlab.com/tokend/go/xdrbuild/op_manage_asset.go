package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateAsset struct {
	Code                string
	MaxIssuanceAmount   uint64
	TrailingDigitsCount uint32
	SecurityType        uint32
	State               uint32
	CreatorDetails      json.Marshaler
}

func (op *CreateAsset) XDR() (*xdr.Operation, error) {
	details, err := op.CreatorDetails.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal creator details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateAsset,
			CreateAssetOp: &xdr.CreateAssetOp{
				Code:                xdr.AssetCode(op.Code),
				MaxIssuanceAmount:   xdr.Uint64(op.MaxIssuanceAmount),
				Details:             xdr.Longstring(details),
				SecurityType:        xdr.Uint32(op.SecurityType),
				State:               xdr.Uint32(op.State),
				TrailingDigitsCount: xdr.Uint32(op.TrailingDigitsCount),
			},
		},
	}, nil
}
