package timelog

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
)

func loadConfig(path string) (*Config, error) {
	configBytes := readfile(path)
	return Parse(configBytes)
}

func loadData(path string) ([]entry, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	rows, _ := r.ReadAll()

	e := make([]entry, 0, 20)

	for _, row := range rows {
		fromLogtime := logtimeDefaultFactory{}.NewLogTime(true)
		from, err := ParseDateTime(row[0])
		if err != nil {
			return e, err
		}
		fromLogtime.t = ToLocal(from)

		toLogtime := logtimeDefaultFactory{}.NewLogTime(true)
		to, err := ParseDateTime(row[1])
		if row[1] == "" {
			toLogtime.finished = false
		} else if err != nil {
			return e, err
		} else {
			toLogtime.t = ToLocal(to)
		}

		e = append(e, entry{
			from:    fromLogtime,
			to:      toLogtime,
			comment: row[2],
		})
	}

	return e, nil
}

func readfile(path string) []byte {
	file, err := os.Open(path)
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
