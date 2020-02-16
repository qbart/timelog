package timelog

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeTime(value string) time.Time {
	parsedTime, _ := time.Parse("2006-01-02 15:04", value)
	return parsedTime
}

func trimHeredoc(s string) string {
	ss := strings.Split(s, "\n")
	lines := []string{}
	for _, line := range ss {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	return strings.Join(lines, "\n")

}

func Test_trimHeredoc(t *testing.T) {
	text := `
	some
		random
			text
	`
	lines := trimHeredoc(text)
	expected := "some\nrandom\ntext"
	assert.Equal(t, expected, lines)
}
