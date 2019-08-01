/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type PollParticipation struct {
	Key
	Relationships PollParticipationRelationships `json:"relationships"`
}
type PollParticipationResponse struct {
	Data     PollParticipation `json:"data"`
	Included Included          `json:"included"`
}

type PollParticipationListResponse struct {
	Data     []PollParticipation `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustPollParticipation - returns PollParticipation from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPollParticipation(key Key) *PollParticipation {
	var pollParticipation PollParticipation
	if c.tryFindEntry(key, &pollParticipation) {
		return &pollParticipation
	}
	return nil
}
