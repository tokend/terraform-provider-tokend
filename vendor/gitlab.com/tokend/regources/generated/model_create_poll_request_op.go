/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type CreatePollRequestOp struct {
	// tasks set on request creation
	AllTasks       *uint32 `json:"all_tasks,omitempty"`
	CreatorDetails Details `json:"creator_details"`
	// the date until which voting in the poll will be allowed
	EndTime time.Time `json:"end_time"`
	// count of allowed choices
	NumberOfChoices uint64 `json:"number_of_choices"`
	// is used to restrict using of poll through rules
	PermissionType uint64   `json:"permission_type"`
	PollData       PollData `json:"poll_data"`
	// id of account who could close the poll
	ResultProviderId string `json:"result_provider_id"`
	// the date from which voting in the poll will be allowed
	StartTime time.Time `json:"start_time"`
	// defines if result provider signature is required to participate in poll voting
	VoteConfirmationRequired bool `json:"vote_confirmation_required"`
}
