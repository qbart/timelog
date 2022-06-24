package timelog

import (
	"fmt"
	"strings"

	"github.com/wzshiming/ctc"
)

// TextPrinter - stdout printer.
type TextPrinter struct {
	timelogger *TimeLogger
}

// Print outputs timelog to stdout.
func (p *TextPrinter) Print() {
	analytics := calcAnalytics(p.timelogger)
	fmt.Println(analytics.EntryNum, "row(s)")
	fmt.Println("---")
	fmt.Println(p.String())

	fmt.Println(strings.Repeat("-", 22))
	for _, prefix := range analytics.PrefixOrder {
		duration := analytics.PrefixDuration[prefix].TotalString()
		fmt.Println(fmt.Sprintf("%22s", duration), prefix)
	}
	fmt.Println(fmt.Sprintf("%22s", "------"))
	total := analytics.Duration.TotalString()
	fmt.Println(fmt.Sprint(ctc.ForegroundYellow, fmt.Sprintf("%22s", total), ctc.Reset))
}

// String returns text representation of timelog.
func (p *TextPrinter) String() string {
	var sb strings.Builder

	for _, token := range p.timelogger.Tokenize(true) {
		sb.WriteString(token.str)
	}

	return sb.String()
}
