## Logid

[![GoDoc Widget](https://godoc.org/github.com/go-courier/goid/log_id?status.svg)](https://godoc.org/github.com/go-courier/goid/log_id)
[![Build Status](https://travis-ci.org/go-courier/goid.svg?branch=master)](https://travis-ci.org/go-courier/goid)
[![codecov](https://codecov.io/gh/go-courier/goid/branch/master/graph/badge.svg)](https://codecov.io/gh/go-courier/goid)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-courier/goid)](https://goreportcard.com/report/github.com/go-courier/goid)

Hacking runtime to get goroutine id for caching logid 

### Usage

```bash
# Patch
go get -u github.com/go-courier/goid/patch-runtime && patch-runtime
```

For go module user:

need run blow

```
cd $GOPATH/src
mkdir -p global-tools
cd global-tools
go mod init
```


```go
package log_id_test

import (
	"time"
	"fmt"
	"math/rand"

	"github.com/go-courier/goid/log_id"
)

func ExampleLogIDMap() {
	for i := 0; i < 100; i ++ {
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
	// Output:
	//100
	//0
}
```
