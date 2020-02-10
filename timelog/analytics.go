package timelog

import (
	"math"
	"time"
)

// Analytics for timelog.
type Analytics struct {
	EntryNum int
	Hours    int
	Minutes  int
}

func calcAnalytics(ee []entry) Analytics {
	h, m := calcDuration(ee)
	return Analytics{
		EntryNum: len(ee),
		Hours:    h,
		Minutes:  m,
	}
}

func calcDuration(ee []entry) (int, int) {
	var sum time.Duration = 0
	for _, e := range ee {
		duration := e.to.t.Sub(e.from.t)
		sum += duration
	}
	return int(sum.Hours()), int(math.Mod(sum.Minutes(), 60))
}
