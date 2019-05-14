package comfig

import (
	"sync"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

type Logger interface {
	Log() *logan.Entry
}

type logger struct {
	getter  kv.Getter
	once    sync.Once
	value   *logan.Entry
	options LoggerOpts
	err     error
}

type LoggerOpts struct {
	Release string
}

func NewLogger(getter kv.Getter, options LoggerOpts) Logger {
	return &logger{
		getter:  getter,
		options: options,
	}
}

func (l *logger) Log() *logan.Entry {
	l.once.Do(func() {
		var config = struct {
			Level         logan.Level `fig:"level"`
			DisableSentry bool        `fig:"disable_sentry"`
		}{
			Level: logan.ErrorLevel,
		}

		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(l.getter, "log")).
			Please()
		if err != nil {
			l.err = errors.Wrap(err, "failed to figure out")
			return
		}

		entry := logan.New().Level(config.Level)

		if !config.DisableSentry {
			sentry := NewSentrier(l.getter, SentryOpts{Release: l.options.Release}).Sentry()

			// TODO set sentry level?
			levels := []logrus.Level{
				logrus.ErrorLevel,
				logrus.FatalLevel,
				logrus.PanicLevel,
			}

			hook, err := logrus_sentry.NewWithClientSentryHook(sentry, levels)
			if err != nil {
				l.err = errors.Wrap(err, "failed to init sentry hook")
				return
			}
			hook.Timeout = 1 * time.Second
			entry.AddLogrusHook(hook)
		}
		l.value = entry
	})
	if l.err != nil {
		panic(l.err)
	}
	return l.value
}
