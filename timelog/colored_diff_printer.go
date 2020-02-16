package timelog

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// ColoredDiffPrinter - timelogger diff printer with color.
type ColoredDiffPrinter struct {
	diffPrinter *DiffPrinter
}

// Print outputs timelog diff to stdout.
func (p *ColoredDiffPrinter) Print() {
	fmt.Println(p.String())
}

// String returns colored diff representation of two timelogs.
func (p *ColoredDiffPrinter) String() string {
	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)

	diff := p.diffPrinter.String()
	lines := strings.Split(diff, "\n")

	re := regexp.MustCompile(`\[([\d:]+)\]`)

	if len(diff) > 0 {
		for i, line := range lines {
			switch line[0] {
			case '-':
				line = strings.Replace(line, "-", red.Sprint("-"), 1)
				line = re.ReplaceAllString(line, red.Sprint("$1"))
			case '+':
				line = strings.Replace(line, "+", green.Sprint("+"), 1)
				line = re.ReplaceAllString(line, green.Sprint("$1"))
			}
			lines[i] = line
		}
	}

	return strings.Join(lines, "\n")
}
