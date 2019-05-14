/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManagePollOpAttributes struct {
	// * 0 - close
	Action xdr.ManagePollAction `json:"action"`
	Close  *ClosePollOp         `json:"close,omitempty"`
}
