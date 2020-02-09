package timelog

import (
	"encoding/csv"
	"log"
	"os"
)

// Service provides timeloggin functionallity and syncs data to files.
type Service struct {
	timelogger *TimeLogger
}

// NewService creates timelog service.
func NewService(timelogger *TimeLogger) *Service {
	return &Service{
		timelogger: timelogger,
	}
}

// Load reads config and all time entries from files.
func (s *Service) Load() (bool, error) {
	c, err := loadConfig(s.timelogger.config.ConfigPath())
	if err != nil {
		return false, err
	}

	e, err := loadData(s.timelogger.config.DataPath())
	if err != nil {
		return false, err
	}

	s.timelogger.entries = e
	s.timelogger.config.Quicklist = c.Quicklist

	return true, nil
}

// Start timelog and sync.
func (s *Service) Start(comment string) {
	s.timelogger.Start(comment)
	s.writeToFile()
}

// Stop timelog and sync.
func (s *Service) Stop() {
	s.timelogger.Stop()
	s.writeToFile()
}

// Export timelog and sync.
func (s *Service) Export() {
	s.timelogger.Stop()
	s.timelogger.Export()
	s.writeToFile()
}

// String returns timelog.
func (s *Service) String() string {
	return s.timelogger.String()
}

// CalculateAnalytics returns generated stats.
func (s *Service) CalculateAnalytics() Analytics {
	return calcAnalytics(s.timelogger.entries)
}

// writeToFile saves entries in UTC format.
func (s *Service) writeToFile() {
	f, err := os.Create(s.timelogger.config.DataPath())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	for _, e := range s.timelogger.entries {
		to := ""
		if e.to.finished {
			to = FormatDateTime(e.to.t.UTC())
		}
		w.Write([]string{
			FormatDateTime(e.from.t.UTC()),
			to,
			e.comment,
		})
	}
	w.Flush()
}
