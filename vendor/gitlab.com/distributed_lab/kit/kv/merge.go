package kv

import (
	"fmt"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func MergeGetters(backends ...Getter) Getter {
	return &getter{
		backends: backends,
	}
}

type getter struct {
	backends []Getter
}

func (g getter) GetStringMap(key string) (map[string]interface{}, error) {
	for _, backend := range g.backends {
		value, err := backend.GetStringMap(key)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get key", logan.F{
				"backend": fmt.Sprintf("%T", backend),
			})
		}
		if value != nil {
			return value, nil
		}
	}
	return nil, nil
}
