package signed

import (
	"net/http"
	"net/url"

	"gitlab.com/tokend/connectors/keyer"
	"gitlab.com/tokend/keypair"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/keypair/figurekeypair"
)

type Clienter interface {
	Client() *Client
	UnthrottledClient() *RawClient
}

type clienter struct {
	getter kv.Getter
	keyer.Keyer
	once            comfig.Once
	onceUnthrottled comfig.Once
}

func NewClienter(getter kv.Getter) *clienter {
	return &clienter{
		getter: getter,
		Keyer:  keyer.NewKeyer(getter),
	}
}

type config struct {
	Endpoint *url.URL        `fig:"endpoint,required"`
	Signer   keypair.Full    `fig:"signer"`
	Source   keypair.Address `fig:"source"`
}

func (h *clienter) UnthrottledClient() *RawClient {
	return h.onceUnthrottled.Do(func() interface{} {
		config := mustPleaseConfig(h.getter)

		var keys keyer.Keys

		if config.Signer != nil {
			keys = keyer.Keys{
				Signer: config.Signer,
				Source: config.Source,
			}
		} else {
			keys = h.Keyer.Keys()
		}

		cli := NewRawClient(http.DefaultClient, config.Endpoint)
		if keys.Signer != nil {
			cli = cli.WithSigner(keys.Signer)
		}
		if keys.Source != nil {
			cli = cli.WithSource(keys.Source)
		}

		return cli
	}).(*RawClient)
}

func (h *clienter) Client() *Client {
	return h.once.Do(func() interface{} {
		config := mustPleaseConfig(h.getter)

		var keys keyer.Keys

		if config.Signer != nil {
			keys = keyer.Keys{
				Signer: config.Signer,
				Source: config.Source,
			}
		} else {
			keys = h.Keyer.Keys()
		}

		cli := NewClient(http.DefaultClient, config.Endpoint)
		if keys.Signer != nil {
			cli = cli.WithSigner(keys.Signer)
		}
		if keys.Source != nil {
			cli = cli.WithSource(keys.Source)
		}

		return cli
	}).(*Client)
}

func mustPleaseConfig(getter kv.Getter) *config {
	var cfg config

	err := figure.
		Out(&cfg).
		With(figure.BaseHooks, figurekeypair.Hooks).
		From(kv.MustGetStringMap(getter, "client")).
		Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to figure out client"))
	}

	return &cfg
}
