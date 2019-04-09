package goid

import (
	"runtime"
	"sync"

	"github.com/go-courier/metax"
)

var Default = &GoIDMetaMap{}

type GoIDMetaMap struct {
	m sync.Map
}

func (m *GoIDMetaMap) Clear() {
	m.m.Delete(runtime.GoID())
}

func (m *GoIDMetaMap) Get() metax.Meta {
	if logID, ok := m.m.Load(runtime.GoID()); ok {
		return logID.(metax.Meta)
	}
	return metax.Meta{}
}

func (m *GoIDMetaMap) Set(meta metax.Meta) {
	m.m.Store(runtime.GoID(), meta)
}

func (m *GoIDMetaMap) With(cb func(), metas ...metax.Meta) func() {
	meta := metax.Meta{}

	if len(metas) == 0 {
		meta = m.Get()
	} else {
		meta = meta.Merge(metas...)
	}

	return func() {
		m.Set(meta)
		defer m.Clear()
		cb()
	}
}

func (m *GoIDMetaMap) All() map[int64]metax.Meta {
	results := map[int64]metax.Meta{}

	m.m.Range(func(key, value interface{}) bool {
		results[key.(int64)] = value.(metax.Meta)
		return true
	})

	return results
}
