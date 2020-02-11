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
	n := len(clone.entries)
	const notChanged = -1
	for i := 0; i <= n; i++ {
		d := adjustments[i]
		if d != 0 {
			from := notChanged
			to := notChanged
			if i == 0 {
				from = 0
			} else if i == 1 {
				to = i - 1
				from = i
			} else if i == n {
				to = i - 1
				if !clone.entries[to].to.finished {
					to = notChanged
				}
			}

			if from != notChanged {
				clone.entries[from].from.t = clone.entries[from].from.t.Add(minutes(d))
			}
			if to != notChanged {
				clone.entries[to].to.t = clone.entries[to].to.t.Add(minutes(d))
			}
		}
	}

	return clone, nil
}

func minutes(d int) time.Duration {
	return time.Duration(d * 60_000_000_000)
}

func (e entry) String() string {
  var sb strings.Builder

  sb.WriteString(FormatDateTime(e.from.t))
  sb.WriteString(" ")

  if e.to.finished {
    if  e.from.t.Day() != e.to.t.Day() {
      sb.WriteString("23:59 ")
      sb.WriteString(e.comment)

      sb.WriteString("\n")
      sb.WriteString(FormatDate(e.to.t))
      sb.WriteString(" 00:00 ")
      sb.WriteString(FormatTime(e.to.t))
      sb.WriteString(" ")
    } else {
      sb.WriteString(FormatTime(e.to.t))
      sb.WriteString(" ")
    }
  } else {
    sb.WriteString("...   ")
  }

  sb.WriteString(e.comment)
  return sb.String()
}
