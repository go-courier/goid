package log_id_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/go-courier/goid/log_id"
)

func ExampleLogIDMap() {
	for i := 0; i < 20; i++ {
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
	//20
	//0
	//
}

func TestLogIDMapGo(t *testing.T) {
	for i := 0; i < 3; i++ {
		do := log_id.Default.With(func() {
			id := log_id.Default.Get()
			fmt.Println("in goroutinue", id)

			for i := 0; i < 3; i++ {
				doInGoroutine := log_id.Default.With(func() {
					id := log_id.Default.Get()
					fmt.Println("goroutinue in goroutinue", id)
				})
				go doInGoroutine()
			}
		}, fmt.Sprintf("%d", rand.Int()))

		go do()
	}

	time.Sleep(500 * time.Millisecond)
}
