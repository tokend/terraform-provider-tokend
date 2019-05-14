package kv

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Getter interface {
	// GetStringMap expected to return nil if key does not exist
	GetStringMap(key string) (map[string]interface{}, error)
}

func MustGetStringMap(getter Getter, key string) map[string]interface{} {
	value, err := getter.GetStringMap(key)
	if err != nil {
		panic(errors.Wrap(err, "failed to get key", logan.F{
			"key": key,
		}))
	}
	return value
}
