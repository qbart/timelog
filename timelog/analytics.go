package timelog

// Analytics for timelog.
type Analytics struct {
	EntryNum int
	Hours    int
	Minutes  int
}

func calcAnalytics(ee []event) Analytics {
	h, m := calcDuration(ee)
	return Analytics{
		EntryNum: calcLen(ee),
		Hours:    h,
		Minutes:  m,
	}
}

func calcLen(ee []event) int {
	return 0
}

func calcDuration(ee []event) (int, int) {
	// var sum time.Duration = 0
	// for _, e := range ee {
	// 	duration := e.to.t.Sub(e.from.t)
	// 	sum += duration
	// }
	// return int(sum.Hours()), int(math.Mod(sum.Minutes(), 60))
	return 0, 0
}
