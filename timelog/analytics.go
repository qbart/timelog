package timelog

import (
	"math"
	"time"
)

type Duration struct {
	Hours   int
	Minutes int
	Comment string
}

// Analytics for timelog.
type Analytics struct {
	EntryNum int
	Duration
	LastDuration   Duration
	PrefixDuration map[string]Duration
}

func calcAnalytics(t *TimeLogger) Analytics {
	h, m := calcDuration(t.events, t.factory.NewTime())
	lh, lm := calcLastDuration(t.events, t.factory.NewTime())
	prefix := calcPrefixDuration(t.events, t.factory.NewTime())
	return Analytics{
		EntryNum:       calcLen(t.events),
		Duration:       Duration{h, m, ""},
		LastDuration:   Duration{lh, lm, ""},
		PrefixDuration: prefix,
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

func calcPrefixDuration(ee []event, stopTime time.Time) map[string]Duration {
	prefixes := make(map[string]Duration, 0)
	return prefixes
}

func extractPrefix(comment string) string {
	return comment
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
