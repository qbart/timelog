package timelog

import (
	"strings"
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

// NewTimeLogger creates new time logger.
func NewTimeLogger(c *Config) *TimeLogger {
	return &TimeLogger{
		config:  c,
		entries: make([]entry, 0),
		factory: logtimeDefaultFactory{},
	}
}

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

// Export clears all entries.
func (t *TimeLogger) Export() {
	t.entries = make([]entry, 0, 10)
}

// Adjust takes adjustments map and applies time modifications based on provided values in minutes.
func (t *TimeLogger) Adjust(adjustments map[int]int) (*TimeLogger, error) {
	clone := &TimeLogger{
		config:  t.config,
		entries: make([]entry, len(t.entries)),
		factory: t.factory,
	}
	copy(clone.entries, t.entries)
	return clone, nil
}

func (e entry) String() string {
	var sb strings.Builder
	sb.WriteString(FormatDateTime(e.from.t))
	sb.WriteString(" ")
	if e.to.finished {
		sb.WriteString(FormatTime(e.to.t))
		sb.WriteString(" ")
	} else {
		sb.WriteString("...   ")
	}
	sb.WriteString(e.comment)
	return sb.String()
}
