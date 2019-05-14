package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateAccountRule struct {
	Resource xdr.AccountRuleResource
	Action   xdr.AccountRuleAction
	Forbid   bool
	Details  json.Marshaler
}

func (op *CreateAccountRule) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAccountRule,
			ManageAccountRuleOp: &xdr.ManageAccountRuleOp{
				Data: xdr.ManageAccountRuleOpData{
					Action: xdr.ManageAccountRuleActionCreate,
					CreateData: &xdr.CreateAccountRuleData{
						Resource: op.Resource,
						Action:   op.Action,
						Forbids:  op.Forbid,
						Details:  xdr.Longstring(details),
					},
				},
			},
		},
	}, nil
}

type UpdateAccountRule struct {
	ID       uint64
	Resource xdr.AccountRuleResource
	Action   xdr.AccountRuleAction
	Forbid   bool
	Details  json.Marshaler
}

func (op *UpdateAccountRule) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAccountRule,
			ManageAccountRuleOp: &xdr.ManageAccountRuleOp{
				Data: xdr.ManageAccountRuleOpData{
					Action: xdr.ManageAccountRuleActionUpdate,
					UpdateData: &xdr.UpdateAccountRuleData{
						RuleId:   xdr.Uint64(op.ID),
						Resource: op.Resource,
						Action:   op.Action,
						Forbids:  op.Forbid,
						Details:  xdr.Longstring(details),
					},
				},
			},
		},
	}, nil
}
