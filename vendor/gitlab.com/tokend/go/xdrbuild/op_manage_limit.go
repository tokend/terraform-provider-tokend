package xdrbuild

import (
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateLimit struct {
	Action     xdr.ManageLimitsAction
	Role       *xdr.Uint64
	Id         *xdr.AccountId
	Type       xdr.StatsOpType
	Code       xdr.AssetCode
	Convert    bool
	DailyOut   xdr.Uint64
	WeeklyOut  xdr.Uint64
	MonthlyOut xdr.Uint64
	AnnualOut  xdr.Uint64
}

func (op *CreateLimit) XDR() (*xdr.Operation, error) {
	details := xdr.LimitsCreateDetails{
		StatsOpType:     op.Type,
		AssetCode:       op.Code,
		IsConvertNeeded: op.Convert,
		DailyOut:        op.DailyOut,
		WeeklyOut:       op.WeeklyOut,
		MonthlyOut:      op.MonthlyOut,
		AnnualOut:       op.AnnualOut,
	}

	if op.Id != nil {
		details.AccountId = op.Id
	}

	if op.Role != nil {
		details.AccountRole = op.Role
	}

	if op.Role == nil && op.Id == nil {
		return nil, errors.New("invalid role and id values")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageLimits,
			ManageLimitsOp: &xdr.ManageLimitsOp{
				Details: xdr.ManageLimitsOpDetails{
					Action:              xdr.ManageLimitsActionCreate,
					LimitsCreateDetails: &details,
				},
			},
		},
	}, nil
}

type RemoveLimit struct {
	ID uint64
}

func (op *RemoveLimit) XDR() (*xdr.Operation, error) {
	Id := xdr.Uint64(op.ID)

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageLimits,
			ManageLimitsOp: &xdr.ManageLimitsOp{
				Details: xdr.ManageLimitsOpDetails{
					Action: xdr.ManageLimitsActionRemove,
					Id:     &Id,
				},
			},
		},
	}, nil
}
