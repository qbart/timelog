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

func (s *Service) writeToFile() {
	f, err := os.Create(DataPath())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	for _, e := range s.timelogger.entries {
		to := ""
		if e.to.finished {
			to = FormatDateTime(e.to.t)
		}
		w.Write([]string{
			FormatDateTime(e.from.t),
			to,
			e.comment,
		})
	}
	w.Flush()
}
