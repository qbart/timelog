package timelog

import (
	"os"
	"path/filepath"
)

// Config data.
type Config struct {
	Dir       string
	Quicklist []string
}

// NewConfig creates new config.
func NewConfig(dir string) *Config {
	config := &Config{
		Dir:       dir,
		Quicklist: make([]string, 0),
	}
	config.init()
	return config
}

// Init initialize config and data files.
func (c *Config) init() {
	touchFile(c.ConfigPath())
	touchFile(c.DataPath())
}

// ConfigPath returns configuration file path.
func (c *Config) ConfigPath() string {
	return filepath.Join(
		c.Dir,
		string(os.PathSeparator),
		"timelog.ini",
	)
}

// DataPath returns data file path.
func (c *Config) DataPath() string {
	return filepath.Join(
		c.Dir,
		string(os.PathSeparator),
		".timelog.csv",
	)
}

// HomeDir returns user home dir path.
func HomeDir() string {
	dir, _ := os.UserHomeDir()
	return dir
}

func touchFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(path)
	}
}
