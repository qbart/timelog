package timelog

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCalcAnalytics(t *testing.T) {
// 	ee := []entry{
// 		entry{
// 			comment: "hello",
// 			from: logtime{
// 				t: makeTime("2020-01-15 22:01"),
// 			},
// 			to: logtime{
// 				t: makeTime("2020-01-15 22:05"),
// 			},
// 		},
// 		entry{
// 			comment: "world",
// 			from: logtime{
// 				t: makeTime("2020-01-15 22:05"),
// 			},
// 			to: logtime{
// 				t: makeTime("2020-01-15 23:16"),
// 			},
// 		},
// 	}
// 	analytics := calcAnalytics(ee)

// 	assert.Equal(t, analytics.EntryNum, 2)
// 	assert.Equal(t, analytics.Hours, 1)
// 	assert.Equal(t, analytics.Minutes, 15)
// }
