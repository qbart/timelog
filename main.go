package main

import (
	"flag"
	"log"

	"github.com/qbart/timelog/timelog"
	"github.com/qbart/timelog/ui"
)

func main() {
	flag.Parse()

	// initialize service
	config := timelog.NewConfig(timelog.HomeDir())
	timelogger := timelog.NewTimeLogger(config)
	service := timelog.NewService(timelogger)

	// load data
	ok, err := service.Load()
	if !ok {
		log.Fatal(err)
	}

	app := ui.NewConsoleApp(service)
	app.Run()
}
