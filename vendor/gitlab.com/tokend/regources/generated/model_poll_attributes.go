/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type PollAttributes struct {
	CreatorDetails Details `json:"creator_details"`
	// the date until which voting in the poll will be allowed
	EndTime time.Time `json:"end_time"`
	// count of allowed choices
	NumberOfChoices uint32 `json:"number_of_choices"`
	// is used to restrict using of poll through rules
	PermissionType uint32   `json:"permission_type"`
	PollData       PollData `json:"poll_data"`
	// * 1 - open * 2 - passed * 3 - failed * 4 - cancelled
	PollState PollState `json:"poll_state"`
	// the date from which voting in the poll will be allowed
	StartTime time.Time `json:"start_time"`
	// defines if result provider is required to participate in poll voting
	VoteConfirmationRequired bool `json:"vote_confirmation_required"`
}
