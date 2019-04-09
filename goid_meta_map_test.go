package goid_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/go-courier/goid"
	"github.com/go-courier/metax"
)

func ExampleLogIDMap() {
	for i := 0; i < 20; i++ {
		go func() {
			// set logid at begin of goroutine
			goid.Default.Set(metax.Meta{
				"id": {fmt.Sprintf("%d", rand.Int())},
			})
			// clear at end of goroutine
			defer goid.Default.Clear()

			// do something with the cached logid
			_ = goid.Default.Get()

			time.Sleep(10 * time.Millisecond)
		}()
	}

	time.Sleep(5 * time.Millisecond)
	fmt.Println(len(goid.Default.All()))

	time.Sleep(50 * time.Millisecond)
	fmt.Println(len(goid.Default.All()))

	// missing
	fmt.Println(goid.Default.Get())
	// Output:
	//20
	//0
	//
}

func TestLogIDMapGo(t *testing.T) {
	for i := 0; i < 3; i++ {
		do := goid.Default.With(func() {
			meta := goid.Default.Get()
			fmt.Println("in goroutinue", meta)

			for i := 0; i < 3; i++ {
				doInGoroutine := goid.Default.With(func() {
					id := goid.Default.Get()
					fmt.Println("goroutinue in goroutinue", id)
				})
				go doInGoroutine()
			}
		}, metax.Meta{
			"id": {fmt.Sprintf("%d", rand.Int())},
		})

		go do()
	}

	time.Sleep(500 * time.Millisecond)
}

func BenchmarkLogIDMap(b *testing.B) {
	b.Run("log map", func(b *testing.B) {
		goid.Default.Set(metax.Meta{
			"id": {"1"},
		})

		for i := 0; i < b.N; i++ {
			goid.Default.Get()
		}
	})

	b.Run("meta context", func(b *testing.B) {
		ctx := metax.ContextWithMeta(context.Background(), metax.Meta{
			"id": {"1"},
		})

		for i := 0; i < b.N; i++ {
			metax.MetaFromContext(ctx)
		}
	})
}
