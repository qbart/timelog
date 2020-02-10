package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
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
