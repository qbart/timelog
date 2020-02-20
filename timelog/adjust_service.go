package timelog

import (
	"fmt"

	"github.com/qbart/timelog/cli"
)

// AdjustService data.
type AdjustService struct {
	timelogger *TimeLogger
}

// Run handles the service logic.
func (p *AdjustService) Run(success func()) {
	if len(p.timelogger.entries) == 0 {
		fmt.Println("Nothing to adjust")
		return
	}

	adjustPrinter := &AdjustPrinter{
		timelogger: p.timelogger,
		selected:   0,
		adjust:     make(chan AdjustEvent),
	}
	adjustPrinter.Print()

	fmt.Println("")
	diffPrinter := &ColoredDiffPrinter{
		diffPrinter: &DiffPrinter{
			timeloggerOriginal: p.timelogger,
			timeloggerModified: adjustPrinter.timelogger,
		},
	}
	diffPrinter.Print()

	cli.AreYouSure("Are you sure to apply changes?", func() {
		*p.timelogger = *adjustPrinter.timelogger
		success()
		fmt.Println("Changes applied")
	}, func() {
		fmt.Println("Cancelled")
	})
}
