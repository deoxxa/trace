package trace

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testEvent1 struct {
	A string `json:"a"`
}

func (e testEvent1) String() string { return "1 " + e.A }

func (e testEvent1) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type": "test1",
		"a":    e.A,
	})
}

type testEvent2 struct {
	B string `json:"b"`
}

func (e testEvent2) String() string { return "2 " + e.B }

func (e testEvent2) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type": "test2",
		"b":    e.B,
	})
}

func init() {
	RegisterType("test1", func(d []byte) (Event, error) {
		var ev testEvent1
		if err := json.Unmarshal(d, &ev); err != nil {
			return nil, err
		}
		return &ev, nil
	})

	RegisterType("test2", func(d []byte) (Event, error) {
		var ev testEvent2
		if err := json.Unmarshal(d, &ev); err != nil {
			return nil, err
		}
		return &ev, nil
	})
}

func TestLogAddCallback(t *testing.T) {
	var l Log
	var called bool
	l.OnEvent = func(e Event) { called = true }
	l.Add(NewGenericEvent("x", nil))
	assert.True(t, called)
}

func TestLogMarshalGeneric(t *testing.T) {
	var l Log
	l.Add(NewGenericEvent("x", nil))
	l.Add(NewGenericEvent("y", nil))
	d, err := json.Marshal(l)
	assert.NoError(t, err)
	assert.JSONEq(t, `[{"type":"x"},{"type":"y"}]`, string(d))
}

func TestLogUnmarshalGeneric(t *testing.T) {
	var l Log
	assert.NoError(t, json.Unmarshal([]byte(`[{"type":"x"},{"type":"y"}]`), &l))
	assert.Len(t, l.Events, 2)
	assert.Equal(t, Log{Events: []Event{
		GenericEvent{Type: "x", Meta: map[string]interface{}{}},
		GenericEvent{Type: "y", Meta: map[string]interface{}{}},
	}}, l)
}

func TestLogMarshalCustom(t *testing.T) {
	var l Log
	l.Add(&testEvent1{A: "a"})
	l.Add(&testEvent2{B: "b"})
	d, err := json.Marshal(l)
	assert.NoError(t, err)
	assert.JSONEq(t, `[{"type":"test1","a":"a"},{"type":"test2","b":"b"}]`, string(d))
}

func TestLogUnmarshalCustom(t *testing.T) {
	var l Log
	assert.NoError(t, json.Unmarshal([]byte(`[{"type":"test1","a":"a"},{"type":"test2","b":"b"}]`), &l))
	assert.Len(t, l.Events, 2)
	assert.Equal(t, Log{Events: []Event{
		&testEvent1{A: "a"},
		&testEvent2{B: "b"},
	}}, l)
}
