package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/qbart/timelog/timelog"
)

func main() {
	flag.Parse()
	config := timelog.NewConfig(timelog.HomeDir())
	timelogger := timelog.NewTimeLogger(config)
	service := timelog.NewService(timelogger)
	ok, err := service.Load()
	if !ok {
		log.Fatal(err)
	}

	if flag.NArg() > 0 {
		switch flag.Arg(0) {
		case "start":
			service.Start(getComment())
			print(service)

		case "stop":
			service.Stop()
			print(service)

		case "export":
			print(service)
			areYouSureToExport(func() {
				service.Export()
			})
		}
	} else {
		print(service)
	}
}

func print(s *timelog.Service) {
	analytics := s.CalculateAnalytics()
	fmt.Println(analytics.EntryNum, " row(s)")
	fmt.Println("---")
	fmt.Println(s.String())
	fmt.Println("---")
	fmt.Print(int64(analytics.Duration.Hours()), "h", int64(analytics.Duration.Minutes()), "m")
}

func areYouSureToExport(yes func()) {
	red := color.New(color.FgRed)
	red.Print("Are you sure to export (local data will be cleared)? y/N: ")

	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	s = string(s[0])
	if s == "y" || s == "Y" {
		yes()
	}
}

func getComment() string {
	return strings.Join(flag.Args()[1:], " ")
}
