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

type UpdateAsset struct {
	Code              string
	MaxIssuanceAmount *uint64
	State             *uint32
	Details           json.Marshaler
}

func (op *UpdateAsset) XDR() (*xdr.Operation, error) {
	var details *xdr.Longstring
	if op.Details != nil {
		tmpDetails, err := op.Details.MarshalJSON()
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal details")
		}
		longDetails := xdr.Longstring(tmpDetails)
		details = &longDetails
	}

	var maxIssuance *xdr.Uint64
	if op.MaxIssuanceAmount != nil {
		tmpMax := xdr.Uint64(*op.MaxIssuanceAmount)
		maxIssuance = &tmpMax
	}

	var state *xdr.Uint32
	if op.State != nil {
		tmpState := xdr.Uint32(*op.State)
		state = &tmpState
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeUpdateAsset,
			UpdateAssetOp: &xdr.UpdateAssetOp{
				Code:              xdr.AssetCode(op.Code),
				MaxIssuanceAmount: maxIssuance,
				Details:           details,
				State:             state,
			},
		},
	}, nil
}
