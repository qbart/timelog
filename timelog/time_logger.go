package timelog

import (
	"time"
)

// TimeLogger data.
type TimeLogger struct {
	config  *Config
	entries []entry
}

type entry struct {
	comment string
	from    time.Time
	to      time.Time
}

func NewTimeLogger(config *Config) *TimeLogger {
	return &TimeLogger{
		config:  config,
		entries: make([]entry, 0, 10),
	}
}

func (t *TimeLogger) Start(comment string) {

}

func (t *TimeLogger) Stop() {

}

func (t *TimeLogger) String() {

}
