package timelog

import (
	"math"
	"time"
)

// Analytics for timelog.
type Analytics struct {
	EntryNum    int
	Hours       int
	Minutes     int
	LastHours   int
	LastMinutes int
}

func calcAnalytics(t *TimeLogger) Analytics {
	h, m := calcDuration(t.events, t.factory.NewTime())
	lh, lm := calcLastDuration(t.events, t.factory.NewTime())
	return Analytics{
		EntryNum:    calcLen(t.events),
		Hours:       h,
		Minutes:     m,
		LastHours:   lh,
		LastMinutes: lm,
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

func calcDuration(ee []event, stopTime time.Time) (int, int) {
	var sum time.Duration

	clone := make([]event, len(ee))
	copy(clone, ee)

	if len(clone) > 0 {
		if clone[len(clone)-1].name != "stop" {
			clone = append(clone, event{
				name: "stop",
				at:   stopTime,
			})
		}
	}

	for i := len(clone) - 2; i >= 0; i-- {
		if clone[i].name == "stop" {
			continue
		}
		d := clone[i+1].at.Sub(clone[i].at)
		sum += d
	}

	return int(sum.Hours()), int(math.Mod(sum.Minutes(), 60))
}

func calcLastDuration(ee []event, stopTime time.Time) (int, int) {
	var sum time.Duration

	clone := make([]event, len(ee))
	copy(clone, ee)

	if len(clone) > 0 {
		if clone[len(clone)-1].name != "stop" {
			clone = append(clone, event{
				name: "stop",
				at:   stopTime,
			})
		}
	}

	for i := len(clone) - 2; i >= 0; i-- {
		if clone[i].name == "stop" {
			continue
		}
		d := clone[i+1].at.Sub(clone[i].at)
		sum += d
		break
	}

	return int(sum.Hours()), int(math.Mod(sum.Minutes(), 60))
}
