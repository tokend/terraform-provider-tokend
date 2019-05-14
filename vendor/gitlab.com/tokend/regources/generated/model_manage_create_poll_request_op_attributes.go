/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManageCreatePollRequestOpAttributes struct {
	// * 0 - create * 1 - cancel
	Action xdr.ManageCreatePollRequestAction `json:"action"`
	Create *CreatePollRequestOp              `json:"create,omitempty"`
}
