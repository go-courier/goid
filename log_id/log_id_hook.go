package log_id

import (
	"github.com/sirupsen/logrus"
)

func NewLogIDHook(logIDMap *LogIDMap) *LogIDHook {
	return &LogIDHook{
		LogIDMap: logIDMap,
	}
}

type LogIDHook struct {
	*LogIDMap
}

func (LogIDHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *LogIDHook) Fire(entry *logrus.Entry) error {
	entry.Data["log_id"] = h.Get()
	return nil
}
