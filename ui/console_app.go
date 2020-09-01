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
		Short: "Stops active time entry",
		Run: func(cmd *cobra.Command, args []string) {
			app.service.Stop()
			app.print()
		},
	}
	root.AddCommand(stop)

	clear := &cobra.Command{
		Use:   "clear",
		Short: "Clears all entries",
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

Example:
  timelog polybar format "{{if .CountNotZero }}%{F#011814}%{B#24f5bf} {{.Comment}} %{B-}%{B#0adba6} {{.Duration}} %{B-}{{ if .TotalGtDuration}}%{B#08aa81} {{.Total}} %{B-}{{ end }}%{F-}{{ end }}"
		`,
	}
	polybar.AddCommand(polybarFormat)

	autocomplete := &cobra.Command{
		Use: "complete",
	}
	root.AddCommand(autocomplete)
	acBash := &cobra.Command{
		Use: "bash",
		Run: func(cmd *cobra.Command, args []string) {
			root.GenBashCompletion(os.Stdout)
		},
	}
	autocomplete.AddCommand(acBash)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (app *ConsoleApp) print() {
	app.service.TextPrinter().Print()
}
