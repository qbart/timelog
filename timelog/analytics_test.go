package timelog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalcAnalytics(t *testing.T) {
	ee := []entry{
		entry{
			comment: "hello",
			from: logtime{
				t: makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				t: makeTime("2020-01-15 22:05"),
			},
		},
		entry{
			comment: "world",
			from: logtime{
				t: makeTime("2020-01-15 22:05"),
			},
			to: logtime{
				t: makeTime("2020-01-15 23:01"),
			},
		},
	}
	analytics := calcAnalytics(ee)

	assert.Equal(t, analytics.EntryNum, 2)
	assert.Equal(t, analytics.Duration, time.Duration((3600+60)*1000*1000*1000))
}
