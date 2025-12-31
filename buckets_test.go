package timex_test

import (
	"slices"
	"testing"
	"time"

	"github.com/ndx-technologies/timex"
)

func TestFindBucket(t *testing.T) {
	buckets := []time.Time{
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
	}

	tests := []struct {
		name   string
		ts     time.Time
		bucket time.Time
		ok     bool
	}{
		{
			name:   "bucket start",
			ts:     time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			bucket: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			ok:     true,
		},
		{
			name:   "within a bucket",
			ts:     time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC),
			bucket: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			ok:     true,
		},
		{
			name:   "within a bucket different timezone",
			ts:     time.Date(2024, 1, 3, 0, 0, 0, 0, time.FixedZone("UTC+3", 3*60*60)),
			bucket: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			ok:     true,
		},
		{
			name: "before all buckets",
			ts:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			name:   "within last bucket",
			ts:     time.Date(2024, 1, 3, 12, 0, 0, 0, time.UTC),
			bucket: time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			ok:     true,
		},
		{
			name: "edge of last bucket",
			ts:   time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "after all buckets",
			ts:   time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			bucket, ok := timex.FindBucket(buckets, tc.ts)

			if ok != tc.ok {
				t.Error(tc.ok, ok)
			}
			if !bucket.Equal(tc.bucket) {
				t.Error(tc.bucket, bucket)
			}
		})
	}

	t.Run("empty buckets", func(t *testing.T) {
		bucket, ok := timex.FindBucket([]time.Time{}, time.Now())
		if ok {
			t.Error("FindBucket() ok = true, want false")
		}
		if !bucket.IsZero() {
			t.Error(bucket)
		}
	})
}

func TestNewBuckets(t *testing.T) {
	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)
	timeGrain := time.Hour * 24

	buckets := timex.NewBuckets(from, until, timeGrain)

	expectedBuckets := []time.Time{
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
	}
	if !slices.Equal(buckets, expectedBuckets) {
		t.Error(buckets)
	}
}
