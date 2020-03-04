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
	fmt.Println(analytics.EntryNum, "row(s)")
	fmt.Println("---")
	fmt.Println(p.String())
	fmt.Println("---")
	fmt.Print(analytics.Hours, "h", analytics.Minutes, "m", "\n")
}

// String returns text representation of timelog.
func (p *TextPrinter) String() string {
	var sb strings.Builder

	for _, token := range p.timelogger.Tokenize() {
		sb.WriteString(token.str)
	}

	return sb.String()
}
