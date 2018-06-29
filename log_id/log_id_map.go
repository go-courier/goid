package log_id

import (
	"runtime"
	"sync"
)

var Default = &LogIDMap{}

type LogIDMap struct {
	m sync.Map
}

func (m *LogIDMap) Clear() {
	m.m.Delete(runtime.GoID())
}

func (m *LogIDMap) Get() string {
	if logID, ok := m.m.Load(runtime.GoID()); ok {
		return logID.(string)
	}
	return ""
}

func (m *LogIDMap) Set(id string) {
	m.m.Store(runtime.GoID(), id)
}

func (m *LogIDMap) All() map[int64]string {
	results := map[int64]string{}

	m.m.Range(func(key, value interface{}) bool {
		results[key.(int64)] = value.(string)
		return true
	})

	return results
}
