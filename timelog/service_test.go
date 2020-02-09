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
		timelogger.factory = logtimeMockFactory{
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

		if assert.Equal(t, len(service.timelogger.entries), 2) {
			// 0
			assert.Equal(t, tl.entries[0].comment, "hello")
			assert.Equal(t, tl.entries[0].from, logtime{
				finished: true,
				t:        expectedTime,
			})
			assert.Equal(t, tl.entries[0].to, logtime{
				finished: true,
				t:        expectedTime,
			})
			// 1
			assert.Equal(t, tl.entries[1].comment, "world")
			assert.Equal(t, tl.entries[1].from, logtime{
				finished: true,
				t:        expectedTime,
			})
			assert.Equal(t, tl.entries[1].to, logtime{
				finished: true,
				t:        expectedTime,
			})
		}
	})
}

func withTmpDir(fn func(dir string)) {
	dir, _ := ioutil.TempDir("", "config-dir")
	defer os.RemoveAll(dir)
	fn(dir)
}
