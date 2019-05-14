/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ReviewRequestOpAttributes struct {
	// * 1: \"approve\" * 2: \"reject\" * 3: \"permanent_reject\"
	Action xdr.ReviewRequestOpAction `json:"action"`
	// Tasks that were added on the request review
	AddedTasks      uint32  `json:"added_tasks"`
	ExternalDetails Details `json:"external_details"`
	// Whether request being reviewed was fulfilled
	IsFulfilled bool `json:"is_fulfilled"`
	// Reject reason
	Reason string `json:"reason"`
	// Tasks that were removed on the request review
	RemovedTasks uint32 `json:"removed_tasks"`
	// Hash of the request being reviewed
	RequestHash string `json:"request_hash"`
	// ID of the request being reviewed
	RequestId int64 `json:"request_id"`
}
