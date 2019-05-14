package resources

type Signer struct {
	PublicKey string `json:"public_key"`
	AccountID string `json:"account_id"`
	Weight    int    `json:"weight"`
	Role      uint64 `json:"role"`
	Identity  uint32 `json:"signer_identity"`
}
