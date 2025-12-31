package timex

import "time"

// IsTimeZoneEqual in a sense of reporting the same time at a time point.
func IsTimeZoneEqual(a, b *time.Location, t time.Time) bool {
	if a == nil || b == nil {
		return a == b
	}

	if a.String() == b.String() {
		return true
	}

	_, offset1 := t.In(a).Zone()
	_, offset2 := t.In(b).Zone()

	return offset1 == offset2
}
