package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateRole struct {
	Details json.Marshaler
	Rules   []uint64
}

func (op *CreateRole) XDR() (*xdr.Operation, error) {
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
			Type: xdr.OperationTypeCreateRole,
			CreateRoleOp: &xdr.CreateRoleOp{
				Details: xdr.Longstring(details),
				RuleIDs: rules,
			},
		},
	}, nil
}

type UpdateRole struct {
	ID      uint64
	Details json.Marshaler
	Rules   []uint64
}

func (op *UpdateRole) XDR() (*xdr.Operation, error) {
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
			Type: xdr.OperationTypeUpdateRole,
			UpdateRoleOp: &xdr.UpdateRoleOp{
				RoleId:  xdr.Uint64(op.ID),
				Details: xdr.Longstring(details),
				RuleIDs: rules,
			},
		},
	}, nil
}

type RemoveRole struct {
	ID uint64
}

func (op *RemoveRole) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeRemoveRole,
			RemoveRoleOp: &xdr.RemoveRoleOp{
				RoleId: xdr.Uint64(op.ID),
			},
		},
	}, nil
}
