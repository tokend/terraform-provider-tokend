package regources

import "gitlab.com/tokend/go/xdr"

//ManageExternalSystemPoolOp - details of corresponding op
type ManageExternalSystemPoolOp struct {
	Key
	Attributes ManageExternalSystemPoolOpAttrs `json:"attributes"`
}

//ManageExternalSystemPoolOpAttrs - details of corresponding op
type ManageExternalSystemPoolOpAttrs struct {
	Action xdr.ManageExternalSystemAccountIdPoolEntryAction `json:"action"`
	Create *CreateExternalSystemPoolOp                      `json:"create"`
	Remove *RemoveExternalSystemPoolOp                      `json:"remove"`
}

//CreateExternalSystemPoolOp - details of corresponding op
type CreateExternalSystemPoolOp struct {
	PoolID             uint64 `json:"pool_id"`
	Data               string `json:"data"`
	Parent             uint64 `json:"parent"`
	ExternalSystemType int32  `json:"external_system_type"`
}

//RemoveExternalSystemPoolOp - details of corresponding op
type RemoveExternalSystemPoolOp struct {
	PoolID uint64 `json:"pool_id"`
}

//BindExternalSystemAccountOpAttrs - details of corresponding op
type BindExternalSystemAccountOp struct {
	Key
	Attributes BindExternalSystemAccountOpAttrs `json:"attributes"`
}

//BindExternalSystemAccountOpAttrs - details of corresponding op
type BindExternalSystemAccountOpAttrs struct {
	ExternalSystemType int32 `json:"external_system_type"`
}
