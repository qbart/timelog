package timelog

import (
	"fmt"
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
	if err := mkdir(c.Dir); err != nil {
		panic(fmt.Sprintf("Failed to initialize dir: %s", c.Dir))
	}

	if err := touchFile(c.ConfigPath()); err != nil {
		panic(fmt.Sprintf("Failed to touch file: %s", c.ConfigPath()))
	}

	if err := touchFile(c.DataPath()); err != nil {
		panic(fmt.Sprintf("Failed to touch file: %s", c.DataPath()))
	}
}

// ConfigPath returns configuration file path.
func (c *Config) ConfigPath() string {
	return filepath.Join(c.Dir, "config.ini")
}

// DataPath returns data file path.
func (c *Config) DataPath() string {
	return filepath.Join(c.Dir, "data-default.csv")
}
