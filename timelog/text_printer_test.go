package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TextPrinter_String(t *testing.T) {
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
				finished: false,
				t:        makeTime("2020-01-15 22:05"),
			},
		},
	}

	p := TextPrinter{
		timelogger: &TimeLogger{entries: entries},
	}

	result := p.String()

	expectedResult := "2020-01-15 22:00 22:05 hello\n2020-01-15 22:05 ...   world"

	assert.Equal(t, result, expectedResult)
}

func Test_TextPrinter_String_split_days(t *testing.T) {
	entries := []entry{
		entry{
			comment: "hello",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-16 01:05"),
			},
		},
	}

	p := TextPrinter{
		timelogger: &TimeLogger{entries: entries},
	}

	result := p.String()

        expectedResult := "2020-01-15 22:00 23:59 hello\n2020-01-16 00:00 01:05 hello"

	assert.Equal(t, result, expectedResult)
}
