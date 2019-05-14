package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateSignerRule struct {
	Resource xdr.SignerRuleResource
	Action   xdr.SignerRuleAction
	Forbid   bool
	Default  bool
	ReadOnly bool
	Details  json.Marshaler
}

func (op *CreateSignerRule) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSignerRule,
			ManageSignerRuleOp: &xdr.ManageSignerRuleOp{
				Data: xdr.ManageSignerRuleOpData{
					Action: xdr.ManageSignerRuleActionCreate,
					CreateData: &xdr.CreateSignerRuleData{
						Resource:   op.Resource,
						Action:     op.Action,
						Forbids:    op.Forbid,
						IsDefault:  op.Default,
						IsReadOnly: op.ReadOnly,
						Details:    xdr.Longstring(details),
					},
				},
			},
		},
	}, nil
}
