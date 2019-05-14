package resources

type Blob struct {
	ID            string                      `json:"id,omitempty"`
	Type          string                      `json:"type"`
	Attributes    BlobAttributes              `json:"attributes"`
	Relationships map[string]BlobRelationship `json:"relationships,omitempty"`
}

func (b Blob) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"id":         b.ID,
		"type":       b.Type,
		"attributes": b.Attributes,
	}
}

type BlobAttributes struct {
	Value string `json:"value"`
}

func (a BlobAttributes) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"value": a.Value,
	}
}


func (b *Blob) AddRelationship(key, value string) {
	if b.Relationships == nil {
		b.Relationships = make(map[string]BlobRelationship)
	}

	b.Relationships[key] = BlobRelationship{
		Data: BlobRelationshipData{ID: value},
	}
}

type BlobRelationship struct {
	Data BlobRelationshipData `json:"data"`
}

type BlobRelationshipData struct {
	ID string `json:"id"`
}
