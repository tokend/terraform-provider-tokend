package xdrbuild

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateAccountSaleRule struct {
	SaleID    uint64
	Forbids   bool
	AccountID string
}

func (op *CreateAccountSaleRule) XDR() (*xdr.Operation, error) {
	destination := &xdr.AccountId{}
	if op.AccountID != "" {
		if err := destination.SetAddress(op.AccountID); err != nil {
			return nil, errors.Wrap(err, "failed to set account id address")
		}
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAccountSpecificRule,
			ManageAccountSpecificRuleOp: &xdr.ManageAccountSpecificRuleOp{
				Data: xdr.ManageAccountSpecificRuleOpData{
					Action: xdr.ManageAccountSpecificRuleActionCreate,
					CreateData: &xdr.CreateAccountSpecificRuleData{
						LedgerKey: xdr.LedgerKey{
							Type: xdr.LedgerEntryTypeSale,
							Sale: &xdr.LedgerKeySale{
								SaleId: xdr.Uint64(op.SaleID),
							},
						},
						Forbids:   op.Forbids,
						AccountId: destination,
					},
				},
			},
		},
	}, nil
}

type RemoveAccountSaleRule struct {
	RuleID uint64
}

func (op *RemoveAccountSaleRule) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAccountSpecificRule,
			ManageAccountSpecificRuleOp: &xdr.ManageAccountSpecificRuleOp{
				Data: xdr.ManageAccountSpecificRuleOpData{
					Action: xdr.ManageAccountSpecificRuleActionRemove,
					RemoveData: &xdr.RemoveAccountSpecificRuleData{
						RuleId: xdr.Uint64(op.RuleID),
					},
				},
			},
		},
	}, nil
}
