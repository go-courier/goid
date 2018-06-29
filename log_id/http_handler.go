package log_id

import (
	"net/http"
)

func LogIDHttpHandler(idMap *LogIDMap) func(handler http.Handler) http.Handler {
	if idMap == nil {
		idMap = Default
	}

	return func(handler http.Handler) http.Handler {
		return &logIdHandler{
			logIDMap:    idMap,
			nextHandler: handler,
		}
	}
}

type logIdHandler struct {
	logIDMap    *LogIDMap
	nextHandler http.Handler
}

func (h *logIdHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.logIDMap.Set(req.Header.Get("X-Request-ID"))
	defer h.logIDMap.Clear()

	h.nextHandler.ServeHTTP(rw, req)
}
