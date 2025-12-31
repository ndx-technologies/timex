package timex

import (
	"testing"
	"time"
)

func TestFloorDay(t *testing.T) {
	tests := []struct {
		name  string
		ts    time.Time
		tsDay time.Time
	}{
		{
			name:  "midnight",
			ts:    time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			tsDay: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "end of day",
			ts:    time.Date(2023, time.January, 1, 23, 59, 59, 999999999, time.UTC),
			tsDay: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "arbitrary time",
			ts:    time.Date(2023, time.January, 1, 12, 34, 56, 789, time.UTC),
			tsDay: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "different timezone",
			ts:    time.Date(2023, time.January, 1, 14, 0, 0, 0, time.FixedZone("Test", -8*60*60)),
			tsDay: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.FixedZone("Test", -8*60*60)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := FloorDay(tt.ts)
			if !d.Equal(tt.tsDay) {
				t.Error(d)
			}
		})
	}
}
