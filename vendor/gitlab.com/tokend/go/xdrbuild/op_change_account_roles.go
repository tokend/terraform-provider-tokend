package xdrbuild

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type ChangeAccountRoles struct {
	DestinationAccount string
	RoleIDsToSet       []uint64
	Details            json.Marshaler
}

func (op *ChangeAccountRoles) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	changeRolesOp := xdr.ChangeAccountRolesOp{
		Details:    xdr.Longstring(details),
		RolesToSet: make([]xdr.Uint64, 0, len(op.RoleIDsToSet)),
		Ext:        xdr.EmptyExt{},
	}

	err = changeRolesOp.DestinationAccount.SetAddress(op.DestinationAccount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set destination account id")
	}

	for _, roleID := range op.RoleIDsToSet {
		changeRolesOp.RolesToSet = append(changeRolesOp.RolesToSet, xdr.Uint64(roleID))
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:                 xdr.OperationTypeChangeAccountRoles,
			ChangeAccountRolesOp: &changeRolesOp,
		},
	}, nil
}
