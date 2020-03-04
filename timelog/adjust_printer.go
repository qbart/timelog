package timelog

import (
	"fmt"
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	defer ui.Close()

	w, h := ui.TerminalDimensions()
	widget := widgets.NewParagraph()
	widget.WrapText = false
	widget.Border = false
	widget.Text = p.String()
	widget.Title = "[j][k][up][down] select, [h][l][left][right] -/+ minute"
	widget.SetRect(0, 0, w, h)

	ui.Render(widget)
	uiEvents := ui.PollEvents()

	quit := false
	for {
		select {
		case ae := <-p.adjust:
			p.adjustTime(ae)
			widget.Text = p.String()
			ui.Render(widget)

		case e := <-uiEvents:
			switch e.ID {
			case "<C-c>", "<Enter>":
				quit = true
			case "j", "<Down>":
				p.selected++
			case "k", "<Up>":
				p.selected--
			case "h", "<Left>":
				go func() {
					p.adjust <- AdjustEvent{p.selected, -1}
				}()
			case "l", "<Right>":
				go func() {
					p.adjust <- AdjustEvent{p.selected, +1}
				}()
			}

			n := len(p.timelogger.events)

			if p.selected < 0 {
				p.selected = 0
			} else if p.selected >= n {
				p.selected = n - 1
			}

			widget.Text = p.String()
			ui.Render(widget)
		}

		if quit {
			break
		}
	}
}

// String returns text representation of timelog.
func (p AdjustPrinter) String() string {
	var sb strings.Builder
	tokens := p.timelogger.Tokenize(false)

	begin := 0
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token.token == tkDate {
			begin = i
		} else if token.token == tkNewLine || token.token == tkEnd {
			prev := Token{}
			for j := begin; j <= i; j++ {
				t := tokens[j]
				selected := p.selected == t.eventIndex
				s := wrapBrackets(t.str, selected)
				if t.token == tkFromTime || t.token == tkToTime {
					s = fmt.Sprintf("%6v", s)
				}
				if t.token == tkSpace &&
					prev.token == tkFromTime &&
					prev.eventIndex == p.selected {
					continue
				}
				sb.WriteString(
					colorize(
						s,
						selected,
					),
				)
				if t.token == tkToTime && !selected {
					sb.WriteString(" ")
				}
				prev = t
			}
		}
	}

	return sb.String()
}

func (p *AdjustPrinter) adjustTime(ae AdjustEvent) {
	p.timelogger = p.timelogger.Adjust(map[int]int{ae.selected: ae.minutes})
}

func colorize(s string, colorize bool) string {
	if colorize {
		return fmt.Sprint("[", s, "]", "(fg:yellow,bg:black)")
	}

	return s
}
