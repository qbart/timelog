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

func TestAdjust_Adjust_ValidClone(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 22:00"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      makeTime("2020-01-15 22:05"),
		},
	}

	tl := TimeLogger{events: events}
	clone := tl.Adjust(map[int]int{})
	assert.Len(t, clone.events, 2)
	assert.Equal(t, tl.config, clone.config)
	assert.Equal(t, tl.factory, clone.factory)
}

func TestAdjust_DurationDoesNotCrossOver(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 22:00"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      makeTime("2020-01-15 22:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      makeTime("2020-01-15 22:10"),
		},
	}

	tl := TimeLogger{events: events}
	clone := tl.Adjust(map[int]int{
		0: -3, // adjust time in "0" point -3 minutes
		1: -6, // adjust time in "1" point -6 minutes
		2: 5,  // adjust time in "2" point +5 minutes
	})

	// original events are not modified
	assert.Equal(t, tl.events[0].at, makeTime("2020-01-15 22:00"))
	assert.Equal(t, tl.events[1].at, makeTime("2020-01-15 22:05"))
	assert.Equal(t, tl.events[2].at, makeTime("2020-01-15 22:10"))

	// cloned objects contains modified events
	assert.Equal(t, clone.events[0].at, makeTime("2020-01-15 21:57"))
	assert.Equal(t, clone.events[1].at, makeTime("2020-01-15 21:59"))
	assert.Equal(t, clone.events[2].at, makeTime("2020-01-15 22:15"))
}

func TestAdjust_DurationNegativeCrossOver(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 22:00"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      makeTime("2020-01-15 22:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      makeTime("2020-01-15 22:10"),
		},
	}

	tl := TimeLogger{events: events}
	clone := tl.Adjust(map[int]int{
		1: -6, // adjust time in "1" point -6 minutes
	})

	// cloned objects contains modified events
	assert.Equal(t, clone.events[0].at, makeTime("2020-01-15 22:00"))
	assert.Equal(t, clone.events[1].at, makeTime("2020-01-15 22:00")) // -6m -> 22:00 (can't go lower than previous)
	assert.Equal(t, clone.events[2].at, makeTime("2020-01-15 22:10"))
}

func TestAdjust_DurationPositiveCrossOver(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 22:00"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      makeTime("2020-01-15 22:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      makeTime("2020-01-15 22:10"),
		},
	}

	tl := TimeLogger{events: events}
	clone := tl.Adjust(map[int]int{
		1: 6, // adjust time in "1" point +6 minutes
	})

	// cloned objects contains modified events
	assert.Equal(t, clone.events[0].at, makeTime("2020-01-15 22:00"))
	assert.Equal(t, clone.events[1].at, makeTime("2020-01-15 22:10")) // +6m -> 22:10 (can't go higher than next)
	assert.Equal(t, clone.events[2].at, makeTime("2020-01-15 22:10"))
}

func Test_Timelog_Tokenize(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 22:00"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      makeTime("2020-01-15 22:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      makeTime("2020-01-15 22:10"),
		},
	}

	tl := TimeLogger{events: events}

	result := tl.Tokenize()

	expectedResult := []Token{
		//
		Token{tkDate, "2020-01-15", -1},
		Token{tkSpace, " ", -1},
		Token{tkFromTime, "22:00", 0},
		Token{tkSpace, " ", -1},
		Token{tkToTime, "22:05", 1},
		Token{tkSpace, " ", -1},
		Token{tkComment, "hello", -1},
		Token{tkNewLine, "\n", -1},
		//
		Token{tkDate, "2020-01-15", -1},
		Token{tkSpace, " ", -1},
		Token{tkFromTime, "22:05", 1},
		Token{tkSpace, " ", -1},
		Token{tkToTime, "22:10", 2},
		Token{tkSpace, " ", -1},
		Token{tkComment, "world", -1},
		//
		Token{tkEnd, "", -1},
	}

	assert.Equal(t, expectedResult, result)
}
