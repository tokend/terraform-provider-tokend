package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateSignerRole struct {
	Details    json.Marshaler
	IsReadOnly bool
	Rules      []uint64
}

func (op *CreateSignerRole) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	rules := make([]xdr.Uint64, 0, len(op.Rules))
	for _, rule := range op.Rules {
		rules = append(rules, xdr.Uint64(rule))
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSignerRole,
			ManageSignerRoleOp: &xdr.ManageSignerRoleOp{
				Data: xdr.ManageSignerRoleOpData{
					Action: xdr.ManageSignerRoleActionCreate,
					CreateData: &xdr.CreateSignerRoleData{
						Details:    xdr.Longstring(details),
						IsReadOnly: op.IsReadOnly,
						RuleIDs:    rules,
					},
				},
			},
		},
	}, nil
}

type UpdateSignerRole struct {
	ID      uint64
	Details json.Marshaler
	Rules   []uint64
}

func (op *UpdateSignerRole) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	rules := make([]xdr.Uint64, 0, len(op.Rules))
	for _, rule := range op.Rules {
		rules = append(rules, xdr.Uint64(rule))
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSignerRole,
			ManageSignerRoleOp: &xdr.ManageSignerRoleOp{
				Data: xdr.ManageSignerRoleOpData{
					Action: xdr.ManageSignerRoleActionUpdate,
					UpdateData: &xdr.UpdateSignerRoleData{
						RoleId:  xdr.Uint64(op.ID),
						RuleIDs: rules,
						Details: xdr.Longstring(details),
					},
				},
			},
		},
	}, nil
}
