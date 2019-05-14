package kv

import (
	"os"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const EnvViperConfigFile = "KV_VIPER_FILE"

var ErrNoBackends = errors.New("no backends configured")

func FromEnv() (Getter, error) {
	getters := []Getter{}
	if viperFn := os.Getenv(EnvViperConfigFile); viperFn != "" {
		getter := NewViperFile(viperFn)
		if err := pingGetter(getter); err != nil {
			return nil, errors.Wrap(err, "viper backend seems unavailable", logan.F{
				"file": viperFn,
			})
		}
		getters = append(getters, NewViperFile(viperFn))
	}

	if len(getters) == 0 {
		return nil, ErrNoBackends
	}

	return MergeGetters(getters...), nil
}

func MustFromEnv() Getter {
	getter, err := FromEnv()
	if err != nil {
		panic(errors.Wrap(err, "kv.FromEnv panicked"))
	}
	return getter
}

// pingGetter checks if getter backend is responsive and available
func pingGetter(getter Getter) error {
	_, err := getter.GetStringMap("ping")
	return err
}
