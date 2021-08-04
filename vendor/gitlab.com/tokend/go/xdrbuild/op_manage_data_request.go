package xdrbuild

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateDataRequest struct {
	RequestID      uint64
	AllTasks       *uint32
	CreatorDetails interface{}

	Type  uint64
	Value interface{}
	Owner string
}

func (r CreateDataRequest) XDR() (*xdr.Operation, error) {
	value, err := json.Marshal(r.Value)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal value")
	}

	var owner xdr.AccountId
	err = owner.SetAddress(r.Owner)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set owner address")
	}

	request := &xdr.CreateDataCreationRequestOp{
		RequestId: xdr.Uint64(r.RequestID),
		DataCreationRequest: xdr.DataCreationRequest{
			Type:           xdr.Uint64(r.Type),
			Value:          xdr.Longstring(value),
			CreatorDetails: xdr.Longstring("{}"),
			Owner:          owner,
		},
		Ext: xdr.EmptyExt{},
	}

	if r.CreatorDetails != nil {
		details, err := json.Marshal(r.CreatorDetails)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal creator details")
		}
		request.DataCreationRequest.CreatorDetails = xdr.Longstring(details)
	}

	if r.AllTasks != nil {
		allTasks := xdr.Uint32(*r.AllTasks)
		request.AllTasks = &allTasks
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:                        xdr.OperationTypeCreateDataCreationRequest,
			CreateDataCreationRequestOp: request,
		},
	}, err
}

type UpdateDataRequest struct {
	RequestID      uint64
	AllTasks       *uint32
	CreatorDetails interface{}

	DataId uint64
	Value  interface{}
}

func (r UpdateDataRequest) XDR() (*xdr.Operation, error) {
	value, err := json.Marshal(r.Value)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal value")
	}

	request := &xdr.CreateDataUpdateRequestOp{
		RequestId: xdr.Uint64(r.RequestID),
		DataUpdateRequest: xdr.DataUpdateRequest{
			Id:             xdr.Uint64(r.DataId),
			Value:          xdr.Longstring(value),
			CreatorDetails: xdr.Longstring("{}"),
		},
		Ext: xdr.EmptyExt{},
	}

	if r.CreatorDetails != nil {
		details, err := json.Marshal(r.CreatorDetails)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal creator details")
		}
		request.DataUpdateRequest.CreatorDetails = xdr.Longstring(details)
	}

	if r.AllTasks != nil {
		allTasks := xdr.Uint32(*r.AllTasks)
		request.AllTasks = &allTasks
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:                      xdr.OperationTypeCreateDataUpdateRequest,
			CreateDataUpdateRequestOp: request,
		},
	}, err
}

type RemoveDataRequest struct {
	RequestID      uint64
	AllTasks       *uint32
	CreatorDetails interface{}

	DataId uint64
}

func (r RemoveDataRequest) XDR() (*xdr.Operation, error) {
	request := &xdr.CreateDataRemoveRequestOp{
		RequestId: xdr.Uint64(r.RequestID),
		DataRemoveRequest: xdr.DataRemoveRequest{
			Id:             xdr.Uint64(r.DataId),
			CreatorDetails: xdr.Longstring("{}"),
		},
		Ext: xdr.EmptyExt{},
	}

	if r.CreatorDetails != nil {
		details, err := json.Marshal(r.CreatorDetails)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal creator details")
		}
		request.DataRemoveRequest.CreatorDetails = xdr.Longstring(details)
	}

	if r.AllTasks != nil {
		allTasks := xdr.Uint32(*r.AllTasks)
		request.AllTasks = &allTasks
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:                      xdr.OperationTypeCreateDataRemoveRequest,
			CreateDataRemoveRequestOp: request,
		},
	}, nil
}
