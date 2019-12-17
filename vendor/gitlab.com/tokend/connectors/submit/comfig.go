package submit

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/json-api-connector/signed"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/keypair"
	"gitlab.com/tokend/keypair/figurekeypair"
	"net/http"
	"net/url"
)

type Submitter interface {
	Submit() *Transactor
}

type submitter struct {
	getter kv.Getter
	once   comfig.Once
}

func NewSubmitter(getter kv.Getter) *submitter {
	return &submitter{getter: getter}
}

func (h *submitter) Submit() *Transactor {
	return h.once.Do(func() interface{} {
		var config struct {
			Endpoint *url.URL        `fig:"endpoint,required"`
			Signer   keypair.Full    `fig:"signer"`
			Source   keypair.Address `fig:"source"`
		}

		err := figure.
			Out(&config).
			With(figure.BaseHooks, figurekeypair.Hooks).
			From(kv.MustGetStringMap(h.getter, "submit")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out api"))
		}

		cli := signed.NewClient(http.DefaultClient, config.Endpoint)
		if config.Signer != nil && config.Source != nil {
			cli = cli.WithSigner(config.Source, config.Signer)
		}

		return New(cli)
	}).(*Transactor)
}
