/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type ExternalSystemIdAttributes struct {
	// the time when the external system ID was binded
	BindedAt time.Time `json:"binded_at"`
	// identifier of an account in the external system.
	Data string `json:"data"`
	// this ID can be binded to another account in the system after the expiration time
	ExpiresAt time.Time `json:"expires_at"`
	// type of the external system
	ExternalSystemType int32 `json:"external_system_type"`
	// if true, this external system ID will not be released back to bool after the expiration but will rather be removed
	IsDeleted bool `json:"is_deleted"`
}
