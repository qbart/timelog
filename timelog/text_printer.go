package timelog

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
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
	fmt.Println()

	data := make([][]string, 0)

	for _, p := range analytics.PrefixOrder {
		data = append(data, []string{p, analytics.PrefixDuration[p].TotalString()})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"prefix", "⌚"})
	table.SetFooter([]string{"", analytics.Duration.TotalString()})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}

// String returns text representation of timelog.
func (p *TextPrinter) String() string {
	var sb strings.Builder

	for _, token := range p.timelogger.Tokenize(true) {
		sb.WriteString(token.str)
	}

	return sb.String()
}
