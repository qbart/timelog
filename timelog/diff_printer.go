package timelog

import (
	"fmt"
	"strings"
)

// DiffPrinter - timelogger diff printer.
type DiffPrinter struct {
	timeloggerOriginal *TimeLogger
	timeloggerModified *TimeLogger
}

// Print outputs timelog diff to stdout.
func (p *DiffPrinter) Print() {
	fmt.Println(p.String())
}

// String returns diff representation of timelog.
func (p *DiffPrinter) String() string {
	var sb strings.Builder

	for i, o := range p.timeloggerOriginal.entries {
		m := p.timeloggerModified.entries[i]

		fromIsDifferent := m.from.t != o.from.t
		toIsDifferent := m.to.t != o.to.t

		appendEntryString := func(
			sb *strings.Builder,
			e *entry,
			ch rune,
			last bool,
			fromIsDifferent, toIsDifferent bool,
		) {
			sb.WriteRune(ch)
			sb.WriteString(" ")
			sb.WriteString(e.FromDateString())
			sb.WriteString(" ")
			sb.WriteString(wrapDiff(e.FromTimeString(), fromIsDifferent))
			sb.WriteString(" ")
			sb.WriteString(wrapDiff(e.ToTimeString(), toIsDifferent))
			sb.WriteString(" ")
			sb.WriteString(e.comment)
			if !last {
				sb.WriteString("\n")
			}
		}

		last := i == len(p.timeloggerOriginal.entries)-1
		if fromIsDifferent || toIsDifferent {
			appendEntryString(&sb, &o, '-', false, fromIsDifferent, toIsDifferent)
			appendEntryString(&sb, &m, '+', last, fromIsDifferent, toIsDifferent)
		} else {
			appendEntryString(&sb, &m, ' ', last, false, false)
		}
	}

	return sb.String()
}

func wrapDiff(s string, diff bool) string {
	if diff {
		return fmt.Sprint("{", s, "}")
	}

	return s
}