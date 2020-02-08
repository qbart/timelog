package timelog

import (
	"os"
	"path/filepath"
)

// Config data.
type Config struct {
	Quicklist []string
}

// ConfigPath returns configuration file path.
func ConfigPath() string {
	return filepath.Join(
		home(),
		string(os.PathSeparator),
		"timelog.ini",
	)
}

// DataPath returns data file path.
func DataPath() string {
	return filepath.Join(
		home(),
		string(os.PathSeparator),
		".timelog.csv",
	)
}

func home() string {
	dir, _ := os.UserHomeDir()
	return dir
}
