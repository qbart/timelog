package timelog

import (
	"time"
)

const (
	formatDateTime = "2006-01-02 15:04"
	formatTime     = "15:04"
)

// ParseDateTime parses time in app-default format.
func ParseDateTime(value string) (time.Time, error) {
	return time.Parse(formatDateTime, value)
}

// FormatDateTime formats datetime to string.
func FormatDateTime(value time.Time) string {
	return value.Format(formatDateTime)
}

// FormatTime formats time to string.
func FormatTime(value time.Time) string {
	return value.Format(formatTime)
}
