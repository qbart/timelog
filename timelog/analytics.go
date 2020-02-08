package timelog

import (
	"time"
)

// Analytics for timelog.
type Analytics struct {
	EntryNum int
	Duration time.Duration
}

func calcAnalytics(ee []entry) Analytics {
	return Analytics{
		EntryNum: len(ee),
		Duration: calcDuration(ee),
	}

}

func calcDuration(ee []entry) time.Duration {
	var sum time.Duration = 0
	for _, e := range ee {
		duration := e.to.t.Sub(e.from.t)
		sum += duration
	}
	return sum
}
