package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolybarPrinter_String(t *testing.T) {
	events := []event{
		{
			name:    "start",
			comment: "hello",
			at:      _time("2020-01-15 22:00"),
		},
		{
			name:    "start",
			comment: "world",
			at:      _time("2020-01-15 22:05"),
		},
		{
			name:    "stop",
			comment: "",
			at:      _time("2020-01-15 22:10"),
		},
	}

	p := PolybarPrinter{
		timelogger: &TimeLogger{
			events:  events,
			factory: timelogDefaultFactory{},
		},
		format: "{{.Count}} {{.Comment}} {{.Duration}} {{.Total}} {{.CountNotZero}} {{.TotalGtDuration}}",
	}

	result := p.String()

	expectedResult := "2 world 0h5m 0h10m true true"

	assert.Equal(t, expectedResult, result)
}

func Test_PolybarPrinter_String_Empty(t *testing.T) {
	events := []event{}

	p := PolybarPrinter{
		timelogger: &TimeLogger{
			events:  events,
			factory: timelogDefaultFactory{},
		},
		format: "{{.Count}} {{.Comment}} {{.Duration}} {{.Total}} {{.CountNotZero}} {{.TotalGtDuration}}",
	}

	result := p.String()

	expectedResult := "0    false false"

	assert.Equal(t, expectedResult, result)
}
