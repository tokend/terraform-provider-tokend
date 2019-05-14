package comfig

import (
	"sync"

	raven "github.com/getsentry/raven-go"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type Sentrier interface {
	Sentry() *raven.Client
}

type sentrier struct {
	getter  kv.Getter
	once    sync.Once
	value   *raven.Client
	options SentryOpts
	err     error
}

type SentryOpts struct {
	Release string
}

func NewSentrier(getter kv.Getter, options SentryOpts) Sentrier {
	return &sentrier{
		getter:  getter,
		options: options,
	}
}

func (s *sentrier) Sentry() *raven.Client {
	s.once.Do(func() {
		var config = struct {
			DSN     string `fig:"dsn,required"`
			Release string `fig:"release"`
		}{
			Release: s.options.Release,
		}

		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(s.getter, "sentry")).
			Please()
		if err != nil {
			s.err = errors.Wrap(err, "failed to figure out")
			return
		}
		if config.Release == "" {
			s.err = errors.New("please, set release for sentry")
			return
		}

		client, err := raven.New(config.DSN)
		if err != nil {
			s.err = errors.Wrap(err, "failed to init sentry client")
			return
		}
		client.SetRelease(config.Release)

		// TODO tags

		s.value = client
	})
	if s.err != nil {
		panic(s.err)
	}
	return s.value
}
