package __old

import (
	"encoding/json"

	. "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type (
	AssetDetails struct {
		ExternalSystemType int    `json:"external_system_type,string"`
		Name               string `json:"name"`
		Logo               *Logo  `json:"logo,omitempty"`
	}

	Logo struct {
		Key  string `json:"key,omitempty"`
		Type string `json:"type,omitempty"`
		URL  string `json:"url,omitempty"`
	}

	CreateAssetOp struct {
		AssetSigner       string
		MaxIssuanceAmount uint64
		PreIssuanceAmount uint64
		Policies          uint32
		Code              string
		Details           AssetDetails
	}

	UpdateAsset struct {
		Code     string
		Policies uint32
		Details  AssetDetails
	}
)

func (ca CreateAssetOp) Validate() error {
	return ValidateStruct(&ca,
		Field(&ca.Code, Required),
		Field(&ca.AssetSigner, Required),
		Field(&ca.Policies, Required),
		Field(&ca.PreIssuanceAmount, Required),
		Field(&ca.MaxIssuanceAmount, Required),
	)
}

func (ca CreateAssetOp) XDR() (*xdr.Operation, error) {
	var signer xdr.AccountId
	err := signer.SetAddress(ca.AssetSigner)
	if err != nil {
		return nil, errors.Wrap(err, "invalid signer")
	}

	details, err := json.Marshal(ca.Details)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal details")
	}

	op := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAsset,
			ManageAssetOp: &xdr.ManageAssetOp{
				RequestId: 0,
				Request: xdr.ManageAssetOpRequest{
					Action: xdr.ManageAssetActionCreateAssetCreationRequest,
					CreateAssetCreationRequest: &xdr.ManageAssetOpCreateAssetCreationRequest{
						CreateAsset: xdr.AssetCreationRequest{
							PreissuedAssetSigner:   signer,
							CreatorDetails:         xdr.Longstring(details),
							MaxIssuanceAmount:      xdr.Uint64(ca.MaxIssuanceAmount),
							InitialPreissuedAmount: xdr.Uint64(ca.PreIssuanceAmount),
							Policies:               xdr.Uint32(ca.Policies),
							Code:                   xdr.AssetCode(ca.Code),
						},
					},
				},
			},
		},
	}

	return op, nil
}

func (ca UpdateAsset) Validate() error {
	return ValidateStruct(&ca,
		Field(&ca.Code, Required),
		Field(&ca.Policies, Required),
	)
}

func (ca UpdateAsset) XDR() (*xdr.Operation, error) {
	details, err := json.Marshal(ca.Details)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal details")
	}

	op := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAsset,
			ManageAssetOp: &xdr.ManageAssetOp{
				RequestId: 0,
				Request: xdr.ManageAssetOpRequest{
					Action: xdr.ManageAssetActionCreateAssetUpdateRequest,
					CreateAssetUpdateRequest: &xdr.ManageAssetOpCreateAssetUpdateRequest{
						UpdateAsset: xdr.AssetUpdateRequest{
							Code:           xdr.AssetCode(ca.Code),
							CreatorDetails: xdr.Longstring(details),
							Policies:       xdr.Uint32(ca.Policies),
							Ext: xdr.AssetUpdateRequestExt{
								V: xdr.LedgerVersionEmptyVersion,
							},
						},
					},
				},
			},
		},
	}

	return op, nil
}
