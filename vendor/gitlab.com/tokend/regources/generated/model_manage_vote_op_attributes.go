/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManageVoteOpAttributes struct {
	// * 0 - create * 1 - remove
	Action xdr.ManageVoteAction `json:"action"`
	Create *CreateVoteOp        `json:"create,omitempty"`
	Remove *RemoveVoteOp        `json:"remove,omitempty"`
}
