package timelog

import (
	"fmt"
	"strings"
)

// TextPrinter - stdout printer.
type TextPrinter struct {
	timelogger *TimeLogger
}

// Print outputs timelog to stdout.
func (p *TextPrinter) Print() {
	analytics := calcAnalytics(p.timelogger.events)
	fmt.Println(analytics.EntryNum, " row(s)")
	fmt.Println("---")
	fmt.Println(p.String())
	fmt.Println("---")
	fmt.Print(analytics.Hours, "h", analytics.Minutes, "m", "\n")
}

// String returns text representation of timelog.
func (p *TextPrinter) String() string {
	var sb strings.Builder
	entries := []string{}
	events := p.timelogger.events
	last := len(p.timelogger.events) - 1

	for i := 0; i <= last; i++ {
		sb.Reset()
		curr := events[i]
		next := event{
			name: "",
		}
		if i+1 <= last {
			next = events[i+1]
		}
		if curr.name == "stop" {
			continue
		}

		sb.WriteString(curr.DateString())
		sb.WriteString(" ")
		sb.WriteString(curr.TimeString())
		sb.WriteString(" ")
		if next.name == "" {
			sb.WriteString("...  ")
		} else {
			sb.WriteString(next.TimeString())
		}
		sb.WriteString(" ")
		sb.WriteString(curr.comment)

		entries = append(entries, sb.String())
	}

	return strings.Join(entries, "\n")
}
