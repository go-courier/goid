package log_id

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type hook struct {
}

func (hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook) Fire(entry *logrus.Entry) error {
	entry.Time = time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)
	s, _ := entry.String()
	fmt.Print(s)
	return nil
}

func ExampleLogIDHook() {
	buf := bytes.NewBuffer(nil)

	l := logrus.New()
	l.Out = buf

	m := &LogIDMap{}

	l.AddHook(NewLogIDHook(m))
	l.AddHook(&hook{})

	m.Set("id")
	defer m.Clear()

	l.Println("1")
	// Output:
	// time="0001-01-01T01:01:01Z" level=info msg=1 log_id=id
}
