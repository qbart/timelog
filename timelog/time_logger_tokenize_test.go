package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Timelog_Tokenize(t *testing.T) {
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

	tl := TimeLogger{events: events}

	result := tl.Tokenize()

	expectedResult := []Token{
		//
		Token{tkDate, "2020-01-15", -1},
		Token{tkSpace, " ", -1},
		Token{tkFromTime, "22:00", 0},
		Token{tkSpace, " ", -1},
		Token{tkToTime, "22:05", 1},
		Token{tkSpace, " ", -1},
		Token{tkComment, "hello", -1},
		Token{tkNewLine, "\n", -1},
		//
		Token{tkDate, "2020-01-15", -1},
		Token{tkSpace, " ", -1},
		Token{tkFromTime, "22:05", 1},
		Token{tkSpace, " ", -1},
		Token{tkToTime, "22:10", 2},
		Token{tkSpace, " ", -1},
		Token{tkComment, "world", -1},
		//
		Token{tkEnd, "", -1},
	}

	assert.Equal(t, expectedResult, result)
}
