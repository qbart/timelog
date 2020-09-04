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
	assert.Equal(t, 1, analytics.Duration.Hours)
	assert.Equal(t, 15, analytics.Duration.Minutes)
	assert.Equal(t, 1, analytics.LastDuration.Hours)
	assert.Equal(t, 11, analytics.LastDuration.Minutes)
}

func TestCalcAnalyticsCommonPrefixes(t *testing.T) {
	tl := TimeLogger{
		events: []event{
			{
				name:    "start",
				comment: "aaa-task-1",
				at:      _time("2020-01-15 22:00"),
			},
			{
				name:    "start",
				comment: "aaa-task-2",
				at:      _time("2020-01-15 22:05"),
			},
			{
				name:    "start",
				comment: "bbb-task-1",
				at:      _time("2020-01-15 22:10"),
			},
			{
				name:    "stop",
				comment: "",
				at:      _time("2020-01-15 22:15"),
			},
		},
		factory: &timelogMockFactory{},
	}
	analytics := calcAnalytics(&tl)

	assert.Equal(t, []string{"aaa", "bbb"}, analytics.PrefixOrder)
	assert.Equal(t, 0, analytics.PrefixDuration["aaa"].Hours)
	assert.Equal(t, 10, analytics.PrefixDuration["aaa"].Minutes)
	assert.Equal(t, 0, analytics.PrefixDuration["bbb"].Hours)
	assert.Equal(t, 5, analytics.PrefixDuration["bbb"].Minutes)

	tl = TimeLogger{
		events: []event{
			{
				name:    "start",
				comment: "ccc task-1",
				at:      _time("2020-01-15 22:00"),
			},
			{
				name:    "start",
				comment: "ddd",
				at:      _time("2020-01-15 22:05"),
			},
			{
				name:    "start",
				comment: "ccc_dev-1",
				at:      _time("2020-01-15 23:10"),
			},
			{
				name:    "start",
				comment: "ddd-xxx",
				at:      _time("2020-01-15 23:20"),
			},
		},
		factory: &timelogMockFactory{
			now: _time("2020-01-15 23:30"),
		},
	}
	analytics = calcAnalytics(&tl)

	assert.Equal(t, []string{"ccc", "ddd", "ccc_dev"}, analytics.PrefixOrder)
	assert.Equal(t, 0, analytics.PrefixDuration["ccc"].Hours)
	assert.Equal(t, 5, analytics.PrefixDuration["ccc"].Minutes)
	assert.Equal(t, 1, analytics.PrefixDuration["ddd"].Hours)
	assert.Equal(t, 15, analytics.PrefixDuration["ddd"].Minutes)
	assert.Equal(t, 0, analytics.PrefixDuration["ccc_dev"].Hours)
	assert.Equal(t, 10, analytics.PrefixDuration["ccc_dev"].Minutes)
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
	assert.Equal(t, 2, analytics.Duration.Hours)
	assert.Equal(t, 0, analytics.Duration.Minutes)
	assert.Equal(t, 1, analytics.LastDuration.Hours)
	assert.Equal(t, 0, analytics.LastDuration.Minutes)
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
	assert.Equal(t, 1, analytics.Duration.Hours)
	assert.Equal(t, 5, analytics.Duration.Minutes)
	assert.Equal(t, 0, analytics.LastDuration.Hours)
	assert.Equal(t, 5, analytics.LastDuration.Minutes)
}
