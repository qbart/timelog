package ui

import (
	"bufio"
	"flag"
        "fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/qbart/timelog/timelog"
)

// ConsoleApp implementation of App interface.
type ConsoleApp struct {
	service *timelog.Service
}

// NewConsoleApp creates new CLI app.
func NewConsoleApp(service *timelog.Service) App {
	return &ConsoleApp{
		service: service,
	}
}

// Run CLI app.
func (app *ConsoleApp) Run() {
	if flag.NArg() > 0 {
		switch flag.Arg(0) {
		case "start":
			app.service.Start(app.getComment())
			app.print()

		case "stop":
			app.service.Stop()
			app.print()

		case "export":
			app.print()
			app.areYouSure("Are you sure to export (local data will be cleared)?", func() {
				app.service.Export()
			})

		case "adjust":
			app.service.AdjustPrinter().Print()

		case "version":
                        fmt.Println("Version ", timelog.Version)
		}
	} else {
		app.print()
	}
}

func (app *ConsoleApp) print() {
	app.service.TextPrinter().Print()
}

func (ConsoleApp) areYouSure(msg string, yes func()) {
	red := color.New(color.FgRed)
	red.Print(msg, " y/N: ")

	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	s = string(s[0])
	if s == "y" || s == "Y" {
		yes()
	}
}

func (ConsoleApp) getComment() string {
	return strings.Join(flag.Args()[1:], " ")
}
