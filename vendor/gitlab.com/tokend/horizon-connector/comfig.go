package horizon

import (
	"net/url"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/keypair"
	"gitlab.com/tokend/keypair/figurekeypair"
)

type Horizoner struct {
	getter kv.Getter
	once   comfig.Once
	value  *Connector
}

func NewHorizoner(getter kv.Getter) *Horizoner {
	return &Horizoner{getter: getter}
}

func (h *Horizoner) Horizon() *Connector {
	return h.once.Do(func() interface{} {
		var config struct {
			Endpoint *url.URL     `fig:"endpoint,required"`
			Signer   keypair.Full `fig:"signer"`
		}

		err := figure.
			Out(&config).
			With(figure.BaseHooks, figurekeypair.Hooks).
			From(kv.MustGetStringMap(h.getter, "horizon")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out horizon"))
		}

		hrz := NewConnector(config.Endpoint)
		if config.Signer != nil {
			hrz = hrz.WithSigner(config.Signer)
		}
		return hrz
	}).(*Connector)
}
