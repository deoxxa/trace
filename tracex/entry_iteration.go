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

type EntryIteration struct {
	ID         uuid.UUID `json:"id"`
	Time       time.Time `json:"time"`
	ObjectType string    `json:"objectType"`
	ObjectID   uuid.UUID `json:"objectId"`
	Number     int       `json:"number"`
}

func (e EntryIteration) String() string {
	return fmt.Sprintf("Iteration [%s] %d on %s (%s)", e.Time, e.Number, e.ObjectType, e.ObjectID.String())
}

func (e EntryIteration) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"event":      "iteration",
		"id":         e.ID,
		"time":       e.Time,
		"objectType": e.ObjectType,
		"objectId":   e.ObjectID,
		"number":     e.Number,
	})
}
