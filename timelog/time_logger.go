package timelog

import (
	"time"

	"github.com/google/uuid"
)

// TimeLogger data.
type TimeLogger struct {
	config  *Config
	events  []event
	factory timelogFactory
}

type event struct {
	uuid    uuid.UUID
	name    string
	at      time.Time
	comment string
}

type timelogFactory interface {
	NewTime() time.Time
	NewUUID() uuid.UUID
}

type timelogDefaultFactory struct{}

// NewTimeLogger creates new time logger.
func NewTimeLogger(c *Config) *TimeLogger {
	return &TimeLogger{
		config:  c,
		events:  make([]event, 0),
		factory: timelogDefaultFactory{},
	}
}

func (timelogDefaultFactory) NewTime() time.Time {
	return time.Now()
}

func (timelogDefaultFactory) NewUUID() uuid.UUID {
	return uuid.New()
}

// Start appends new time log entry closing last unclosed entry.
func (t *TimeLogger) Start(comment string) {
	evt := event{
		uuid:    t.factory.NewUUID(),
		name:    "start",
		at:      t.factory.NewTime(),
		comment: comment,
	}
	t.events = append(t.events, evt)
}

// Stop closes existing unfinished entry.
func (t *TimeLogger) Stop() {
	evt := event{
		uuid:    t.factory.NewUUID(),
		name:    "stop",
		at:      t.factory.NewTime(),
		comment: "",
	}

	if len(t.events) > 0 {
		prev := t.events[len(t.events)-1]
		if prev.name != "stop" {
			t.events = append(t.events, evt)
		}
	}
}

// Clear clears all entries.
func (t *TimeLogger) Clear() {
	t.events = make([]event, 0, 10)
}

// Adjust takes adjustments map and applies time modifications based on provided values in minutes.
func (t *TimeLogger) Adjust(adjustments map[int]int) *TimeLogger {
	const notChanged = -1

	clone := &TimeLogger{
		config:  t.config,
		events:  make([]event, len(t.events)),
		factory: t.factory,
	}
	copy(clone.events, t.events)
	n := len(clone.events)

	// apply adjustments
	for i := 0; i < n; i++ {
		d := adjustments[i]
		if d != 0 {
			clone.events[i].at = clone.events[i].at.Add(minutes(d))
		}
	}

	// keep adjustments within allowed range
	for i := 0; i < n; i++ {
		d := adjustments[i]
		if d != 0 {
			prevDate := clone.events[i].at.Add(-minutes(60 * 24 * 7))
			nextDate := time.Now()
			if i-1 >= 0 {
				prevDate = clone.events[i-1].at
			}
			if i+1 < n {
				nextDate = clone.events[i+1].at
			}

			if clone.events[i].at.Before(prevDate) {
				clone.events[i].at = prevDate
			}
			if clone.events[i].at.After(nextDate) {
				clone.events[i].at = nextDate
			}
		}
	}

	return clone
}

const (
	tkDate = iota
	tkFromTime
	tkToTime
	tkComment
	tkNewLine
	tkSpace
	tkEnd
)

// Token represents timelog single part.
type Token struct {
	token      int
	str        string
	eventIndex int
}

// Equals checks if two tokens are identical.
func (t Token) Equals(other Token) bool {
	return t.token == other.token && t.str == other.str
}

// Tokenize generate list of tokens of how timelog output can be split.
func (t *TimeLogger) Tokenize(split bool) []Token {
	tokens := make([]Token, 0, 20)
	last := len(t.events) - 1

	for i := 0; i <= last; i++ {
		curr := t.events[i]
		if curr.name == "stop" {
			continue
		}

		next := event{
			name: "",
			at:   time.Now(),
		}
		if i+1 <= last {
			next = t.events[i+1]
		}

		crossDaySplit := split && next.name != "" && next.at.Day() != curr.at.Day()
		if crossDaySplit {
			tokens = append(tokens, Token{tkDate, curr.DateString(), -1})
			tokens = append(tokens, Token{tkSpace, " ", -1})
			tokens = append(tokens, Token{tkFromTime, curr.TimeString(), i})
			tokens = append(tokens, Token{tkSpace, " ", -1})
			tokens = append(tokens, Token{tkToTime, "23:59", -1})
			tokens = append(tokens, Token{tkSpace, " ", -1})
			tokens = append(tokens, Token{tkComment, curr.comment, -1})
			tokens = append(tokens, Token{tkNewLine, "\n", -1})
			//
			tokens = append(tokens, Token{tkDate, next.DateString(), -1})
			tokens = append(tokens, Token{tkSpace, " ", -1})
			tokens = append(tokens, Token{tkFromTime, "00:00", -1})
			tokens = append(tokens, Token{tkSpace, " ", -1})
			tokens = append(tokens, Token{tkToTime, next.TimeString(), i + 1})
			tokens = append(tokens, Token{tkSpace, " ", -1})
			tokens = append(tokens, Token{tkComment, curr.comment, -1})
			tokens = append(tokens, Token{tkNewLine, "\n", -1})

		} else {
			tokens = append(tokens, Token{tkDate, curr.DateString(), -1})
			tokens = append(tokens, Token{tkSpace, " ", -1})
			tokens = append(tokens, Token{tkFromTime, curr.TimeString(), i})
			tokens = append(tokens, Token{tkSpace, " ", -1})
			if next.name == "" {
				tokens = append(tokens, Token{tkToTime, "...  ", -1})
			} else {
				tokens = append(tokens, Token{tkToTime, next.TimeString(), i + 1})
			}
			tokens = append(tokens, Token{tkSpace, " ", -1})
			tokens = append(tokens, Token{tkComment, curr.comment, -1})
			tokens = append(tokens, Token{tkNewLine, "\n", -1})
		}
	}

	if last >= 0 {
		tokens[len(tokens)-1] = Token{tkEnd, "", -1}
	} else {
		tokens = append(tokens, Token{tkEnd, "", -1})
	}

	return tokens
}

func minutes(d int) time.Duration {
	return time.Duration(d * 60_000_000_000)
}

func (e event) DateString() string {
	return FormatDate(e.at)
}

func (e event) TimeString() string {
	return FormatTime(e.at)
}

func (e event) ToCsvRecord() []string {
	return []string{
		e.uuid.String(),
		FormatDateTime(e.at.UTC()),
		e.name,
		e.comment,
	}
}
