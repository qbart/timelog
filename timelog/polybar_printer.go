package timelog

import (
	"bytes"
	"fmt"
	"html/template"
)

// PolybarPrinter - timelogger formatted for polybar output.
type PolybarPrinter struct {
	timelogger *TimeLogger
	format     string
}

type polybarItem struct {
	Comment         string // task comment
	Duration        string // last task duration
	Total           string // tasks total duration
	Count           int    // task count
	CountNotZero    bool   //
	TotalGtDuration bool   // true when total > duration
}

// Print outputs timelog diff to stdout.
func (p *PolybarPrinter) Print() {
	fmt.Println(p.String())
}

// String returns colored diff representation of two timelogs.
func (p *PolybarPrinter) String() string {
	var buf bytes.Buffer
	item := polybarItem{}

	t := p.timelogger
	if len(t.events) > 0 {
		for i := len(t.events) - 1; i >= 0; i-- {
			if t.events[i].name != "stop" {
				item.Comment = t.events[i].comment
				break
			}
		}
	}
	a := calcAnalytics(t)
	item.Count = a.EntryNum
	item.CountNotZero = a.EntryNum != 0
	if a.EntryNum > 0 {
		item.Duration = fmt.Sprint(a.LastDuration.Hours, "h", a.LastDuration.Minutes, "m")
		item.Total = fmt.Sprint(a.Hours, "h", a.Minutes, "m")
		item.TotalGtDuration = !(a.LastDuration.Hours == a.Hours && a.LastDuration.Minutes == a.Minutes)
	}

	tmpl, err := template.New("polybarItem").Parse(p.format)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&buf, item)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
