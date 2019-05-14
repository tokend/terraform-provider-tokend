/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateIssuanceRequestOpAttributes struct {
	// tasks set on request creation
	AllTasks       *int64  `json:"all_tasks,omitempty"`
	Amount         Amount  `json:"amount"`
	CreatorDetails Details `json:"creator_details"`
	Fee            Fee     `json:"fee"`
	// reference of the request
	Reference string `json:"reference"`
}
