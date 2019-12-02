package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateSigner struct {
	SignerData
}

type SignerData struct {
	PublicKey string
	RoleIDs   []uint64
	Weight    uint32
	Identity  uint32
	Details   json.Marshaler
}

func (d *SignerData) XDR() (*xdr.SignerData, error) {
	details, err := d.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	var publicKey xdr.AccountId
	if err := publicKey.SetAddress(d.PublicKey); err != nil {
		return nil, errors.Wrap(err, "failed to set public key")
	}

	roleIDs := make([]xdr.Uint64, 0, len(d.RoleIDs))
	for _, roleID := range d.RoleIDs {
		roleIDs = append(roleIDs, xdr.Uint64(roleID))
	}

	return &xdr.SignerData{
		PublicKey: xdr.PublicKey(publicKey),
		RoleIDs:   roleIDs,
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
			Type: xdr.OperationTypeCreateSigner,
			CreateSignerOp: &xdr.CreateSignerOp{
				Data: *data,
			},
		},
	}, nil
}

type UpdateSigner struct {
	SignerData
}

func (op *UpdateSigner) XDR() (*xdr.Operation, error) {
	data, err := op.SignerData.XDR()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build signer data")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeUpdateSigner,
			UpdateSignerOp: &xdr.UpdateSignerOp{
				Data: *data,
			},
		},
	}, nil
}

type RemoveSigner struct {
	PublicKey string
}

func (op *RemoveSigner) XDR() (*xdr.Operation, error) {
	publicKey := xdr.PublicKey{}
	err := publicKey.FromString(op.PublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set public key from string")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeRemoveSigner,
			RemoveSignerOp: &xdr.RemoveSignerOp{
				PublicKey: publicKey,
			},
		},
	}, nil
}
