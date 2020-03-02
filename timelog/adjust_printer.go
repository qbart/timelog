package timelog

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// AdjustEvent sent when adjusting minutes.
type AdjustEvent struct {
	selected int
	minutes  int
}

// AdjustPrinter - stdout printer.
type AdjustPrinter struct {
	timelogger *TimeLogger
	selected   int
	adjust     chan AdjustEvent
}

// Print adjust lines for timelog to stdout.
func (p *AdjustPrinter) Print() {
	// if err := ui.Init(); err != nil {
	// 	log.Fatalf("Failed to initialize app: %v", err)
	// }
	// defer ui.Close()

	// w, h := ui.TerminalDimensions()
	// widget := widgets.NewParagraph()
	// widget.WrapText = false
	// widget.Border = false
	// widget.Text = p.String()
	// widget.Title = "[j][k][up][down] select, [h][l][left][right] -/+ minute"
	// widget.SetRect(0, 0, w, h)

	// ui.Render(widget)
	// uiEvents := ui.PollEvents()

	// quit := false
	// for {
	// 	select {
	// 	case ae := <-p.adjust:
	// 		p.adjustTime(ae)
	// 		widget.Text = p.String()
	// 		ui.Render(widget)

	// 	case e := <-uiEvents:
	// 		switch e.ID {
	// 		case "<C-c>", "<Enter>":
	// 			quit = true
	// 		case "j", "<Down>":
	// 			p.selected++
	// 		case "k", "<Up>":
	// 			p.selected--
	// 		case "h", "<Left>":
	// 			go func() {
	// 				p.adjust <- AdjustEvent{p.selected, -1}
	// 			}()
	// 		case "l", "<Right>":
	// 			go func() {
	// 				p.adjust <- AdjustEvent{p.selected, +1}
	// 			}()
	// 		}

	// 		n := len(p.timelogger.entries)
	// 		if n-1 >= 0 {
	// 			if !p.timelogger.entries[n-1].to.finished {
	// 				n = n - 1
	// 			}
	// 		}

	// 		if p.selected < 0 {
	// 			p.selected = 0
	// 		} else if p.selected > n {
	// 			p.selected = n
	// 		}

	// 		widget.Text = p.String()
	// 		ui.Render(widget)
	// 	}

	// 	if quit {
	// 		break
	// 	}
	// }
}

// String returns text representation of timelog.
func (p AdjustPrinter) String() string {
	var sb strings.Builder
	// style := color.New(color.BgBlue, color.FgYellow)
	// last := len(p.timelogger.entries) - 1

	// for i, e := range p.timelogger.entries {
	// 	begin := 2
	// 	mid := 2
	// 	end := 2
	// 	wrapFrom := false
	// 	wrapTo := false

	// 	if i == p.selected {
	// 		wrapFrom = true
	// 		begin = 1
	// 		mid = 1
	// 	}
	// 	if i == p.selected-1 {
	// 		wrapTo = true
	// 		mid = 1
	// 		end = 1
	// 	}

	// 	sb.WriteString(e.FromDateString())
	// 	sb.WriteString(strings.Repeat(" ", begin))
	// 	sb.WriteString(
	// 		colorize(
	// 			wrapBrackets(e.FromTimeString(), wrapFrom),
	// 			wrapFrom,
	// 			style,
	// 		),
	// 	)
	// 	sb.WriteString(strings.Repeat(" ", mid))
	// 	sb.WriteString(
	// 		colorize(
	// 			wrapBrackets(e.ToTimeString(), wrapTo),
	// 			wrapTo,
	// 			style,
	// 		),
	// 	)
	// 	sb.WriteString(strings.Repeat(" ", end))
	// 	sb.WriteString(e.comment)
	// 	if i != last {
	// 		sb.WriteString("\n")
	// 	}
	// }
	return sb.String()
}

func (p *AdjustPrinter) adjustTime(ae AdjustEvent) {
	tl, _ := p.timelogger.Adjust(map[int]int{ae.selected: ae.minutes})
	p.timelogger = tl
}

func colorize(s string, colorize bool, color *color.Color) string {
	if colorize {
		return fmt.Sprint("[", s, "]", "(fg:yellow,bg:black)")
	}

	return s
}
