package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcAnalytics(t *testing.T) {
	tl := TimeLogger{
		events: []event{
			{
				name:    "start",
				comment: "hello",
				at:      _time("2020-01-15 22:01"),
			},
			{
				name:    "start",
				comment: "world",
				at:      _time("2020-01-15 22:05"),
			},
			{
				name:    "stop",
				comment: "",
				at:      _time("2020-01-15 23:16"),
			},
		},
		factory: &timelogMockFactory{},
	}
	analytics := calcAnalytics(&tl)

	assert.Equal(t, 2, analytics.EntryNum)
	assert.Equal(t, 1, analytics.Hours)
	assert.Equal(t, 15, analytics.Minutes)
	assert.Equal(t, 1, analytics.LastHours)
	assert.Equal(t, 11, analytics.LastMinutes)
}

func TestCalcAnalytics_MultipleStops(t *testing.T) {
	tl := TimeLogger{
		events: []event{
			{
				name:    "start",
				comment: "hello",
				at:      _time("2020-01-15 22:00"),
			},
			{
				name:    "stop",
				comment: "",
				at:      _time("2020-01-15 23:00"),
			},
			{
				name:    "start",
				comment: "hello 2",
				at:      _time("2020-01-16 22:00"),
			},
			{
				name:    "stop",
				comment: "",
				at:      _time("2020-01-16 23:00"),
			},
		},
		factory: &timelogMockFactory{},
	}
	analytics := calcAnalytics(&tl)

	assert.Equal(t, 2, analytics.EntryNum)
	assert.Equal(t, 2, analytics.Hours)
	assert.Equal(t, 0, analytics.Minutes)
	assert.Equal(t, 1, analytics.LastHours)
	assert.Equal(t, 0, analytics.LastMinutes)
}

func TestCalcAnalytics_OneStopInTheMiddle(t *testing.T) {
	tl := TimeLogger{
		events: []event{
			{
				name:    "start",
				comment: "hello",
				at:      _time("2020-01-15 22:00"),
			},
			{
				name:    "stop",
				comment: "",
				at:      _time("2020-01-15 23:00"),
			},
			{
				name:    "start",
				comment: "hello 2",
				at:      _time("2020-01-16 22:00"),
			},
		},
		factory: &timelogMockFactory{
			now: _time("2020-01-16 22:05"),
		},
	}
	analytics := calcAnalytics(&tl)

	assert.Equal(t, 2, analytics.EntryNum)
	assert.Equal(t, 1, analytics.Hours)
	assert.Equal(t, 5, analytics.Minutes)
	assert.Equal(t, 0, analytics.LastHours)
	assert.Equal(t, 5, analytics.LastMinutes)
}
