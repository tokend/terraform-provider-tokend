package logan

import "github.com/sirupsen/logrus"

// DEPRECATED: Use logan/v3 instead
type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}

// helper struct implementing logrus.Hook for feeding logan.Hook to logrus
//
// DEPRECATED
type hookConnector struct {
	levels []logrus.Level
	fireFn func(*logrus.Entry) error
}

// DEPRECATED: Use logan/v3 instead
func (h *hookConnector) Levels() []logrus.Level {
	return h.levels
}

// DEPRECATED: Use logan/v3 instead
func (h *hookConnector) Fire(entry *logrus.Entry) error {
	return h.fireFn(entry)
}

// DEPRECATED
func castLevels(levels []Level) []logrus.Level {
	result := make([]logrus.Level, len(levels))
	for _, level := range levels {
		result = append(result, logrus.Level(level))
	}
	return result
}

// DEPRECATED
func castFireFn(fire func(*Entry) error) func(*logrus.Entry) error {
	return func(entry *logrus.Entry) error {
		return fire(&Entry{entry})
	}
}

// DEPRECATED: Use logan/v3 instead
func (e *Entry) AddHook(hook Hook) {
	hooks := e.Entry.Logger.Hooks
	connector := &hookConnector{
		levels: castLevels(hook.Levels()),
		fireFn: castFireFn(hook.Fire),
	}
	for _, level := range hook.Levels() {
		llevel := logrus.Level(level)
		hooks[llevel] = append(hooks[llevel], connector)
	}
}

// DEPRECATED: Use logan/v3 instead
func (e *Entry) AddLogrusHook(hook logrus.Hook) {
	e.Entry.Logger.Hooks.Add(hook)
}
