package signed

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

type Clienter interface {
	Client() *Client
}

type clienter struct {
	getter kv.Getter
	once   comfig.Once
}

func NewClienter(getter kv.Getter) *clienter {
	return &clienter{getter: getter}
}

func (h *clienter) Client() *Client {
	return h.once.Do(func() interface{} {
		var config struct {
			Endpoint *url.URL        `fig:"endpoint,required"`
			Signer   keypair.Full    `fig:"signer"`
			Source   keypair.Address `fig:"source"`
		}

		err := figure.
			Out(&config).
			With(figure.BaseHooks, figurekeypair.Hooks).
			From(kv.MustGetStringMap(h.getter, "client")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out client"))
		}

		cli := NewClient(http.DefaultClient, config.Endpoint)
		if config.Signer != nil && config.Source != nil {
			cli = cli.WithSigner(config.Source, config.Signer)
		}

		return cli
	}).(*Client)
}
