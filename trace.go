package trace

import (
	"encoding/json"
	"fmt"
)

type TypeFactoryFunc func(d []byte) (Event, error)

var typeRegistry map[string]TypeFactoryFunc

func init() {
	typeRegistry = make(map[string]TypeFactoryFunc)
}

func RegisterType(name string, fn TypeFactoryFunc) { typeRegistry[name] = fn }

type Event interface {
	fmt.Stringer
}

type Log struct {
	Events  []Event
	OnEvent func(e Event)
}

func (a *Log) Add(e Event) {
	a.Events = append(a.Events, e)

	if a.OnEvent != nil {
		a.OnEvent(e)
	}
}

func (a Log) MarshalJSON() ([]byte, error) { return json.Marshal(a.Events) }

func (a *Log) UnmarshalJSON(d []byte) error {
	var arr []json.RawMessage
	if err := json.Unmarshal(d, &arr); err != nil {
		return err
	}

	for _, e := range arr {
		var v struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(e, &v); err != nil {
			return err
		}

		if fn, ok := typeRegistry[v.Type]; ok {
			ev, err := fn([]byte(e))
			if err != nil {
				return err
			}
			a.Events = append(a.Events, ev)
		} else {
			var ev GenericEvent
			if err := json.Unmarshal([]byte(e), &ev); err != nil {
				return err
			}
			a.Events = append(a.Events, ev)
		}
	}

	return nil
}

type GenericEvent struct {
	Type string                 `json:"type"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

func NewGenericEvent(eventType string, meta map[string]interface{}) GenericEvent {
	return GenericEvent{Type: eventType, Meta: meta}
}

func (e GenericEvent) String() string {
	return fmt.Sprintf("%s %v", e.Type, e.Meta)
}

func (e GenericEvent) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})

	for k, v := range e.Meta {
		m[k] = v
	}

	m["type"] = e.Type

	return json.Marshal(m)
}

func (e *GenericEvent) UnmarshalJSON(d []byte) error {
	var partial struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(d, &partial); err != nil {
		return err
	}

	e.Type = partial.Type

	if err := json.Unmarshal(d, &e.Meta); err != nil {
		return err
	}

	delete(e.Meta, "type")

	return nil
}

type contextKey struct{}

var (
	ContextKey = &contextKey{}
)
