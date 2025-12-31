package timex

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestTimeParser(t *testing.T) {
	tests := []struct {
		ts time.Time
		s  string
	}{
		{
			ts: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			s:  "2023-01-01T00:00:00Z",
		},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			var v time.Time

			f := TimeParser(&v)
			if f == nil {
				t.Fatal("parser is nil")
			}

			if err := f(tc.s); err != nil {
				t.Error(err)
			}

			if !v.Equal(tc.ts) {
				t.Error(v, tc.ts)
			}
		})
	}

	t.Run("now", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			ts := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
			time.Sleep(time.Until(ts))

			var v time.Time

			f := TimeParser(&v)
			if f == nil {
				t.Fatal("parser is nil")
			}

			if err := f("now"); err != nil {
				t.Error(err)
			}

			if !v.Equal(ts) {
				t.Error(v, ts)
			}
		})
	})
}
