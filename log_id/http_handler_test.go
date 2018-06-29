package log_id

import (
	"fmt"
	"net/http"
	"time"
)

type handler struct {
}

func (h *handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	time.Sleep(10 * time.Millisecond)
}

func ExampleLogIDHttpHandler() {
	h := &handler{}

	// could nil
	_ = LogIDHttpHandler(nil)

	logIDMap := &LogIDMap{}

	for i := 0; i < 100; i++ {
		go func() {
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			LogIDHttpHandler(logIDMap)(h).ServeHTTP(http.ResponseWriter(nil), req)
		}()
	}

	time.Sleep(5 * time.Millisecond)
	fmt.Println(len(logIDMap.All()))

	time.Sleep(50 * time.Millisecond)
	fmt.Println(len(logIDMap.All()))
	// Output:
	//100
	//0
}
