package regources

// PublicKeyEntryResponse - response on /public_key request
type PublicKeyEntryResponse struct {
	Data     PublicKeyEntry `json:"data"`
	Included Included       `json:"included"`
}

// PublicKeyEntry - Resource object representing "public key" resource
type PublicKeyEntry struct {
	Key
	Relationships PublicKeyEntryRelationships `json:"relationships"`
}

// PublicKeyEntryRelationships - the relationships of the PublicKey resource
type PublicKeyEntryRelationships struct {
	Accounts *RelationCollection `json:"accounts"`
}
