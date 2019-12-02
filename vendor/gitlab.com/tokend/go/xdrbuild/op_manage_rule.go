package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateRule struct {
	Resource xdr.RuleResource
	Action   xdr.RuleAction
	Forbid   bool
	Details  json.Marshaler
}

func (op *CreateRule) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateRule,
			CreateRuleOp: &xdr.CreateRuleOp{
				Resource: op.Resource,
				Action:   op.Action,
				Forbids:  op.Forbid,
				Details:  xdr.Longstring(details),
			},
		},
	}, nil
}

type UpdateRule struct {
	ID       uint64
	Resource xdr.RuleResource
	Action   xdr.RuleAction
	Forbid   bool
	Details  json.Marshaler
}

func (op *UpdateRule) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeUpdateRule,
			UpdateRuleOp: &xdr.UpdateRuleOp{
				RuleId:   xdr.Uint64(op.ID),
				Resource: op.Resource,
				Action:   op.Action,
				Forbids:  op.Forbid,
				Details:  xdr.Longstring(details),
			},
		},
	}, nil
}

type RemoveRule struct {
	ID uint64
}

func (op *RemoveRule) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeRemoveRule,
			RemoveRuleOp: &xdr.RemoveRuleOp{
				RuleId: xdr.Uint64(op.ID),
			},
		},
	}, nil
}
