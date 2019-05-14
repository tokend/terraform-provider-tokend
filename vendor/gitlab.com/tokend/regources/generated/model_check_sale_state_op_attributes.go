/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type CheckSaleStateOpAttributes struct {
	// * 1: \"canceled\" * 2: \"closed\" * 3: \"updated\"
	Effect xdr.CheckSaleStateEffect `json:"effect"`
}
