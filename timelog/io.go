package timelog

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func mkdir(dirs ...string) error {
	dir := filepath.Join(dirs...)
	return os.MkdirAll(dir, 0755)
}

func touchFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		return f.Close()
	}

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// HomeDir returns user home dir path.
func HomeDir() string {
	dir, _ := os.UserHomeDir()
	return dir
}

// WriteTextFile writes text file to given path ensuring folders exist.
func WriteTextFile(dir string, file string, data string) error {
	err := mkdir(dir)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(dir, file), []byte(data), 0640)
	if err != nil {
		return err
	}
	return nil
}

// ReadTextFile reads text file from given.
func ReadTextFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	return string(b), err
}
