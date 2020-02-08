package timelog

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	touchFile(ConfigPath())
	touchFile(DataPath())
}

func loadConfig() *Config {
	configBytes := readfile(ConfigPath())

	config, err := Parse(configBytes)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func loadData() []entry {
	f, err := os.Open(DataPath())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	rows, _ := r.ReadAll()

	e := make([]entry, 0, 20)

	for _, row := range rows {
		fromLogtime := logtimeDefaultFactory{}.NewLogTime(true)
		toLogtime := logtimeDefaultFactory{}.NewLogTime(true)

		from, err := ParseDateTime(row[0])
		if err != nil {
			log.Fatal(err)
		}
		fromLogtime.t = from

		to, err := ParseDateTime(row[1])
		if row[1] == "" {
			toLogtime.finished = false
		} else if err != nil {
			log.Fatal(err)
		} else {
			toLogtime.t = to
		}

		e = append(e, entry{
			from:    fromLogtime,
			to:      toLogtime,
			comment: row[2],
		})
	}

	return e
}

// Load reads config and all time entries from files.
func Load() *TimeLogger {
	config := loadConfig()
	entries := loadData()

	return &TimeLogger{
		config:  config,
		entries: entries,
		factory: logtimeDefaultFactory{},
	}
}

func readfile(path string) []byte {
	file, err := os.Open(ConfigPath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func touchFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(path)
	}
}
