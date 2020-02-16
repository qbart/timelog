package timelog

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// AdjustPrinter - stdout printer.
type AdjustPrinter struct {
	timelogger *TimeLogger
	selected   int
}

// Print adjust lines for timelog to stdout.
func (p AdjustPrinter) Print() {
	p.selected = 0

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(p.String())
		key, _ := reader.ReadByte()

		if key == 'j' {

		}
	}
}

// String returns text representation of timelog.
func (p AdjustPrinter) String() string {
	style := color.New(color.BgBlue, color.FgYellow)
	var sb strings.Builder
	last := len(p.timelogger.entries) - 1

	for i, e := range p.timelogger.entries {
		begin := 2
		mid := 2
		end := 2
		wrapFrom := false
		wrapTo := false

		if i == p.selected {
			wrapFrom = true
			begin = 1
			mid = 1
		}
		if i == p.selected-1 {
			wrapTo = true
			mid = 1
			end = 1
		}

		sb.WriteString(e.FromDateString())
		sb.WriteString(strings.Repeat(" ", begin))
		sb.WriteString(
			colorize(
				wrapBrackets(e.FromTimeString(), wrapFrom),
				wrapFrom,
				style,
			),
		)
		sb.WriteString(strings.Repeat(" ", mid))
		sb.WriteString(
			colorize(
				wrapBrackets(e.ToTimeString(), wrapTo),
				wrapTo,
				style,
			),
		)
		sb.WriteString(strings.Repeat(" ", end))
		sb.WriteString(e.comment)
		if i == last {
		} else {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func colorize(s string, colorize bool, color *color.Color) string {
	if colorize {
		return color.Sprint(s)
	}

	return s
}
