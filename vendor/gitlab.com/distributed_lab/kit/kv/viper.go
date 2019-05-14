package kv

import (
	"sync"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/spf13/viper"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type viperbackend struct {
	viper *viper.Viper

	read    sync.Once
	readErr error
}

func NewViperFile(fn string) Getter {
	v := viper.New()
	v.SetConfigFile(fn)
	return &viperbackend{
		viper: v,
	}
}

func (v *viperbackend) ensureRead() error {
	v.read.Do(func() {
		if err := v.viper.ReadInConfig(); err != nil {
			v.readErr = errors.Wrap(err, "failed to read config", logan.F{
				"config_file": v.viper.ConfigFileUsed(),
			})
		}
	})
	return v.readErr
}

func (v *viperbackend) GetStringMap(key string) (map[string]interface{}, error) {
	if err := v.ensureRead(); err != nil {
		return nil, err
	}
	return v.viper.GetStringMap(key), nil
}
