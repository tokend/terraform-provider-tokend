/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManageLimitsOpAttributes struct {
	// * 0: \"create\", * 1: \"remove\"
	Action xdr.ManageLimitsAction  `json:"action"`
	Create *ManageLimitsCreationOp `json:"create,omitempty"`
	Remove *ManageLimitsRemovalOp  `json:"remove,omitempty"`
}
