package regources

import "gitlab.com/tokend/go/xdr"

//ReviewRequestOpAttrs - details of corresponding op
type ReviewRequestOp struct {
	Key
	Attributes ReviewRequestOpAttrs `json:"attributes"`
	//TODO: add review request details as relation
}

//ReviewRequestOpAttrs - details of corresponding op
type ReviewRequestOpAttrs struct {
	Action          xdr.ReviewRequestOpAction `json:"action"`
	Reason          string                    `json:"reason"`
	RequestHash     string                    `json:"request_hash"`
	RequestID       int64                     `json:"request_id"`
	IsFulfilled     bool                      `json:"is_fulfilled"`
	AddedTasks      uint32                    `json:"added_tasks"`
	RemovedTasks    uint32                    `json:"removed_tasks"`
	ExternalDetails Details                   `json:"external_details"`
}
