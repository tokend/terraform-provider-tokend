package xdrbuild

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/strkey"
	"gitlab.com/tokend/go/xdr"
)

type CreateData struct {
	Type  uint64
	Value json.Marshaler
}

func (c CreateData) XDR() (*xdr.Operation, error) {
	value, err := c.Value.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateData,
			CreateDataOp: &xdr.CreateDataOp{
				Type:  xdr.Uint64(c.Type),
				Value: xdr.Longstring(value),
				Ext:   xdr.EmptyExt{},
			},
		},
	}, nil
}

type UpdateData struct {
	ID    uint64
	Value json.Marshaler
}

func (u UpdateData) XDR() (*xdr.Operation, error) {
	value, err := u.Value.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeUpdateData,
			UpdateDataOp: &xdr.UpdateDataOp{
				DataId: xdr.Uint64(u.ID),
				Value:  xdr.Longstring(value),
				Ext:    xdr.EmptyExt{},
			},
		},
	}, nil
}

type RemoveData struct {
	ID uint64
}

func (r RemoveData) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeRemoveData,
			RemoveDataOp: &xdr.RemoveDataOp{
				DataId: xdr.Uint64(r.ID),
				Ext:    xdr.EmptyExt{},
			},
		},
	}, nil
}

type UpdateDataOwner struct {
	ID       uint64
	NewOwner string
}

func (u UpdateDataOwner) XDR() (*xdr.Operation, error) {
	var newOwner xdr.AccountId
	err := newOwner.SetAddress(u.NewOwner)
	if err != nil {
		return nil, errors.Wrap(err, "invalid new owner")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeUpdateDataOwner,
			UpdateDataOwnerOp: &xdr.UpdateDataOwnerOp{
				DataId:   xdr.Uint64(u.ID),
				NewOwner: newOwner,
				Ext:      xdr.EmptyExt{},
			},
		},
	}, nil
}

func convert(accountId string) xdr.AccountId {

	raw, _ := strkey.Decode(strkey.VersionByteAccountID, accountId)

	var ui xdr.Uint256
	copy(ui[:], raw)

	result, _ := xdr.NewAccountId(xdr.CryptoKeyTypeKeyTypeEd25519, ui)

	return result
}
