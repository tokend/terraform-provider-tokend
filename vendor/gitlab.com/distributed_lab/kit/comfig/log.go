package comfig

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	sentryhook "github.com/xr9kayu/logrus/sentry"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

type Logger interface {
	Log() *logan.Entry
}

type logger struct {
	getter  kv.Getter
	once    Once
	options LoggerOpts
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
	return l.once.Do(func() interface{} {
		config, err := parseLogConfig(kv.MustGetStringMap(l.getter, "log"))
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out logger"))
		}

		entry := logan.New().Level(config.Level)

		if !config.DisableSentry {
			sentrier := NewSentrier(l.getter, SentryOpts{Release: l.options.Release})
			sentry := sentrier.Sentry()
			sentryConfig := sentrier.SentryConfig()

			var selectedLevel logrus.Level
			if sentryConfig.Level != nil {
				selectedLevel = logrus.Level(*sentryConfig.Level)
			} else {
				selectedLevel = logrus.Level(config.Level)
			}
			levels := make([]logrus.Level, 0)
			for level := logrus.PanicLevel; level <= selectedLevel; level++ {
				levels = append(levels, level)
			}
			hook, err := sentryhook.NewHook(sentry, levels...)
			if err != nil {
				panic(errors.Wrap(err, "failed to init sentry hook"))
			}

			entry.AddLogrusHook(hook)
		}

		return entry
	}).(*logan.Entry)
}

type loggerConfig struct {
	Level         logan.Level `fig:"level"`
	DisableSentry bool        `fig:"disable_sentry"`
}

func parseLogConfig(raw map[string]interface{}) (*loggerConfig, error) {
	config := loggerConfig{
		Level: logan.ErrorLevel,
	}

	err := figure.
		Out(&config).
		From(raw).
		Please()
	if err != nil {
		return nil, errors.Wrap(err, "failed to figure out")
	}

	return &config, nil
}
