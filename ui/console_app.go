package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/qbart/timelog/cli"
	"github.com/qbart/timelog/timelog"
	"github.com/spf13/cobra"
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
	root := &cobra.Command{
		Use:   "timelog",
		Short: "Time logging in CLI",
		Run: func(cmd *cobra.Command, args []string) {
			app.print()
		},
	}

	start := &cobra.Command{
		Use:   "start [comment]",
		Short: "Starts a new time entry",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			comment := strings.Join(args, " ")
			app.service.Start(comment)
			app.print()
		},
	}
	root.AddCommand(start)

	stop := &cobra.Command{
		Use:   "stop",
		Short: "Stops given time entry",
		Run: func(cmd *cobra.Command, args []string) {
			app.service.Stop()
			app.print()
		},
	}
	root.AddCommand(stop)

	clear := &cobra.Command{
		Use:   "clear",
		Short: "Clears all entries. No backup.",
		Run: func(cmd *cobra.Command, args []string) {
			app.print()
			cli.AreYouSure("Are you sure to clear all data?", func() {
				app.service.Clear()
				fmt.Println("Done")
			}, func() {
				fmt.Println("Cancelled")
			})
		},
	}
	root.AddCommand(clear)

	adjust := &cobra.Command{
		Use:   "adjust",
		Short: "Adjusts time between entries",
		Run: func(cmd *cobra.Command, args []string) {
			app.service.RunAdjustService()
		},
	}
	root.AddCommand(adjust)

	archive := &cobra.Command{
		Use:   "archive",
		Short: "Archive data file",
		Long:  "File is moved to archive subfolder in config dir",
		Run: func(cmd *cobra.Command, args []string) {
			app.print()
			cli.AreYouSure("Sure to archive?", func() {
				path, _ := app.service.Archiver().Archive()
				fmt.Println("Archived to", path)
			}, func() {
				fmt.Println("Cancelled")
			})
		},
	}
	root.AddCommand(archive)

	qlist := &cobra.Command{
		Use:   "qlist",
		Short: "Prints all quicklist entries",
		Run: func(cmd *cobra.Command, args []string) {
			for _, s := range app.service.Quicklist() {
				fmt.Println(s)
			}
		},
	}
	root.AddCommand(qlist)

	version := &cobra.Command{
		Use:   "version",
		Short: "Prints software version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version", timelog.Version)
		},
	}
	root.AddCommand(version)

	polybar := &cobra.Command{
		Use:   "polybar",
		Short: "Polybar configuration",
	}
	root.AddCommand(polybar)

	polybarFormat := &cobra.Command{
		Use:   "format",
		Short: "Polybar configuration",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			app.service.PolybarPrinter(args[0]).Print()
		},
		Long: `
Comment         string // task comment
Duration        string // last task duration
Total           string // tasks total duration
Count           int    // task count
CountNotZero    bool   //
TotalGtDuration bool   // true when total > duration
		`,
	}
	polybar.AddCommand(polybarFormat)

	autocomplete := &cobra.Command{
		Use:   "autocomplete",
		Short: "Autocomplete for entries",
	}
	root.AddCommand(autocomplete)

	autocompleteCommands := &cobra.Command{
		Use:   "commands",
		Short: "List commands for autcomplete",
		Run: func(cmd *cobra.Command, args []string) {
			//TODO: read this from cobra
			fmt.Println("start")
			fmt.Println("stop")
			fmt.Println("adjust")
			fmt.Println("qlist")
			fmt.Println("clear")
			fmt.Println("version")
		},
	}
	autocomplete.AddCommand(autocompleteCommands)

	autocompleteQlist := &cobra.Command{
		Use:   "qlist",
		Short: "Quicklist for autocomplete",
		Run: func(cmd *cobra.Command, args []string) {
			for _, s := range app.service.Quicklist() {
				fmt.Println(s)
			}
		},
	}
	autocomplete.AddCommand(autocompleteQlist)

	autocompleteInstall := &cobra.Command{
		Use:   "install",
		Short: "Installation script for autocomplete",
		Run: func(cmd *cobra.Command, args []string) {
			app.service.InstallAutocomplete()
		},
	}
	autocomplete.AddCommand(autocompleteInstall)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (app *ConsoleApp) print() {
	app.service.TextPrinter().Print()
}
