package timelog

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
)

func loadConfig(path string) (*Config, error) {
	configBytes := readfile(path)
	return Parse(configBytes)
}

func loadData(path string) ([]event, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	rows, _ := r.ReadAll()

	e := make([]event, 0, 20)

	for _, row := range rows {
		at, err := ParseDateTime(row[2])
		if err != nil {
			return e, err
		}
		at = ToLocal(at)
		e = append(e, event{
			workspace: row[0],
			uuid:      uuid.MustParse(row[1]),
			name:      row[3],
			at:        at,
			comment:   row[4],
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
