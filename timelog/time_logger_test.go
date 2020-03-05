package timelog

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Timelog_Start(t *testing.T) {
	tl := TimeLogger{
		events: make([]event, 0),
		factory: &timelogMockFactory{
			now: _time("2020-01-15 22:00"),
			uuids: []uuid.UUID{
				_uuid("111aa398-5f30-11ea-b48d-4cedfb79ac39"),
			},
		},
	}

	tl.Start("hello")

	if assert.Equal(t, len(tl.events), 1) {
		assert.Equal(t, tl.events[0].uuid, _uuid("111aa398-5f30-11ea-b48d-4cedfb79ac39"))
		assert.Equal(t, tl.events[0].name, "start")
		assert.Equal(t, tl.events[0].comment, "hello")
		assert.Equal(t, tl.events[0].at, _time("2020-01-15 22:00"))
	}
}

func Test_Timelog_Start_WithExistingEvent(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			at:      _time("2020-01-15 22:00"),
			comment: "hello",
		},
	}
	tl := TimeLogger{
		events:  events,
		factory: &timelogMockFactory{now: _time("2020-01-15 22:02")},
	}

	tl.Start("world")

	if assert.Equal(t, len(tl.events), 2) {
		assert.Equal(t, tl.events[0].name, "start")
		assert.Equal(t, tl.events[0].comment, "hello")
		assert.Equal(t, tl.events[0].at, _time("2020-01-15 22:00"))

		assert.Equal(t, tl.events[1].name, "start")
		assert.Equal(t, tl.events[1].comment, "world")
		assert.Equal(t, tl.events[1].at, _time("2020-01-15 22:02"))
	}
}

func Test_Timelog_Stop_WhenNoEventsExist(t *testing.T) {
	events := []event{}
	tl := TimeLogger{
		events:  events,
		factory: &timelogMockFactory{now: _time("2020-01-15 22:02")},
	}

	tl.Stop()

	assert.Equal(t, len(tl.events), 0)
}

func Test_Timelog_Stop_WhenPreviousIsStart(t *testing.T) {
	events := []event{
		event{
			uuid:    _uuid("111aa398-5f30-11ea-b48d-4cedfb79ac39"),
			name:    "start",
			at:      _time("2020-01-15 22:00"),
			comment: "hello",
		},
	}
	tl := TimeLogger{
		events: events,
		factory: &timelogMockFactory{
			now: _time("2020-01-15 22:02"),
			uuids: []uuid.UUID{
				_uuid("222aa398-5f30-11ea-b48d-4cedfb79ac39"),
			},
		},
	}

	tl.Stop()

	if assert.Equal(t, len(tl.events), 2) {
		assert.Equal(t, tl.events[0].uuid, _uuid("111aa398-5f30-11ea-b48d-4cedfb79ac39"))
		assert.Equal(t, tl.events[0].name, "start")
		assert.Equal(t, tl.events[0].comment, "hello")
		assert.Equal(t, tl.events[0].at, _time("2020-01-15 22:00"))

		assert.Equal(t, tl.events[1].uuid, _uuid("222aa398-5f30-11ea-b48d-4cedfb79ac39"))
		assert.Equal(t, tl.events[1].name, "stop")
		assert.Equal(t, tl.events[1].comment, "")
		assert.Equal(t, tl.events[1].at, _time("2020-01-15 22:02"))
	}
}

func Test_Timelog_Stop_WhenPreviousIsStop(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			at:      _time("2020-01-15 22:00"),
			comment: "hello",
		},
		event{
			name:    "stop",
			at:      _time("2020-01-15 22:02"),
			comment: "",
		},
	}
	tl := TimeLogger{
		events:  events,
		factory: &timelogMockFactory{now: _time("2020-01-15 22:02")},
	}

	tl.Stop()

	if assert.Equal(t, len(tl.events), 2) {
		assert.Equal(t, tl.events[0].name, "start")
		assert.Equal(t, tl.events[0].comment, "hello")
		assert.Equal(t, tl.events[0].at, _time("2020-01-15 22:00"))

		assert.Equal(t, tl.events[1].name, "stop")
		assert.Equal(t, tl.events[1].comment, "")
		assert.Equal(t, tl.events[1].at, _time("2020-01-15 22:02"))
	}
}

func Test_Timelog_Clear(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			at:      _time("2020-01-15 22:00"),
			comment: "hello",
		},
		event{
			name:    "stop",
			at:      _time("2020-01-15 22:02"),
			comment: "",
		},
	}

	tl := TimeLogger{events: events}

	assert.Len(t, tl.events, 2)
	tl.Clear()
	assert.Len(t, tl.events, 0)
}
