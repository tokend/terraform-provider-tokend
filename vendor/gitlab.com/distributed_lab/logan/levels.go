package logan

import "github.com/sirupsen/logrus"

// DEPRECATED: Use logan/v3 instead
type Level logrus.Level

const (
	// Use logan/v3 instead

	// DEPRECATED
	PanicLevel Level = iota
	// DEPRECATED
	FatalLevel
	// DEPRECATED
	ErrorLevel
	// DEPRECATED
	WarnLevel
	// DEPRECATED
	InfoLevel
	// DEPRECATED
	DebugLevel
)

// DEPRECATED: Use logan/v3 instead
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
}

// DEPRECATED: Use logan/v3 instead
func ParseLevel(level string) (Level, error) {
	l, err := logrus.ParseLevel(level)
	return Level(l), err
}
