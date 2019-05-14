package regources

import "encoding/json"

type Relation struct {
	Links *Links `json:"links,omitempty"`
	Data  *Key   `json:"data,omitempty"`
}

type RelationCollection struct {
	Links *Links `json:"links,omitempty"`
	Data  []Key  `json:"data,omitempty"`
}

func (r RelationCollection) MarshalJSON() ([]byte, error) {
	if r.Data == nil {
		r.Data = []Key{}
	}

	type temp RelationCollection
	return json.Marshal(temp(r))
}
