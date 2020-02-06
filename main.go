package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/qbart/timelog/timelog"
)

func main() {
	flag.Parse()
	config, err := timelog.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	timelogger := timelog.NewTimeLogger(config)
	if flag.NArg() > 0 {
		switch flag.Arg(0) {
		case "start":
			timelogger.Start(getComment())
		case "stop":
			timelogger.Stop()
		}
		fmt.Println(timelogger.String())
	}
}

func getComment() string {
	return strings.Join(flag.Args()[1:], " ")
}
