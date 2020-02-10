package timelog

import (
	"fmt"
	"strconv"
	"strings"
)

// AdjustPrinter - stdout printer.
type AdjustPrinter struct {
	timelogger *TimeLogger
}

// Print adjust lines for timelog to stdout.
func (p AdjustPrinter) Print() {
	fmt.Println(p.String())
}

// String returns text representation of timelog.
func (p AdjustPrinter) String() string {
	var sb strings.Builder
	last := len(p.timelogger.entries) - 1
	splitPoint := 0
	for i, e := range p.timelogger.entries {
		sb.WriteString("- ")
		sb.WriteString(strconv.Itoa(splitPoint))
		sb.WriteString(" -\n")
		splitPoint++

		sb.WriteString(e.String())
		if i == last {
			if e.to.finished {
				sb.WriteString("\n- ")
				sb.WriteString(strconv.Itoa(splitPoint))
				sb.WriteString(" -")
			}
		} else {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
