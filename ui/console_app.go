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
	n := flag.NArg()
	if n > 0 {
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
				app.service.Clear()
				fmt.Println("Done")
			}, func() {
				fmt.Println("Cancelled")
			})

		case "autocomplete":
			switch flag.Arg(1) {
			case "install":
				app.service.InstallAutocomplete()

			case "commands":
				fmt.Println("start")
				fmt.Println("stop")
				fmt.Println("adjust")
				fmt.Println("qlist")
				fmt.Println("clear")
				fmt.Println("version")

			case "qlist":
				for _, s := range app.service.Quicklist() {
					fmt.Println(s)
				}
			}

		case "qlist":
			for _, s := range app.service.Quicklist() {
				fmt.Println(s)
			}

		case "adjust":
			app.service.RunAdjustService()

		case "archive":
			app.print()
			cli.AreYouSure("Sure to archive?", func() {
				path, _ := app.service.Archiver().Archive()
				fmt.Println("Archived to", path)
			}, func() {
				fmt.Println("Cancelled")
			})

		case "version":
			fmt.Println("Version ", timelog.Version)

		case "polybar":
			switch flag.Arg(1) {
			case "format":
				app.service.PolybarPrinter(flag.Arg(2)).Print()
			}
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
