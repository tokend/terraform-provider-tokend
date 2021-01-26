package __old

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateUpdateKYCRequestOp struct {
	RequestID          uint64
	AccountToUpdateKYC string
	AccountTypeToSet   xdr.AccountType
	KYCLevelToSet      uint32
	KYCData            string
	AllTasks           *uint32
}

func (op CreateUpdateKYCRequestOp) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.AccountTypeToSet, validation.Required),
		validation.Field(&op.KYCData, validation.Required),
		validation.Field(&op.AccountToUpdateKYC, validation.Required),
	)
}

func (op CreateUpdateKYCRequestOp) XDR() (*xdr.Operation, error) {
	var accountToUpdateKYC xdr.AccountId
	if err := accountToUpdateKYC.SetAddress(op.AccountToUpdateKYC); err != nil {
		return nil, errors.Wrap(err, "failed to set updated account")
	}

	var allTasksXDR xdr.Uint32
	var allTasksXDRPointer *xdr.Uint32

	if op.AllTasks != nil {
		allTasksXDR = xdr.Uint32(*op.AllTasks)
		allTasksXDRPointer = &allTasksXDR
	} else {
		allTasksXDRPointer = nil
	}

	xdrop := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateKycRequest,
			CreateUpdateKycRequestOp: &xdr.CreateUpdateKycRequestOp{
				RequestId: xdr.Uint64(op.RequestID),
				UpdateKycRequestData: xdr.UpdateKycRequestData{
					AccountToUpdateKyc: accountToUpdateKYC,
					AccountTypeToSet:   op.AccountTypeToSet,
					KycLevelToSet:      xdr.Uint32(op.KYCLevelToSet),
					KycData:            xdr.Longstring(op.KYCData),
				},
				AllTasks: allTasksXDRPointer,
			},
		},
	}
	return xdrop, nil
}
