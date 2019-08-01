package regources

import "encoding/json"

//ParticipantEffectsResponse - response for history and movements handlers
type ParticipantEffectsResponse struct {
	Links    *Links              `json:"links"`
	Data     []ParticipantEffect `json:"data"`
	Included Included            `json:"included"`
}

func (r ParticipantEffectsResponse) MarshalJSON() ([]byte, error) {
	if r.Data == nil {
		r.Data = []ParticipantEffect{}
	}

	type temp ParticipantEffectsResponse
	return json.Marshal(temp(r))
}

//ParticipantEffect - represent account effected by operation
type ParticipantEffect struct {
	Key
	Relationships ParticipantEffectRelation `json:"relationships"`
}

//ParticipantEffectRelation - represents relations of resource
type ParticipantEffectRelation struct {
	Operation *Relation `json:"operation"`
	Effect    *Relation `json:"effect"`
}
