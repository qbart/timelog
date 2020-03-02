package timelog

import (
	"time"
)

// TimeLogger data.
type TimeLogger struct {
	config  *Config
	events  []event
	factory timeFactory
}

type event struct {
	name    string
	at      time.Time
	comment string
}

type timeFactory interface {
	New() time.Time
}

type timeDefaultFactory struct{}

// NewTimeLogger creates new time logger.
func NewTimeLogger(c *Config) *TimeLogger {
	return &TimeLogger{
		config:  c,
		events:  make([]event, 0),
		factory: timeDefaultFactory{},
	}
}

func (timeDefaultFactory) New() time.Time {
	return time.Now()
}

// Start appends new time log entry closing last unclosed entry.
func (t *TimeLogger) Start(comment string) {
	evt := event{
		name:    "start",
		at:      t.factory.New(),
		comment: comment,
	}
	t.events = append(t.events, evt)
}

// Stop closes existing unfinished entry.
func (t *TimeLogger) Stop() {
	evt := event{
		name:    "stop",
		at:      t.factory.New(),
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
func (t *TimeLogger) Adjust(adjustments map[int]int) (*TimeLogger, error) {
	const notChanged = -1

	clone := &TimeLogger{
		config:  t.config,
		events:  make([]event, len(t.events)),
		factory: t.factory,
	}
	// copy(clone.entries, t.entries)

	// n := len(clone.entries)
	// if n == 0 {
	// 	return clone, nil
	// }

	// for i := 0; i <= n; i++ {
	// 	d := adjustments[i]
	// 	if d != 0 {
	// 		from := notChanged
	// 		to := notChanged
	// 		if i == 0 {
	// 			from = 0
	// 		} else if i == n {
	// 			to = i - 1
	// 			if !clone.entries[to].to.finished {
	// 				to = notChanged
	// 			}
	// 		} else {
	// 			to = i - 1
	// 			from = i
	// 		}

	// 		if from != notChanged {
	// 			clone.entries[from].from.t = clone.entries[from].from.t.Add(minutes(d))
	// 		}

	// 		if to != notChanged {
	// 			clone.entries[to].to.t = clone.entries[to].to.t.Add(minutes(d))
	// 		}
	// 	}
	// }

	// for i := 0; i <= n; i++ {
	// 	d := adjustments[i]
	// 	if d != 0 {
	// 		// A,[from,to],A,[from,to],A
	// 		p1 := i*3 - 1
	// 		n1 := i*3 + 1

	// 		if d < 0 {
	// 			if p1 >= 0 {
	// 				if clone.entries[p1/3].to.t.Before(clone.entries[p1/3].from.t) {
	// 					t := clone.entries[p1/3].from.t
	// 					clone.entries[p1/3].to.t = t
	// 					clone.entries[n1/3].from.t = t
	// 				}
	// 			}
	// 		} else {
	// 			if n1/3 < n {
	// 				if clone.entries[n1/3].from.t.After(clone.entries[n1/3].to.t) {
	// 					t := clone.entries[n1/3].to.t
	// 					clone.entries[n1/3].from.t = t
	// 					clone.entries[p1/3].to.t = t
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return clone, nil
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
	token int
	str   string
}

// Equals checks if two tokens are identical.
func (t Token) Equals(other Token) bool {
	return t.token == other.token && t.str == other.str
}

// Tokenize generate list of tokens of how timelog output can be split.
func (t *TimeLogger) Tokenize() []Token {
	tokens := make([]Token, 0, 20)
	last := len(t.events) - 1

	for i := 0; i <= last; i++ {
		curr := t.events[i]
		next := event{
			name: "",
		}
		if i+1 <= last {
			next = t.events[i+1]
		}
		if curr.name == "stop" {
			continue
		}
		tokens = append(tokens, Token{tkDate, curr.DateString()})
		tokens = append(tokens, Token{tkSpace, " "})
		tokens = append(tokens, Token{tkFromTime, curr.TimeString()})
		tokens = append(tokens, Token{tkSpace, " "})

		if next.name == "" {
			tokens = append(tokens, Token{tkToTime, "...  "})
		} else {
			tokens = append(tokens, Token{tkToTime, next.TimeString()})
		}
		tokens = append(tokens, Token{tkSpace, " "})
		tokens = append(tokens, Token{tkComment, curr.comment})
		tokens = append(tokens, Token{tkNewLine, "\n"})
	}

	if last >= 0 {
		tokens[len(tokens)-1] = Token{tkEnd, ""}
	} else {
		tokens = append(tokens, Token{tkEnd, ""})
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
		FormatDateTime(e.at.UTC()),
		e.name,
		e.comment,
	}
}
