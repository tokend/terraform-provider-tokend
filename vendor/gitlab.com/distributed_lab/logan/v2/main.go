package logan

import (
	"os"

	"github.com/sirupsen/logrus"
	"io"
)

// DEPRECATED: Use Entry from logan/v3 instead
func New() *Entry {
	lastLogLevel := AllLevels[len(AllLevels)-1]

	logger := logrus.New()
	logger.Level = logrus.Level(lastLogLevel)

	return &Entry{
		logrus.NewEntry(logger).WithField("pid", os.Getpid()),
	}
}

// DEPRECATED: Use logan/v3 instead
func (e *Entry) Out(writer io.Writer) *Entry {
	logger := copyLogger(e.Entry.Logger)
	logger.Out = writer

	return &Entry{
		logrus.NewEntry(logger).WithField("pid", os.Getpid()),
	}
}

// DEPRECATED: Use logan/v3 instead
func (e *Entry) Formatter(formatter Formatter) *Entry {
	logger := copyLogger(e.Entry.Logger)
	logger.Formatter = formatter

	return &Entry{
		logrus.NewEntry(logger).WithField("pid", os.Getpid()),
	}
}

// DEPRECATED: Use logan/v3 instead
func (e *Entry) Level(level Level) *Entry {
	logger := copyLogger(e.Entry.Logger)
	logger.Level = logrus.Level(level)

	return &Entry{
		logrus.NewEntry(logger).WithField("pid", os.Getpid()),
	}
}

// DEPRECATED
func copyLogger(l *logrus.Logger) *logrus.Logger {
	result := logrus.New()

	result.Out = l.Out
	result.Hooks = l.Hooks
	result.Formatter = l.Formatter
	result.Level = l.Level

	return result
}
