package log_id

import (
	"net/http"
)

func LogIDRoundTripper(idMap *LogIDMap) func(rt http.RoundTripper) http.RoundTripper {
	if idMap == nil {
		idMap = Default
	}

	return func(rt http.RoundTripper) http.RoundTripper {
		return &logIdRoundTripper{
			logIDMap: idMap,
			next:     rt,
		}
	}
}

type logIdRoundTripper struct {
	logIDMap *LogIDMap
	next     http.RoundTripper
}

func (h *logIdRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Request-ID", h.logIDMap.Get())
	if h.next != nil {
		return h.next.RoundTrip(req)
	}
	return nil, nil
}
