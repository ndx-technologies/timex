package timex

import "time"

func FloorDay(ts time.Time) time.Time {
	return time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location())
}
