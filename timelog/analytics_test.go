package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcAnalytics(t *testing.T) {
	ee := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 22:01"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      makeTime("2020-01-15 22:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      makeTime("2020-01-15 23:16"),
		},
	}
	analytics := calcAnalytics(ee)

	assert.Equal(t, 2, analytics.EntryNum)
	assert.Equal(t, 1, analytics.Hours)
	assert.Equal(t, 15, analytics.Minutes)
}
