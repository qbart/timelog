package timelog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DiffPrinter_String_WithoutStop(t *testing.T) {
	eventsOriginal := []event{
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
	eventsModified := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 21:59"),
		},
		event{
			name:    "start",
			comment: "world",
			at:      makeTime("2020-01-15 22:10"),
		},
	}

	p := DiffPrinter{
		timeloggerOriginal: &TimeLogger{events: eventsOriginal},
		timeloggerModified: &TimeLogger{events: eventsModified},
	}

	result := p.String()

	expectedResult := trimHeredoc(`
	-2020-01-15 [22:00] [22:05] hello
	+2020-01-15 [21:59] [22:10] hello
	-2020-01-15 [22:05] ...   world
	+2020-01-15 [22:10] ...   world
	`)

	assert.Equal(t, expectedResult, result)
}

func Test_DiffPrinter_String_WithStop(t *testing.T) {
	eventsOriginal := []event{
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
	eventsModified := []event{
		event{
			name:    "start",
			comment: "hello",
			at:      makeTime("2020-01-15 21:59"),
		},
		event{
			name:    "stop",
			comment: "",
			at:      makeTime("2020-01-15 22:10"),
		},
	}

	p := DiffPrinter{
		timeloggerOriginal: &TimeLogger{events: eventsOriginal},
		timeloggerModified: &TimeLogger{events: eventsModified},
	}

	result := p.String()

	expectedResult := trimHeredoc(`
	-2020-01-15 [22:00] [22:05] hello
	+2020-01-15 [21:59] [22:10] hello
	`)

	assert.Equal(t, expectedResult, result)
}

func Test_DiffPrinter_String_Empty(t *testing.T) {
	p := DiffPrinter{
		timeloggerOriginal: &TimeLogger{events: []event{}},
		timeloggerModified: &TimeLogger{events: []event{}},
	}

	result := p.String()

	expectedResult := ""

	assert.Equal(t, expectedResult, result)
}

func Test_DiffPrinter_String_NoChanges(t *testing.T) {
	eventsOriginal := []event{
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
	eventsModified := []event{
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

	p := DiffPrinter{
		timeloggerOriginal: &TimeLogger{events: eventsOriginal},
		timeloggerModified: &TimeLogger{events: eventsModified},
	}

	result := p.String()

	expectedResult := fmt.Sprint(
		" 2020-01-15 22:00 22:05 hello\n",
		" 2020-01-15 22:05 ...   world",
	)

	assert.Equal(t, expectedResult, result)
}
