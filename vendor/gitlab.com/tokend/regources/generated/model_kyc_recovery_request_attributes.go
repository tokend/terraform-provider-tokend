/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type KycRecoveryRequestAttributes struct {
	CreatorDetails Details `json:"creator_details"`
	// Sequence number
	SequenceNumber uint32             `json:"sequence_number"`
	SignersData    []UpdateSignerData `json:"signers_data"`
}
