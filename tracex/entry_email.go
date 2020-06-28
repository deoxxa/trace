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

type EntryEmail struct {
	ID       uuid.UUID `json:"id"`
	Time     time.Time `json:"time"`
	To       []string  `json:"to"`
	CC       []string  `json:"cc"`
	BCC      []string  `json:"bcc"`
	Subject  string    `json:"subject"`
	BodyHTML string    `json:"bodyHTML"`
	BodyText string    `json:"bodyText"`
}

func (e EntryEmail) String() string {
	return fmt.Sprintf("Email [%s] To=%v CC=%v BCC=%v Subject=%q BodyHTML=%q BodyText=%q", e.Time, e.To, e.CC, e.BCC, e.Subject, e.BodyHTML, e.BodyText)
}

func (e EntryEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"event":    "email",
		"id":       e.ID,
		"time":     e.Time,
		"to":       e.To,
		"cc":       e.CC,
		"bcc":      e.BCC,
		"subject":  e.Subject,
		"bodyHTML": e.BodyHTML,
		"bodyText": e.BodyText,
	})
}
