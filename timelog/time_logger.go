package timelog

import (
	"time"
)

// TimeLogger data.
type TimeLogger struct {
	config  *Config
	entries []entry
	factory logtimeFactory
}

type entry struct {
	comment string
	from    logtime
	to      logtime
}

type logtimeFactory interface {
	NewLogTime(finished bool) logtime
}

type logtimeDefaultFactory struct{}

func (logtimeDefaultFactory) NewLogTime(finished bool) logtime {
	return logtime{
		t:        time.Now(),
		finished: finished,
	}
}

type logtime struct {
	t        time.Time
	finished bool
}

func NewTimeLogger(config *Config) *TimeLogger {
	return &TimeLogger{
		config:  config,
		entries: make([]entry, 0, 10),
		factory: logtimeDefaultFactory{},
	}
}

// Start appends new time log entry closing last unclosed entry.
func (t *TimeLogger) Start(comment string) {
	entry := entry{
		comment: comment,
		from:    t.factory.NewLogTime(true),
		to:      t.factory.NewLogTime(false),
	}
	t.entries = append(t.entries, entry)
	if len(t.entries) >= 2 {
		prev := len(t.entries) - 2
		curr := len(t.entries) - 1
		if !t.entries[prev].to.finished {
			t.entries[prev].to.t = t.entries[curr].from.t
			t.entries[prev].to.finished = true
		}
	}
}

// Stop closes existing unfinished entry.
func (t *TimeLogger) Stop() {
	if len(t.entries) > 0 {
		t.entries[len(t.entries)-1].to.finished = true
	}
}
