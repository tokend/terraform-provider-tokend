/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type SignerAttributes struct {
	Details Details `json:"details"`
	// identity of the signer (signatures of signers with same identities are considered as one signature with max weight)
	Identity uint32 `json:"identity"`
	// weight of the signature created by the signer
	Weight uint32 `json:"weight"`
}
