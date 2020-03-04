package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TextPrinter_String_OneEntry(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 22:00"),
		},
	}

	p := TextPrinter{
		timelogger: &TimeLogger{events: events},
	}

	result := p.String()

	expectedResult := "2020-01-15 22:00 ...   hello"

	assert.Equal(t, expectedResult, result)
}

func Test_TextPrinter_String_WithStartOnly(t *testing.T) {
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
	}

	p := TextPrinter{
		timelogger: &TimeLogger{events: events},
	}

	result := p.String()

	expectedResult := "2020-01-15 22:00 22:05 hello\n2020-01-15 22:05 ...   world"

	assert.Equal(t, expectedResult, result)
}

func Test_TextPrinter_String_WithStop(t *testing.T) {
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

	p := TextPrinter{
		timelogger: &TimeLogger{events: events},
	}

	result := p.String()

	expectedResult := "2020-01-15 22:00 22:05 hello\n2020-01-15 22:05 22:10 world"

	assert.Equal(t, expectedResult, result)
}

func Test_TextPrinter_String_SplitDaysWithStop(t *testing.T) {
	events := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 22:00"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      makeTime("2020-01-16 01:05"),
		},
	}

	p := TextPrinter{
		timelogger: &TimeLogger{events: events},
	}

	result := p.String()

	expectedResult := "2020-01-15 22:00 23:59 hello\n2020-01-16 00:00 01:05 hello"

	assert.Equal(t, expectedResult, result)
}
