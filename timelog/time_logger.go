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

// String returns text representation of timelog.
func (t *TimeLogger) String() string {
	var sb strings.Builder
	last := len(t.entries) - 1
	for i, e := range t.entries {
		sb.WriteString(e.from.t.Format("2006-01-02 15:04 "))
		if e.to.finished {
			sb.WriteString(e.to.t.Format("15:04 "))
		} else {
			sb.WriteString(e.to.t.Format("...   "))
		}
		sb.WriteString(e.comment)
		if i != last {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
