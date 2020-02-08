package main

import (
	"flag"
	"fmt"
	"strings"

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
		case "stop":
			service.Stop()
			// case "export":
		}
		fmt.Println(service.String())
	}
}

func getComment() string {
	return strings.Join(flag.Args()[1:], " ")
}
