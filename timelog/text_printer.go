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
	analytics := calcAnalytics(p.timelogger.entries)
	fmt.Println(analytics.EntryNum, " row(s)")
	fmt.Println("---")
	fmt.Println(p.String())
	fmt.Println("---")
	fmt.Print(analytics.Hours, "h", analytics.Minutes, "m", "\n")
}

// String returns text representation of timelog.
func (p *TextPrinter) String() string {
	var sb strings.Builder
	last := len(p.timelogger.entries) - 1
	for i, e := range p.timelogger.entries {
		sb.WriteString(e.String())
		if i != last {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
