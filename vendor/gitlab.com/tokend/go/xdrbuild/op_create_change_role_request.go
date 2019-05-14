package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateChangeRoleRequest struct {
	RequestID          uint64
	DestinationAccount string
	RoleToSet          uint64
	KYCData            map[string]interface{}
	AllTasks           *uint32
}

func (op *CreateChangeRoleRequest) XDR() (*xdr.Operation, error) {
	var destination xdr.AccountId
	err := destination.SetAddress(op.DestinationAccount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set destination address")
	}

	kycData, err := json.Marshal(op.KYCData)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal map"))
	}

	changeRole := xdr.CreateChangeRoleRequestOp{
		RequestId:          xdr.Uint64(op.RequestID),
		DestinationAccount: destination,
		AccountRoleToSet:   xdr.Uint64(op.RoleToSet),
		CreatorDetails:     xdr.Longstring(kycData),
	}

	if op.AllTasks != nil {
		allTasks := xdr.Uint32(*op.AllTasks)
		changeRole.AllTasks = &allTasks
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateChangeRoleRequest,
			CreateChangeRoleRequestOp: &changeRole,
		},
	}, nil
}
