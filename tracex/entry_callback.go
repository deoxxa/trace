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

type Change struct {
	Name     string      `json:"name"`
	OldValue interface{} `json:"oldValue"`
	NewValue interface{} `json:"newValue"`
}

type EntryCallback struct {
	ID        uuid.UUID     `json:"id"`
	Time      time.Time     `json:"time"`
	Duration  time.Duration `json:"duration"`
	Name      string        `json:"name"`
	Triggered []Change      `json:"triggered"`
	Changed   []Change      `json:"changed"`
	Skipped   bool          `json:"skipped"`
	Forced    bool          `json:"forced"`
}

func (e EntryCallback) String() string {
	return fmt.Sprintf("Callback [%s] %s %s (skipped=%v) (forced=%v) %v -> %v", e.Time, e.Duration, e.Name, e.Skipped, e.Forced, e.Triggered, e.Changed)
}

func (e EntryCallback) MarshalJSON() ([]byte, error) {
	triggered := e.Triggered
	if triggered == nil {
		triggered = make([]Change, 0)
	}

	changed := e.Changed
	if changed == nil {
		changed = make([]Change, 0)
	}

	return json.Marshal(map[string]interface{}{
		"event":     "callback",
		"id":        e.ID,
		"time":      e.Time,
		"duration":  e.Duration,
		"name":      e.Name,
		"skipped":   e.Skipped,
		"forced":    e.Forced,
		"triggered": triggered,
		"changed":   changed,
	})
}
