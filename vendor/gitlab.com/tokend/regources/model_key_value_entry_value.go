/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type KeyValueEntryValue struct {
	Str  *string               `json:"str,omitempty"`
	Type xdr.KeyValueEntryType `json:"type"`
	U32  *uint32               `json:"u32,omitempty"`
	U64  *uint64               `json:"u64,omitempty"`
}
