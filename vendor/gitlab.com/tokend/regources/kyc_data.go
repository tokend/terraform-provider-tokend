package regources

type KYCData struct {
	BlobID string `json:"blob_id"`
}

func (d KYCData) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"blob_id": d.BlobID,
	}
}
