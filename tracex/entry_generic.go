package tracex

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/satori/go.uuid"

	"fknsrs.biz/p/trace"
)

func init() {
	trace.RegisterType("callback", func(d []byte) (trace.Event, error) {
		var ev EntryCallback
		if err := json.Unmarshal(d, &ev); err != nil {
			return nil, err
		}
		return &ev, nil
	})
}

type Meta map[string]interface{}

type EntryGeneric struct {
	Event    string        `json:"event"`
	ID       uuid.UUID     `json:"id"`
	Time     time.Time     `json:"time"`
	Duration time.Duration `json:"duration"`
	Meta     Meta          `json:"meta"`
}

func (e EntryGeneric) String() string {
	return fmt.Sprintf("%s [%s] (%s) %s %v", e.ID, e.Event, e.Time, e.Duration, e.Meta)
}

func (e EntryGeneric) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})

	for k, v := range e.Meta {
		m[k] = v
	}

	m["event"] = e.Event
	m["id"] = e.ID
	m["time"] = e.Time
	m["duration"] = e.Duration

	return json.Marshal(m)
}
