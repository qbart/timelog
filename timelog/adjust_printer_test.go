package timelog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdjustPrinter_String_DefaultSelection(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      _time("2020-01-15 22:00"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      _time("2020-01-15 22:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      _time("2020-01-15 22:10"),
		},
	}
	p := AdjustPrinter{
		timelogger: &TimeLogger{events: events},
	}

	result := p.String()

	expectedResult := fmt.Sprint(
		"2020-01-15 [[22:00]](fg:yellow,bg:black) 22:05  hello\n",
		"2020-01-15  22:05  22:10  world",
	)

	assert.Equal(t, expectedResult, result)
}

func Test_AdjustPrinter_String_SelectionBetween(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      _time("2020-01-15 22:00"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      _time("2020-01-15 22:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      _time("2020-01-15 22:10"),
		},
	}
	p := AdjustPrinter{
		timelogger: &TimeLogger{events: events},
		selected:   1,
	}

	result := p.String()

	expectedResult := fmt.Sprint(
		"2020-01-15  22:00 [[22:05]](fg:yellow,bg:black) hello\n",
		"2020-01-15 [[22:05]](fg:yellow,bg:black) 22:10  world",
	)

	assert.Equal(t, expectedResult, result)
}

func Test_AdjustPrinter_String_SelectionLast(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      _time("2020-01-15 22:00"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      _time("2020-01-15 22:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      _time("2020-01-15 22:10"),
		},
	}
	p := AdjustPrinter{
		timelogger: &TimeLogger{events: events},
		selected:   2,
	}

	result := p.String()

	expectedResult := fmt.Sprint(
		"2020-01-15  22:00  22:05  hello\n",
		"2020-01-15  22:05 [[22:10]](fg:yellow,bg:black) world",
	)

	assert.Equal(t, expectedResult, result)
}

func Test_AdjustPrinter_String_SelectionWithMultipleStops(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      _time("2020-01-15 22:00"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      _time("2020-01-15 22:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      _time("2020-01-15 22:10"),
		},
		event{
			name:    "start",
			comment: "xd",
			at:      _time("2020-01-15 23:05"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      _time("2020-01-15 23:10"),
		},
	}
	p := AdjustPrinter{
		timelogger: &TimeLogger{events: events},
		selected:   3,
	}

	result := p.String()

	expectedResult := fmt.Sprint(
		"2020-01-15  22:00  22:05  hello\n",
		"2020-01-15  22:05  22:10  world\n",
		"2020-01-15 [[23:05]](fg:yellow,bg:black) 23:10  xd",
	)

	assert.Equal(t, expectedResult, result)
}

func Test_AdjustPrinter_String_Empty(t *testing.T) {
	p := AdjustPrinter{
		timelogger: &TimeLogger{events: []event{}},
	}

	result := p.String()

	expectedResult := ""

	assert.Equal(t, expectedResult, result)
}
