package timelog

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Archiver_Archive(t *testing.T) {
	withTmpDir(func(dir string) {
		config := NewConfig(dir)
		WriteTextFile(dir, "data-default.csv", "hello,world")

		arch := Archiver{config: config}
		filename, _ := arch.Archive()
		path := filepath.Join(config.ArchiveDir(), filename)

		content, _ := ReadTextFile(path)

		assert.Equal(t, "hello,world", content)
	})
}
