package comfig

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

type Sentrier interface {
	Sentry() *sentry.Client
	SentryConfig() SentryConfig
}

type sentrier struct {
	getter       kv.Getter
	once         sync.Once
	client       *sentry.Client
	sentryConfig SentryConfig
	options      SentryOpts
	err          error
}

type SentryOpts struct {
	Release string
}

type SentryConfig struct {
	DSN     string       `fig:"dsn,required"`
	Release string       `fig:"release"`
	Level   *logan.Level `fig:"level"`
}

func NewSentrier(getter kv.Getter, options SentryOpts) Sentrier {
	return &sentrier{
		getter:  getter,
		options: options,
	}
}

func (s *sentrier) Sentry() *sentry.Client {
	s.readConfig()
	if s.err != nil {
		panic(s.err)
	}
	return s.client
}

func (s *sentrier) SentryConfig() SentryConfig {
	s.readConfig()
	if s.err != nil {
		panic(s.err)
	}
	return s.sentryConfig
}

func (s *sentrier) readConfig() {
	s.once.Do(func() {
		var config = SentryConfig{
			Release: s.options.Release,
		}

		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(s.getter, "sentry")).
			With(figure.Merge(figure.BaseHooks, figure.Hooks{
				"*logan.Level": func(value interface{}) (reflect.Value, error) {
					if value == nil {
						return reflect.ValueOf(nil), nil
					}
					switch v := value.(type) {
					case string:
						lvl, err := logan.ParseLevel(v)
						if err != nil {
							return reflect.Value{}, errors.Wrap(err, "failed to parse log level")
						}
						return reflect.ValueOf(&lvl), nil
					default:
						return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
					}
				},
			})).
			Please()
		if err != nil {
			s.err = errors.Wrap(err, "failed to figure out")
			return
		}
		if config.Release == "" {
			s.err = errors.New("please, set release for sentry")
			return
		}

		client, err := sentry.NewClient(sentry.ClientOptions{Dsn: config.DSN})
		if err != nil {
			s.err = errors.Wrap(err, "failed to init sentry client")
			return
		}
		// TODO tags
		s.client = client
		s.sentryConfig = config
	})
}
