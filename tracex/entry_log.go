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

type EntryLog struct {
	ID       uuid.UUID     `json:"id"`
	Time     time.Time     `json:"time"`
	Duration time.Duration `json:"duration"`
	Level    string        `json:"level"`
	Data     []interface{} `json:"data"`
}

func (e EntryLog) String() string {
	return fmt.Sprintf("Log [%s] %s %s: %v", e.Time, e.Duration, e.Level, e.Data)
}

func (e EntryLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"event":    "log",
		"id":       e.ID,
		"time":     e.Time,
		"duration": e.Duration,
		"level":    e.Level,
		"data":     e.Data,
	})
}
