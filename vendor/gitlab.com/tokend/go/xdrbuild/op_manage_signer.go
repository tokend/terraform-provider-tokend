package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateSigner struct {
	SignerData
}

type UpdateSigner struct {
	SignerData
}

type RemoveSigner struct {
	PublicKey string
}

type SignerData struct {
	PublicKey string
	RoleID    uint64
	Weight    uint32
	Identity  uint32
	Details   json.Marshaler
}

func (d *SignerData) XDR() (*xdr.UpdateSignerData, error) {
	details, err := d.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	var publicKey xdr.AccountId
	if err := publicKey.SetAddress(d.PublicKey); err != nil {
		return nil, errors.Wrap(err, "failed to set public key")
	}

	return &xdr.UpdateSignerData{
		PublicKey: xdr.PublicKey(publicKey),
		RoleId:    xdr.Uint64(d.RoleID),
		Weight:    xdr.Uint32(d.Weight),
		Identity:  xdr.Uint32(d.Identity),
		Details:   xdr.Longstring(details),
	}, nil
}

func (op *CreateSigner) XDR() (*xdr.Operation, error) {
	data, err := op.SignerData.XDR()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build signer data")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSigner,
			ManageSignerOp: &xdr.ManageSignerOp{
				Data: xdr.ManageSignerOpData{
					Action:     xdr.ManageSignerActionCreate,
					CreateData: data,
				},
			},
		},
	}, nil
}

func (op *UpdateSigner) XDR() (*xdr.Operation, error) {
	data, err := op.SignerData.XDR()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build signer data")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSigner,
			ManageSignerOp: &xdr.ManageSignerOp{
				Data: xdr.ManageSignerOpData{
					Action:     xdr.ManageSignerActionUpdate,
					UpdateData: data,
				},
			},
		},
	}, nil
}

func (op *RemoveSigner) XDR() (*xdr.Operation, error) {
	var publicKey xdr.AccountId
	if err := publicKey.SetAddress(op.PublicKey); err != nil {
		return nil, errors.Wrap(err, "failed to set public key")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSigner,
			ManageSignerOp: &xdr.ManageSignerOp{
				Data: xdr.ManageSignerOpData{
					Action: xdr.ManageSignerActionRemove,
					RemoveData: &xdr.RemoveSignerData{
						PublicKey: xdr.PublicKey(publicKey),
					},
				},
			},
		},
	}, nil
}
