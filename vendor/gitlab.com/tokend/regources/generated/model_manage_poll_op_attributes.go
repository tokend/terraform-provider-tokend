/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManagePollOpAttributes struct {
	// * 0 - close * 1 - update end time * 2 - cancel
	Action xdr.ManagePollAction `json:"action"`
	Close  *ClosePollOp         `json:"close,omitempty"`
	// ID of the poll to manage
	PollId        int64                `json:"poll_id"`
	UpdateEndTime *UpdatePollEndTimeOp `json:"update_end_time,omitempty"`
}
