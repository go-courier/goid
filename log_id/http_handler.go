package log_id

import (
	"net/http"

	"github.com/google/uuid"
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
	requestID := req.Header.Get("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
	}

	do := h.logIDMap.With(func() {
		h.nextHandler.ServeHTTP(rw, req)
	}, requestID)

	do()
}
