package timelog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	config, err := Parse([]byte(`
	[quicklist]
	project
	tag
	`))

	assert.Nil(t, err)

	quicklist := config.Quicklist
	if assert.Equal(t, len(quicklist), 2) {
		assert.Equal(t, quicklist[0], "project")
		assert.Equal(t, quicklist[1], "tag")
	}
}
