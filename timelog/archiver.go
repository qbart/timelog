package timelog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Archiver
type Archiver struct {
	config *Config
}

// Archive moves current data file to archive dir.
// archive dir is created if does not exist.
func (a *Archiver) Archive() (string, error) {
	err := mkdir(a.config.ArchiveDir())
	if err != nil {
		return "", err
	}

	ts := time.Now().Format("20060102150405-")
	name := fmt.Sprint(ts, filepath.Base(a.config.DataPath()))

	return name,
		os.Rename(
			a.config.DataPath(),
			filepath.Join(a.config.ArchiveDir(), name),
		)
}
