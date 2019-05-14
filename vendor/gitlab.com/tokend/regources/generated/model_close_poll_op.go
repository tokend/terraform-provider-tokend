/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ClosePollOp struct {
	Details Details `json:"details"`
	// id of the closed poll
	PollId int64 `json:"poll_id"`
	// * 0 - passed * 1 - failed
	PollResult xdr.PollResult `json:"poll_result"`
}
