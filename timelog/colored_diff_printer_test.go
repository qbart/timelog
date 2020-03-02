package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ColoredDiffPrinter_String_Empty(t *testing.T) {
	p := &ColoredDiffPrinter{
		diffPrinter: &DiffPrinter{
			timeloggerOriginal: &TimeLogger{events: []event{}},
			timeloggerModified: &TimeLogger{events: []event{}},
		}}

	result := p.String()

	expectedResult := ""

	assert.Equal(t, expectedResult, result)
}
