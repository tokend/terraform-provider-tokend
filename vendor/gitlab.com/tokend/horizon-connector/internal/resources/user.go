package resources

type User struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes UserAttributes
}

type UserAttributes struct {
	Email           string `json:"email"`
	LastIPAddr      string `json:"last_ip_address"`
	State           string `json:"state"`
	KYCSequence     int    `json:"kyc_sequence"`
	RejectReason    string `json:"reject_reason"`
	RecoveryAddress string `json:"recovery_address"`
	CreatedAt       string `json:"created_at"`
	AirdropState    string `json:"airdrop_state"`
}

func (u User) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"type":             u.Type,
		"id":               u.ID,
		"email":            u.Attributes.Email,
		"last_ip_addr":     u.Attributes.LastIPAddr,
		"state":            u.Attributes.State,
		"kyc_sequence":     u.Attributes.KYCSequence,
		"reject_reason":    u.Attributes.RejectReason,
		"recovery_address": u.Attributes.RecoveryAddress,
		"created_at":       u.Attributes.CreatedAt,
		"airdrop_state":    u.Attributes.AirdropState,
	}
}
