package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
