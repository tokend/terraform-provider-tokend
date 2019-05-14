/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type VoteData struct {
	// type of the poll
	PollType     xdr.PollType `json:"poll_type"`
	SingleChoice *uint64      `json:"single_choice,omitempty"`
}
