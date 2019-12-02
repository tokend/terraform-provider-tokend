package horizon

import (
	"net/http"
	"net/url"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/keypair"
	"gitlab.com/tokend/keypair/figurekeypair"
)

type Connectorer interface {
	Connector() *Connector
}

type connectorer struct {
	getter kv.Getter
	once   comfig.Once
}

func NewConnectorer(getter kv.Getter) *connectorer {
	return &connectorer{getter: getter}
}

func (h *connectorer) Connector() *Connector {
	return h.once.Do(func() interface{} {
		var config struct {
			Endpoint *url.URL     `fig:"endpoint,required"`
			Signer   keypair.Full `fig:"signer"`
			UseCache bool 		  `fig:"use_cache"`
		}

		err := figure.
			Out(&config).
			With(figure.BaseHooks, figurekeypair.Hooks).
			From(kv.MustGetStringMap(h.getter, "horizon")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out horizon"))
		}

		cli := NewClient(http.DefaultClient, config.Endpoint)
		if config.Signer != nil {
			cli = cli.WithSigner(config.Signer)
		}

		return NewConnector(cli, config.UseCache)
	}).(*Connector)
}
