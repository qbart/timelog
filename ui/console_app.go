package ui

import (
	"flag"
	"fmt"
	"strings"

	"github.com/qbart/timelog/cli"
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

		case "clear":
			app.print()
			cli.AreYouSure("Are you sure to clear all data?", func() {
				app.service.Export()
				fmt.Println("Done")
			}, func() {
				fmt.Println("Cancelled")
			})

		case "adjust":
			app.service.RunAdjustService()

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

func (ConsoleApp) getComment() string {
	return strings.Join(flag.Args()[1:], " ")
}
