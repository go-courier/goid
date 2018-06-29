package log_id_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-courier/goid/log_id"
)

func ExampleLogIDMap() {
	for i := 0; i < 100; i++ {
		go func() {
			// set logid at begin of goroutine
			log_id.Default.Set(fmt.Sprintf("%d", rand.Int()))
			// clear at end of goroutine
			defer log_id.Default.Clear()

			// do something with the cached logid
			_ = log_id.Default.Get()

			time.Sleep(10 * time.Millisecond)
		}()
	}

	time.Sleep(5 * time.Millisecond)
	fmt.Println(len(log_id.Default.All()))

	time.Sleep(50 * time.Millisecond)
	fmt.Println(len(log_id.Default.All()))

	// missing
	fmt.Println(log_id.Default.Get())
	// Output:
	//100
	//0
	//
}
