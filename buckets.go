package timex

import (
	"sort"
	"time"
)

// FindBucket whose [start, end) contains time.
// Buckets are defined as periods in-between consecutive time values in an array.
// The buckets are assumed to be sorted in ascending order.
func FindBucket(buckets []time.Time, ts time.Time) (time.Time, bool) {
	if len(buckets) < 2 {
		return time.Time{}, false
	}

	if ts.Before(buckets[0]) || !ts.Before(buckets[len(buckets)-1]) {
		return time.Time{}, false
	}

	if idx := sort.Search(len(buckets), func(i int) bool { return !buckets[i].Before(ts) }); idx < len(buckets) {
		if buckets[idx].Equal(ts) {
			return buckets[idx], true
		} else {
			if idx > 0 {
				return buckets[idx-1], true
			}
		}
	}

	return time.Time{}, false
}

// NewBuckets for time range [from, until) with time grain.
func NewBuckets(from, until time.Time, timeGrain time.Duration) []time.Time {
	buckets := make([]time.Time, 0, until.Sub(from)/timeGrain+1)

	for ts := from; !ts.After(until); ts = ts.Add(timeGrain) {
		buckets = append(buckets, ts)
	}

	return buckets
}
