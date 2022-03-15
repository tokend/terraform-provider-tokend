package keyer

import (
	"fmt"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/keypair"
	"gitlab.com/tokend/keypair/figurekeypair"
)

type Keys struct {
	Signer keypair.Full    `fig:"signer"`
	Source keypair.Address `fig:"source"`
}

type Keyer interface {
	Keys() Keys
}

type keyer struct {
	once   comfig.Once
	getter kv.Getter
	slug   string
}

func NewKeyer(getter kv.Getter) *keyer {
	return &keyer{
		getter: getter,
		slug:   "keys",
	}
}

func NewNamedKeyer(getter kv.Getter, slug string) *keyer {
	return &keyer{
		getter: getter,
		slug:   slug,
	}
}

func (k *keyer) Keys() Keys {
	return k.once.Do(func() interface{} {
		var config Keys
		err := figure.
			Out(&config).
			With(figure.BaseHooks, figurekeypair.Hooks).
			From(kv.MustGetStringMap(k.getter, k.slug)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, fmt.Sprintf("failed to figure out %s", k.slug)))
		}

		return config
	}).(Keys)
}
