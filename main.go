package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/qbart/timelog/timelog"
)

func main() {
	flag.Parse()
	timelogger := timelog.Load()
	service := timelog.NewService(timelogger)

	if flag.NArg() > 0 {
		switch flag.Arg(0) {
		case "start":
			service.Start(getComment())
			fmt.Println(service.String())

		case "stop":
			service.Stop()
			fmt.Println(service.String())

		case "export":
			fmt.Println(service.String())
			areYouSureToExport(func() {
				service.Export()
			})
		}
	} else {
		fmt.Println(service.String())
	}
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
