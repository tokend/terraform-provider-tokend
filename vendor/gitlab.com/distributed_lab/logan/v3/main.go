package logan

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"io"
)

func New() *Entry {
	lastLogLevel := AllLevels[len(AllLevels)-1]

	logger := logrus.New()
	logger.Level = logrus.Level(lastLogLevel)

	// need default formatter, but with nano seconds for time
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = time.RFC3339Nano
	logger.Formatter = formatter

	return &Entry{
		logrus.NewEntry(logger).WithField("pid", os.Getpid()),
	}
}

func (e *Entry) Out(writer io.Writer) *Entry {
	logger := copyLogger(e.entry.Logger)
	logger.Out = writer

	return &Entry{
		logrus.NewEntry(logger).WithField("pid", os.Getpid()),
	}
}

func (e *Entry) Formatter(formatter Formatter) *Entry {
	logger := copyLogger(e.entry.Logger)
	logger.Formatter = formatter

	return &Entry{
		logrus.NewEntry(logger).WithField("pid", os.Getpid()),
	}
}

func (e *Entry) Level(level Level) *Entry {
	logger := copyLogger(e.entry.Logger)
	logger.Level = logrus.Level(level)

	e.entry.Logger = logger

	return &Entry{
		logrus.NewEntry(logger).WithField("pid", os.Getpid()),
	}
}

func copyLogger(l *logrus.Logger) *logrus.Logger {
	result := logrus.New()

	result.Out = l.Out
	result.Hooks = l.Hooks
	result.Formatter = l.Formatter
	result.Level = l.Level

	return result
}
