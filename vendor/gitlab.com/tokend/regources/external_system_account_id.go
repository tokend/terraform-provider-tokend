package regources

type ExternalSystemAccountID struct {
	Type      Flag    `json:"type"`
	Data      string  `json:"data"`
	AssetCode string  `json:"asset_code,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
}
