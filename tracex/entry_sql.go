package tracex

import (
	"encoding/json"
	"fmt"
	"strings"
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

type EntrySQL struct {
	ID       uuid.UUID     `json:"id"`
	Time     time.Time     `json:"time"`
	Duration time.Duration `json:"duration"`
	Query    string        `json:"query"`
	Stack    []string      `json:"stack"`
}

func (e EntrySQL) String() string {
	return fmt.Sprintf("SQL [%s] %s %s (%s)", e.Time, e.Duration, e.Query, strings.Join(e.Stack, " -> "))
}

func (e EntrySQL) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"event":    "sql",
		"id":       e.ID,
		"time":     e.Time,
		"duration": e.Duration,
		"query":    e.Query,
		"stack":    e.Stack,
	})
}
