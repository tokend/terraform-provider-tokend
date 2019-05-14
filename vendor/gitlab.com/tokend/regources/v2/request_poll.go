package regources

import (
	"time"
)

// CreatePollRequest - represents details of the `issuance` reviewable request
type CreatePollRequest struct {
	Key
	Attributes    CreatePollRequestAttrs     `json:"attributes"`
	Relationships CreatePollRequestRelations `json:"relationships"`
}

// CreatePollRequestAttrs - attributes of the `issuance` reviewable request
type CreatePollRequestAttrs struct {
	PermissionType           uint32    `json:"permission_type"`
	NumberOfChoices          uint32    `json:"number_of_choices"`
	PollData                 PollData  `json:"poll_data"`
	VoteConfirmationRequired bool      `json:"vote_confirmation_required"`
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
	CreatorDetails           Details   `json:"creator_details"`
}

// CreatePollRequestRelations - relationships of the `issuance` reviewable request
type CreatePollRequestRelations struct {
	ResultProvider *Relation `json:"result_provider"`
}
