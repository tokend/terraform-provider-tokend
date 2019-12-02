/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type SignerData struct {
	Details   Details  `json:"details"`
	Identity  uint32   `json:"identity"`
	PublicKey string   `json:"public_key"`
	RoleIds   []uint64 `json:"role_ids"`
	// Weight of the signer of the account
	Weight uint32 `json:"weight"`
}
