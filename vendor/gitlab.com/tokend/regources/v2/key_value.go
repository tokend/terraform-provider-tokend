package regources

import "gitlab.com/tokend/go/xdr"

type KeyValueEntryResponse struct {
	Data     KeyValueEntry `json:"data"`
	Included Included      `json:"included"`
}

type KeyValueEntriesResponse struct {
	Links    *Links          `json:"links"`
	Data     []KeyValueEntry `json:"data"`
	Included Included        `json:"included"`
}

type KeyValueEntry struct {
	Key
	Attributes KeyValueEntryAttrs `json:"attributes"`
}

type KeyValueEntryAttrs struct {
	Value KeyValueEntryValue `json:"value"`
}

//KeyValueEntryValue - represents xdr.KeyValueEntryValue
type KeyValueEntryValue struct {
	Type xdr.KeyValueEntryType `json:"type"`
	U32  *uint32               `json:"u_32,omitempty"`
	Str  *string               `json:"str,omitempty"`
	U64  *uint64               `json:"u_64,omitempty"`
}
