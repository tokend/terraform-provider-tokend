/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ParticipantsEffect struct {
	Key
	Relationships ParticipantsEffectRelationships `json:"relationships"`
}
type ParticipantsEffectResponse struct {
	Data     ParticipantsEffect `json:"data"`
	Included Included           `json:"included"`
}

type ParticipantsEffectListResponse struct {
	Data     []ParticipantsEffect `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *ParticipantsEffectListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ParticipantsEffectListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustParticipantsEffect - returns ParticipantsEffect from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustParticipantsEffect(key Key) *ParticipantsEffect {
	var participantsEffect ParticipantsEffect
	if c.tryFindEntry(key, &participantsEffect) {
		return &participantsEffect
	}
	return nil
}
