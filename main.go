package main

import (
	"github.com/qbart/timelog/timelog"
	"log"
)

func main() {
	config, err := timelog.Load()
	mgr := timelog.NewTimeLogger(config)
	log.Println(mgr)
	log.Println(err)
}
