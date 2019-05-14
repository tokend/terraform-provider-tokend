package responses

import "gitlab.com/tokend/horizon-connector/internal/resources/operations"

type ReviewRequestOperationIndex struct {
	Embedded struct {
		Records []operations.ReviewRequest `json:"records"`
	} `json:"_embedded"`
}
