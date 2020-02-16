package timelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ColoredDiffPrinter_String_Empty(t *testing.T) {
	p := &ColoredDiffPrinter{
		diffPrinter: &DiffPrinter{
			timeloggerOriginal: &TimeLogger{entries: []entry{}},
			timeloggerModified: &TimeLogger{entries: []entry{}},
		}}

	result := p.String()

	expectedResult := ""

	assert.Equal(t, expectedResult, result)
}
