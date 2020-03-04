package timelog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type logtimeMockFactory struct {
	now time.Time
}

type timeMockFactory struct {
	now time.Time
}

func (mock timeMockFactory) New() time.Time {
	return mock.now
}

func Test_Timelog_Start(t *testing.T) {
	tl := TimeLogger{
		events:  make([]event, 0),
		factory: timeMockFactory{now: makeTime("2020-01-15 22:00")},
	}

	tl.Start("hello")

	if assert.Equal(t, len(tl.events), 1) {
		assert.Equal(t, tl.events[0].name, "start")
		assert.Equal(t, tl.events[0].comment, "hello")
		assert.Equal(t, tl.events[0].at, makeTime("2020-01-15 22:00"))
	}
}

func Test_Timelog_Start_WithExistingEvent(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			at:      makeTime("2020-01-15 22:00"),
			comment: "hello",
		},
	}
	tl := TimeLogger{
		events:  events,
		factory: timeMockFactory{now: makeTime("2020-01-15 22:02")},
	}

	tl.Start("world")

	if assert.Equal(t, len(tl.events), 2) {
		assert.Equal(t, tl.events[0].name, "start")
		assert.Equal(t, tl.events[0].comment, "hello")
		assert.Equal(t, tl.events[0].at, makeTime("2020-01-15 22:00"))

		assert.Equal(t, tl.events[1].name, "start")
		assert.Equal(t, tl.events[1].comment, "world")
		assert.Equal(t, tl.events[1].at, makeTime("2020-01-15 22:02"))
	}
}

func Test_Timelog_Stop_WhenNoEventsExist(t *testing.T) {
	events := []event{}
	tl := TimeLogger{
		events:  events,
		factory: timeMockFactory{now: makeTime("2020-01-15 22:02")},
	}

	tl.Stop()

	assert.Equal(t, len(tl.events), 0)
}

func Test_Timelog_Stop_WhenPreviousIsStart(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			at:      makeTime("2020-01-15 22:00"),
			comment: "hello",
		},
	}
	tl := TimeLogger{
		events:  events,
		factory: timeMockFactory{now: makeTime("2020-01-15 22:02")},
	}

	tl.Stop()

	if assert.Equal(t, len(tl.events), 2) {
		assert.Equal(t, tl.events[0].name, "start")
		assert.Equal(t, tl.events[0].comment, "hello")
		assert.Equal(t, tl.events[0].at, makeTime("2020-01-15 22:00"))

		assert.Equal(t, tl.events[1].name, "stop")
		assert.Equal(t, tl.events[1].comment, "")
		assert.Equal(t, tl.events[1].at, makeTime("2020-01-15 22:02"))
	}
}

func Test_Timelog_Stop_WhenPreviousIsStop(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			at:      makeTime("2020-01-15 22:00"),
			comment: "hello",
		},
		event{
			name:    "stop",
			at:      makeTime("2020-01-15 22:02"),
			comment: "",
		},
	}
	tl := TimeLogger{
		events:  events,
		factory: timeMockFactory{now: makeTime("2020-01-15 22:02")},
	}

	tl.Stop()

	if assert.Equal(t, len(tl.events), 2) {
		assert.Equal(t, tl.events[0].name, "start")
		assert.Equal(t, tl.events[0].comment, "hello")
		assert.Equal(t, tl.events[0].at, makeTime("2020-01-15 22:00"))

		assert.Equal(t, tl.events[1].name, "stop")
		assert.Equal(t, tl.events[1].comment, "")
		assert.Equal(t, tl.events[1].at, makeTime("2020-01-15 22:02"))
	}
}

func Test_Timelog_Clear(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			at:      makeTime("2020-01-15 22:00"),
			comment: "hello",
		},
		event{
			name:    "stop",
			at:      makeTime("2020-01-15 22:02"),
			comment: "",
		},
	}

	tl := TimeLogger{events: events}

	assert.Len(t, tl.events, 2)
	tl.Clear()
	assert.Len(t, tl.events, 0)
}
