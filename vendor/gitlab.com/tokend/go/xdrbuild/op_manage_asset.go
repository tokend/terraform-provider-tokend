package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateAsset struct {
	Code                     string
	MaxIssuanceAmount        uint64
	PreIssuanceSigner        string
	InitialPreIssuanceAmount uint64
	TrailingDigitsCount      uint32
	Policies                 uint32
	Type                     uint64
	CreatorDetails           json.Marshaler
	AllTasks                 *uint32
}

func (op *CreateAsset) XDR() (*xdr.Operation, error) {
	details, err := op.CreatorDetails.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal creator details")
	}

	var tasks *xdr.Uint32
	if op.AllTasks != nil {
		tasks = (*xdr.Uint32)(op.AllTasks)
	}

	var preIssuanceSigner xdr.AccountId
	if err := preIssuanceSigner.SetAddress(op.PreIssuanceSigner); err != nil {
		return nil, errors.Wrap(err, "failed to encode pre-issuance signer")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAsset,
			ManageAssetOp: &xdr.ManageAssetOp{
				Request: xdr.ManageAssetOpRequest{
					Action: xdr.ManageAssetActionCreateAssetCreationRequest,
					CreateAssetCreationRequest: &xdr.ManageAssetOpCreateAssetCreationRequest{
						CreateAsset: xdr.AssetCreationRequest{
							Code:                   xdr.AssetCode(op.Code),
							PreissuedAssetSigner:   preIssuanceSigner,
							MaxIssuanceAmount:      xdr.Uint64(op.MaxIssuanceAmount),
							InitialPreissuedAmount: xdr.Uint64(op.InitialPreIssuanceAmount),
							Policies:               xdr.Uint32(op.Policies),
							CreatorDetails:         xdr.Longstring(details),
							Type:                   xdr.Uint64(op.Type),
							TrailingDigitsCount:    xdr.Uint32(op.TrailingDigitsCount),
						},
						AllTasks: tasks,
					},
				},
			},
		},
	}, nil
}
