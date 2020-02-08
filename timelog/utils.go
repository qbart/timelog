package timelog

import (
	"time"
)

// ParseTime parses time in given format.
func ParseTime(value string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04", value)
}
