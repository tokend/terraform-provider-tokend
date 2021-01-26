/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateKycRecoveryRequestOpAttributes struct {
	// tasks set on request creation
	AllTasks       *uint32            `json:"all_tasks,omitempty"`
	CreatorDetails Details            `json:"creator_details"`
	SignersData    []UpdateSignerData `json:"signers_data"`
}
