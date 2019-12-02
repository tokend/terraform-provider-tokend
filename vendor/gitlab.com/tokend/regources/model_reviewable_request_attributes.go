/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "time"

type ReviewableRequestAttributes struct {
	// All tasks that have been set for a request
	AllTasks uint32 `json:"all_tasks"`
	// Time when a request has been submitted
	CreatedAt       time.Time `json:"created_at"`
	ExternalDetails Details   `json:"external_details"`
	// Hash of a particular request
	Hash string `json:"hash"`
	// Tasks that have not been removed yet
	PendingTasks uint32 `json:"pending_tasks"`
	// Details on why a request has been rejected
	RejectReason string `json:"reject_reason"`
	SecurityType uint32 `json:"security_type"`
	// String representation of the request's state
	State string `json:"state"`
	// Integer representation of the request's state
	StateI int32 `json:"state_i"`
	// Last time when a request has been updated
	UpdatedAt time.Time `json:"updated_at"`
}
