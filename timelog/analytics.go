package timelog

import (
	"fmt"
	"math"
	"time"
)

type Duration struct {
	Hours   int
	Minutes int
	Comment string
}

func (d Duration) TotalString() string {
	return fmt.Sprint(d.Hours, "h", d.Minutes, "m")
}

// Analytics for timelog.
type Analytics struct {
	EntryNum int
	Duration
	LastDuration   Duration
	PrefixDuration map[string]Duration
	PrefixOrder    []string
}

func calcAnalytics(t *TimeLogger) Analytics {
	h, m := calcDuration(t.events, t.factory.NewTime())
	lh, lm := calcLastDuration(t.events, t.factory.NewTime())
	prefix, order := calcPrefixDuration(t.events, t.factory.NewTime())
	return Analytics{
		EntryNum:       calcLen(t.events),
		Duration:       Duration{h, m, ""},
		LastDuration:   Duration{lh, lm, ""},
		PrefixDuration: prefix,
		PrefixOrder:    order,
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

func calcPrefixDuration(ee []event, stopTime time.Time) (map[string]Duration, []string) {
	prefixes := make(map[string]time.Duration, 0)
	order := make([]string, 0)

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

	for i := 0; i < len(clone)-1; i++ {
		if clone[i].name == "stop" {
			continue
		}

		d := clone[i+1].at.Sub(clone[i].at)
		prefix := extractPrefix(clone[i].comment)
		prefixes[prefix] += d

		found := false
		for _, x := range order {
			if x == prefix {
				found = true
				break
			}
		}
		if !found {
			order = append(order, prefix)
		}
	}
	durations := make(map[string]Duration, 0)
	for k, v := range prefixes {
		h, m := int(v.Hours()), int(math.Mod(v.Minutes(), 60))
		durations[k] = Duration{h, m, ""}
	}

	return durations, order
}

func extractPrefix(comment string) string {
	for i, c := range comment {
		if c == '-' || c == ' ' {
			return comment[:i]
		}
	}

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
