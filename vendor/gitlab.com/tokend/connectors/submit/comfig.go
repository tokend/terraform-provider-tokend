package submit

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/connectors/keyer"
	"gitlab.com/tokend/connectors/signed"
	"gitlab.com/tokend/keypair/figurekeypair"
	"net/http"
	"net/url"
)

type Submission interface {
	Submit() *Submitter
}

type submission struct {
	getter kv.Getter
	once   comfig.Once

	keyer.Keyer
}

func NewSubmission(getter kv.Getter) *submission {
	return &submission{
		getter: getter,
		Keyer:  keyer.NewKeyer(getter),
	}
}

func (h *submission) Submit() *Submitter {
	return h.once.Do(func() interface{} {

		keys := h.Keyer.Keys()
		var config struct {
			Endpoint *url.URL `fig:"endpoint,required"`
		}

		err := figure.
			Out(&config).
			With(figure.BaseHooks, figurekeypair.Hooks).
			From(kv.MustGetStringMap(h.getter, "submit")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out submit"))
		}

		cli := signed.NewClient(http.DefaultClient, config.Endpoint)
		if keys.Signer != nil {
			cli = cli.WithSigner(keys.Signer)
		}

		return New(cli)
	}).(*Submitter)
}
