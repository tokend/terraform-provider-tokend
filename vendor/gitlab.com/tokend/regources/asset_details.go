package regources

// AssetDetails is asset details JSON set by clients
type AssetDetails struct {
	// ExternalSystemType supposed external system type used for deposit
	ExternalSystemType int32 `json:"external_system_type,string"`
}
