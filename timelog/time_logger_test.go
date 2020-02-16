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

func TestAdjust_Adjust_ValidClone(t *testing.T) {
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
	}

	tl := TimeLogger{entries: entries}
	clone, err := tl.Adjust(map[int]int{})
	assert.Nil(t, err)
	assert.NotNil(t, clone)
	assert.Len(t, clone.entries, 1)
	assert.Equal(t, tl.config, clone.config)
	assert.Equal(t, tl.factory, clone.factory)
}

func TestAdjust_DurationDoesNotCrossOver_FinishedEntry(t *testing.T) {
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
				finished: true,
				t:        makeTime("2020-01-15 22:10"),
			},
		},
	}

	tl := TimeLogger{entries: entries}
	clone, err := tl.Adjust(map[int]int{
		0: -3, // adjust time in "0" point -3 minutes
		1: -6, // adjust time in "1" point -6 minutes
		2: 5,  // adjust time in "2" point +5 minutes
	})

	assert.Nil(t, err)
	assert.NotNil(t, clone)

	// original entries are not modified
	assert.Equal(t, tl.entries[0].from.t, makeTime("2020-01-15 22:00"))
	assert.Equal(t, tl.entries[0].to.t, makeTime("2020-01-15 22:05"))
	assert.Equal(t, tl.entries[1].from.t, makeTime("2020-01-15 22:05"))
	assert.Equal(t, tl.entries[1].to.t, makeTime("2020-01-15 22:10"))

	// cloned objects contains modified entries
	assert.Equal(t, clone.entries[0].from.t, makeTime("2020-01-15 21:57"))
	assert.Equal(t, clone.entries[0].to.t, makeTime("2020-01-15 21:59"))
	assert.Equal(t, clone.entries[1].from.t, makeTime("2020-01-15 21:59"))
	assert.Equal(t, clone.entries[1].to.t, makeTime("2020-01-15 22:15"))
}

func TestAdjust_DurationDoesNotCrossOver_UnfinishedEntry(t *testing.T) {
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
				t:        makeTime("2020-01-15 22:10"),
			},
		},
	}

	tl := TimeLogger{entries: entries}
	clone, err := tl.Adjust(map[int]int{
		0: -3, // adjust time in "0" point -3 minutes
		1: -6, // adjust time in "1" point -6 minutes
		2: 5,  // adjust time in "2" point +5 minutes
	})

	assert.Nil(t, err)
	assert.NotNil(t, clone)

	// original entries are not modified
	assert.Equal(t, tl.entries[0].from.t, makeTime("2020-01-15 22:00"))
	assert.Equal(t, tl.entries[0].to.t, makeTime("2020-01-15 22:05"))
	assert.Equal(t, tl.entries[1].from.t, makeTime("2020-01-15 22:05"))
	assert.Equal(t, tl.entries[1].to.t, makeTime("2020-01-15 22:10"))

	// cloned objects contains modified entries
	assert.Equal(t, clone.entries[0].from.t, makeTime("2020-01-15 21:57"))
	assert.Equal(t, clone.entries[0].to.t, makeTime("2020-01-15 21:59"))
	assert.Equal(t, clone.entries[1].from.t, makeTime("2020-01-15 21:59"))
	assert.Equal(t, clone.entries[1].to.t, makeTime("2020-01-15 22:10"))
}

func TestAdjust_DurationNegativeCrossOver_FinishedEntry(t *testing.T) {
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
				finished: true,
				t:        makeTime("2020-01-15 22:10"),
			},
		},
	}

	tl := TimeLogger{entries: entries}
	clone, err := tl.Adjust(map[int]int{
		1: -6, // adjust time in "1" point -6 minutes
	})

	assert.Nil(t, err)
	assert.NotNil(t, clone)

	// cloned objects contains modified entries
	assert.Equal(t, clone.entries[0].from.t, makeTime("2020-01-15 22:00"))
	assert.Equal(t, clone.entries[0].to.t, makeTime("2020-01-15 22:00"))   // -6m -> 22:00 (can't go lower than previous)
	assert.Equal(t, clone.entries[1].from.t, makeTime("2020-01-15 22:00")) // same as above
	assert.Equal(t, clone.entries[1].to.t, makeTime("2020-01-15 22:10"))
}

func TestAdjust_DurationNegativeCrossOver_UnfinishedEntry(t *testing.T) {
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
				t:        makeTime("2020-01-15 22:10"),
			},
		},
	}

	tl := TimeLogger{entries: entries}
	clone, err := tl.Adjust(map[int]int{
		2: -6, // adjust time in "2" point -6 minutes
	})

	assert.Nil(t, err)
	assert.NotNil(t, clone)

	// cloned objects contains modified entries
	assert.Equal(t, clone.entries[0].from.t, makeTime("2020-01-15 22:00"))
	assert.Equal(t, clone.entries[0].to.t, makeTime("2020-01-15 22:05"))
	assert.Equal(t, clone.entries[1].from.t, makeTime("2020-01-15 22:05"))
	assert.Equal(t, clone.entries[1].to.t, makeTime("2020-01-15 22:10"))
}

func TestAdjust_DurationPositiveCrossOver_FinishedEntry(t *testing.T) {
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
				finished: true,
				t:        makeTime("2020-01-15 22:10"),
			},
		},
	}

	tl := TimeLogger{entries: entries}
	clone, err := tl.Adjust(map[int]int{
		1: 6, // adjust time in "1" point +6 minutes
	})

	assert.Nil(t, err)
	assert.NotNil(t, clone)

	// cloned objects contains modified entries
	assert.Equal(t, clone.entries[0].from.t, makeTime("2020-01-15 22:00"))
	assert.Equal(t, clone.entries[0].to.t, makeTime("2020-01-15 22:10"))   // +6m -> 22:10 (can't go higher than next)
	assert.Equal(t, clone.entries[1].from.t, makeTime("2020-01-15 22:10")) // same as above
	assert.Equal(t, clone.entries[1].to.t, makeTime("2020-01-15 22:10"))
}

func TestAdjust_DurationPositiveCrossOver_UninishedEntry(t *testing.T) {
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
				t:        makeTime("2020-01-15 22:10"),
			},
		},
	}

	tl := TimeLogger{entries: entries}
	clone, err := tl.Adjust(map[int]int{
		1: 6, // adjust time in "1" point +6 minutes
	})

	assert.Nil(t, err)
	assert.NotNil(t, clone)

	// cloned objects contains modified entries
	assert.Equal(t, clone.entries[0].from.t, makeTime("2020-01-15 22:00"))
	assert.Equal(t, clone.entries[0].to.t, makeTime("2020-01-15 22:10"))   // +6m -> 22:10 (can't go higher than next)
	assert.Equal(t, clone.entries[1].from.t, makeTime("2020-01-15 22:10")) // same as above
	assert.Equal(t, clone.entries[1].to.t, makeTime("2020-01-15 22:10"))
}
