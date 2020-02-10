package timelog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type logtimeMockFactory struct {
	now time.Time
}

func (mock logtimeMockFactory) NewLogTime(finished bool) logtime {
	return logtime{
		t:        mock.now,
		finished: finished,
	}
}

func TestStart(t *testing.T) {
	tl := TimeLogger{
		entries: make([]entry, 0),
		factory: logtimeMockFactory{
			now: makeTime("2020-01-15 22:00"),
		},
	}

	tl.Start("hello")

	if assert.Equal(t, len(tl.entries), 1) {
		assert.Equal(t, tl.entries[0].comment, "hello")
		assert.Equal(t, tl.entries[0].from, logtime{
			finished: true,
			t:        makeTime("2020-01-15 22:00"),
		})
		assert.Equal(t, tl.entries[0].to, logtime{
			finished: false,
			t:        makeTime("2020-01-15 22:00"),
		})
	}
}

func TestStart_WithExistingFinishedEntry(t *testing.T) {
	entries := []entry{
		entry{
			comment: "hello",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:01"),
			},
		},
	}
	tl := TimeLogger{
		entries: entries,
		factory: logtimeMockFactory{
			now: makeTime("2020-01-15 22:02"),
		},
	}

	tl.Start("world")

	if assert.Equal(t, len(tl.entries), 2) {
		// 0
		assert.Equal(t, tl.entries[0].comment, "hello")
		assert.Equal(t, tl.entries[0].from, logtime{
			finished: true,
			t:        makeTime("2020-01-15 22:00"),
		})
		assert.Equal(t, tl.entries[0].to, logtime{
			finished: true,
			t:        makeTime("2020-01-15 22:01"),
		})
		// 1
		assert.Equal(t, tl.entries[1].comment, "world")
		assert.Equal(t, tl.entries[1].from, logtime{
			finished: true,
			t:        makeTime("2020-01-15 22:02"),
		})
		assert.Equal(t, tl.entries[1].to, logtime{
			finished: false,
			t:        makeTime("2020-01-15 22:02"),
		})
	}
}

func TestStart_WithExistingUnfinishedEntry(t *testing.T) {
	entries := []entry{
		entry{
			comment: "hello",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				finished: false,
				t:        makeTime("2020-01-15 22:01"),
			},
		},
	}
	tl := TimeLogger{
		entries: entries,
		factory: logtimeMockFactory{
			now: makeTime("2020-01-15 22:02"),
		},
	}

	tl.Start("world")

	if assert.Equal(t, len(tl.entries), 2) {
		// 0
		assert.Equal(t, tl.entries[0].comment, "hello")
		assert.Equal(t, tl.entries[0].from, logtime{
			finished: true,
			t:        makeTime("2020-01-15 22:00"),
		})
		assert.Equal(t, tl.entries[0].to, logtime{
			finished: true,
			t:        makeTime("2020-01-15 22:02"),
		})
		// 1
		assert.Equal(t, tl.entries[1].comment, "world")
		assert.Equal(t, tl.entries[1].from, logtime{
			finished: true,
			t:        makeTime("2020-01-15 22:02"),
		})
		assert.Equal(t, tl.entries[1].to, logtime{
			finished: false,
			t:        makeTime("2020-01-15 22:02"),
		})
	}
}

func TestStop(t *testing.T) {
	entries := []entry{
		entry{
			comment: "hello",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:01"),
			},
		},
	}
	tl := TimeLogger{
		entries: entries,
		factory: logtimeMockFactory{
			now: makeTime("2020-01-15 22:02"),
		},
	}

	tl.Stop()

	assert.Equal(t, len(tl.entries), 1)
	assert.Equal(t, tl.entries[0].comment, "hello")
	assert.Equal(t, tl.entries[0].from, logtime{
		finished: true,
		t:        makeTime("2020-01-15 22:00"),
	})
	assert.Equal(t, tl.entries[0].to, logtime{
		finished: true,
		t:        makeTime("2020-01-15 22:01"),
	})
}

func TestStop_WithUnfinishedEntry(t *testing.T) {
	entries := []entry{
		entry{
			comment: "hello",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				finished: false,
				t:        makeTime("2020-01-15 22:01"),
			},
		},
	}
	tl := TimeLogger{
		entries: entries,
		factory: logtimeMockFactory{
			now: makeTime("2020-01-15 22:02"),
		},
	}

	tl.Stop()

	assert.Equal(t, len(tl.entries), 1)
	assert.Equal(t, tl.entries[0].comment, "hello")
	assert.Equal(t, tl.entries[0].from, logtime{
		finished: true,
		t:        makeTime("2020-01-15 22:00"),
	})
	assert.Equal(t, tl.entries[0].to, logtime{
		finished: true,
		t:        makeTime("2020-01-15 22:01"),
	})
}

func TestExport(t *testing.T) {
	entries := []entry{
		entry{
			comment: "hello",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:00"),
			},
			to: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:05"),
			},
		},
		entry{
			comment: "world",
			from: logtime{
				finished: true,
				t:        makeTime("2020-01-15 22:05"),
			},
			to: logtime{
				finished: false,
				t:        makeTime("2020-01-15 22:05"),
			},
		},
	}

	tl := TimeLogger{entries: entries}

	assert.Len(t, tl.entries, 2)
	tl.Export()
	assert.Len(t, tl.entries, 0)
}

func makeTime(value string) time.Time {
	parsedTime, _ := time.Parse("2006-01-02 15:04", value)
	return parsedTime
}
