package __old

import (
	"gitlab.com/tokend/go/xdr"
)

type manageExternalPoolEntryOp struct {
	Action xdr.ManageExternalSystemAccountIdPoolEntryAction
	Create *createExternalPoolEntryInput
}

func (op manageExternalPoolEntryOp) Validate() error {
	// TODO implement
	return nil
}

func (op manageExternalPoolEntryOp) XDR() (*xdr.Operation, error) {
	mop := xdr.ManageExternalSystemAccountIdPoolEntryOp{
		ActionInput: xdr.ManageExternalSystemAccountIdPoolEntryOpActionInput{
			Action: op.Action,
		},
	}

	switch op.Action {
	case xdr.ManageExternalSystemAccountIdPoolEntryActionCreate:
		mop.ActionInput.CreateExternalSystemAccountIdPoolEntryActionInput = &xdr.CreateExternalSystemAccountIdPoolEntryActionInput{
			ExternalSystemType: xdr.Int32(op.Create.ExternalSystemType),
			Data:               xdr.Longstring(op.Create.Data),
			Parent:             xdr.Uint64(op.Create.Parent),
		}
	default:
		panic("not implemented")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageExternalSystemAccountIdPoolEntry,
			ManageExternalSystemAccountIdPoolEntryOp: &mop,
		},
	}, nil
}

type createExternalPoolEntryInput struct {
	ExternalSystemType int32
	Data               string
	Parent             uint64
}

func CreateExternalPoolEntry(externalSystemType int32, data string, parent uint64) manageExternalPoolEntryOp {
	return manageExternalPoolEntryOp{
		Action: xdr.ManageExternalSystemAccountIdPoolEntryActionCreate,
		Create: &createExternalPoolEntryInput{
			ExternalSystemType: externalSystemType,
			Data:               data,
			Parent:             parent,
		},
	}
}
