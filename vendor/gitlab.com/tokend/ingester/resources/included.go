/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Resource interface {
	//GetKey - returns key of the Resource
	GetKey() Key
}

// Included - an array of Resource objects that are related to the primary data and/or
// each other (“included resources”).
type Included struct {
	includes map[Key]json.RawMessage
}

// Add - adds new include into collection. If one already present - skips it
func (c *Included) Add(includes ...Resource) {
	for i := range includes {
		c.add(includes[i])
	}
}

func (c *Included) add(include Resource) {
	if c.includes == nil {
		c.includes = make(map[Key]json.RawMessage)
	}

	_, ok := c.includes[include.GetKey()]
	if ok {
		return
	}

	data, err := json.Marshal(include)
	if err != nil {
		panic(errors.Wrap(err, "failed to add into includes"))
	}

	c.includes[include.GetKey()] = json.RawMessage(data)
}

//MarshalJSON - marshals include collection as array of json objects
func (c Included) MarshalJSON() ([]byte, error) {
	uniqueEntries := make([]json.RawMessage, 0, len(c.includes))
	for _, value := range c.includes {
		uniqueEntries = append(uniqueEntries, value)
	}

	return json.Marshal(uniqueEntries)
}

//UmarshalJSON - unmarshal array of json objects into include collection
func (c *Included) UnmarshalJSON(data []byte) error {
	var keys []Key
	err := json.Unmarshal(data, &keys)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal keys for include")
	}

	var entries []json.RawMessage
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal entries for include")
	}

	c.includes = make(map[Key]json.RawMessage)
	// we assume that json.Unmarshal guaranties the same order for arrays
	// in case of broken order for arrays - we would panic on get of specific entry
	for i := range keys {
		c.includes[keys[i]] = entries[i]
	}

	return nil
}

// tryFindEntry - tries to find entry in include collection and unmarshal it
// if entry does not exists - returns false
// if entry exists but fails to unmarshal or key mismatches - panics
func (c *Included) tryFindEntry(key Key, entry Resource) bool {
	rawEntry, exist := c.includes[key]
	if !exist {
		return false
	}

	err := json.Unmarshal(rawEntry, entry)
	if err != nil {
		panic(errors.Wrap(err, "failed to unmarshal entry from include collection"))
	}

	if entry.GetKey() != key {
		panic(errors.From(errors.New("keys mismatched"), map[string]interface{}{
			"expected_key": key,
			"actual_key":   entry.GetKey(),
		}))
	}

	return true
}
