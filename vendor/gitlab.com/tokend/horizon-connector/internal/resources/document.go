package resources

type Document struct {
	URL         string `json:"url"`
}

func (d Document) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"url": d.URL,
	}
}
