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

func calcAnalytics(ee []event) Analytics {
	h, m := calcDuration(ee)
	return Analytics{
		EntryNum: calcLen(ee),
		Hours:    h,
		Minutes:  m,
	}
}

func calcLen(ee []event) int {
	sum := 0
	for _, e := range ee {
		if e.name == "start" {
			sum++
		}
	}
	return sum
}

func calcDuration(ee []event) (int, int) {
	var sum time.Duration = 0
	for i := len(ee) - 2; i >= 0; i-- {
		sum += ee[i+1].at.Sub(ee[i].at)
	}
	return int(sum.Hours()), int(math.Mod(sum.Minutes(), 60))
}
