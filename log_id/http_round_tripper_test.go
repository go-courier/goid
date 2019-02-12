package log_id

import (
	"fmt"
	"net/http"
	"testing"
)

type roundTrigger struct {
}

func (h *roundTrigger) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func TestLogIDRoundTripper(t *testing.T) {
	logIDMap := &LogIDMap{}

	for i := 0; i < 1000; i++ {
		rid := fmt.Sprintf("%d", i)
		go func() {
			logIDMap.Set(rid)
			req, _ := http.NewRequest(http.MethodGet, "/", nil)

			LogIDRoundTripper(logIDMap)(nil).RoundTrip(req)

			if req.Header.Get("X-Request-ID") != rid {
				t.Fatal()
			}
		}()
	}
}
