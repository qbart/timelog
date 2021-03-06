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
	org := p.timeloggerOriginal.Tokenize(false)
	mod := p.timeloggerModified.Tokenize(false)

	//todo: simplify
	if len(org) == 1 && org[0].token == tkEnd {
		return ""
	}

	diff := false
	begin := 0

	for i := 0; i < len(org); i++ {
		o := org[i]
		m := mod[i]
		if !o.Equals(m) {
			diff = true
		}

		if o.token == tkDate {
			diff = false
			begin = i
		} else if o.token == tkNewLine || o.token == tkEnd {
			if diff {
				sb.WriteRune('-')
			} else {
				sb.WriteRune(' ')
			}
			for j := begin; j <= i; j++ {
				sb.WriteString(wrapBrackets(
					org[j].str,
					!org[j].Equals(mod[j]),
				))
			}
			if diff {
				if o.token == tkEnd {
					sb.WriteString("\n")
				}
				sb.WriteRune('+')
				for j := begin; j <= i; j++ {
					modded := !mod[j].Equals(org[j])
					if modded {
						sb.WriteString(wrapBrackets(
							mod[j].str,
							true,
						))

					} else {
						switch mod[j].token {
						case tkNewLine:
							sb.WriteString(mod[j].str)
						case tkEnd:
							// nothing
						default:
							sb.WriteString(strings.Repeat(" ", len(mod[j].str)))
						}

					}
				}
			}
		}
	}

	return sb.String()
}

func wrapBrackets(s string, wrap bool) string {
	if wrap {
		return fmt.Sprint("[", s, "]")
	}

	return s
}
