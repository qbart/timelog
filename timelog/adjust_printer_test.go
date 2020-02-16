package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdjustPrinter_String_DefaultSelection(t *testing.T) {
	entries := []entry{
		entry{
			comment: "hello",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:05"),
			},
		},
		entry{
			comment: "world",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:05"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:10"),
			},
		},
	}
	p := AdjustPrinter{
		timelogger: &TimeLogger{entries: entries},
	}

	result := p.String()

	expectedResult := trimHeredoc(`
		2020-01-15 [22:00] 22:05  hello
		2020-01-15  22:05  22:10  world
	`)

	assert.Equal(t, expectedResult, result)
}

func Test_AdjustPrinter_String_SelectionBetween(t *testing.T) {
	entries := []entry{
		entry{
			comment: "hello",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:05"),
			},
		},
		entry{
			comment: "world",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:05"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:10"),
			},
		},
	}
	p := AdjustPrinter{
		timelogger: &TimeLogger{entries: entries},
		selected:   1,
	}

	result := p.String()

	expectedResult := trimHeredoc(`
		2020-01-15  22:00 [22:05] hello
		2020-01-15 [22:05] 22:10  world
	`)

	assert.Equal(t, expectedResult, result)
}

func Test_AdjustPrinter_String_SelectionLast(t *testing.T) {
	entries := []entry{
		entry{
			comment: "hello",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:05"),
			},
		},
		entry{
			comment: "world",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:05"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:10"),
			},
		},
	}
	p := AdjustPrinter{
		timelogger: &TimeLogger{entries: entries},
		selected:   2,
	}

	result := p.String()

	expectedResult := trimHeredoc(`
		2020-01-15  22:00  22:05  hello
		2020-01-15  22:05 [22:10] world
	`)

	assert.Equal(t, expectedResult, result)
}

func Test_AdjustPrinter_String_Empty(t *testing.T) {
	p := AdjustPrinter{
		timelogger: &TimeLogger{entries: []entry{}},
	}

	result := p.String()

	expectedResult := ""

	assert.Equal(t, expectedResult, result)
}
