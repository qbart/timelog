package timelog

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationService(t *testing.T) {
	withTmpDir(func(dir string) {
		config := NewConfig(dir)

		now := time.Now()
		timelogger := NewTimeLogger(config)
		timelogger.factory = timeMockFactory{
			now: now,
		}
		setupService := NewService(timelogger)
		setupService.Load()
		setupService.Start("hello")
		setupService.Start("world")
		setupService.Stop()

		tl := NewTimeLogger(config)
		service := NewService(tl)
		service.Load()

		expectedTime := time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			0,
			0,
			now.Location(),
		)

		if assert.Equal(t, len(service.timelogger.events), 3) {
			// 0
			assert.Equal(t, tl.events[0].name, "start")
			assert.Equal(t, tl.events[0].comment, "hello")
			assert.Equal(t, tl.events[0].at, expectedTime)
			// 1
			assert.Equal(t, tl.events[1].name, "start")
			assert.Equal(t, tl.events[1].comment, "world")
			assert.Equal(t, tl.events[1].at, expectedTime)
			// 2
			assert.Equal(t, tl.events[2].name, "stop")
			assert.Equal(t, tl.events[2].comment, "")
			assert.Equal(t, tl.events[2].at, expectedTime)
		}
	})
}

func withTmpDir(fn func(dir string)) {
	dir, _ := ioutil.TempDir("", "config-dir")
	defer os.RemoveAll(dir)
	fn(dir)
}
